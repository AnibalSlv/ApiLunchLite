package manager

import (
	"fmt"
	"os"
)

func (m *ApiManager) CheckLife() error {
	result, err := m.Db.GetAll()

	if err != nil {
		fmt.Println("Error Connection Db: ", err)
		return err
	}

	for _, api := range result {
		if api.Pid == 0 {
			continue
		}

		_, err := os.FindProcess(api.Pid)

		if err != nil {
			m.Db.UpdatePID(0, api.Id)
			m.Db.UpdateState("stop", api.Id)
			continue
		}

	}

	return nil
}
