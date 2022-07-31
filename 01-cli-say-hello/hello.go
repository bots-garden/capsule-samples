package main

import hf "github.com/bots-garden/capsule/capsulemodule/hostfunctions"

// main is required.
func main() {
	hf.SetHandle(Handle)
}

func Handle(params []string) (string, error) {
	var err error
	for _, param := range params {
		hf.Log("- parameter is: " + param)
	}

	ret := "The first parameter is: " + params[0]

	return ret, err // err = nil
}
