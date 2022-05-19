package internal

import (
	"encoding/json"
	"log"
)

func PrettyJson(v any) []byte {
	convertedJson, err := json.MarshalIndent(v, "", "    ")
	LogFatal(err)
	return convertedJson
}

func PrintJSONLog(v any) {
	log.Println(string(PrettyJson(v)))
}

func LogFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
