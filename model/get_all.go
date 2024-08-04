package model

import "apisimples/db"

func GetAll() (indices []Indice, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM indice`)
	if err != nil {
		return
	}

	for rows.Next() {
		var indice Indice

		err = rows.Scan(&indice.ID_Indice, &indice.Descricao_Indice)
		if err != nil {
			continue
		}

		indices = append(indices, indice)
	}

	return
}
