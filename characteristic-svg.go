package template

// Handlebars.registerHelper('characteristic_svg', function(id, isMatch, options) {
//   let svg_name = isMatch == 'true' ? `${id}-match` : id;
//   return `https://s3.amazonaws.com/assets.pairin.com/icons/characteristics/reports/hiring_readiness_report/${svg_name}.svg`;
// });

import (
	"fmt"
	"github.com/aymerick/raymond"
)

func characteristicSvg(id string, isMatch bool, options *raymond.Options) interface{} {
	name := id
	if isMatch {
		name = fmt.Sprintf("%s-match", id)
	}

	return fmt.Sprintf("https://s3.amazonaws.com/assets.pairin.com/icons/characteristics/reports/hiring_readiness_report/%s.svg", name)
}
