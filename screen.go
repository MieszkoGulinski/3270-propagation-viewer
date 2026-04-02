package main

import (
	"github.com/racingmars/go3270"
)

func getColorForHFConditions(value string) go3270.Color {
	switch value {
	case "Poor":
		return go3270.Red
	case "Fair":
		return go3270.Yellow
	case "Good":
		return go3270.Green
	default:
		return go3270.White
	}
}

func getColorForVHFConditions(value string) go3270.Color {
	switch value {
	case "Band Closed":
		return go3270.Red
	default:
		return go3270.Green
	}
}

// Convert PropagationConditions to TN3270 screen
func toScreen(conditions PropagationConditions) go3270.Screen {
	scr := go3270.Screen{
		{Row: 0, Col: 0, Content: "Propagation Data - Source: N0NBH http://www.hamqsl.com/solar.html"},

		// HF Calculated values
		{Row: 2, Col: 0, Content: "Band", Color: go3270.Pink},
		{Row: 2, Col: 10, Content: "Day", Color: go3270.Pink},
		{Row: 2, Col: 16, Content: "Night", Color: go3270.Pink},

		{Row: 3, Col: 0, Content: "80m-40m", Color: go3270.White},
		{Row: 4, Col: 0, Content: "30m-20m", Color: go3270.White},
		{Row: 5, Col: 0, Content: "17m-15m", Color: go3270.White},
		{Row: 6, Col: 0, Content: "12m-10m", Color: go3270.White},

		{Row: 3, Col: 10, Content: conditions.SolarData.Calculated.Bands[0].Value, Color: getColorForHFConditions(conditions.SolarData.Calculated.Bands[0].Value)},
		{Row: 4, Col: 10, Content: conditions.SolarData.Calculated.Bands[1].Value, Color: getColorForHFConditions(conditions.SolarData.Calculated.Bands[1].Value)},
		{Row: 5, Col: 10, Content: conditions.SolarData.Calculated.Bands[2].Value, Color: getColorForHFConditions(conditions.SolarData.Calculated.Bands[2].Value)},
		{Row: 6, Col: 10, Content: conditions.SolarData.Calculated.Bands[3].Value, Color: getColorForHFConditions(conditions.SolarData.Calculated.Bands[3].Value)},

		{Row: 3, Col: 16, Content: conditions.SolarData.Calculated.Bands[4].Value, Color: getColorForHFConditions(conditions.SolarData.Calculated.Bands[4].Value)},
		{Row: 4, Col: 16, Content: conditions.SolarData.Calculated.Bands[5].Value, Color: getColorForHFConditions(conditions.SolarData.Calculated.Bands[5].Value)},
		{Row: 5, Col: 16, Content: conditions.SolarData.Calculated.Bands[6].Value, Color: getColorForHFConditions(conditions.SolarData.Calculated.Bands[6].Value)},
		{Row: 6, Col: 16, Content: conditions.SolarData.Calculated.Bands[7].Value, Color: getColorForHFConditions(conditions.SolarData.Calculated.Bands[7].Value)},

		// VHF calculated values
		{Row: 8, Col: 0, Content: "Aurora", Color: go3270.White},
		{Row: 8, Col: 10, Content: conditions.SolarData.CalculatedVHF.Phenomenon[0].Value, Color: getColorForVHFConditions(conditions.SolarData.CalculatedVHF.Phenomenon[0].Value)},
		{Row: 9, Col: 0, Content: "Aur lat", Color: go3270.White},
		{Row: 9, Col: 10, Content: conditions.SolarData.AuroraLat},
		{Row: 10, Col: 0, Content: "ES EU 2m", Color: go3270.White},
		{Row: 10, Col: 10, Content: conditions.SolarData.CalculatedVHF.Phenomenon[1].Value, Color: getColorForVHFConditions(conditions.SolarData.CalculatedVHF.Phenomenon[1].Value)},
		{Row: 11, Col: 0, Content: "ES NA 2m", Color: go3270.White},
		{Row: 11, Col: 10, Content: conditions.SolarData.CalculatedVHF.Phenomenon[2].Value, Color: getColorForVHFConditions(conditions.SolarData.CalculatedVHF.Phenomenon[2].Value)},
		{Row: 12, Col: 0, Content: "ES EU 6m", Color: go3270.White},
		{Row: 12, Col: 10, Content: conditions.SolarData.CalculatedVHF.Phenomenon[3].Value, Color: getColorForVHFConditions(conditions.SolarData.CalculatedVHF.Phenomenon[3].Value)},
		{Row: 13, Col: 0, Content: "ES EU 4m", Color: go3270.White},
		{Row: 13, Col: 10, Content: conditions.SolarData.CalculatedVHF.Phenomenon[4].Value, Color: getColorForVHFConditions(conditions.SolarData.CalculatedVHF.Phenomenon[4].Value)},

		// Remaining conditions
		{Row: 15, Col: 0, Content: "Geomagn. field", Color: go3270.White},
		{Row: 15, Col: 16, Content: conditions.SolarData.GeomagField}, // QUIET / UNSETTLD / ACTIVE
		{Row: 16, Col: 0, Content: "Signal/noise", Color: go3270.White},
		{Row: 16, Col: 16, Content: conditions.SolarData.SignalNoise}, // possible values e.g. S1-S2, S2-S3, S3-S4 etc.
		{Row: 17, Col: 0, Content: "MUF US Boulder", Color: go3270.White},
		{Row: 17, Col: 16, Content: conditions.SolarData.Muf},

		// Measurements of solar conditions
		{Row: 2, Col: 50, Content: "Solar flux", Color: go3270.White},
		{Row: 2, Col: 64, Content: conditions.SolarData.SolarFlux},

		{Row: 3, Col: 50, Content: "A index", Color: go3270.White},
		{Row: 3, Col: 64, Content: conditions.SolarData.AIndex},

		{Row: 4, Col: 50, Content: "K index", Color: go3270.White},
		{Row: 4, Col: 64, Content: conditions.SolarData.KIndex},

		{Row: 5, Col: 50, Content: "X-ray", Color: go3270.White},
		{Row: 5, Col: 64, Content: conditions.SolarData.XRay},

		{Row: 6, Col: 50, Content: "Helium line", Color: go3270.White},
		{Row: 6, Col: 64, Content: conditions.SolarData.HeliumLine},

		{Row: 7, Col: 50, Content: "Proton flux", Color: go3270.White},
		{Row: 7, Col: 64, Content: conditions.SolarData.ProtonFlux},

		{Row: 8, Col: 50, Content: "Electron flux", Color: go3270.White},
		{Row: 8, Col: 64, Content: conditions.SolarData.ElectronFlux},

		{Row: 9, Col: 50, Content: "Aurora", Color: go3270.White},
		{Row: 9, Col: 64, Content: conditions.SolarData.Aurora + " / " + conditions.SolarData.Normalization},

		{Row: 10, Col: 50, Content: "Magn. field", Color: go3270.White},
		{Row: 10, Col: 64, Content: conditions.SolarData.MagneticField},

		{Row: 11, Col: 50, Content: "Solar wind", Color: go3270.White},
		{Row: 11, Col: 64, Content: conditions.SolarData.SolarWind},
		
		// Time
		{Row: 19, Col: 0, Content: "Last updated:", Color: go3270.White},
		{Row: 19, Col: 16, Content: parseTime(conditions.SolarData.Updated).Format("2006-01-02 15:04")},
	}

	return scr
}