/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"apiLunchLite/internal/platform"
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Ejecuta el reverse proxy (EN PROCESO)",
	Long:  `Recomendado no tocar`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Levantando servicios en segundo plano...")

		// La ruta de tu propia CLI instalada
		pathEjecutable, _ := os.Executable()

		//  Llama a la misma CLI pero pasándole el comando oculto "core"
		clon := exec.Command(pathEjecutable, "core")

		// Se desvinculan el proceso de la terminal actual (dependiendo del SO)
		platform.StartDaemon(clon)

		// Los log se van a un archivo para no perderlos
		logFile, err := apiMgr.Logger.GetLogFile("aplDaemon")
		if err != nil {
			fmt.Println(err)
		}
		clon.Stderr = logFile
		clon.Stdout = logFile

		if err := clon.Start(); err != nil {
			fmt.Printf("Error al lanzar el core: %v\n", err)
			return
		}

		// Guarda el PID en un archivo
		// []byte convierte el texto en bytes porque es lo que necesita la funcion para escrbir en el disco
		// 0644 son los permisos (en Linux/Unix)
		os.WriteFile("internal/logs/aplDaemon.pid", []byte(fmt.Sprintf("%d", clon.Process.Pid)), 0644)

		fmt.Println("✨ ¡Entorno listo! Puedes seguir usando tu terminal.")

	},
}

func init() {
	rootCmd.AddCommand(startCmd)

}
