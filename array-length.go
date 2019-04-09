package template

import (
	// "fmt"
	// util "github.com/Pairin/go-util"
	"github.com/aymerick/raymond"
	// "io/ioutil"
	"reflect"
)

func arrayLength(input interface{}, options *raymond.Options) int {
	s := reflect.ValueOf(input)

	if s.Kind() != reflect.Slice {
		return 0
	}

	return s.Len()
}
