package template

import (
	// "fmt"
	// util "github.com/Pairin/go-util"
	"github.com/aymerick/raymond"
	// "io/ioutil"
	"reflect"
	"strconv"
)

func distinctArrayLength(input interface{}, options *raymond.Options) int {
	s := reflect.ValueOf(input)

	if s.Kind() != reflect.Slice {
		return 0
	}

	arr := make([]string, s.Len())
	for i := 0; i < s.Len(); i++ {
		f, _ := getFloat(s.Index(i).Interface())
		arr[i] = strconv.FormatFloat(f, 'E', -1, 64)
	}

	uniq := make([]string, 0, len(arr))
	found := make(map[string]bool)

	for _, v := range arr {

		if _, ok := found[v]; !ok {
			found[v] = true
			uniq = append(uniq, v)
		}
	}

	return len(uniq)
}
