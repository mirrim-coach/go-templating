package template

import (
	// "fmt"
	// util "github.com/Pairin/go-util"
	"github.com/aymerick/raymond"
	// log "github.com/sirupsen/logrus"
	// "io/ioutil"
	// "reflect"
	// "strconv"
)

func values(input map[string]interface{}, options *raymond.Options) []interface{} {
	out := make([]interface{}, 0, len(input))

	for _, v := range input {
		out = append(out, v)
	}

	return out
}
