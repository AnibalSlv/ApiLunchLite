package manager

import (
	"apiLunchLite/internal/utils"
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

func (m *ApiManager) Check() error {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	result, err := m.Db.GetAll()

	if err != nil {
		log.Fatal("Error:", err)
		return err
	}

	if len(result) == 0 {
		fmt.Println("No hay APIs registradas. Usa el comando 'addApi' para empezar.")
		return nil
	}

	t.AppendHeader(table.Row{"ID", "Nombre", "Host", "State", "PID"})

	for _, api := range result {
		displayState := utils.Capitalize(api.State)
		switch api.State {
		case "stop":
			displayState = color.RedString(displayState)
		case "run":
			displayState = color.GreenString(displayState)
		default:
			displayState = color.YellowString(displayState)
		}

		t.AppendRow(table.Row{
			api.Id,
			api.Name,
			fmt.Sprintf("%s:%d", api.Host, api.Port),
			displayState,
			api.Pid,
		})
	}

	t.SetStyle(table.StyleRounded)

	// Convierte el texto en el header de la tabla en negrita
	t.Style().Color.Header = text.Colors{text.Bold}
	t.Render()

	return nil
}
