package manager

import (
	"apiLunchLite/models"
)

func (m *ApiManager) AddApi(item models.ApiConfig) {
	m.Db.Save(item)
}
