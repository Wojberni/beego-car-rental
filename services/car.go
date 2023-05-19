package services

import (
	"beego-car-rental/models"
	"beego-car-rental/utils"
	"errors"
)

func GetCar(car *models.Car, uuid string) error {
	if uuid == "" {
		return errors.New("uuid is empty")
	}
	car.Uuid = uuid
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

func UpdateCar(car *models.Car, uuid string) error {
	if uuid == "" {
		return errors.New("uuid is empty")
	}
	car.Uuid = uuid
	if err := car.Read(); err != nil {
		return err
	}
	return nil
}

func DeleteCar(uuid string) error {
	car := &models.Car{Uuid: uuid}
	if uuid == "" {
		return errors.New("uuid is empty")
	}
	if err := car.Delete(); err != nil {
		return err
	}
	return nil
}

func CreateCar(car *models.Car) error {
	if car.RegPlate == "" {
		return errors.New("empty field, please fill it")
	}
	car.Uuid = utils.GenerateUuid()

	return car.Insert()
}
