package template

import "testing"

//
// Those tests come from:
//   https://github.com/wycats/handlebars.js/blob/master/spec/builtin.js
//

var numTimesTests = []Test{
	Test{
		name:     "#numTimes - each with @index",
		input:    "{{#numTimes 2}}{{@index}}({{#each @root.goodbyes}}{{this}}{{/each}}){{/numTimes}}cruel !",
		data:     map[string]interface{}{"goodbyes": []int{1, 2, 3}},
		privData: map[string]interface{}{"exclaim": numTimes},
		helpers:  map[string]interface{}{"numTimes": numTimes},
		partials: nil,
		output:   "0(123)1(123)cruel !",
	},
}

func TestNumTimes(t *testing.T) {
	launchTests(t, numTimesTests)
}
