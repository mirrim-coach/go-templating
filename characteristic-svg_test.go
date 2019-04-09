package template

// Handlebars.registerHelper('characteristic_svg', function(id, isMatch, options) {
//   let svg_name = isMatch == 'true' ? `${id}-match` : id;
//   return `https://s3.amazonaws.com/assets.pairin.com/icons/characteristics/reports/hiring_readiness_report/${svg_name}.svg`;
// });

import (
	"testing"
)

func TestCharacteristicSvgMatch(t *testing.T) {

	ret := characteristicSvg("1", true, nil)

	if ret != "https://s3.amazonaws.com/assets.pairin.com/icons/characteristics/reports/hiring_readiness_report/1-match.svg" {
		t.Errorf("Expected characteristicSvg(1, true) to return the match url. Got: %s", ret)
	}
}
func TestCharacteristicSvgNoMatch(t *testing.T) {
	ret := characteristicSvg("1", false, nil)

	if ret != "https://s3.amazonaws.com/assets.pairin.com/icons/characteristics/reports/hiring_readiness_report/1.svg" {
		t.Errorf("Expected characteristicSvg(1, false) to return the icon url. Got: %s", ret)
	}
}
