package errors

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	e := A()
	fmt.Printf("%s\n\n", e)
	fmt.Printf("%v\n\n", e)
}

func A() error {
	return B()
}

func B() error {
	return C()
}

func C() error {
	return InternalServer("internal server error").
		WithError(fmt.Errorf("db connection error")).
		WithStack()
}
