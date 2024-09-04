package persistence

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	domainchallenge "talentpitch/src/modules/challenges/domain"
	"talentpitch/src/modules/challenges/infra/persistence/entityData"
)

type challengeRepository struct {
	db *gorm.DB
}

func NewChallengesRepository(db *gorm.DB) domainchallenge.ChallengesRepository {
	db.AutoMigrate(entityData.Challenges{})
	return &challengeRepository{
		db: db,
	}
}

func (v *challengeRepository) Create(challenge domainchallenge.Challenges) error {
	tx := v.db.Create(&entityData.Challenges{
		ID:          uuid.New().String(),
		Title:       challenge.Title,
		Description: challenge.Description,
	})

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (v *challengeRepository) GetChallengesByID(Id string) (*domainchallenge.Challenges, error) {
	challenge := entityData.Challenges{}
	result := v.db.First(&challenge, "id = ?", Id)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("challenge data not found")
	}

	return challenge.ToEntity(), nil
}

func (v *challengeRepository) Update(challengeEntity domainchallenge.Challenges) error {
	challenge := entityData.Challenges{}

	result := v.db.Model(challenge).Where("id = ?", challengeEntity.ID).Updates(entityData.Challenges{
		ID:          challengeEntity.ID,
		Title:       challengeEntity.Title,
		Description: challengeEntity.Description,
	})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("challenge data not update")
	}

	return nil
}

func (v *challengeRepository) DeleteByID(Id string) error {
	var challenge entityData.Challenges
	result := v.db.Where("id = ?", Id).Delete(&challenge)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("challenge data not update")
	}

	return nil
}

func (v *challengeRepository) GetChallenges() ([]*domainchallenge.Challenges, error) {
	challenge := []entityData.Challenges{}
	result := v.db.Find(&challenge)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return []*domainchallenge.Challenges{}, nil
	}

	response := make([]*domainchallenge.Challenges, len(challenge))

	for i := 0; i < len(challenge); i++ {
		response[i] = challenge[i].ToEntity()
	}

	return response, nil
}
