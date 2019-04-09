package template

import (
	// "fmt"
	// util "github.com/Pairin/go-util"
	"github.com/aymerick/raymond"
	// log "github.com/sirupsen/logrus"
	// "io/ioutil"
	// "math"
	// "reflect"
	// "strconv"
)

func basicOperation(in1 interface{}, operator string, in2 interface{}, options *raymond.Options) float64 {
	v1, _ := getFloat(in1)
	v2, _ := getFloat(in2)

	switch operator {
	case "+":
		return v1 + v2
	case "-":
		return v1 - v2
	case "*":
		return v1 * v2
	case "/":
		return v1 / v2
	}
	return 0
}
