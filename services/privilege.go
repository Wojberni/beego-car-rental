package services

import (
	"beego-car-rental/dtos"
	"beego-car-rental/models"
	"errors"
)

// todo: better error messages when row is not present in db

func GetPrivilege(privilege *models.Privilege, id int) error {
	if id < 1 {
		return errors.New("id is less than one")
	}
	privilege.Id = id
	if err := privilege.Read(); err != nil {
		return err
	}
	return nil
}

// add to controller later?
func GetPrivilegeByName(privilege *models.Privilege, name string) error {
	if name == "" {
		return errors.New("name is empty")
	}
	privilege.Name = name
	if err := privilege.Read("name"); err != nil {
		return err
	}
	return nil
}

func GetAllPrivileges(privileges *models.PrivilegeList) error {
	if err := privileges.ReadAll(); err != nil {
		return err
	}
	return nil
}

func UpdatePrivilege(privilegeDto *dtos.PrivilegeDto, id int) error {
	if id < 1 {
		return errors.New("id is less than one")
	}

	privilege := &models.Privilege{
		Id:   id,
		Name: privilegeDto.Name,
	}

	if err := privilege.Update(); err != nil {
		return err
	}
	return nil
}

func DeletePrivilege(id int) error {
	privilege := &models.Privilege{Id: id}
	if id < 1 {
		return errors.New("id is less than one")
	}
	if err := privilege.Delete(); err != nil {
		return err
	}
	return nil
}

func CreatePrivilege(privilegeDto *dtos.PrivilegeDto) error {
	if privilegeDto.Name == "" {
		return errors.New("empty field, please fill it")
	}

	privilege := &models.Privilege{
		Name: privilegeDto.Name,
	}

	return privilege.Insert()
}
