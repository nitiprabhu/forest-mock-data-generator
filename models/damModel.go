package models

import (
	"encoding/json"
	"math/rand"
	"time"
)

type DamModel struct {
	Temperature     int  `json:"temperature"`
	AirPressure     int  `json:"air_pressure"`
	Humidity        int  `json:"humidity"`
	Raining         bool `json:"raining"`
	WindSpeedNumber int  `json:"wind_speed_number"`
	CrackMeter      int  `json:"crack_meter"`
	WaterLevel      int  `json:"water_level"`
	RiverWaterLevel int  `json:"river_water_level"`
}

func GenerateRandomData() string {

	damModel := &DamModel{
		Temperature:     rand.Intn(70-1) + 1,
		AirPressure:     rand.Intn(70-1) + 1,
		Humidity:        rand.Intn(60-5) + 1,
		Raining:         RandBool(),
		WindSpeedNumber: rand.Intn(10-0) + 1,
		CrackMeter:      rand.Intn(100-0) + 1,
		WaterLevel:      rand.Intn(100-0) + 1,
		RiverWaterLevel: rand.Intn(100-0) + 1,
	}

	bytes, e := json.Marshal(damModel)

	if e != nil {
		return ""
	}

	return string(bytes)
}

func RandBool() bool {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(2) == 1
}
