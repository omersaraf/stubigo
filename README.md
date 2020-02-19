# Simple stub packages for Go

## Usage
### The stub
Let's say we want to stub our FileReader
```golang
type FileReader interface {
    Read(path string) string
}    
``` 
After choosing the interface you want to mock create new implementation
```golang
type fileReaderStub struct {
    stubigo.Stub
}
```
Now implement the interface methods using `stubigo.Called`
```golang
func (reader *fileReaderStub) Read(path string) (string, error) {
    res := reader.Called(path)
    return res.String(0, "default string"), res.Error(1, nil)
}
```
### The test

```golang
func NewFileReaderStub() *fileReaderStub {
    return &fileReaderStub{ NewStub() }
}

func TestStub(t *testing.T){
    stub := NewFileReaderStub()
    stub.With(stub.Read).Returning("test!")

    result, err := stub.Read("path")

    stub.Assert(t, stub.Read).CalledOnce()
    stub.Assert(t, stub.Read).CalledWith("path")

    // This would be the result after calling the stub 
    // assert.Equal(t, "test!", result)
    // assert.NoError(t, err)
}
```