package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

//db.collection.createIndex({ "location": "2dsphere" })

/*
db.collection.find({
  "location": {
    "$near": {
      "$geometry": {
        "type": "Point",
        "coordinates": [-73.97, 40.77] // longitude, latitude order
      },
      "$maxDistance": 1000 // distance in meters
    }
  }
})
*/

type Location struct {
	Type        string    `bson:"type"`
	Coordinates []float64 `bson:"coordinates"`
}

func NewLocation(coordinates []float64) *Location {
	return &Location{
		Type:        "Point",
		Coordinates: coordinates,
	}
}

type Advertisement struct {
	ID                      primitive.ObjectID `bson:"_id"`
	ActiveFlag              *bool              `bson:"ActiveFlag"`
	AdditionalInfo          *string            `bson:"AdditionalInfo"`
	AdvertisementDate       *time.Time         `bson:"AdvertisementDate,omitempty"` // Pointer to time.Time for optional dates
	AdvertisementType       *string            `bson:"AdvertisementType,omitempty"`
	AvailabilityDate        *time.Time         `bson:"AvailabilityDate,omitempty"`
	AdvertiserId            *int64             `bson:"AdvertiserId,omitempty"`
	City                    *string            `bson:"City,omitempty"`
	ClientId                *int64             `bson:"ClientId,omitempty"`
	DailyPrice              *float32           `bson:"DailyPrice,omitempty"` // Using float32 for prices
	Description             string             `bson:"Description"`
	FavoriteFlag            *bool              `bson:"FavoriteFlag"`
	HourlyPrice             *float32           `bson:"HourlyPrice,omitempty"`
	Latitude                *float64           `bson:"Latitude,omitempty"`
	Limit                   *int               `bson:"Limit"`
	Longitude               *float64           `bson:"Longitude,omitempty"`
	Miles                   *float64           `bson:"Miles,omitempty"`
	MonthlyPrice            *float32           `bson:"MonthlyPrice,omitempty"`
	Municipality            *string            `bson:"Municipality,omitempty"`
	Neighborhood            *string            `bson:"Neighborhood,omitempty"`
	Number                  *string            `bson:"Number,omitempty"`
	Period                  *string            `bson:"Period,omitempty"`
	PeriodPrice             *float32           `bson:"PeriodPrice,omitempty"`
	Photos                  []string           `bson:"Photos,omitempty"` // Assuming Photos is an array of string. Adjust if the structure is different.
	PriceFrom               *float32           `bson:"PriceFrom,omitempty"`
	PriceUntil              *float32           `bson:"PriceUntil,omitempty"`
	Size                    *string            `bson:"Size,omitempty"`
	State                   *string            `bson:"State,omitempty"`
	Street                  *string            `bson:"Street,omitempty"`
	Title                   string             `bson:"Title"`
	ValidAdvertisementFrom  *time.Time         `bson:"ValidAdvertisementFrom,omitempty"`
	ValidAdvertisementUntil *time.Time         `bson:"ValidAdvertisementUntil,omitempty"`
	WeeklyPrice             *float32           `bson:"WeeklyPrice,omitempty"`
	YearlyPrice             *float32           `bson:"YearlyPrice,omitempty"`
	ZipCode                 *string            `bson:"ZipCode,omitempty"`
	Keywords                []string           `bson:"Keywords"`
	Location                *Location          `bson:"location,omitempty"` // Replaces separate Latitude and Longitude fields
}
