package persistence

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	cohere "github.com/cohere-ai/cohere-go/v2"
	"github.com/cohere-ai/cohere-go/v2/client"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
	"os"
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

func (v *challengeRepository) MassiveCreate() {
	co := client.NewClient(client.WithToken(os.Getenv("TOKEN_AI")))

	resp, err := co.Chat(
		context.TODO(),
		&cohere.ChatRequest{
			Message: "dame un array en formato json con 20 objectos que cumplan esta estructura {'title': '','description': ''} sin espacios y todo en una sola linea ",
		},
	)

	if err != nil {
		log.Fatal(err)
	}

	challenges := []domainchallenge.Challenges{}

	for i := 0; i < len(resp.ChatHistory); i++ {
		if resp.ChatHistory[i].Role == "CHATBOT" {
			message := resp.ChatHistory[i].Chatbot.Message

			err = json.Unmarshal([]byte(message), &challenges)
			if err != nil {
				fmt.Println("Error generating data by challenges flow")
			}
			break
		}
	}

	for i := 0; i < len(challenges); i++ {
		err = v.Create(challenges[i])
		if err != nil {
			fmt.Println("error creating challenges massive: ", err.Error())
		}
	}
}
