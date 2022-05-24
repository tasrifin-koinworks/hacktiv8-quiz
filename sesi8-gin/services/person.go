package services

import (
	"net/http"
	"sesi8-gin/models"
	"sesi8-gin/params"
	"sesi8-gin/repositories"
)

type PersonService struct {
	personRepo repositories.PersonRepo
}

// var db = gorm.DB{}
// var repo = repositories.NewPersonRepo(&db)

func NewPersonServices(repo repositories.PersonRepo) *PersonService {
	return &PersonService{
		personRepo: repo,
	}
}
func (p *PersonService) CreatePerson(request params.CreatePerson) *params.Response {
	model := models.Person{
		FirstName: request.FirstName,
		LastName:  request.LastName,
	}

	err := p.personRepo.CreatePerson(&model)

	if err != nil {
		return &params.Response{
			Status:         400,
			Error:          "BAD REQUEST",
			AdditionalInfo: err.Error(),
		}
	}

	return &params.Response{
		Status:  200,
		Error:   "CREATE SUKSES",
		Payload: request,
	}
}

func (p *PersonService) GetAllPersons() *params.Response {
	response, err := p.personRepo.GetAllPersons()

	if err != nil {
		return &params.Response{
			Status:         http.StatusInternalServerError,
			Error:          "INTERNAL SERVER ERROR",
			AdditionalInfo: err.Error(),
		}
	}

	return &params.Response{
		Status:  http.StatusOK,
		Payload: response,
	}

}
