package database

import (
	"fmt"
	"neocheckin_cache/database/models"
	"neocheckin_cache/shared"
)

type MemoryDatabase struct {
	employees []models.Employee
	options   []models.Option
	actions   []models.Action
}

func findEmployee(a []models.Employee, f func(models.Employee) bool) (int, *models.Employee, error) {
	for i, v := range a {
		if f(v) {
			return i, &v, nil
		}
	}
	return -1, nil, fmt.Errorf("employee not found")
}

func findOption(a []models.Option, f func(models.Option) bool) (int, *models.Option, error) {
	for i, v := range a {
		if f(v) {
			return i, &v, nil
		}
	}
	return -1, nil, fmt.Errorf("option not found")
}

func findAction(a []models.Action, f func(models.Action) bool) (int, *models.Action, error) {
	for i, v := range a {
		if f(v) {
			return i, &v, nil
		}
	}
	return -1, nil, fmt.Errorf("action not found")
}

func (db *MemoryDatabase) GetEmployeeWithRfid(rfid string) (models.Employee, error) {
	_, empl, err := findEmployee(db.employees, func(e models.Employee) bool {
		return e.Rfid == rfid
	})

	if err == nil {
		return *empl, nil
	}

	return models.Employee{}, fmt.Errorf("could not find Employee with rfid '%s'", rfid)
}

func (db *MemoryDatabase) GetEmployeeWithDatabaseId(id string) (models.Employee, error) {
	_, empl, err := findEmployee(db.employees, func(e models.Employee) bool {
		return e.DatabaseId == id
	})

	if err == nil {
		return *empl, nil
	}

	return models.Employee{}, fmt.Errorf("could not find Employee with database id '%s'", id)
}

func (db *MemoryDatabase) InsertEmployee(empl models.Employee) error {
	_, oldEmpl, err := findEmployee(db.employees, func(e models.Employee) bool {
		return e.DatabaseId == empl.DatabaseId
	})

	if err == nil && oldEmpl.DatabaseId == empl.DatabaseId {
		return fmt.Errorf("employee with database id '%s' already exists", oldEmpl.DatabaseId)
	}

	db.employees = append(db.employees, empl)
	return nil
}

func (db *MemoryDatabase) UpdateEmployeeWithDatabaseId(id string, props models.Employee) error {
	_, empl, err := findEmployee(db.employees, func(e models.Employee) bool {
		return e.DatabaseId == id
	})

	if err == nil {
		empl.Rfid = props.Rfid
		empl.Name = props.Name
		empl.Flex = props.Flex
		empl.Working = props.Working
		empl.Department = props.Department
		empl.Photo = props.Photo

		return nil
	}

	return fmt.Errorf("could not find Employee with database id '%s'", id)
}

func (db *MemoryDatabase) DeleteEmployeeWithDatabaseId(id string) error {
	i, _, err := findEmployee(db.employees, func(e models.Employee) bool {
		return e.DatabaseId == id
	})

	if err == nil {
		db.employees[i] = db.employees[len(db.employees)-1]
		db.employees = db.employees[:len(db.employees)-1]
		return nil
	}

	return fmt.Errorf("could not find Employee with database id '%s'", id)
}

func (db *MemoryDatabase) GetOptionWithWrapperId(id shared.WrapperEnum) (models.Option, error) {
	_, opt, err := findOption(db.options, func(o models.Option) bool {
		return o.WrapperId == id
	})

	if err == nil {
		return *opt, nil
	}

	return models.Option{}, fmt.Errorf("could not find Option with wrapper id '%d'", id)
}

func (db *MemoryDatabase) GetOptionWithDatabaseId(id string) (models.Option, error) {
	_, opt, err := findOption(db.options, func(o models.Option) bool {
		return o.DatabaseId == id
	})

	if err == nil {
		return *opt, nil
	}

	return models.Option{}, fmt.Errorf("could not find Option with database id '%s'", id)
}

func (db *MemoryDatabase) InsertOption(opt models.Option) error {
	_, oldOpt, err := findOption(db.options, func(o models.Option) bool {
		return o.DatabaseId == opt.DatabaseId
	})

	if err == nil && opt.DatabaseId == oldOpt.DatabaseId {
		return fmt.Errorf("option with database id '%s' already exists", oldOpt.DatabaseId)
	}

	db.options = append(db.options, opt)
	return nil
}

func (db *MemoryDatabase) UpdateOptionWithDatabaseId(id string, props models.Option) error {
	_, opt, err := findOption(db.options, func(o models.Option) bool {
		return o.DatabaseId == id
	})

	if err == nil {
		opt.Name = props.Name
		opt.WrapperId = props.WrapperId
		opt.Available = props.Available
	}

	return fmt.Errorf("could not find Option with database id '%s'", id)
}

func (db *MemoryDatabase) DeleteOptionWithDatabaseId(id string) error {
	i, _, err := findOption(db.options, func(e models.Option) bool {
		return e.DatabaseId == id
	})

	if err == nil {
		db.options[i] = db.options[len(db.options)-1]
		db.options = db.options[:len(db.options)-1]
		return nil
	}

	return fmt.Errorf("could not find Option with database id '%s'", id)
}

func (db *MemoryDatabase) AddAction(action models.Action) error {
	_, oldAction, err := findAction(db.actions, func(e models.Action) bool {
		return e.DatabaseId == action.DatabaseId
	})

	if err == nil && oldAction.DatabaseId == action.DatabaseId {
		return fmt.Errorf("action with database id '%s' already exists", oldAction.DatabaseId)
	}

	db.actions = append(db.actions, action)
	return nil
}
