package template

import (
	"math"
	"sort"

	"github.com/aymerick/raymond"
)

// Handlebars.registerHelper('sort_by', function(array_of_objects, property_name, direction, options) {
// 	if(direction == 'asc') {
// 		return array_of_objects.sort((a, b) => (a[property_name] > b[property_name]) ? 1 : -1)
// 	} else {
// 		return array_of_objects.sort((a, b) => (a[property_name] < b[property_name]) ? 1 : -1)
// 	}
// })

func sortBy(object []interface{}, key string, direction string, options *raymond.Options) []map[string]interface{} {
	object_map := interfaceReflection(object)

	sort.Slice(object_map, func(i, j int) bool {
		i_value, _ := getFloat(object_map[i][key])
		j_value, _ := getFloat(object_map[j][key])

		if math.IsNaN(i_value) == true {
			i_value = float64(0)
			object_map[i][key] = i_value
		}
		if math.IsNaN(j_value) == true {
			j_value = float64(0)
			object_map[j][key] = i_value
		}
		if direction == "asc" {
			return i_value < j_value
		} else {
			return i_value > j_value
		}
	})
	return object_map
}
