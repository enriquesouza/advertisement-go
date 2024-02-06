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

func (a AdvertisementRepository) Insert(items []domain.Advertisement) {
	database.Init()
	defer database.Disconnect()

	database.InsertMany[domain.Advertisement]("anuncio", items)
}

func (a AdvertisementRepository) ListByGeoLocation(longitude, latitude float64) []domain.Advertisement {
	database.Init()
	defer database.Disconnect()
	users := database.Find[domain.Advertisement]("anuncio", bson.D{
		{"location", bson.D{
			{"$near", bson.D{
				{"$geometry", bson.D{
					{"type", "Point"},
					{"coordinates", bson.A{longitude, latitude}},
				}},
				{"$maxDistance", 1000},
			}},
		}},
	})

	for i := 0; i < len(users); i++ {
		user := users[i]
		fmt.Println(user)
	}

	return users
}
