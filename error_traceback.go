package ripo

import (
	"runtime"
	"strings"
)

type TracebackRecord interface {
	File() string
	Function() string
	FunctionLocal() string
	Line() int
}

type Traceback interface {
	SetHandlerName(handlerName string) Traceback
	Records() []TracebackRecord
	MapRecords() []map[string]interface{}
}

type tracebackRecordImp struct {
	file     string
	function string
	line     int
}

func (tr *tracebackRecordImp) File() string {
	return tr.file
}

func (tr *tracebackRecordImp) Function() string {
	return tr.function
}

func (tr *tracebackRecordImp) FunctionLocal() string {
	full := tr.function
	if full == "" {
		return ""
	}
	parts := strings.Split(full, ".")
	return parts[len(parts)-1]
}

func (tr *tracebackRecordImp) Line() int {
	return tr.line
}

type tracebackImp struct {
	handlerName string
	callers     []uintptr         // must set initially
	records     []TracebackRecord // will be set on-demand
}

func (t *tracebackImp) SetHandlerName(handlerName string) Traceback {
	if handlerName != t.handlerName {
		t.handlerName = handlerName
		t.records = nil
	}
	return t
}

func (t *tracebackImp) Callers() []uintptr {
	return t.callers
}

func (t *tracebackImp) Records() []TracebackRecord {
	if t.records != nil {
		return t.records
	}
	frames := runtime.CallersFrames(t.callers)
	records := []TracebackRecord{}
	processFrame := func(frame runtime.Frame) bool {
		if frame.Func == nil {
			return true
		}
		frameRecord := &tracebackRecordImp{
			file:     frame.File,
			function: frame.Function,
			line:     frame.Line,
		}
		records = append(records, frameRecord)
		if frame.Function == t.handlerName {
			return false
		}
		return true
	}
	for {
		frame, more := frames.Next()
		if !processFrame(frame) || !more {
			break
		}
	}
	t.records = records
	return records
}

func (t *tracebackImp) MapRecords() []map[string]interface{} {
	records := t.Records()
	mapRecords := make([]map[string]interface{}, len(records))
	for index, record := range records {
		mapRecords[index] = map[string]interface{}{
			"file":          record.File(),
			"function":      record.Function(),
			"functionLocal": record.FunctionLocal(),
			"line":          record.Line(),
		}
	}
	return mapRecords
}
