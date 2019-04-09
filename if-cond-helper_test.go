package template

import (
	"testing"
)

var ifCondTests = []Test{
	Test{
		name:     "#ifCond - each with @index",
		input:    "{{#ifCond (toFloat 0) '!=' (toFloat '0')}}I should not show up!{{/ifCond}}",
		data:     map[string]interface{}{"goodbyes": []int{1, 2, 3}},
		privData: map[string]interface{}{"exclaim": ifConditionHelper},
		helpers:  map[string]interface{}{"ifCond": ifConditionHelper, "toFloat": toFloat},
		partials: nil,
		output:   "",
	},
	Test{
		name:     "#ifCond - each with @index",
		input:    "{{#ifCond (toFloat '1') '!=' (toFloat '0')}}I should show up!{{/ifCond}}",
		data:     map[string]interface{}{"goodbyes": []int{1, 2, 3}},
		privData: map[string]interface{}{"exclaim": ifConditionHelper},
		helpers:  map[string]interface{}{"ifCond": ifConditionHelper, "toFloat": toFloat},
		partials: nil,
		output:   "I should show up!",
	},
}

func TestIfCond(t *testing.T) {
	launchTests(t, ifCondTests)
}
