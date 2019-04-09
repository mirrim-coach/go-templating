package template

import (
	// "fmt"
	// util "github.com/Pairin/go-util"
	"github.com/aymerick/raymond"
	// log "github.com/sirupsen/logrus"
	// "io/ioutil"
	// "reflect"
	// "strconv"
	"fmt"
	"math"
)

func placeScore(input interface{}, options *raymond.Options) string {
	score, _ := getFloat(input)

	if score == 0 {
		return "zero"
	}

	v := "positive"

	if score < 0 {
		v = "negative"
	}

	s := math.Abs(float64(score))

	return fmt.Sprintf("%s score-%d", v, int(s))
}
