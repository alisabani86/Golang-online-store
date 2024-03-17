package repository

import (
	"server/db"
)

type repository struct {
	db db.DBTX
}

func NewRepository(db db.DBTX) Repository {
	return &repository{db: db}
}
