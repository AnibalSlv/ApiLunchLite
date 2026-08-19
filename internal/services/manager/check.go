package manager

import (
	"apiLunchLite/models"
)

func (m *ApiManager) Check() ([]models.ApiType, error) {

	result, err := m.Db.GetAll()

	if err != nil {
		return nil, err
	}

	return result, nil
}
