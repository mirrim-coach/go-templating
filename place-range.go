package template

// Handlebars.registerHelper('place_range', function(range, options) {
//   console.log(range)
//   let graph_width = 588;
//   let styles = '';
//
//   let min = range.start
//   let max = range.end
//
//   if ((max - min) < 3) {
//     min = min - 2;
//     max = max + 2;
//   }
//   let left_circle_border = ((min) * 0.01) * 588;
//   let right_circle_border = ((max) * 0.01) * 588;
//   let size = Math.abs(right_circle_border - left_circle_border);
//   let half_size = size / 2;
//
//   styles += `width: ${size}pt;`
//   styles += `height: ${size}pt;`
//   styles += `border-radius: ${half_size}pt;`
//
//   styles += `left: ${50 + min}%;`
//   styles += `top: -${half_size - 13}pt;`
//   return styles;
// });

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

func placeRange(r map[string]interface{}, options *raymond.Options) string {
	_, sOk := r["start"]
	_, eOk := r["end"]

	if !sOk || !eOk {
		return ""
	}

	width := float64(588.0)
	min, _ := getFloat(r["start"])
	max, _ := getFloat(r["end"])

	size := math.Abs((min * 0.01 * width) - (max * 0.01 * width))

	if min-max < 3 {
		min -= 2.0
		max += 2.0
	}

	min += 50

	return fmt.Sprintf("width: %fpt; height: %fpt; border-radius: 50%%; left: %f%%; top: -%fpt;", size, size, min, (size/2)-13)
}
