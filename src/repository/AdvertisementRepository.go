package repository

import (
	"fmt"
	"udemy/src/database"
	"udemy/src/domain"

	"go.mongodb.org/mongo-driver/bson"
)

type AdvertisementRepository struct {
}

func (a AdvertisementRepository) List() []domain.Advertisement {
	database.Init()
	defer database.Disconnect()
	users := database.Find[domain.Advertisement]("anuncio", bson.D{})

	for i := 0; i < len(users); i++ {
		user := users[i]
		fmt.Println(user)
	}

	return users
}
