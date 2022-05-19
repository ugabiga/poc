package internal

import "encoding/json"

func PrettyJson(v any) []byte {
	convertedJson, err := json.MarshalIndent(v, "", "    ")
	PanicErr(err)
	return convertedJson
}

func PanicErr(err error) {
	if err != nil {
		panic(err)
	}
}
