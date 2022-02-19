package main

import (
	"encoding/json"
	"fmt"
	"github.com/golang/snappy"
	"io/ioutil"
	"net/http"
)

func httpHandleRoot() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		compressed, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Println(string(compressed))

		// JSON code taken from https://pkg.go.dev/encoding/json#Unmarshal
		var jsonBlob = []byte(compressed)
		type Animal struct {
			Name  string
			Order string
		}
		var animals []Animal
		err2 := json.Unmarshal(jsonBlob, &animals)
		if err2 != nil {
			fmt.Println("error:", err)
		}
		fmt.Printf("%+v\n", animals)
		// End of json code

		fmt.Println(animals[0].Name)
		fmt.Println(animals[0].Order)

		w.Header().Set("Content-Type", "text/html")
		w.Header().Set("Content-Encoding", "UTF-8")

		response := "The animal name is " + animals[0].Name + " and it's order is " + animals[0].Order + "!\n"

		compressed = snappy.Encode(nil, []byte(response))
		if _, err := w.Write(compressed); err != nil {
			fmt.Println("Error writing response", "err", err)
		}

	}
}
