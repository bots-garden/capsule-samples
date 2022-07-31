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
    // add a key, value
    res1, err := hf.RedisSet("greetings", "Hello World")
    if err != nil {
        hf.Log(err.Error())
    } else {
        hf.Log("Result: " + res1)
    }

    // read the value
    res2, err := hf.RedisGet("greetings")
    if err != nil {
        hf.Log(err.Error())
    } else {
        hf.Log("Value: " + res2)
    }

	return "OK", err
}
