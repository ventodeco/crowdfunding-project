package travel

import "errors"

type Service interface {

	// RegisterUser(request RegisterUserRequest) (User, error)
	// LoginUser(request LoginRequest) (User, error)
	// IsEmailAvailable(request CheckEmailRequest) (bool, error)
	// SaveAvatar(Id int, fileLocation string) (User, error)
	GetTravelById(Id int) (TravelLocation, error)
	GetTravelLocations() ([]TravelLocation, error)
	GetFavoriteTravel() ([]TravelLocation, error)
	UpdateTravel(Id int, like bool) (TravelLocation, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

// func (s *service) RegisterUser(request RegisterUserRequest) (User, error) {
// 	user := User{}
// 	user.Name = request.Name
// 	user.Email = request.Email
// 	user.Occupation = request.Occupation
// 	passwordHash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.MinCost)

// 	if err != nil {
// 		return user, err
// 	}

// 	user.PasswordHash = string(passwordHash)
// 	user.Role = "user"

// 	newUser, err := s.repository.Save(user)

// 	if err != nil {
// 		return newUser, err
// 	}

// 	return newUser, nil
// }

// func (s *service) LoginUser(request LoginRequest) (User, error) {
// 	email := request.Email
// 	password := request.Password

// 	user, err := s.repository.FindByEmail(email)

// 	if err != nil {
// 		return user, err
// 	}

// 	if user.ID == 0 {
// 		return user, errors.New("User not found on that email")
// 	}

// 	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))

// 	if err != nil {
// 		return user, err
// 	}

// 	return user, nil
// }

// func (s *service) IsEmailAvailable(request CheckEmailRequest) (bool, error) {
// 	email := request.Email

// 	userByEmail, err := s.repository.FindByEmail(email)

// 	if err != nil {
// 		return false, err
// 	}

// 	if userByEmail.ID == 0 {
// 		return true, nil
// 	}

// 	return false, nil
// }

// func (s *service) SaveAvatar(Id int, fileLocation string) (User, error) {

// 	user, err := s.repository.FindById(Id)
// 	if err != nil {
// 		return user, err
// 	}

// 	user.AvatarFileName = fileLocation

// 	updatedUser, err := s.repository.Update(user)
// 	if err != nil {
// 		return updatedUser, err
// 	}

// 	return updatedUser, nil
// }

func (s *service) GetTravelById(Id int) (TravelLocation, error) {
	user, err := s.repository.FindById(Id)

	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("travel tidak ditemukan")
	}

	return user, nil
}

func (s *service) GetTravelLocations() ([]TravelLocation, error) {
	travels, err := s.repository.GetAll()

	if err != nil {
		return travels, err
	}

	return travels, nil
}

func (s *service) GetFavoriteTravel() ([]TravelLocation, error) {
	travels, err := s.repository.GetByLiked()

	if err != nil {
		return travels, err
	}

	return travels, nil
}

func (s *service) UpdateTravel(Id int, like bool) (TravelLocation, error) {

	travel, err := s.repository.FindById(Id)
	if err != nil {
		return travel, err
	}

	travel.IsLiked = like

	updatedTravel, err := s.repository.Update(travel)
	if err != nil {
		return updatedTravel, err
	}

	return updatedTravel, nil
}
