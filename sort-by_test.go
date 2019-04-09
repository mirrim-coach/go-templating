package template

import (
	"testing"
)

func TestSortByKeyAsc(t *testing.T) {

	key, data := testSortByKeyDataSet()
	ret := sortBy(data, key, "asc", nil)

	expectation := []float64{0, 0, 15, 30}

	for i := 0; i < len(ret); i++ {
		if ret[i]["age"] != expectation[i] {
			t.Errorf("The %d st value should have been %f, but instead got %f", i, expectation[i], ret[i]["age"])
		}
	}

}

func TestSortByKeyDesc(t *testing.T) {

	key, data := testSortByKeyDataSet()
	ret := sortBy(data, key, "desc", nil)

	expectation := []float64{30, 15, 0, 0}

	for i := 0; i < len(ret); i++ {
		if ret[i]["age"] != expectation[i] {
			t.Errorf("The %d st value should have been %f, but instead got %f", i, expectation[i], ret[i]["age"])
		}
	}

}

func testSortByKeyDataSet() (string, []interface{}) {
	key := "age"
	data := []interface{}{
		map[string]interface{}{
			"name": "bob",
			"age":  float64(15),
		},
		map[string]interface{}{
			"name": "joe",
			"age":  nil,
		},
		map[string]interface{}{
			"name": "susan",
			"age":  float64(30),
		},
		map[string]interface{}{
			"name": "joe",
			"age":  nil,
		},
	}

	return key, data
}
