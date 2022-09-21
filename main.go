package main

import (
	"github.com/fontainecoutino/CelebsDontCare/database"
	"github.com/fontainecoutino/CelebsDontCare/database/retrieve"
)

const retrieveNewData = false

func main() {
	// get the data from twitter in order to store in DB
	if retrieveNewData {
		retrieve.GetData()
	}
	database.Main()
}
