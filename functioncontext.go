package stubigo

import "fmt"

type functionContext struct {
	returnValues  *ConcurrentInterfaceArrayMap
	callArguments *ConcurrentInterfaceArrayMap
	name          string
}

func (f *functionContext) Returning(returnValues ...interface{}) {
	f.returnValues.Set(f.name, returnValues)
}

func (f *functionContext) GetArgumentCalledAt(index int) (interface{}, error) {
	calledArguments := f.callArguments.Get(f.name)

	if len(calledArguments) <= index {
		return nil, fmt.Errorf("requested parameter at index that is out of bounds")
	}

	return calledArguments[index], nil
}
