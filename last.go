package template

import (
	// "fmt"
	// util "github.com/Pairin/go-util"
	"github.com/aymerick/raymond"
	// log "github.com/sirupsen/logrus"
	// "io/ioutil"
	"reflect"
	// "strconv"
)

func last(input interface{}, options *raymond.Options) interface{} {
	s := reflect.ValueOf(input)

	if s.Kind() != reflect.Slice || s.Len() == 0 {
		return ""
	}

	return s.Index(s.Len() - 1).Interface()
}

// Handlebars.registerHelper('last', function(array, options) {
//   return array[array.length - 1]
// });
