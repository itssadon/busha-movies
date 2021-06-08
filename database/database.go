package database

import (
	"log"

	"github.com/go-pg/pg/v10"
)

func NewDBConn() (con *pg.DB) {
	options := &pg.Options{
		User:     "hrfqxsaoxawmwk",
		Password: "eea5979f55da9ed7b33f07e47fdd5f73bd3edf75f47b26b1941bc101aacf5c8e",
		Addr:     "ec2-54-197-100-79.compute-1.amazonaws.com:5432",
		Database: "d2u0p50e5eoprb",
	}

	con = pg.Connect(options)
	if con == nil {
		log.Println("Cannot connect to postgres")
	}

	return
}
