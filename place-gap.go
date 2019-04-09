package template

// Handlebars.registerHelper('place_gap', function(real, target, range, options) {
// 	let styles = ""
// 	let left_fix = null;
// 	if (range) {
//
// 		if ((range.end - range.start) < 3) {
// 			range = [range.start - 2, range.end + 2]
// 		}
//
// 		if (real < range.start) {
// 			target_score = range.start
// 			offset = 2
// 		}
// 		else {
// 			target_score = range.end
// 			offset = 2
// 			left_fix = 2
// 		}
// 	}
// 	else {
// 		target_score = target
// 		offset = 4.25
// 	}
//
// 	let left_offset = Math.min(real, target_score);
// 	left_fix = left_fix ? left_fix : 4.25
//
// 	let left_style = 47.75 + left_offset + left_fix;
// 	let right_style = Math.abs(real - target_score) - offset;
// 	styles += `left: ${left_style}%;`
// 	styles += `width: ${right_style}%;`
// 	return styles;
// })

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

func placeGap(realIn interface{}, targetIn interface{}, r map[string]interface{}, options *raymond.Options) string {
	_, sOk := r["start"]
	_, eOk := r["end"]

	real, _ := getFloat(realIn)
	target, _ := getFloat(targetIn)

	targetScore := 0.0
	offset := 4.0
	leftFix := 4.0

	if sOk && eOk {
		start, _ := getFloat(r["start"])
		end, _ := getFloat(r["end"])

		if (end - start) < 3 {
			start -= 2
			end += 2
		}

		if real < start {
			targetScore = start
		} else {
			targetScore = end
			leftFix = 2.0
		}
	} else {
		targetScore = target
		offset = 4.0
	}

	leftOffset := math.Min(float64(real), float64(targetScore))

	left := 47.75 + leftOffset + leftFix
	width := math.Abs(float64(real-targetScore)) - offset

	return fmt.Sprintf("left: %f%%; width: %f%%;", left, width)
}
