package trip

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/fontainecoutino/CelebsDontCare/database"
)

/**
 * Gets the specific trip from the database based on the id. Puts data into
 * struct and returns with no error. If an error occurs then it returns an
 * empty obj and the error.
 */
func getTrip(tripID int) (*Trip, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	row := database.DB.QueryRowContext(ctx,
		`SELECT * FROM trips WHERE id = $1`, tripID)

	trip := &Trip{}
	err := row.Scan(
		&trip.TripID,
		&trip.TimeStamp,
		&trip.Name,
		&trip.Distance,
		&trip.GallonsUsed,
		&trip.CostOfFuel,
		&trip.Flight,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		log.Println(err)
		return nil, err
	}
	return trip, nil
}

/**
 * Gets all the trips from the database. Puts data into a slice of
 * structs and returns with no error. If an error occurs then it returns an
 * empty obj and the error.
 */
func getTripList() ([]Trip, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	results, err := database.DB.QueryContext(ctx, `SELECT * FROM trips `)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer results.Close()
	trips := make([]Trip, 0)
	for results.Next() {
		var trip Trip
		results.Scan(
			&trip.TripID,
			&trip.TimeStamp,
			&trip.Name,
			&trip.Distance,
			&trip.GallonsUsed,
			&trip.CostOfFuel,
			&trip.Flight)

		trips = append(trips, trip)
	}
	return trips, nil
}

/**
 * Inserts a given trip into the database. If an error occurs then it returns
 * the error.
 */
func insertTrip(trip Trip) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_, err := database.DB.ExecContext(ctx,
		`INSERT INTO trips
		(time_stamp, name, distance, gallons_used, cost_of_fuel, flight) 
		VALUES 
		($1, $2, $3, $4, $5, $6)`,
		trip.TimeStamp,
		trip.Name,
		trip.Distance,
		trip.GallonsUsed,
		trip.CostOfFuel,
		trip.Flight)

	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

/**
 * Removes specific trip from database based on id. Returns error if one
 * occurred.
 */
func removeTrip(tripID int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_, err := database.DB.ExecContext(ctx,
		`DELETE FROM trips where id = $1`, tripID)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

/**
 * Gets the data to be displayed from the database. Puts data into
 * struct and returns with no error. If an error occurs then it returns an
 * empty obj and the error.
 */
func getDataDisplay() ([]DataDisplay, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	results, err := database.DB.QueryContext(ctx, `
	SELECT 
			name, 
			FLOOR(SUM(distance)*1.609) distance_traveled_km,
			ROUND((SUM(gallons_used)*3.125*3.16)/1000, 1) co2_produced_tons,
			ROUND(((SUM(gallons_used)*3.125*3.16)/.00171)/1000000, 1) million_plastic_straws_equivalent
		FROM trips
		GROUP BY name
		ORDER BY sum(gallons_used) desc;`)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer results.Close()

	tableData := make([]DataDisplay, 0)
	for results.Next() {
		var data DataDisplay
		results.Scan(
			&data.Name,
			&data.DistanceTraveled,
			&data.CO2Produced,
			&data.PlasticStrawsUsed)

		tableData = append(tableData, data)
	}
	return tableData, nil
}
