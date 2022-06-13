package travel

import (
	"time"
)

type TravelLocation struct {
	ID           int
	Name         string
	ImageUrl     string
	Subtitle     string
	ShortAddress string
	LongAddress  string
	PhoneNumber  string
	IsLiked      bool
	Rating       int
	MapUrl       string
	Destinations []Destination
	ImagesTravel []ImageTravel
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Destination struct {
	ID               uint
	Name             string
	Price            float32
	TravelLocationId int
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type ImageTravel struct {
	ID               uint
	ImageUrl         string
	TravelLocationId int
}
