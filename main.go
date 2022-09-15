package main

import (
	"github.com/fontainecoutino/CelebsDontCare/database/retrieve"
)

func main() {
	// get the data from twitter in order to store in DB
	retrieve.GetData()
}
