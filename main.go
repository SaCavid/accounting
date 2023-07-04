package main

import (
	"log"
	"net/http"
)

func main() {
	client := http.Client{}

	req, err := http.NewRequest(http.MethodGet, "https://webhook.site/4933b8d2-8c3e-4b15-ba23-b61fb8bad85c", nil)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
}
