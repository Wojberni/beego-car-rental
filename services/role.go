package services

import (
	"beego-car-rental/dtos"
	"beego-car-rental/models"
	"errors"
	"strings"
)

// todo: better error messages when row is not present in db

func GetRole(role *models.Role, id int) error {
	if id < 1 {
		return errors.New("id is less than one")
	}
	role.Id = id
	if err := role.Read(); err != nil {
		return err
	}
	return nil
}

func GetAllRoles(roles *models.RoleList) error {
	if err := roles.ReadAll(); err != nil {
		return err
	}
	return nil
}

func UpdateRole(roleDto *dtos.RoleDto, id int) error {
	if id < 1 {
		return errors.New("id is less than one")
	}

	privilegeList := strings.Split(roleDto.PrivilegeNames, ",")

	role := &models.Role{
		Id:   id,
		Name: roleDto.Name,
	}

	for _, value := range privilegeList {
		privilege := &models.Privilege{}
		if err := GetPrivilegeByName(privilege, value); err != nil {
			return err
		}
		role.Privileges = append(role.Privileges, privilege)
	}

	if err := role.Update(); err != nil {
		return err
	}

	return nil
}

func DeleteRole(id int) error {
	role := &models.Role{Id: id}
	if id < 1 {
		return errors.New("id is less than one")
	}
	if err := role.Delete(); err != nil {
		return err
	}
	return nil
}

func CreateRole(roleDto *dtos.RoleDto) error {
	if roleDto.Name == "" {
		return errors.New("empty field, please fill it")
	}

	privilegeList := strings.Split(roleDto.PrivilegeNames, ",")

	role := &models.Role{
		Name: roleDto.Name,
	}

	for _, value := range privilegeList {
		privilege := &models.Privilege{}
		if err := GetPrivilegeByName(privilege, value); err != nil {
			return err
		}
		role.Privileges = append(role.Privileges, privilege)
	}

	return role.Insert()
}
