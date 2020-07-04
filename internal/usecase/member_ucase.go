package usecase

import "main/internal/models"

func (u *useCase) AddMemberIntoRoom(member *models.Member) error {
	return u.repository.AddMemberIntoRoom(member)
}

func (u *useCase) DeleteMemberFromRoom(member *models.Member) error {
	return u.repository.DeleteMemberFromRoom(member)
}

func (u *useCase) GetWhoBuy(roomID int) (*models.Member, error) {
	room, err := u.repository.GetInfoAboutRoom(roomID)
	if err != nil {
		return nil, err
	}
	/* Algorithm of God */

	WhoBuy := new(models.Member)
	var MinCountToiletPaper int
	for i, v := range room.Members {
		if v.Valuable && v.PayAbility {
			WhoBuy = room.Members[i]
			MinCountToiletPaper = room.Members[i].ToiletPaperCount
			break
		}
	}
	for i, v := range room.Members {
		if v.Valuable && v.PayAbility {
			if v.ToiletPaperCount < MinCountToiletPaper {
				WhoBuy = room.Members[i]
				MinCountToiletPaper = room.Members[i].ToiletPaperCount
			}
		}
	}
	return WhoBuy, nil
}

func (u *useCase) UpdatePaperToilet(paper *models.Paper) (int, error) {
	return u.repository.UpdatePaperToilet(paper)
}

func (u *useCase) ChangeValuable(valuable *models.Valuable) (bool, error) {
	return u.repository.ChangeValuable(valuable)
}

func (u *useCase) ChangePayAbility(ability *models.PayAbility) (bool, error) {
	return u.repository.ChangePayAbility(ability)
}

