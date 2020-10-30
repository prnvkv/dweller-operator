package controller

import (
	"github.com/prnvkv/operators/test-operator/pkg/controller/dweller"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, dweller.Add)
}
