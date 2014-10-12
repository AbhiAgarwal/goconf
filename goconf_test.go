package goconf

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	var returnData map[string]interface{}
	returnData = Parse("./test/go.conf")

	var expectedData map[string]interface{}
	expectedData = make(map[string]interface{})
	expectedData["HELLO_KEY"] = 4
	expectedData["API_NESS"] = 128328
	expectedData["GOOGLIE"] = "string"

	if !reflect.DeepEqual(returnData, expectedData) {
		t.Error("\n Got: ", returnData, "\n Expected: ", expectedData)
	}
}
