package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// HandleRoot ...
func HandleRoot(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello world")
}

// HandleHome ...
func HandleHome(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "This is the API EndPoint")
}

// postRequest ...
func postRequest(writer http.ResponseWriter, request *http.Request) {
	// fmt.Fprintf(writer, "This is the API EndPoint")
	decoder := json.NewDecoder(request.Body)
	var metadata MetaData
	error := decoder.Decode(&metadata)
	if error != nil {
		fmt.Fprintf(writer, "error: %v", error)
		return
	}
	fmt.Fprintf(writer, "Payload %v \n", metadata)
}

// userPostRequest ...
func userPostRequest(writer http.ResponseWriter, request *http.Request) {
	// fmt.Fprintf(writer, "This is the API EndPoint")
	decoder := json.NewDecoder(request.Body)
	var user User
	error := decoder.Decode(&user)
	if error != nil {
		fmt.Fprintf(writer, "error: %v", error)
		return
	}

	response, error := user.toJSON()
	if error != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(response)
}
