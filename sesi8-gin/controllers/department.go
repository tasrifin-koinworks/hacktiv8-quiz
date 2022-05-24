package controllers

import (
	"net/http"
	"sesi8-gin/params"
	"sesi8-gin/services"

	"github.com/gin-gonic/gin"
)

type DepartmentController struct {
	departmentService services.DepartmentService
}

func NewDepartmentController(service *services.DepartmentService) *DepartmentController {
	return &DepartmentController{
		departmentService: *service,
	}
}

func (d *DepartmentController) CreateNewDepartment(c *gin.Context) {
	var req params.CreateDepartment

	err := c.ShouldBind(&req)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status:         http.StatusBadRequest,
			Error:          "BAD REQUEST",
			AdditionalInfo: err,
		})

		return
	}

	response := d.departmentService.CreateDepartment(req)

	c.JSON(response.Status, response)
}

func (d *DepartmentController) GetAllDepartments(c *gin.Context) {
	response := d.departmentService.GetAllDepartments()
	c.JSON(response.Status, response)
}
