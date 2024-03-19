package repository

import (
	"online_store/db"
)

type repository struct {
	db db.DBTX
}

func NewRepository(db db.DBTX) Repository {
	return &repository{db: db}
}
