package manager

import (
	"apiLunchLite/models"
	"crypto/rand"
)

func (m *ApiManager) AddApi(item models.ApiConfig) (models.ApiConfig, error, bool) {
	if item.PathFolder == "" {
		return item, nil, false
	}

	if item.Name == "" {
		// Crea una lista o slice de tipo binario con una longitud de 5
		randomBytes := make([]byte, 5)

		// Llena el slice de bytes
		_, err := rand.Read(randomBytes)
		if err != nil {
			return item, err, false
		}

		result := make([]byte, len(randomBytes))

		letters := "abcdefghijklmnopqrstuvwxyz"

		for i, b := range randomBytes {
			// Le asigna los valores a la lista de bytes diviendo b con el tamano en bytes de letters y obtiene el residuo
			// esa es la letra que almacenara
			result[i] = letters[b%byte(len(letters))]
		}

		item.Name = string(result)
	}

	m.Db.Save(item)

	return item, nil, true
}
