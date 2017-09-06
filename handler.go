package restpc

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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

func TranslateHandler(handler Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r != nil && r.Body != nil {
				r.Body.Close()
			}
		}()
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "error in parsing form", http.StatusBadRequest)
			return
		}
		res, err := callHandler(handler, &requestImp{r: r})
		if err != nil {
			code := Unknown
			errorMsg := "Unknown" // FIXME: "unknown"
			rpcErr, isRpcErr := err.(RPCError)
			if isRpcErr {
				code = rpcErr.Code()
				errorMsg = rpcErr.Error() // FIXME: use a mapping or make it space-separated
			} else {
				log.Printf(
					"myrpc.TranslateHandler: handler '%v' returned non-rpc error: %#v\n",
					getFunctionName(handler),
					err,
				)
				rpcErr = NewError(Unknown, "", err)
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
			errorDispatcher(rpcErr)
			return
		}
		if res == nil {
			panic("TranslateHandler: func: err == nil && res == nil ")
		}
		wh := w.Header()
		if res.Header != nil {
			for key, values := range res.Header {
				for _, value := range values {
					wh.Add(key, value)
				}
			}
		}
		if res.Data != nil {
			resBodyBytes, err := json.Marshal(res.Data)
			if err != nil {
				log.Println("error in json.Marshal(res.Data):", err)
			} else {
				_, err := w.Write(resBodyBytes)
				if err != nil {
					log.Println("error in w.Write(resBodyBytes):", err)
				}
			}
		}
	}
}
