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

func LogFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
