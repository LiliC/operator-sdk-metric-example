package controller

import (
	"github.com/example-inc/metric-operator/pkg/controller/metricservice"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, metricservice.Add)
}
