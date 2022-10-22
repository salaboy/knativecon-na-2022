package function

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

)


// Handle an HTTP Request.
func Handle(ctx context.Context, res http.ResponseWriter, req *http.Request) {

	resp, err := http.Get("http://input-source.default.svc.cluster.local")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Fprintf(res, string(body))
}

