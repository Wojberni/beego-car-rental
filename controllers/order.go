package controllers

import (
	"beego-car-rental/dtos"
	"beego-car-rental/models"
	"beego-car-rental/services"
	"encoding/json"
	"fmt"
)

// OrderController operations for Car
type OrderController struct {
	BaseController
}

// @Title CreateOrder
// @Description Create Order
// @Param	body		body 	dtos.OrderDto	true		"Body for Order content"
// @Success 201 {string} message: "Created order!"
// @Failure 403 {string} error: "message"
// @Accept json
// @router / [post]
func (c *OrderController) Post() {
	order := &dtos.OrderDto{}
	json.Unmarshal(c.Ctx.Input.RequestBody, order)
	if err := services.CreateOrder(order); err != nil {
		c.Data["json"] = map[string]string{"error": err.Error()}
	} else {
		message := "Created order!"
		c.Data["json"] = map[string]string{"message": message}
	}
	c.ServeJSON()
}

// @Title GetOrder
// @Description Get Order by id
// @Param	id		path 	string	true		"The id of order to get"
// @Success 200 {object} models.Order
// @Failure 403 {string} error: "message"
// @router /:id [get]
func (c *OrderController) Get() {
	id, _ := c.GetInt(":id")
	order := &models.Order{}
	if err := services.GetOrder(order, id); err != nil {
		c.Data["json"] = map[string]string{"error": err.Error()}
	} else {
		c.Data["json"] = order
	}
	c.ServeJSON()
}

// @Title GetAllOrders
// @Description Get all Orders
// @Success 200 {object} models.Order
// @Failure 404 {string} error: "message"
// @Accept json
// @router / [get]
func (c *OrderController) GetAll() {
	orders := &models.OrderList{}
	if err := services.GetAllOrders(orders); err != nil {
		c.Data["json"] = map[string]string{"error": err.Error()}
	} else {
		c.Data["json"] = orders
	}
	c.ServeJSON()
}

// @Title Put
// @Description update the Order
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	dtos.OrderDto	true		"body for Order content"
// @Success 200 {string} message: "Updated order: id"
// @Failure 403 {string} error: "message"
// @router /:id [put]
func (c *OrderController) Put() {
	id, _ := c.GetInt(":id")
	order := &dtos.OrderDto{}
	json.Unmarshal(c.Ctx.Input.RequestBody, order)
	if err := services.UpdateOrder(order, id); err != nil {
		c.Data["json"] = map[string]string{"error": err.Error()}
	} else {
		message := fmt.Sprintf("Updated order: %v!", id)
		c.Data["json"] = map[string]string{"message": message}
	}
	c.ServeJSON()
}

// @Title Delete
// @Description delete the Order
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} message: "Deleted order: id"
// @Failure 403 {string} error: "message"
// @router /:id [delete]
func (c *OrderController) Delete() {
	id, _ := c.GetInt(":id")
	if err := services.DeleteOrder(id); err != nil {
		c.Data["json"] = map[string]string{"error": err.Error()}
	} else {
		c.Data["json"] = map[string]string{"message": fmt.Sprintf("Deleted order: %v", id)}
	}
	c.ServeJSON()
}
