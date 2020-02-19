package stubigo

type functionContext struct {
	returnValues map[string][]interface{}
	name         string
}

func (f *functionContext) Returning(returnValues ...interface{}) {
	f.returnValues[f.name] = returnValues
}
