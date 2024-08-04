package model

import "apisimples/db"

func Insert(indice Indice) (id_indice int64, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	defer conn.Close()

	sql := `INSERT INTO indice (descricao_indice) VALUES ($1) RETURNING id_indice`

	err = conn.QueryRow(sql, indice.Descricao_Indice).Scan(&id_indice)

	return
}
