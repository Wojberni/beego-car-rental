package controllers

import (
	"beego-car-rental/models"
	"beego-car-rental/services"
	"encoding/json"
	"fmt"

	beego "github.com/beego/beego/v2/server/web"
)

// CarController operations for Car
type CarController struct {
	beego.Controller
}

// @Title CreateCar
// @Description Create Car
// @Param	body		body 	models.Car	true		"Body for Car content"
// @Success 201 {object} models.Car
// @Failure 403 {string} error: error message
// @Accept json
// @router / [post]
func (c *CarController) Post() {
	car := &models.Car{}
	json.Unmarshal(c.Ctx.Input.RequestBody, car)
	if err := services.CreateCar(car); err != nil {
		// message := fmt.Sprintf("Error creating car: %v", err.Error())
		c.Data["json"] = map[string]string{"error": err.Error()}
	} else {
		message := fmt.Sprintf("Created car: %v!", car.RegPlate)
		c.Data["json"] = map[string]string{"message": message}
	}
	c.ServeJSON()
}

// @Title GetCar
// @Description Get Car by id
// @Param	id		path 	string	true		"The uuid of car to get"
// @Success 200 {object} models.Car
// @Failure 403 :id is empty
// @router /:id [get]
func (c *CarController) Get() {
	uid := c.GetString(":uuid")
	car := &models.Car{}
	if err := services.GetCar(car, uid); err != nil {
		c.Data["json"] = map[string]string{"error": err.Error()}
	} else {
		c.Data["json"] = car
	}
	c.ServeJSON()
}

// @Title GetAllCars
// @Description Get all Cars
// @Success 200 {object} models.Car
// @Failure 404 {string} Error retrieving data, please try again later!
// @Accept json
// @router / [get]
func (c *CarController) GetAll() {
	cars := &models.CarList{}
	if err := services.GetAllCars(cars); err != nil {
		c.Data["json"] = map[string]string{"error": err.Error()}
	} else {
		c.Data["json"] = cars
	}
	c.ServeJSON()
}

// @Title Put
// @Description update the Car
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Car	true		"body for Car content"
// @Success 200 {object} models.Car
// @Failure 403 :id is not int
// @router /:id [put]
func (c *CarController) Put() {
	uid := c.GetString(":uuid")
	car := &models.Car{}
	json.Unmarshal(c.Ctx.Input.RequestBody, car)
	if err := services.UpdateCar(car, uid); err != nil {
		c.Data["json"] = map[string]string{"error": err.Error()}
	} else {
		c.Data["json"] = car
	}
	c.ServeJSON()
}

// @Title Delete
// @Description delete the Car
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *CarController) Delete() {
	uid := c.GetString(":uuid")
	if err := services.DeleteCar(uid); err != nil {
		c.Data["json"] = map[string]string{"error": err.Error()}
	} else {
		c.Data["json"] = map[string]string{"message": fmt.Sprintf("Deleted car: %s", uid)}
	}
	c.ServeJSON()
}
