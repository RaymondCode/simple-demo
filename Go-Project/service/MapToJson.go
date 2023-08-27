package service

import "encoding/json"

func MapToJson(Mapstruct interface{}) string {
	// map转 json str
	jsonBytes, _ := json.Marshal(Mapstruct)
	jsonStr := string(jsonBytes)
	return jsonStr
}
