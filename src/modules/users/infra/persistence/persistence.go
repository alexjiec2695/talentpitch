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
	domainuser "talentpitch/src/modules/users/domain"
	"talentpitch/src/modules/users/infra/persistence/entityData"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domainuser.UserRepository {
	db.AutoMigrate(entityData.User{})
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) Create(user domainuser.User) error {
	tx := u.db.Create(&entityData.User{
		ID:    uuid.New().String(),
		Name:  user.Name,
		Email: user.Email,
	})

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (u *userRepository) GetUserByID(Id string) (*domainuser.User, error) {
	user := entityData.User{}
	result := u.db.First(&user, "id = ?", Id)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("user data not found")
	}

	return user.ToEntity(), nil
}

func (u *userRepository) Update(userEntity domainuser.User) error {
	user := entityData.User{}

	result := u.db.Model(user).Where("id = ?", userEntity.ID).Updates(entityData.User{
		ID:    userEntity.ID,
		Name:  userEntity.Name,
		Email: userEntity.Email,
	})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("user data not update")
	}

	return nil
}

func (u *userRepository) DeleteByID(Id string) error {
	var user entityData.User
	result := u.db.Where("id = ?", Id).Delete(&user)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("user data not update")
	}

	return nil
}

func (u *userRepository) GetUsers(pageSize, offset int) ([]*domainuser.User, error) {
	users := []entityData.User{}
	result := u.db.Limit(pageSize).Offset(offset).Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return []*domainuser.User{}, nil
	}

	response := make([]*domainuser.User, len(users))

	for i := 0; i < len(users); i++ {
		response[i] = users[i].ToEntity()
	}

	return response, nil
}

func (u *userRepository) MassiveCreate() {
	co := client.NewClient(client.WithToken(os.Getenv("TOKEN_AI")))

	resp, err := co.Chat(
		context.TODO(),
		&cohere.ChatRequest{
			Message: "dame un array en formato json con 20 objectos que cumplan esta estructura {'name': '','email': ''} sin espacios y todo en una sola linea ",
		},
	)

	if err != nil {
		log.Fatal(err)
	}

	users := []domainuser.User{}

	for i := 0; i < len(resp.ChatHistory); i++ {
		if resp.ChatHistory[i].Role == "CHATBOT" {
			message := resp.ChatHistory[i].Chatbot.Message

			err = json.Unmarshal([]byte(message), &users)
			if err != nil {
				fmt.Println("Error generating data by users flow")
			}
			break
		}
	}

	for i := 0; i < len(users); i++ {
		err = u.Create(users[i])
		if err != nil {
			fmt.Println("error creating users massive: ", err.Error())
		}
	}
}
