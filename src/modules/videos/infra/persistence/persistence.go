package persistence

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	cohere "github.com/cohere-ai/cohere-go/v2"
	client "github.com/cohere-ai/cohere-go/v2/client"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
	"os"
	"talentpitch/src/modules/videos/domain"
	"talentpitch/src/modules/videos/infra/persistence/entityData"
)

type videosRepository struct {
	db *gorm.DB
}

func NewVideosRepository(db *gorm.DB) domainvideos.VideosRepository {
	db.AutoMigrate(entityData.Videos{})
	return &videosRepository{
		db: db,
	}
}

func (v *videosRepository) Create(videos domainvideos.Videos) error {
	tx := v.db.Create(&entityData.Videos{
		ID:   uuid.New().String(),
		Name: videos.Name,
		Url:  videos.Url,
	})

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (v *videosRepository) GetVideosByID(Id string) (*domainvideos.Videos, error) {
	video := entityData.Videos{}
	result := v.db.First(&video, "id = ?", Id)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("video data not found")
	}

	return video.ToEntity(), nil
}

func (v *videosRepository) Update(videoEntity domainvideos.Videos) error {
	video := entityData.Videos{}

	result := v.db.Model(video).Where("id = ?", videoEntity.ID).Updates(entityData.Videos{
		ID:   videoEntity.ID,
		Name: videoEntity.Name,
		Url:  videoEntity.Url,
	})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("video data not update")
	}

	return nil
}

func (v *videosRepository) DeleteByID(Id string) error {
	var video entityData.Videos
	result := v.db.Where("id = ?", Id).Delete(&video)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("video data not update")
	}

	return nil
}

func (v *videosRepository) GetVideos(pageSize, offset int) ([]*domainvideos.Videos, error) {
	videos := []entityData.Videos{}
	result := v.db.Limit(pageSize).Offset(offset).Find(&videos)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return []*domainvideos.Videos{}, nil
	}

	response := make([]*domainvideos.Videos, len(videos))

	for i := 0; i < len(videos); i++ {
		response[i] = videos[i].ToEntity()
	}

	return response, nil
}

func (v *videosRepository) MassiveCreate() {
	co := client.NewClient(client.WithToken(os.Getenv("TOKEN_AI")))

	resp, err := co.Chat(
		context.TODO(),
		&cohere.ChatRequest{
			Message: "dame un array en formato json con 20 objectos que cumplan esta estructura {'name': '','url': ''} sin espacios y todo en una sola linea ",
		},
	)

	if err != nil {
		log.Fatal(err)
	}

	videos := []domainvideos.Videos{}

	for i := 0; i < len(resp.ChatHistory); i++ {
		if resp.ChatHistory[i].Role == "CHATBOT" {
			message := resp.ChatHistory[i].Chatbot.Message

			err = json.Unmarshal([]byte(message), &videos)
			if err != nil {
				fmt.Println("Error generating data by videos flow")
			}
			break
		}
	}

	for i := 0; i < len(videos); i++ {
		err = v.Create(videos[i])
		if err != nil {
			fmt.Println("error creating videos massive: ", err.Error())
		}
	}
}
