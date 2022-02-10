package server

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func NewValidator() *Validator {
	return &Validator{
		validator: validator.New(),
	}
}

type Validator struct {
	validator *validator.Validate
}

func (v *Validator) Validate(i interface{}) error {
	err := v.CustomValidate()
	if err != nil {
		return err
	}
	return v.validator.Struct(i)
}

func (v *Validator) CustomValidate() error {
	//v.validator.VarWithValue(password, confirmpassword, "eqfield")

	if err := v.validator.RegisterValidation("password", validatePassword); err != nil {
		return err
	}
	if err := v.validator.RegisterValidation("uuid", validateUUIDV4); err != nil {
		return err
	}
	return nil
}

func validatePassword(fl validator.FieldLevel) bool {
	if len(fl.Field().String()) < 8 {
		return false
	}
	return true
}

func validateUUIDV4(fl validator.FieldLevel) bool {
	if _, err := uuid.Parse(fl.Field().String()); err == nil {
		return true
	}
	return false
}
