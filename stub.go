package stubigo

import (
	"reflect"
	"runtime"
	"strings"
	"testing"
)

var Any = new(interface{})

type Stub struct {
	callArguments *ConcurrentInterfaceArrayMap
	numberOfCalls *ConcurrentIntMap
	outputs       *ConcurrentInterfaceArrayMap
}

func NewStub() Stub {
	return Stub{
		callArguments: NewConcurrentInterfaceArrayMap(),
		outputs:       NewConcurrentInterfaceArrayMap(),
		numberOfCalls: NewConcurrentIntMap(),
	}
}

func (s *Stub) With(function interface{}) *functionContext {
	name := getFunctionName(reflect.ValueOf(function).Pointer())
	return &functionContext{
		returnValues: s.outputs,
		name:         name,
	}
}

func (s *Stub) Returning(function interface{}, outputs ...interface{}) *Stub {
	name := getFunctionName(reflect.ValueOf(function).Pointer())
	s.outputs.Set(name, outputs)
	return s
}

func (s Stub) Assert(t *testing.T, function interface{}) *AssertionContext {
	functionName := getFunctionName(reflect.ValueOf(function).Pointer())
	inputs := s.callArguments.Get(functionName)
	numberOfCalls := s.numberOfCalls.Get(functionName)
	stub := &AssertionContext{
		called:          numberOfCalls,
		calledArguments: inputs,
		t:               t,
	}
	return stub
}

func (s Stub) Called(inputs ...interface{}) *Return {
	pc, _, _, _ := runtime.Caller(1)
	name := getFunctionName(pc)

	s.callArguments.Set(name, inputs)
	s.numberOfCalls.Increase(name)
	if values := s.outputs.Get(name); values != nil && len(values) > 0 {
		return &Return{values}
	}
	return &Return{make([]interface{}, 0)}
}

func getFunctionName(pointer uintptr) string {
	name := runtime.FuncForPC(pointer).Name()
	dot := strings.LastIndex(name, ".")
	function := name[dot+1:]
	function = strings.Split(function, "-")[0]
	if index := strings.LastIndex(function, "·"); index != -1 {
		return function[:index]
	}
	return function
}
