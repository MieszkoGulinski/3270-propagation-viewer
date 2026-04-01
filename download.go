package main

import (
	"encoding/xml"
	"os"
	"strings"
	"time"
)

type PropagationConditions struct {
	SolarData SolarData `xml:"solardata"`
}

type SolarData struct {
	Updated         string `xml:"updated"`
	SolarFlux       string `xml:"solarflux"`
	AIndex          string `xml:"aindex"`
	KIndex          string `xml:"kindex"`
	KIndexNT        string `xml:"kindexnt"`
	XRay            string `xml:"xray"`
	Sunspots        string `xml:"sunspots"`
	HeliumLine      string `xml:"heliumline"`
	ProtonFlux      string `xml:"protonflux"`
	ElectronFlux    string `xml:"electonflux"` // XML has a typo
	Aurora          string `xml:"aurora"`
	Normalization   string `xml:"normalization"`
	AuroraLat       string `xml:"latdegree"`
	SolarWind       string `xml:"solarwind"`
	MagneticField   string `xml:"magneticfield"`
	GeomagField     string `xml:"geomagfield"`
	SignalNoise     string `xml:"signalnoise"`
	Fof2            string `xml:"fof2"`
	MufFactor       string `xml:"muffactor"`
	Muf             string `xml:"muf"`
	Calculated      CalculatedConditions `xml:"calculatedconditions"`
	CalculatedVHF   CalculatedVHFConditions `xml:"calculatedvhfconditions"`
}

type CalculatedConditions struct {
	Bands []Band `xml:"band"`
}

type CalculatedVHFConditions struct {
	Phenomenon []Phenomenon `xml:"phenomenon"`
}

type Band struct {
	Name  string `xml:"name,attr"`
	Time  string `xml:"time,attr"`
	Value string `xml:",chardata"`
}

type Phenomenon struct {
	Name  string `xml:"name,attr"`
	Location string `xml:"location,attr"`
	Value string `xml:",chardata"`
}

func getConditions() PropagationConditions {
	// Temporary code, read sample.xml file
	sampleFileContent, err := os.ReadFile("sample.xml")
	if err != nil {
		panic("Error reading sample.xml")
	}
	// Parse sampleFileContent
	var conditions PropagationConditions
	err = xml.Unmarshal(sampleFileContent, &conditions)
	if err != nil {
		panic("Error unmarshalling sample.xml")
	}

	return conditions
}

// func getConditions() PropagationConditions {
// 	// Check if cache.xml exists
// 	_, err := os.Stat("cache.xml")
// 	if os.IsNotExist(err) {
// 		// Download and save to cache.xml
// 	}

// 	// Check if cache.xml is older than 3 hours

// 	return PropagationConditions{}
// }

func downloadConditionsFromAPI() {
	// Download XML, save to cache.xml and parse it
}

func parseTime(timeStr string) time.Time {
	timeStr = strings.TrimSpace(timeStr)
	parsedTime, err := time.Parse("02 Jan 2006 1504 MST", timeStr)
	if err != nil {
		panic("Error parsing time: " + err.Error())
	}
	return parsedTime
}