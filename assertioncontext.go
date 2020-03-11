package stubigo

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type AssertionContext struct {
	called          int
	calledArguments []interface{}
	t               *testing.T
}

type equalityFunction func(actual interface{}, expected interface{}) bool

func defaultEqualityFunction(item1 interface{}, item2 interface{}) bool {
	return item1 == item2
}

func (s *AssertionContext) CalledOnce() {
	s.Called(1)
}

func (s *AssertionContext) Called(times int) {
	assert.Equal(s.t, times, s.called, fmt.Sprintf("Expected number of callArguments to be %d but was actually %d", times, s.called))
}

func (s *AssertionContext) NotCalled() {
	s.Called(0)
}

func (s *AssertionContext) CalledWith(arguments ...interface{}) {
	lastInvocationArguments := s.getLastInvocationArguments()

	assert.LessOrEqual(s.t, len(arguments), len(lastInvocationArguments), fmt.Sprintf("Expected too many arguments to be called (expected %d, actual %d)", len(arguments), len(lastInvocationArguments)))
	for i := 0; i < len(arguments) && i < len(lastInvocationArguments); i++ {
		s.CalledWithArgumentAt(i, arguments[i])
	}
}

func (s *AssertionContext) CalledWithArgumentAt(index int, argument interface{}) {
	s.CalledWithArgumentAtWithEqualityFunction(index, argument, defaultEqualityFunction)
}

func (s *AssertionContext) CalledWithArgumentAtWithEqualityFunction(index int, argument interface{}, areEqual equalityFunction) {
	lastInvocationArguments := s.getLastInvocationArguments()
	actualCalledArgument := lastInvocationArguments[index]
	assert.True(s.t, areEqual(actualCalledArgument, argument), fmt.Sprintf("Expected function to be called with %v at index %d but called with %v", argument, index, actualCalledArgument))
}

func (s *AssertionContext) getLastInvocationArguments() []interface{} {
	return s.calledArguments

}
