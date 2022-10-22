package function

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var inputSource = os.Getenv("INPUT_SOURCE_URL")
// Handle an HTTP Request.
func Handle(ctx context.Context, res http.ResponseWriter, req *http.Request) {
	if inputSource == "" {
		inputSource = "http://input-source.default.svc.cluster.local"
	}
	resp, err := http.Get(inputSource)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Fprintf(res, string(body))
}

