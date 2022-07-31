// wasm module
package main

import (
	hf "github.com/bots-garden/capsule/capsulemodule/hostfunctions"
)

// main is required.
func main() {
	hf.SetHandle(Handle)
}

func Handle(params []string) (string, error) {

	// Set the headers
	headers := map[string]string{"Accept": "application/json", "Content-Type": "text/html; charset=UTF-8"}

	response, err := hf.Http("https://httpbin.org/get", "GET", headers, "")
	if err != nil {
		hf.Log("😡 error:" + err.Error())
	} else {
		hf.Log("Response: " + response)
	}

	return response, err
}
