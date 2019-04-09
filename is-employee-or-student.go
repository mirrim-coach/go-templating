package template

// Handlebars.registerHelper('isEmployeeOrStudent', function(recipient_type, options) {
//   return ['Employee', 'Student'].indexOf(recipient_type) !== -1;
// });

import (
	"github.com/aymerick/raymond"
)

func isEmployeeOrStudent(types []string, options *raymond.Options) bool {
	for _, k := range types {
		if k == "Employee" || k == "Student" {
			return true
		}
	}
	return false
}
