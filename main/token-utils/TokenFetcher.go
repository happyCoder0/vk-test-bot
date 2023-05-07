package tokenutils

import (
	"encoding/json"
	"io"
	"os"
	errorhandlers "vk-test-task/main/error-handlers"
)

type TokenData struct {
	Token string
}

func RetrieveToken() string {
	var data TokenData

	stream, err := os.Open("main/token.json")
	errorhandlers.HandleFatal(err)

	text, err := io.ReadAll(stream)
	errorhandlers.HandleFatal(err)

	json.Unmarshal([]byte(text), &data)

	return data.Token
}
