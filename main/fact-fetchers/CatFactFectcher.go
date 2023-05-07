package factfetchers

import (
	"encoding/json"
	"io"
	"net/http"
	errorhandlers "vk-test-task/main/error-handlers"
)

const catFactUrl = "https://catfact.ninja/fact"

type CatFact struct {
	Fact   string
	Length int
}

func FetchCatFact(ch chan string) {
	var fact CatFact
	client := http.DefaultClient
	response, err := client.Get(catFactUrl)
	errorhandlers.HandleFatal(err)

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	errorhandlers.HandleFatal(err)

	err = json.Unmarshal([]byte(body), &fact)
	errorhandlers.HandleFatal(err)

	ch <- fact.Fact
}
