package main

import hf "github.com/bots-garden/capsule/capsulemodule/hostfunctions"

// main is required.
func main() {
	hf.SetHandleHttp(Handle)
}

// Handle open: http://localhost:8080
func Handle(bodyReq string, headersReq map[string]string) (bodyResp string, headersResp map[string]string, errResp error) {

	html, err := hf.ReadFile("index.html")
	if err != nil {
		hf.Log(err.Error())
	}

	headersResp = map[string]string{
		"Content-Type": "text/html; charset=utf-8",
	}

	return html, headersResp, err
}
