package dto

import (
	// "encoding/json"
	"errors"
)

type Record struct {
	ID         int64  `json:"-" sql.field:"id"`
	Name       string `json:"name,omitempty" sql.field:"name"`
	LastName   string `json:"last_name,omitempty" sql.field:"last_name"`
	MiddleName string `json:"middle_name,omitempty" sql.field:"middle_name"`
	Address    string `json:"address,omitempty" sql.field:"address"`
	Phone      string `json:"phone,omitempty" sql.field:"phone"`
}

type Cond struct {
	Lop    string
	PgxInd string
	Field  string
	Value  any
}


func RecordValidation(record Record) (err error) {
    if err != nil {
        return err
    }

    if record.Name == "" {
        return errors.New("Name is empty")
    }

    if record.LastName == "" {
        return errors.New("LastName is empty")
    }

    if record.MiddleName == "" {
        return errors.New("MiddleName is empty")
    }

    if record.Address == "" {
        return errors.New("Address is empty")
    }

    if record.Phone == "" {
        return errors.New("Phone is empty")
    }

    return nil
}