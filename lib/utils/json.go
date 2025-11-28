package utils

import (
	"encoding/json"
	"log"
)

func Jout(v any) {
	log.Println(JsonPrettier(v))
}

func Json(v any) string {
	bs, _ := json.Marshal(v)
	return string(bs)
}

func JsonPrettier(v any) string {
	bs, _ := json.MarshalIndent(v, "", "  ")
	return string(bs)
}
