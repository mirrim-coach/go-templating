package template

import (
	"fmt"
	"math"

	"github.com/aymerick/raymond"
)

//
// Handlebars.registerHelper('place_center_gap', function (score, options) {
//   let styles = ""
//   let width = 0;
//   let left = 0;
//
//   if (score > 0) {
//     styles += "left: 50%; ";
//     width = Math.abs(score) - 2;
//     styles += `width: ${width}%;`;
//   }
//   else if(score < 0) {
//     left = 50 - Math.abs(score) + 2;
//     width = 50 - (50 - Math.abs(score)) - 2;
//     styles += `left: ${left}%; `;
//     styles += `width: ${width}%;`;
//   }
//   return styles
// })

func placeCenterGap(score int, options *raymond.Options) string {
	absScore := math.Abs(float64(score))

	if score == 0 {
		return ""
	}

	if score > 0 {
		width := absScore - 2
		return fmt.Sprintf("left: 50%%; width: %d%%", int(width))
	}

	left := absScore - 2
	width := absScore - 2

	return fmt.Sprintf("left: %d%%; width: %d%%", int(left), int(width))
}
