package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"submission-3/helpers"
	"submission-3/structs"

	"github.com/gin-gonic/gin"
)

func GetMain(ctx *gin.Context) {
	dataJson := getJSON()
	valueWater, statusWater := getStatusByValue("water", dataJson.Status.Water)
	valueWind, statusWind := getStatusByValue("water", dataJson.Status.Wind)

	helpers.IntervalFunction(updateJSON, 15)

	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"statusWater": statusWater,
		"valueWater":  valueWater,
		"statusWind":  statusWind,
		"valueWind":   valueWind,
	})
}

func updateJSON() {
	var dataWeather structs.Data

	dataWeather.Status.Water = rand.Intn(100-1+1) + 1
	dataWeather.Status.Wind = rand.Intn(100-1+1) + 1

	resultMarshal, errMarshaling := json.Marshal(dataWeather)
	if errMarshaling != nil {
		fmt.Println(errMarshaling.Error())
	}

	os.WriteFile("./data/data.json", []byte(resultMarshal), 0644)

}

func getJSON() structs.Data {
	jsonFile, err := os.Open("./data/data.json")
	jsonByte, _ := ioutil.ReadAll(jsonFile)

	if err != nil {
		panic(err)
	}

	var dataWeather structs.Data

	errMarshal := json.Unmarshal(jsonByte, &dataWeather)

	if errMarshal != nil {
		fmt.Println(errMarshal.Error())
	}

	return dataWeather
}

func getStatusByValue(category string, value int) (int, string) {
	var status string = ""
	if category == "water" {
		if value < 5 {
			status = "Aman"
		} else if value >= 6 && value <= 8 {
			status = "Siaga"
		} else if value > 8 {
			status = "Bahaya"
		}
	}

	if category == "wind" {
		if value < 6 {
			status = "Aman"
		} else if value >= 7 && value <= 15 {
			status = "Siaga"
		} else if value > 15 {
			status = "Bahaya"
		}
	}

	return value, status
}
