package database

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func SetupDatabase() {
	host := os.Getenv("dbhost")
	port, _ := strconv.Atoi(os.Getenv("dbport"))
	user := os.Getenv("dbuser")
	password := os.Getenv("dbpassword")
	dbname := os.Getenv("dbname")

	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	var err error
	DB, err = sql.Open("postgres", conn)
	if err != nil {
		fmt.Println("> failed connecting to db")
		panic(err)
	}

	DB.SetMaxOpenConns(3)
	DB.SetMaxIdleConns(3)
	DB.SetConnMaxLifetime(60 * time.Second)
}

/*
	!TODO: This can be used to split flight trips
 	flight := "A-Rod's Jet 1,090 mile (947 NM) flight from OPF to TEB\n \n~ 1,227 gallons (4,646 liters). \n~ 8,225 lbs (3,731 kg) of jet fuel used. \n~ $8,026 cost of fuel. \n~ 13 tons of CO2 emissions."
	flight := "Travis Scott's (Cactus Jack LLC)  Jet 229 mile (199 NM) flight from LAS to VNY\n \n~ 356 gallons (1,347 liters). \n~ 2,384 lbs (1,082 kg) of jet fuel used. \n~ $2,327 cost of fuel. \n~ 4 tons of CO2 emissions."
	flightSplit := strings.Split(flight, "\n")

	celebrity := strings.Split(flightSplit[0], "'s")[0]

	distanceSlice := strings.Split(strings.Split(flightSplit[0], "mile")[0], " ")
	distance := distanceSlice[len(distanceSlice)-2]

	placesSplit := strings.Split(flightSplit[0], "from")
	origin := strings.Split(placesSplit[len(placesSplit)-1], " ")[1]
	destination := strings.Split(placesSplit[len(placesSplit)-1], " ")[3]

	jetFuelSlice := strings.Split(flightSplit[2], "~")[1]
	jetFuel := jetFuelSlice[1 : len(jetFuelSlice)-2]

	costFuelTripSlice := strings.Split(flightSplit[4], "~")[1]
	costFuelTrip := costFuelTripSlice[1 : len(costFuelTripSlice)-2]

	fmt.Println("Celebrity     : " + celebrity)
	fmt.Println("Distance      : " + distance)
	fmt.Println("Origin        : " + origin)
	fmt.Println("Destination   : " + destination)
	fmt.Println("Jet fuel used : " + jetFuel)
	fmt.Println("Cost fuel used: " + costFuelTrip)
	fmt.Println(flightSplit)

*/
