package services

import (
	"beego-car-rental/dtos"
	"beego-car-rental/models"
	"errors"
)

// todo: better error messages when row is not present in db

func GetOrder(order *models.Order, id int) error {
	if id < 1 {
		return errors.New("id is less than one")
	}
	order.Id = id
	if err := order.Read(); err != nil {
		return err
	}
	return nil
}

func GetAllOrders(orders *models.OrderList) error {
	if err := orders.ReadAll(); err != nil {
		return err
	}
	return nil
}

func UpdateOrder(orderDto *dtos.OrderDto, id int) error {
	if id < 1 {
		return errors.New("id is less than one")
	}

	order := &models.Order{
		Id: id,
		// TODO: valid order DTO
		// Cars: orderDto.Cars,
		// Users: orderDto.Users,
	}

	if err := order.Update(); err != nil {
		return err
	}
	return nil
}

func DeleteOrder(id int) error {
	order := &models.Order{Id: id}
	if id < 1 {
		return errors.New("id is less than one")
	}
	if err := order.Delete(); err != nil {
		return err
	}
	return nil
}

func CreateOrder(orderDto *dtos.OrderDto) error {

	car := &models.Order{
		// TODO: valid order DTO
		// Cars: orderDto.Cars,
		// Users: orderDto.Users,
	}

	return car.Insert()
}
