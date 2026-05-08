package manager

import (
	"apiLunchLite/internal/services"
	"apiLunchLite/models"
)

type ApiManager struct {
	Db     models.Db
	Logger *services.Logger
}

func New(db models.Db) *ApiManager {
	return &ApiManager{Db: db}
}
