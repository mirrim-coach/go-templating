package template

import (
	"math"

	"github.com/aymerick/raymond"
)

func numGroupsOf(full_length_group interface{}, group_size interface{}, options *raymond.Options) float64 {

	full_length_group_size, _ := getFloat(full_length_group)
	float_group_size, _ := getFloat(group_size)
	return math.Ceil(full_length_group_size / float_group_size)
}
