package factfetchers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	errorhandlers "vk-test-task/main/error-handlers"
)

const dogFactUrl = "https://dog-api.kinduff.com/api/facts?number=1"

type DogFact struct {
	Facts   []string
	Success bool
}

func FetchDogFact(ch chan string) {
	var fact DogFact
	client := http.DefaultClient
	response, err := client.Get(dogFactUrl)
	errorhandlers.HandleFatal(err)

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	errorhandlers.HandleFatal(err)

	log.Println(response.Body)

	err = json.Unmarshal([]byte(body), &fact)
	errorhandlers.HandleFatal(err)

	ch <- fact.Facts[0]
}
