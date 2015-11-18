package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	db := getInitializedDatabase()

	http.HandleFunc("/zip/", func(w http.ResponseWriter, r *http.Request) {
		handleZipcodeRequest(w, r, db)
	})

	http.ListenAndServe(":8080", nil)
}

func getInitializedDatabase() *ZipcodeDatabase {
	db := NewZipcodeDatabase()
	err := db.LoadFromCSV("./zips.csv")
	if err != nil {
		fmt.Printf("%s\n", err)
	}

	return db
}

func handleZipcodeRequest(response http.ResponseWriter, request *http.Request, db *ZipcodeDatabase) {
	zip := zipcodeForRequest(request)
	if details := db.Find(zip); details == nil {
		http.Error(response, "", 404)
	} else {
		setResponseHeaders(response)
		sendZipcodeDetails(details, response)
	}
}

func setResponseHeaders(response http.ResponseWriter) {
	response.Header().Set("Access-Control-Allow-Origin", "*")
	response.Header().Set("Content-Type", "application/json")
}

func zipcodeForRequest(request *http.Request) string {
	return request.URL.Path[len("/zip/"):]
}

func sendZipcodeDetails(details *ZipcodeDetails, response http.ResponseWriter) {
	data, _ := json.MarshalIndent(details, "", "  ")
	fmt.Fprintf(response, string(data))
}
