package stubigo

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

type SomeInterface interface {
	SomeFunction(arg string) (int, error)
}

type InterfaceStub struct {
	Stub
}

func (stub *InterfaceStub) SomeFunction(arg string) (int, error) {
	called := stub.Called(arg)
	return called.Int(0, 1), called.Error(1, nil)
}

func interfaceUsage(someInterface SomeInterface) (int, error) {
	return someInterface.SomeFunction("test")
}

func TestStubigo(t *testing.T) {
	stub := &InterfaceStub{NewStub()}
	stub.With(stub.SomeFunction).Returning(10, fmt.Errorf("error"))

	v1, err := interfaceUsage(stub)

	stub.Assert(t, stub.SomeFunction).CalledWith("test")
	assert.Error(t, err, "error")
	assert.Equal(t, 10, v1)
}

func TestStubigo_WithContainsEqualityFunction(t *testing.T) {
	stub := &InterfaceStub{NewStub()}
	stub.With(stub.SomeFunction).Returning(10, fmt.Errorf("error"))
	contains := func(actual interface{}, expected interface{}) bool {
		actualString, _ := actual.(string)
		expectedString, _ := expected.(string)
		return strings.Contains(actualString, expectedString)
	}
	v1, err := interfaceUsage(stub)

	stub.Assert(t, stub.SomeFunction).CalledWithArgumentAtWithEqualityFunction(0, "te", contains)
	assert.Error(t, err, "error")
	assert.Equal(t, 10, v1)
}

func TestFunctionContext_GetArgumentCalledAt(t *testing.T) {
	stub := &InterfaceStub{NewStub()}
	stub.With(stub.SomeFunction)

	interfaceUsage(stub)

	firstArgument, _ := stub.With(stub.SomeFunction).GetArgumentCalledAt(0)
	assert.Equal(t, firstArgument, "test")
}
