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
// @Failure 401 {string} error: "Unauthenticated, please log in!"
// @Failure 500 {string} error: "message"
// @Accept json
// @router / [post]
func (o *OrderController) Post() {
	order := &dtos.OrderDto{}
	json.Unmarshal(o.Ctx.Input.RequestBody, order)
	if err := services.CreateOrder(order); err != nil {
		o.Data["json"] = map[string]string{"error": err.Error()}
		o.Ctx.Output.SetStatus(500)
	} else {
		message := "Created order!"
		o.Data["json"] = map[string]string{"message": message}
	}
	o.ServeJSON()
}

// @Title GetOrder
// @Description Get Order by id
// @Param	id		path 	string	true		"The id of order to get"
// @Success 200 {object} models.Order
// @Failure 401 {string} error: "Unauthenticated, please log in!"
// @Failure 500 {string} error: "message"
// @router /:id [get]
func (o *OrderController) Get() {
	id, _ := o.GetInt(":id")
	order := &models.Order{}
	if err := services.GetOrder(order, id); err != nil {
		o.Data["json"] = map[string]string{"error": err.Error()}
		o.Ctx.Output.SetStatus(500)
	} else {
		o.Data["json"] = order
	}
	o.ServeJSON()
}

// @Title GetAllOrders
// @Description Get all Orders
// @Success 200 {object} models.Order
// @Failure 401 {string} error: "Unauthenticated, please log in!"
// @Failure 500 {string} error: "message"
// @Accept json
// @router / [get]
func (o *OrderController) GetAll() {
	orders := &models.OrderList{}
	if err := services.GetAllOrders(orders); err != nil {
		o.Data["json"] = map[string]string{"error": err.Error()}
		o.Ctx.Output.SetStatus(500)
	} else {
		o.Data["json"] = orders
	}
	o.ServeJSON()
}

// @Title Put
// @Description update the Order
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	dtos.OrderDto	true		"body for Order content"
// @Success 200 {string} message: "Updated order: id"
// @Failure 401 {string} error: "Unauthenticated, please log in!"
// @Failure 500 {string} error: "message"
// @router /:id [put]
func (o *OrderController) Put() {
	id, _ := o.GetInt(":id")
	order := &dtos.OrderDto{}
	json.Unmarshal(o.Ctx.Input.RequestBody, order)
	if err := services.UpdateOrder(order, id); err != nil {
		o.Data["json"] = map[string]string{"error": err.Error()}
		o.Ctx.Output.SetStatus(500)
	} else {
		message := fmt.Sprintf("Updated order: %v!", id)
		o.Data["json"] = map[string]string{"message": message}
	}
	o.ServeJSON()
}

// @Title Delete
// @Description delete the Order
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} message: "Deleted order: id"
// @Failure 401 {string} error: "Unauthenticated, please log in!"
// @Failure 500 {string} error: "message"
// @router /:id [delete]
func (o *OrderController) Delete() {
	id, _ := o.GetInt(":id")
	if err := services.DeleteOrder(id); err != nil {
		o.Data["json"] = map[string]string{"error": err.Error()}
		o.Ctx.Output.SetStatus(500)
	} else {
		o.Data["json"] = map[string]string{"message": fmt.Sprintf("Deleted order: %v", id)}
	}
	o.ServeJSON()
}
