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

	hf.Log("First parameter: " + params[0])

	headers := map[string]string{"Accept": "application/json", "Content-Type": "text/html; charset=UTF-8"}

	ret, err := hf.Http("https://httpbin.org/post", "POST", headers, params[0])
	if err != nil {
		hf.Log("😡 error:" + err.Error())
	} else {
		hf.Log("Response: " + ret)
	}

	return ret, err
}
