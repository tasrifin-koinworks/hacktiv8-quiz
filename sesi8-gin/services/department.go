package services

import (
	"net/http"
	"sesi8-gin/models"
	"sesi8-gin/params"
	"sesi8-gin/repositories"
)

type DepartmentService struct {
	departmentRepo repositories.DepartmentRepo
}

// var db = gorm.DB{}
// var repo = repositories.NewDepartmentRepo(&db)

func NewDepartmentServices(repo repositories.DepartmentRepo) *DepartmentService {
	return &DepartmentService{
		departmentRepo: repo,
	}
}

func (d *DepartmentService) CreateDepartment(request params.CreateDepartment) *params.Response {
	model := models.Department{
		DepartmentName: request.DepartmentName,
	}

	err := d.departmentRepo.CreateDepartment(&model)

	if err != nil {
		return &params.Response{
			Status:         400,
			Error:          "BAD REQUEST",
			AdditionalInfo: err.Error(),
		}
	}

	return &params.Response{
		Status:         200,
		Error:          "CREATE SUCCESS",
		AdditionalInfo: request,
	}
}

func (d *DepartmentService) GetAllDepartments() *params.Response {
	response, err := d.departmentRepo.GetAllDepartments()

	if err != nil {
		return &params.Response{
			Status:         http.StatusInternalServerError,
			Error:          "BAD REQUEST",
			AdditionalInfo: err.Error(),
		}
	}

	return &params.Response{
		Status:         http.StatusOK,
		AdditionalInfo: response,
	}

}
