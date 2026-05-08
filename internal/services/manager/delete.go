package manager

import (
	"fmt"

	"github.com/fatih/color"
)

func (m *ApiManager) Delete(apiName string) error {
	red := color.New(color.FgRed).SprintFunc()

	result, err := m.Db.GetName(apiName)

	if err != nil {
		fmt.Print("Error Connection Db: ", err)
		return err
	}

	m.Db.Delete(result.Id)

	fmt.Printf("\n%s %s\n", red(result.Name), red("Eliminado"))

	return nil
}
