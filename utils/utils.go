package utils

import (
	"encoding/json"
	"log"
)

func PrintJson(s interface{}) {
	empJSON, _ := json.MarshalIndent(s, "", "   ")
	log.Println(string(empJSON))
}
