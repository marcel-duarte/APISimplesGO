package model

import "apisimples/db"

func Update(id_indice int64, indice Indice) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}

	defer conn.Close()

	res, err := conn.Exec(`UPDATE indice SET descricao_indice = $1 WHERE id_indice = $2`, indice.Descricao_Indice, id_indice)

	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
