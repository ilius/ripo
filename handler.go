package ripo

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"runtime"
)

type Handler func(req Request) (res *Response, err error)

func callHandler(handler Handler, request Request) (res *Response, err error) {
	defer func() {
		panicMsg := recover()
		if panicMsg != nil {
			err = NewError(
				Internal,
				Internal.String(),
				fmt.Errorf(
					"panic in handler %v: %v",
					getFunctionName(handler),
					panicMsg,
				),
			)
		}
	}()
	res, err = handler(request)
	return
}

func handleError(err error, handlerName string, w http.ResponseWriter, request ExtendedRequest) {
	code := Unknown
	errorMsg := "Unknown" // FIXME: "unknown"
	rpcErr, isRpcErr := err.(RPCError)
	if isRpcErr {
		code = rpcErr.Code()
		errorMsg = rpcErr.Error() // FIXME: use a mapping or make it space-separated
	} else {
		log.Printf(
			"myrpc.TranslateHandler: handler '%v' returned non-rpc error: %#v\n",
			handlerName,
			err,
		)
		rpcErr = NewError(
			Unknown, "", err,
		)
	}
	status := HTTPStatusFromCode(code)
	jsonByte, _ := json.Marshal(map[string]string{
		"code":  code.String(),
		"error": errorMsg,
	})
	http.Error(
		w,
		string(jsonByte),
		status,
	)
	errorDispatcher(request, rpcErr)
}

func TranslateHandler(handler Handler) http.HandlerFunc {
	handlerFuncObj := runtime.FuncForPC(reflect.ValueOf(handler).Pointer())
	handlerName := handlerFuncObj.Name()
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r != nil && r.Body != nil {
				r.Body.Close()
			}
		}()
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "error in parsing form", http.StatusBadRequest)
			return
		}
		request := &requestImp{
			r:           r,
			handlerName: handlerName,
		}
		res, err := callHandler(handler, request)
		if res == nil && err == nil {
			err = NewError(Internal, "", fmt.Errorf("handler %v returned nil response with nil error", handlerName))
		}
		if err != nil {
			handleError(err, handlerName, w, request)
			return
		}
		wh := w.Header()
		if res.Header != nil {
			for key, values := range res.Header {
				for _, value := range values {
					wh.Add(key, value)
				}
			}
		}
		if res.RedirectPath != "" {
			code := res.RedirectStatusCode
			if code == 0 {
				code = http.StatusSeeOther
			}
			http.Redirect(w, r, res.RedirectPath, code)
			return
		}
		var resBodyBytes []byte
		data := res.Data
		switch dataTyped := data.(type) {
		case []byte:
			resBodyBytes = dataTyped
		case string:
			resBodyBytes = []byte(dataTyped)
		default:
			if data == nil {
				data = map[string]any{}
			}
			jsonBytes, err := json.Marshal(data)
			if err != nil {
				log.Println("error in json.Marshal(res.Data):", err)
			} else {
				w.Header().Set("Content-Type", "application/json; charset=UTF-8")
				resBodyBytes = jsonBytes
			}
		}
		_, err = w.Write(resBodyBytes)
		if err != nil {
			log.Println("error in w.Write(resBodyBytes):", err)
		}
	}
}
