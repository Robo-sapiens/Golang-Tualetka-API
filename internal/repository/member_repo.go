package repository

import (
	"main/internal/models"
	"main/internal/tools/errors"
)

func (db *DB) AddMemberIntoRoom (member *models.Member) error {
	result, err := db.DBConnPool.Exec("INSERT INTO members (user_id, room_id, toilet_paper_count, pay_ability, valuable) VALUES($1, $2, $3, $4, $5) ON CONFLICT DO NOTHING", member.UserID, member.RoomID, 0, true, true)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return errors.MemberAlreadyExists
	}
	_, err = db.DBConnPool.Exec("UPDATE users SET status = 'member' WHERE id = $1", member.UserID)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) DeleteMemberFromRoom (member *models.Member) error {
	row, err := db.DBConnPool.Exec("DELETE FROM members WHERE user_id = $1 AND room_id = $2", member.UserID, member.RoomID)
	if err != nil {
		return err
	}
	if row.RowsAffected() == 0 {
		return errors.MemberNotFound
	}
	_, err = db.DBConnPool.Exec("UPDATE users SET status = 'common' WHERE id = $1", member.UserID)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) UpdatePaperToilet(paper *models.Paper) (int, error){
	row := db.DBConnPool.QueryRow("UPDATE members SET toilet_paper_count = toilet_paper_count + $1 WHERE user_id = $2 AND room_id = $3 RETURNING toilet_paper_count", paper.PaperCount, paper.UserID, paper.RoomID)
	var paperCount int
	err := row.Scan(&paperCount)
	if err != nil {
		return paperCount, err
	}
	return paperCount, nil
}

func (db *DB) ChangeValuable(valuable *models.Valuable) (bool, error){
	row := db.DBConnPool.QueryRow("UPDATE members SET valuable = $1 WHERE user_id = $2 AND room_id = $3 RETURNING valuable", valuable.Valuable,valuable.UserID, valuable.RoomID)
	var valueAbility bool
	err := row.Scan(&valueAbility)
	if err != nil {
		return valueAbility, err
	}
	return valueAbility, nil
}

func (db *DB) ChangePayAbility(ability *models.PayAbility) (bool, error) {
	row := db.DBConnPool.QueryRow("UPDATE members SET pay_ability = $1 WHERE user_id = $2 AND room_id = $3 RETURNING pay_ability", ability.PayAbility, ability.UserID, ability.RoomID)
	var payAbility bool
	err := row.Scan(&payAbility)
	if err != nil {
		return payAbility, err
	}
	return payAbility, nil
}
