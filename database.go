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
		db.Insert(row[0], getDetailsFromRow(row))
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

func getFloat(stringVal string) float64 {
	floatVal, err := strconv.ParseFloat(strings.TrimSpace(stringVal), 64)
	if err != nil {
		return 0
	} else {
		return floatVal
	}
}

func getDetailsFromRow(row []string) *ZipcodeDetails {
	return &ZipcodeDetails{
		State:     row[1],
		Latitude:  getFloat(row[2]),
		Longitude: getFloat(row[3]),
		City:      row[4],
	}
}
