package template

import (
	// "fmt"
	// util "github.com/Pairin/go-util"

	"github.com/aymerick/raymond"

	// "io/ioutil"
	"reflect"
)

func ifConditionHelper(v1 interface{}, operator string, v2 interface{}, options *raymond.Options) string {
	switch operator {
	case "==":
		if reflect.DeepEqual(v1, v2) {
			return options.Fn()
		}
		return options.Inverse()
	case "===":
		if reflect.DeepEqual(v1, v2) {
			return options.Fn()
		}
		return options.Inverse()
	case "!=":
		if !reflect.DeepEqual(v1, v2) {
			return options.Fn()
		}
		return options.Inverse()
	case "!==":
		if v1 != v2 {
			return options.Fn()
		}
		return options.Inverse()
	case "<":
		f1, _ := getFloat(v1)
		f2, _ := getFloat(v2)
		if f1 < f2 {
			return options.Fn()
		}
		return options.Inverse()
	case "<=":
		f1, _ := getFloat(v1)
		f2, _ := getFloat(v2)
		if f1 <= f2 {
			return options.Fn()
		}
		return options.Inverse()
	case ">":
		f1, _ := getFloat(v1)
		f2, _ := getFloat(v2)
		if f1 > f2 {
			return options.Fn()
		}
		return options.Inverse()
	case ">=":
		f1, _ := getFloat(v1)
		f2, _ := getFloat(v2)
		if f1 >= f2 {
			return options.Fn()
		}
		return options.Inverse()
	case "&&":
		if raymond.IsTrue(v1) && raymond.IsTrue(v2) {
			return options.Fn()
		}
		return options.Inverse()
	case "||":
		if raymond.IsTrue(v1) || raymond.IsTrue(v2) {
			return options.Fn()
		}
		return options.Inverse()
	}
	return options.Inverse()
}
