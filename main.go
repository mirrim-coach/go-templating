package template

import (
	// "fmt"
	"io/ioutil"
	"os"

	util "github.com/mirrim-coach/go-util"
	"github.com/aymerick/raymond"
)

var helpers = map[string]interface{}{
	"ifCond":                          ifConditionHelper,
	"distinctArrayLength":             distinctArrayLength,
	"keysHelper":                      keys,
	"valuesHelper":                    values,
	"basicOperation":                  basicOperation,
	"place_range":                     placeRange,
	"place_score":                     placeScore,
	"place_center_gap":                placeCenterGap,
	"isEmployeeOrStudent":             isEmployeeOrStudent,
	"last":                            last,
	"joinBy":                          joinBy,
	"characteristic_svg":              characteristicSvg,
	"storedReportIsHiringOrInterview": storedReportIsHiringOrInterview,
	"score_page_classes":              scorePageClasses,
	"categoryIsHiring":                categoryIsHiring,
	"place_gap":                       placeGap,
	"arrayLength":                     arrayLength,
	"numTimes":                        numTimes,
	"numGroupsOf":                     numGroupsOf,
	"isBetween":                       isBetween,
	"sort_by":                         sortBy,
	"toFloat":                         toFloat,
}

//ParseTemplate fetches the requested template from the database and returns it
func ParseTemplate(layout string, data interface{}) (string, error) {
	template, err := raymond.Parse(string(layout))
	if err != nil {
		return "", err
	}

	for name, item := range helpers {
		template.RegisterHelper(name, item)
	}

	result, err := template.Exec(data)
	if err != nil {
		return "", err
	}

	if util.GetEnvVariable("DEBUG_TEMPLATE", "true") == "true" {
		util.Logger().Debug("Writing Template")
		os.MkdirAll("out/", 0777)
		ioutil.WriteFile("./index.html", []byte(result), 0777)
	}

	return string(result), nil
}
