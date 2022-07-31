package main

import hf "github.com/bots-garden/capsule/capsulemodule/hostfunctions"

// main is required.
func main() {
	hf.SetHandleHttp(Handle)
}

func Handle(bodyReq string, headersReq map[string]string) (bodyResp string, headersResp map[string]string, errResp error) {
	hf.Log("📝 body: " + bodyReq)

	// Read the request headers
	hf.Log("Content-Type: " + headersReq["Content-Type"])
	hf.Log("Content-Length: " + headersReq["Content-Length"])
	hf.Log("User-Agent: " + headersReq["User-Agent"])

	// Read the MESSAGE environment variable
	envMessage, err := hf.GetEnv("MESSAGE")
	if err != nil {
		hf.Log("😡 " + err.Error())
	} else {
		hf.Log("Environment variable: " + envMessage)
	}

	// Set the response content type and add a message header
	headersResp = map[string]string{
		"Content-Type": "application/json; charset=utf-8",
		"Message":      "👋 hello world 🌍",
	}

	jsonResponse := `{"message": "hey people!"}`

	return jsonResponse, headersResp, err
}
