package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type DistanceZipcode struct {
	ZipcodeDetails
	Miles      float64 `json:"miles"`
	Kilometers float64 `json:"kilometers"`
}

type DistanceResponse struct {
	ZipcodeDetails
	Distance DistanceZipcode `json:"distance"`
}

func main() {
	db := getInitializedDatabase()
	http.HandleFunc("/zip/", func(w http.ResponseWriter, r *http.Request) {
		handleZipcodeRequest(w, r, db)
	})
	http.ListenAndServe(getListeningAddress(), nil)
}

func getInitializedDatabase() *ZipcodeDatabase {
	db := NewZipcodeDatabase()
	err := db.LoadFromCSV("zips.csv")
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	return db
}

func getListeningAddress() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return strings.Join([]string{
		":",
		port,
	}, "")
}

func handleZipcodeRequest(response http.ResponseWriter, request *http.Request, db *ZipcodeDatabase) {
	zip := zipcodeForPath(request, db)
	if zip == nil {
		http.Error(response, "", 404)
	} else {
		setResponseHeaders(response)
		if zip2 := zipcodeForQuery(request, db); zip2 != nil {
			sendDistanceResponse(zip, zip2, response)
		} else {
			sendJson(response, zip)
		}
	}
}

func sendDistanceResponse(zip1, zip2 *ZipcodeDetails, writer http.ResponseWriter) {
	mi, km := zip1.DistanceTo(zip2)
	response := &DistanceResponse{}
	response.ZipcodeDetails = *zip1
	response.Distance.ZipcodeDetails = *zip2
	response.Distance.Miles = mi
	response.Distance.Kilometers = km
	sendJson(writer, response)
}

func sendJson(writer http.ResponseWriter, data interface{}) {
	bytes, _ := json.MarshalIndent(data, "", "  ")
	fmt.Fprintf(writer, string(bytes))
}

func setResponseHeaders(response http.ResponseWriter) {
	response.Header().Set("Access-Control-Allow-Origin", "*")
	response.Header().Set("Content-Type", "application/json")
}

func extractZipFromPath(request *http.Request) string {
	return request.URL.Path[len("/zip/"):]
}

func zipcodeForPath(request *http.Request, db *ZipcodeDatabase) *ZipcodeDetails {
	return db.Find(extractZipFromPath(request))
}

func extractZipFromQuery(request *http.Request) string {
	if param := request.URL.Query()["distance"]; param != nil {
		return param[0]
	} else {
		return ""
	}
}

func zipcodeForQuery(request *http.Request, db *ZipcodeDatabase) *ZipcodeDetails {
	return db.Find(extractZipFromQuery(request))
}
