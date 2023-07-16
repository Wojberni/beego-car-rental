package services

import (
	"beego-car-rental/dtos"
	"beego-car-rental/models"
	"errors"
	"strconv"
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

	carId, err := strconv.Atoi(orderDto.CarId)
	userId, err2 := strconv.Atoi(orderDto.UserId)

	if err != nil || err2 != nil {
		return errors.New("error converting id to int")
	}
	if carId < 1 || userId < 1 {
		return errors.New("car or user id can't be less than one")
	}

	car := &models.Car{
		Id: carId,
	}
	car.Read()
	user := &models.User{
		Id: userId,
	}
	user.Read()
	order := &models.Order{
		Id:   id,
		Car:  car,
		User: user,
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
	carId, err := strconv.Atoi(orderDto.CarId)
	userId, err2 := strconv.Atoi(orderDto.UserId)

	if err != nil || err2 != nil {
		return errors.New("error converting id to int")
	}
	if carId < 1 || userId < 1 {
		return errors.New("car or user id can't be less than one")
	}

	car := &models.Car{
		Id: carId,
	}
	user := &models.User{
		Id: userId,
	}
	order := &models.Order{
		Car:  car,
		User: user,
	}

	return order.Insert()
}
