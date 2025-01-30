package requests

import (
	"net/http"
)

func GetRequest(url string, result chan<- *http.Response) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	result <- resp
}
