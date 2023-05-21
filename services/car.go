package services

import (
	"beego-car-rental/dtos"
	"beego-car-rental/models"
	"errors"
)

// todo: better error messages when row is not present in db

func GetCar(car *models.Car, id int) error {
	if id < 1 {
		return errors.New("id is less than one")
	}
	car.Id = id
	if err := car.Read(); err != nil {
		return err
	}
	return nil
}

func GetAllCars(cars *models.CarList) error {
	if err := cars.ReadAll(); err != nil {
		return err
	}
	return nil
}

func UpdateCar(carDto *dtos.CarDto, id int) error {
	if id < 1 {
		return errors.New("id is less than one")
	}

	car := &models.Car{
		Id:       id,
		Make:     carDto.Make,
		Model:    carDto.Model,
		RegPlate: carDto.RegPlate,
	}

	if err := car.Update(); err != nil {
		return err
	}
	return nil
}

func DeleteCar(id int) error {
	car := &models.Car{Id: id}
	if id < 1 {
		return errors.New("id is less than one")
	}
	if err := car.Delete(); err != nil {
		return err
	}
	return nil
}

func CreateCar(carDto *dtos.CarDto) error {
	if carDto.RegPlate == "" {
		return errors.New("empty field, please fill it")
	}

	car := &models.Car{
		Make:     carDto.Make,
		Model:    carDto.Model,
		RegPlate: carDto.RegPlate,
	}

	return car.Insert()
}
