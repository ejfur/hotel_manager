package postgres

import (
	_ "crypto/rsa"
	"database/sql"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

func CreateGuest(db *sql.DB, fn, sn string, room uuid.UUID) (uuid.UUID, error) {

	//INSERT GUEST INTO DB RETURNING UUID

	var id uuid.UUID
	err := db.QueryRow(`
				INSERT INTO guest(
					id,
					first_name,
					second_name, 
					current_room_num, 
					created_at,
					checked_in
				) VALUES (
					uuid_generate_v4(),
					$1,
					$2,
					$3,
					CURRENT_TIMESTAMP, 
					CURRENT_TIMESTAMP
				) RETURNING id
			`, fn, sn, room).Scan(&id)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func CheckInGuest(db *sql.DB, room uuid.UUID, fn string) (sql.Result, error) {

	// CHECH IN A GUEST

	result, err := db.Exec(`UPDATE guest 
							set current_room_num = $1, 
							checked_in = current_timestamp, 
							checked_out = null 
							where guest.first_name = $2`, room, fn)

	if err != nil {
		return result, err
	}

	return result, nil
}

func CheckOutGuest(db *sql.DB, room uuid.UUID) (sql.Result, error) {

	// CHECH OUT A GUEST

	result, err := db.Exec("UPDATE guest set checked_out = current_timestamp where current_room_num = $1", room)

	if err != nil {
		return result, err
	}
	rowsAffected, err := result.RowsAffected()
	
	if rowsAffected == 0 || err != nil {
		return result, err
	}
	return result, nil
}

func DeleteGuest(db *sql.DB, fn string) (sql.Result, error) {

	// DELETE GUEST

	result, err := db.Exec("update guest set deleted_at = CURRENT_TIMESTAMP where first_name = $1", fn)
	if err != nil {
		return result, err
	}

	rowsAffected, err := result.RowsAffected()
	
	if rowsAffected == 0 || err != nil {
		return result, err
	}
	return result, nil
}

func ChangeGuestRoom(db *sql.DB, room uuid.UUID, fn string) (sql.Result, error) {

	// UPDATE GUEST ROOM WHILE CHECHED IN

	result, err := db.Exec("UPDATE guest set current_room_num = $1 where first_name = $2", room, fn)
	if err != nil {
		return result, err
	}

	rowsAffected, err := result.RowsAffected()
	
	if rowsAffected == 0 || err != nil {
		return result, err
	}

	return result, nil
}
