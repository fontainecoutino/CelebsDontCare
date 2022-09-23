package trip

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/fontainecoutino/CelebsDontCare/database"
)

func getProduct(productID int) (*Trip, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	row := database.DB.QueryRowContext(ctx,
		`SELECT * FROM trips WHERE id = $1`, productID)

	trip := &Trip{}
	err := row.Scan(
		&trip.TripID,
		&trip.TimeStamp,
		&trip.UserID,
		&trip.Distance,
		&trip.GallonsUsed,
		&trip.CostOfFuel,
		&trip.StartDest,
		&trip.EndDest,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		log.Println(err)
		return nil, err
	}
	return trip, nil
}

func insertTrip(trip Trip) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_, err := database.DB.ExecContext(ctx,
		`INSERT INTO trips
		(time_stamp, user_id, distance, gallons_used, 
			cost_of_fuel, start_dest, end_dest) 
		VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		trip.TimeStamp,
		trip.UserID,
		trip.Distance,
		trip.GallonsUsed,
		trip.CostOfFuel,
		trip.StartDest,
		trip.EndDest)

	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

/*
func removeProduct(productID int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_, err := database.DbConn.ExecContext(ctx, `DELETE FROM products where productId = ?`, productID)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func getProductList() ([]Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	results, err := database.DbConn.QueryContext(ctx, `SELECT
	productId,
	manufacturer,
	sku,
	upc,
	pricePerUnit,
	quantityOnHand,
	productName
	FROM products`)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer results.Close()
	products := make([]Product, 0)
	for results.Next() {
		var product Product
		results.Scan(&product.ProductID,
			&product.Manufacturer,
			&product.Sku,
			&product.Upc,
			&product.PricePerUnit,
			&product.QuantityOnHand,
			&product.ProductName)

		products = append(products, product)
	}
	return products, nil
}

func updateProduct(product Product) error {
	// if the product id is set, update, otherwise add
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if product.ProductID == nil || *product.ProductID == 0 {
		return errors.New("product has invalid ID")
	}
	_, err := database.DbConn.ExecContext(ctx, `UPDATE products SET
		manufacturer=?,
		sku=?,
		upc=?,
		pricePerUnit=CAST(? AS DECIMAL(13,2)),
		quantityOnHand=?,
		productName=?
		WHERE productId=?`,
		product.Manufacturer,
		product.Sku,
		product.Upc,
		product.PricePerUnit,
		product.QuantityOnHand,
		product.ProductName,
		product.ProductID)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}
*/
