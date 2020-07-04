package repository

import (
	"main/internal/models"
	"main/internal/tools/errors"
)

func (db *DB) CreateRoom(room *models.Room) (*models.Room,error){
	err := db.DBConnPool.QueryRow("INSERT INTO rooms (name, about) VALUES($1, $2) ON CONFLICT DO NOTHING RETURNING id", room.Name, room.About).
		Scan(&room.ID)
	if err != nil {
		return nil, errors.RoomAlreadyExists
	}
	return room, nil
}

func (db *DB) DeleteRoom(roomID int) error{
	row, err := db.DBConnPool.Exec("DELETE FROM rooms WHERE id = $1", roomID)
	if err != nil {
		return err
	}
	if row.RowsAffected() == 0 {
		return errors.RoomNotFound
	}
	return nil
}

func (db *DB)  GetInfoAboutRoom(roomID int) (*models.Room, error) {
	room := models.Room{}
	err := db.DBConnPool.QueryRow("SELECT id, name, about FROM rooms WHERE id = $1", roomID).Scan(&room.ID, &room.Name, &room.About)
	if err != nil {
		return nil, err
	}
	rows, err := db.DBConnPool.Query("SELECT u.nickname, u.name, u.phone, u.status, m.user_id, m.toilet_paper_count, m.pay_ability, m.valuable FROM members AS m INNER JOIN users AS u ON (m.user_id = u.id) WHERE room_id = $1", roomID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		member := models.Member{}
		err = rows.Scan(&member.Nickname, &member.Name, &member.Phone, &member.Status, &member.UserID, &member.ToiletPaperCount, &member.PayAbility, &member.Valuable)
		if err != nil {
			rows.Close()
			return nil, err
		}
		room.Members = append(room.Members, &member)
	}
	rows.Close()
	if err != nil {
		return nil, err
	}
	return &room, nil
}