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

func keys(input map[string]interface{}, options *raymond.Options) []string {
	out := make([]string, 0, len(input))

	for k := range input {
		out = append(out, k)
	}

	return out
}
