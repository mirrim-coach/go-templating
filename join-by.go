package template

// Handlebars.registerHelper('joinBy', function(array, delimiter, options) {
//   return array.join(String(delimiter));
// });

import (
	"github.com/aymerick/raymond"
	"reflect"
	"strconv"
	"strings"
)

func joinBy(input interface{}, delimiter string, options *raymond.Options) string {
	s := reflect.ValueOf(input)

	if s.Kind() != reflect.Slice {
		return ""
	}

	arr := make([]string, s.Len())
	for i := 0; i < s.Len(); i++ {
		if s.Index(i).Kind() == reflect.String {
			arr[i] = s.Index(i).String()
		} else if s.Index(i).Kind() == reflect.Int {
			arr[i] = strconv.FormatInt(s.Index(i).Int(), 10)
		}
	}

	return strings.Join(arr, delimiter)
}
