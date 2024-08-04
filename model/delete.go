package model

import "apisimples/db"

func Delete(id_indice int64) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}

	defer conn.Close()

	res, err := conn.Exec(`DELETE FROM indice WHERE id_indice = $1`, id_indice)

	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
