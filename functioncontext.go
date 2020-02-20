package stubigo

type functionContext struct {
	returnValues *ConcurrentInterfaceArrayMap
	name         string
}

func (f *functionContext) Returning(returnValues ...interface{}) {
	f.returnValues.Set(f.name, returnValues)
}
