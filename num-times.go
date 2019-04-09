package template

import (
	"github.com/aymerick/raymond"
)

// Handlebars.registerHelper('times', function(n, block) {
// 	var accum = '';
// 	for(var i = 0; i < n; ++i)
// 			accum += block.fn(i);
// 	return accum;
// });

func numTimes(num interface{}, options *raymond.Options) string {
	n, _ := getFloat(num)
	result := ""
	for i := 0; i < int(n); i++ {
		data_frame := options.NewDataFrame()
		data_frame.Set("index", i)
		data_frame.Set("first", i == 0)
		data_frame.Set("last", i == int(n)-1)
		result += options.FnCtxData(num, data_frame)
	}
	return result
}
