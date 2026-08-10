package database

import (
	"database/sql"
	"fmt"
	"log"

	"apiLunchLite/models"

	_ "modernc.org/sqlite"
)

type ApiType = models.ApiConfig

type SQLite struct {
	DbConn *sql.DB
}

// Crea una tabla en la d.DbConn si no existe
func (d *SQLite) InitDb() error {
	createTableSQL := `CREATE TABLE IF NOT EXISTS APIs (
	"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	"name" TEXT,
	"host" TEXT,
	"port" INTEGER,
	"path_folder" TEXT,
	"state" TEXT DEFAULT "stop",
	"pid"  INTEGER DEFAULT 0
	);`

	_, err := d.DbConn.Exec(createTableSQL)
	if err != nil {
		log.Fatal("Error Create Table")
		return err
	}

	return nil
}

func (d *SQLite) Save(Api ApiType) {
	query := `INSERT INTO APIs (name, host, port, path_folder) VALUES (?, ?, ?, ?)`

	_, err := d.DbConn.Exec(query, Api.Name, Api.Host, Api.Port, Api.PathFolder)
	if err != nil {
		fmt.Println("Error Save Data:", err)
	}
	fmt.Println("Data Save")
}

func (d *SQLite) UpdatePID(pid int, id int) error {
	query := `UPDATE APIs SET pid = ? WHERE id = ?`

	_, err := d.DbConn.Exec(query, pid, id)
	if err != nil {
		return err
	}

	return nil
}

func (d *SQLite) UpdateState(state string, id int) error {
	query := `UPDATE APIs SET state = ? WHERE id = ?`

	_, err := d.DbConn.Exec(query, state, id)
	if err != nil {
		return err
	}

	return nil
}

func (d *SQLite) GetAll() ([]ApiType, error) {
	var Api ApiType
	query := `SELECT id, name, host, port, path_folder, state, pid FROM APIs`

	// ⁡⁢⁣⁣.Query()⁡ Es usado para los SELECT
	rows, err := d.DbConn.Query(query)
	if err != nil {
		return nil, err
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

func (d *SQLite) GetName(name string) (ApiType, error) {
	var Api ApiType

	query := `SELECT id, name, host, port, path_folder, state, pid FROM APIs WHERE name = ?`

	// Busca por nombre y devuelve los datos solicitado
	err := d.DbConn.QueryRow(query, name).Scan(&Api.Id, &Api.Name, &Api.Host, &Api.Port, &Api.PathFolder, &Api.State, &Api.Pid)

	if err != nil {
		return Api, err
	}

	return Api, nil

}

func (d *SQLite) GetId(id int) (ApiType, error) {
	var Api ApiType

	query := `SELECT id, name, host, port, path_folder, state, pid FROM APIs WHERE id = ?`

	err := d.DbConn.QueryRow(query, id).Scan(&Api.Id, &Api.Name, &Api.Host, &Api.Port, &Api.PathFolder, &Api.State, &Api.Pid)

	if err != nil {
		fmt.Println("Error Select ID:", err)
		return Api, err
	}

	return Api, nil
}

func (d *SQLite) Delete(id int) error {
	query := `DELETE FROM APIs WHERE id = ?`

	_, err := d.DbConn.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}
