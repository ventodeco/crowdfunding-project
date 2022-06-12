package travel

import "gorm.io/gorm"

type Repository interface {
	GetAll() ([]TravelLocation, error)
	GetByLiked() ([]TravelLocation, error)
	FindById(travelId int) (TravelLocation, error)
	Update(travel TravelLocation) (TravelLocation, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAll() ([]TravelLocation, error) {
	var travels []TravelLocation

	err := r.db.Preload("Destinations").Preload("ImagesTravel").Find(&travels).Error

	if err != nil {
		return travels, err
	}

	return travels, nil
}

func (r *repository) GetByLiked() ([]TravelLocation, error) {
	var travels []TravelLocation

	err := r.db.Where("is_liked = true").Preload("Destinations").Find(&travels).Error

	if err != nil {
		return travels, err
	}

	return travels, nil
}

func (r *repository) FindById(travelId int) (TravelLocation, error) {
	var travel TravelLocation

	err := r.db.Where("ID = ?", travelId).Find(&travel).Error

	if err != nil {
		return travel, err
	}

	return travel, nil
}

func (r *repository) Update(travel TravelLocation) (TravelLocation, error) {
	err := r.db.Save(&travel).Error

	if err != nil {
		return travel, err
	}

	return travel, nil
}
