package main

import (
	"encoding/csv"
	"os"
	"strconv"
	"strings"
)

type ZipcodeDetails struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	City      string  `json:"city"`
	State     string  `json:"state"`
}

type ZipcodeDatabase struct {
	Zipcodes map[string]*ZipcodeDetails
}

func NewZipcodeDatabase() *ZipcodeDatabase {
	db := &ZipcodeDatabase{}
	db.Zipcodes = make(map[string]*ZipcodeDetails)
	return db
}

func (db *ZipcodeDatabase) Insert(zipcode string, details *ZipcodeDetails) {
	db.Zipcodes[zipcode] = details
}

func (db *ZipcodeDatabase) Find(zipcode string) *ZipcodeDetails {
	return db.Zipcodes[zipcode]
}

func (db *ZipcodeDatabase) LoadFromCSV(filename string) error {
	zips, err := getZipsFromFile(filename)
	if err != nil {
		return err
	}

	for _, row := range zips {
		details, err := getDetailsFromRow(row)
		if err != nil {
			return err
		}
		db.Insert(row[0], details)
	}

	return nil
}

func getZipsFromFile(filename string) ([][]string, error) {
	reader, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	csvReader := csv.NewReader(reader)
	zips, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	return zips, nil
}

func getDetailsFromRow(row []string) (*ZipcodeDetails, error) {
	lat, err := strconv.ParseFloat(strings.TrimSpace(row[2]), 64)
	if err != nil {
		return nil, err
	}
	lon, err := strconv.ParseFloat(strings.TrimSpace(row[3]), 64)
	if err != nil {
		return nil, err
	}

	zip := &ZipcodeDetails{
		State:     row[1],
		Latitude:  lat,
		Longitude: lon,
		City:      row[4],
	}

	return zip, nil
}
