package main

import (
	"encoding/csv"
	"math"
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

const EarthRadiusMiles float64 = 3960
const EarthRadiusKm float64 = 6373
const DegToRad float64 = math.Pi / 180

func (zip1 *ZipcodeDetails) DistanceTo(zip2 *ZipcodeDetails) (float64, float64) {
	// the arc distance ψ = arccos(sin φ_1 sin φ_2 cos(θ_1-θ_2) + cos φ_1 cos φ_2)
	// where
	//		φ_k = 90° - latitude_k
	//		θ_k = longitude_k
	// (both in radians)
	theta1 := zip1.Longitude * DegToRad
	theta2 := zip2.Longitude * DegToRad
	phi1 := (90 - zip1.Latitude) * DegToRad
	phi2 := (90 - zip2.Latitude) * DegToRad
	arc := math.Acos(math.Sin(phi1)*math.Sin(phi2)*math.Cos(theta1-theta2) + math.Cos(phi1)*math.Cos(phi2))
	return arc * EarthRadiusMiles, arc * EarthRadiusKm
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
