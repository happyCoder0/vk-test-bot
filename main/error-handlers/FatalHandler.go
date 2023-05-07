package errorhandlers

import "log"

func HandleFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
