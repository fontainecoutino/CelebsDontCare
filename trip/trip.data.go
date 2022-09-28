package trip

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/fontainecoutino/CelebsDontCare/database"
)

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

func insertTrip(trip Trip) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_, err := database.DB.ExecContext(ctx,
		`INSERT INTO trips
		(time_stamp, name, distance, gallons_used, cost_of_fuel, flight) 
		VALUES ($1, $2, $3, $4, $5, $6)`,
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

func removeTrip(productID int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_, err := database.DB.ExecContext(ctx,
		`DELETE FROM trips where id = $1`, productID)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}
