package controllers

import (
	"beego-car-rental/dtos"
	"beego-car-rental/models"
	"beego-car-rental/services"
	"encoding/json"
	"fmt"
)

// CarController operations for Car
type CarController struct {
	BaseController
}

// @Title CreateCar
// @Description Create Car
// @Param	body		body 	dtos.CarDto	true		"Body for Car content"
// @Success 201 {string} message: "Created car: RegPlate"
// @Failure 401 {string} error: "Unauthenticated, please log in!"
// @Failure 500 {string} error: "message"
// @Accept json
// @router / [post]
func (c *CarController) Post() {
	car := &dtos.CarDto{}
	json.Unmarshal(c.Ctx.Input.RequestBody, car)
	if err := services.CreateCar(car); err != nil {
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.Ctx.Output.SetStatus(500)
	} else {
		message := fmt.Sprintf("Created car: %v!", car.RegPlate)
		c.Data["json"] = map[string]string{"message": message}
	}
	c.ServeJSON()
}

// @Title GetCar
// @Description Get Car by id
// @Param	id		path 	string	true		"The id of car to get"
// @Success 200 {object} models.Car
// @Failure 401 {string} error: "Unauthenticated, please log in!"
// @Failure 500 {string} error: "message"
// @router /:id [get]
func (c *CarController) Get() {
	id, _ := c.GetInt(":id")
	car := &models.Car{}
	if err := services.GetCar(car, id); err != nil {
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.Ctx.Output.SetStatus(500)
	} else {
		c.Data["json"] = car
	}
	c.ServeJSON()
}

// @Title GetAllCars
// @Description Get all Cars
// @Success 200 {object} models.Car
// @Failure 401 {string} error: "Unauthenticated, please log in!"
// @Failure 500 {string} error: "message"
// @Accept json
// @router / [get]
func (c *CarController) GetAll() {
	cars := &models.CarList{}
	if err := services.GetAllCars(cars); err != nil {
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.Ctx.Output.SetStatus(500)
	} else {
		c.Data["json"] = cars
	}
	c.ServeJSON()
}

// @Title Put
// @Description update the Car
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	dtos.CarDto	true		"body for Car content"
// @Success 200 {string} message: "Updated car: id"
// @Failure 401 {string} error: "Unauthenticated, please log in!"
// @Failure 500 {string} error: "message"
// @router /:id [put]
func (c *CarController) Put() {
	id, _ := c.GetInt(":id")
	car := &dtos.CarDto{}
	json.Unmarshal(c.Ctx.Input.RequestBody, car)
	if err := services.UpdateCar(car, id); err != nil {
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.Ctx.Output.SetStatus(500)
	} else {
		message := fmt.Sprintf("Updated car: %v!", car.RegPlate)
		c.Data["json"] = map[string]string{"message": message}
	}
	c.ServeJSON()
}

// @Title Delete
// @Description delete the Car
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} message: "Deleted car: id"
// @Failure 401 {string} error: "Unauthenticated, please log in!"
// @Failure 500 {string} error: "message"
// @router /:id [delete]
func (c *CarController) Delete() {
	id, _ := c.GetInt(":id")
	if err := services.DeleteCar(id); err != nil {
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.Ctx.Output.SetStatus(500)
	} else {
		c.Data["json"] = map[string]string{"message": fmt.Sprintf("Deleted car: %v", id)}
	}
	c.ServeJSON()
}
