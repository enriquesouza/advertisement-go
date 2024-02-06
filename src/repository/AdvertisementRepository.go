package repository

import (
	"advertisement-go/src/database"
	"advertisement-go/src/domain"

	"go.mongodb.org/mongo-driver/bson"
)

type AdvertisementRepository struct {
}

func (a AdvertisementRepository) List() []domain.Advertisement {
	database.Init()
	defer database.Disconnect()
	ads := database.Find[domain.Advertisement]("anuncio", bson.D{})

	return ads
}

func (a AdvertisementRepository) Insert(items []domain.Advertisement) {
	database.Init()
	defer database.Disconnect()

	database.InsertMany[domain.Advertisement]("anuncio", items)
}

func (a AdvertisementRepository) ListByGeoLocation(longitude, latitude float64) []domain.Advertisement {
	database.Init()
	defer database.Disconnect()
	ads := database.Find[domain.Advertisement]("anuncio", bson.D{
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

	return ads
}

func (a AdvertisementRepository) ListByGeoLocationWithMaxDistance(longitude, latitude float64, maxDistance int32) []domain.Advertisement {
	database.Init()
	defer database.Disconnect()

	if maxDistance == 0 {
		maxDistance = 1000
	}

	ads := database.Find[domain.Advertisement]("anuncio", bson.D{
		{"location", bson.D{
			{"$near", bson.D{
				{"$geometry", bson.D{
					{"type", "Point"},
					{"coordinates", bson.A{longitude, latitude}},
				}},
				{"$maxDistance", maxDistance},
			}},
		}},
	})

	return ads
}
