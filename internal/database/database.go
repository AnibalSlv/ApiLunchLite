package database

import (
	"database/sql"
	"fmt"
	"log"

	"apiLunchLite/models"

	_ "modernc.org/sqlite"
)

type ApiType = models.ApiConfig

// Crea una tabla en la db si no existe
func InitDb() {
	db, err := sql.Open("sqlite", "DbAPL.db")
	if err != nil {
		log.Fatal("Error Open Db:", err)
	}
	defer db.Close()

	createTableSQL := `CREATE TABLE IF NOT EXISTS APIs (
	"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	"name" TEXT,
	"host" TEXT,
	"port" INTEGER,
	"path_folder" TEXT,
	"state" TEXT DEFAULT "stop",
	"pid"  INTEGER DEFAULT 0
	);`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal("Error Create Table")
	}

}

func Save(db *sql.DB, Api ApiType) {
	query := `INSERT INTO APIs (name, host, port, path_folder) VALUES (?, ?, ?, ?)`

	_, err := db.Exec(query, Api.Name, Api.Host, Api.Port, Api.PathFolder)
	if err != nil {
		fmt.Println("Error Save Data:", err)
	}
	fmt.Println("Data Save")
}

func UpdatePID(db *sql.DB, pid int, id int) {
	query := `UPDATE APIs SET pid = ? WHERE id = ?`

	_, err := db.Exec(query, pid, id)
	if err != nil {
		fmt.Println("Error Update PID: ", err)
	}
}

func UpdateState(db *sql.DB, state string, id int) {
	query := `UPDATE APIs SET state = ? WHERE id = ?`

	_, err := db.Exec(query, state, id)
	if err != nil {
		fmt.Println("Error Update PID: ", err)
	}
}

func GetAll(db *sql.DB) ([]ApiType, error) {
	var Api ApiType
	query := `SELECT id, name, host, port, path_folder, state, pid FROM APIs`

	// ⁡⁢⁣⁣.Query()⁡ Es usado para los SELECT
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("Error Select Data:", err)
	}
	// ⁡⁢⁣⁣defer:⁡ Cierra el proceso al terminar la funcion
	// Se tiene que cerrar todo proceso que este ligado a un proceso del SO
	// por ejemplo: una aplicacion externa, un sistema, etc.
	// en este caso se cierra porque SQLite es un archivo del disco duro
	defer rows.Close()

	var results []ApiType

	// ⁡⁢⁣⁣.Next()⁡ recorre la lista como si fuera una fila, devuelve True mientras tenga elementos
	// devuelve false cuando ya no tenga mas elementos
	for rows.Next() {
		// ⁡⁢⁣⁣.Scan()⁡ le asisgna los valores a las variables en orden de llegada es decir:
		// Como primero se esta recibiendo Id -> item.Id, luego Name -> item.Name
		err := rows.Scan(&Api.Id, &Api.Name, &Api.Host, &Api.Port, &Api.PathFolder, &Api.State, &Api.Pid)
		if err != nil {
			return nil, err
		}
		results = append(results, Api)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return results, nil

}

func GetName(db *sql.DB, name string) (ApiType, error) {
	var Api ApiType

	query := `SELECT id, name, host, port, path_folder, state, pid FROM APIs WHERE name = ?`

	// Busca por nombre y devuelve los datos solicitado
	err := db.QueryRow(query, name).Scan(&Api.Id, &Api.Name, &Api.Host, &Api.Port, &Api.PathFolder, &Api.State, &Api.Pid)

	if err != nil {
		fmt.Println("Error Search Name:", err)
		return Api, err
	}

	return Api, nil

}

func GetId(db *sql.DB, id int) (ApiType, error) {
	var Api ApiType

	query := `SELECT id, name, host, port, path_folder, state, pid FROM APIs WHERE id = ?`

	err := db.QueryRow(query, id).Scan(&Api.Id, &Api.Name, &Api.Host, &Api.Port, &Api.PathFolder, &Api.State, &Api.Pid)

	if err != nil {
		fmt.Println("Error Select ID:", err)
		return Api, err
	}

	return Api, nil
}
