package utils

import (
	"encoding/json"
)

func ToJson(object interface{}) string {
	b, err := json.Marshal(object)
	if err != nil {
		print("error is: "+err.Error())
		return ""
	} else {
		res:= string(b)
		return res
	}
}