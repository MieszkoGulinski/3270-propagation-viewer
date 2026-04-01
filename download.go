package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
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
	// Check if cache.xml exists
	_, err := os.Stat("cache.xml")
	if os.IsNotExist(err) {
		// Download and save to cache.xml
		conditions, err := downloadConditionsFromAPI()
		if err != nil {
			panic("Error downloading conditions from API")
		}
		return conditions
	}

	cacheFileContent, err := os.ReadFile("cache.xml")
	if err != nil {
		panic("Error reading cache.xml")
	}

	var conditionsFromCache PropagationConditions
	err = xml.Unmarshal(cacheFileContent, &conditionsFromCache)
	if err != nil {
		panic("Error unmarshalling cache.xml")
	}

	// Check if cache.xml is older than 3 hours
	cachedDataTimestamp := parseTime(conditionsFromCache.SolarData.Updated)
	if time.Since(cachedDataTimestamp) > 3*time.Hour {
		conditions, err := downloadConditionsFromAPI()
		if err != nil {
			fmt.Println("Error downloading conditions from API, using cached data")
			return conditionsFromCache
		}
		return conditions
	}
	return conditionsFromCache
}

func downloadConditionsFromAPI() (PropagationConditions, error) {
	fmt.Println("Downloading conditions from API")
	
	resp, err := http.Get("https://www.hamqsl.com/solarxml.php")
	if err != nil {
		return PropagationConditions{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return PropagationConditions{}, fmt.Errorf("Response error status: %s", resp.Status)
	}

	responseContent, err := io.ReadAll(resp.Body)
	if err != nil {
		return PropagationConditions{}, err
	}

	var conditions PropagationConditions
	if err := xml.Unmarshal(responseContent, &conditions); err != nil {
		return PropagationConditions{}, err
	}

	// Save response to cache.xml
	if err := os.WriteFile("cache.xml", responseContent, 0644); err != nil {
		return PropagationConditions{}, err
	}

	return conditions, nil
}

func parseTime(timeStr string) time.Time {
	timeStr = strings.TrimSpace(timeStr)
	parsedTime, err := time.Parse("02 Jan 2006 1504 MST", timeStr)
	if err != nil {
		panic("Error parsing time: " + err.Error())
	}
	return parsedTime
}