package stubigo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type FileReader interface {
	Read(path string) string
}

type fileReaderStub struct {
	Stub
}

func (reader *fileReaderStub) Read(path string) (string, error) {
	res := reader.Called(path)
	return res.String(0, "default string"), res.Error(1, nil)
}

func NewFileReaderStub() *fileReaderStub {
	return &fileReaderStub{NewStub()}
}

func TestStub(t *testing.T) {
	stub := NewFileReaderStub()
	stub.With(stub.Read).Returning("test!")

	result, err := stub.Read("path")

	stub.Assert(t, stub.Read).CalledOnce()
	stub.Assert(t, stub.Read).CalledWith("path")
	assert.Equal(t, "test!", result)
	assert.NoError(t, err)
}
