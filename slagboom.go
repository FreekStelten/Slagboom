package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	//data "git.fhict.nl/I882775/gatekeeper/data"
	datetime "git.fhict.nl/I882775/gatekeeper/datetime"
	logger "git.fhict.nl/I882775/gatekeeper/logging"
	database "git.fhict.nl/I882775/gatekeeper/database"
)

var licensePlate string
var countryCode string

func init() {
	flag.StringVar(&licensePlate, "licensePlate", "", "specify the license plate number")
	flag.Parse()
	if licensePlate == "" {
		fmt.Println("Please specify a license plate number")
		flag.PrintDefaults()
		os.Exit(1)
	}
	var found bool = false
	countryCode, found = os.LookupEnv("fonteyn_parc_country")
	if !found {
		fmt.Println("Couldn't find the environment variable 'fonteyn_parc_country'")
		os.Exit(2)
	}
	err := database.Connect()
	if err != nil {
		logger.LogFatal(fmt.Errorf("failed to connect to db %v", err).Error(), 3)
	}
	logger.LogInfo("Finished initializing app")
}

func main() {
	dayPart := datetime.GetDayPart(time.Now())
	if dayPart == datetime.Night {
		fmt.Println("Sorry, the parkinglot is closed at night.")
		return
	}
	result, err := database.CheckBooking(licensePlate, time.Now(), countryCode)
	if err != nil {
		logger.LogFatal(err.Error(), 4)
	}
	if result.Result {
		fmt.Println("Welcome to Fonteyn Holidayparcs!")
	} else {
		logger.LogWarning(result.Message)
		fmt.Println("Sorry, you cannot access this parking lot.")
	}
}
