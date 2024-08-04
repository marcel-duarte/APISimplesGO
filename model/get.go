package model

import "apisimples/db"

func Get(id_indice int64) (indice Indice, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	defer conn.Close()

	row := conn.QueryRow(`SELECT * FROM indice WHERE id_indice = $1`, id_indice)

	err = row.Scan(&indice.ID_Indice, &indice.Descricao_Indice)

	return
}
