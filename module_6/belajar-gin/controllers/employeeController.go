package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Employee struct {
	ID       int
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Division string `json:"division"`
}

var employess []Employee = []Employee{
	{
		ID:       1,
		Name:     "Labib",
		Age:      20,
		Division: "Operation",
	},
}

func CreateEmployee(ctx *gin.Context) {
	var newEmployee Employee
	ctx.ShouldBindJSON(&newEmployee)

	newEmployee.ID = len(employess) + 1
	employess = append(employess, newEmployee)

	ctx.JSON(http.StatusCreated, newEmployee)
}

func GetAllEmployees(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, employess)
}

func GetAllEmployeesWithView(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "template.html", employess)
	// ctx.JSON(http.StatusOK, employess)
}

func ShowEmployee(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id")
	idEmployee, _ := strconv.Atoi(id)

	var employee Employee

	for _, value := range employess {
		if value.ID == idEmployee {
			employee = value
			break
		}
	}

	if employee.ID == 0 {
		ctx.JSON(http.StatusNotFound, "Data Not Found")
		return
	}

	ctx.JSON(http.StatusOK, employee)
}

func DeleteEmployee(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id")
	idEmployee, _ := strconv.Atoi(id)

	var employee Employee
	var indexFound int

	for index, value := range employess {
		if value.ID == idEmployee {
			employee = value
			indexFound = index
			break
		}
	}

	if employee.ID == 0 {
		ctx.JSON(http.StatusNotFound, "Data Not Found")
		return
	}

	employess = append(employess[:indexFound], employess[indexFound+1:]...)

	ctx.JSON(http.StatusOK, "Successfuly Deleting Data")
}

func UpdateEmployee(ctx *gin.Context) {
	var updateEmployee Employee
	ctx.ShouldBindJSON(&updateEmployee)

	id, _ := ctx.Params.Get("id")
	idEmployee, _ := strconv.Atoi(id)

	var employee Employee
	var indexFound int

	for index, value := range employess {
		if value.ID == idEmployee {
			employee = value
			indexFound = index
			break
		}
	}

	if employee.ID == 0 {
		ctx.JSON(http.StatusNotFound, "Data Not Found")
		return
	}

	employee.Name = updateEmployee.Name
	employee.Age = updateEmployee.Age
	employee.Division = updateEmployee.Division

	employess[indexFound] = employee

	ctx.JSON(http.StatusOK, "Successfuly Updating Data")

}
