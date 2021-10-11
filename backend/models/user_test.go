package models

import (
	"main/public"
	"testing"
)

func TestUserCreateValidation(t *testing.T) {
	//Case => Empty Name
	userForName := User{
		ID:       1,
		Name:     "",
		Email:    "test@hotmail.com",
		Status:   true,
		Password: "test",
	}
	errorForName := userForName.ValidateFor(ValidationStatus.CREATE)
	if errorForName == nil {
		t.Error("Name field cannot be empty")
	}

	//Case => Empty Email
	userForEmail := User{
		ID:       1,
		Name:     "test",
		Email:    "",
		Status:   true,
		Password: "test",
	}
	errorForEmail := userForEmail.ValidateFor(ValidationStatus.CREATE)
	if errorForEmail == nil {
		t.Error("Email field cannot be empty")
	}

	//Case => Invalid Mail
	userForInvalidEmail := User{
		ID:       1,
		Name:     "test",
		Email:    "aaa@.com",
		Status:   true,
		Password: "test",
	}
	errorForInvalidEmail := userForInvalidEmail.ValidateFor(ValidationStatus.CREATE)
	if errorForInvalidEmail == nil {
		t.Error("It's not valid email address")
	}

	//Case => Empty Password
	userForPassword := User{
		ID:       1,
		Name:     "test",
		Email:    "test@hotmail.com",
		Status:   true,
		Password: "",
	}
	errorForPassword := userForPassword.ValidateFor(ValidationStatus.CREATE)
	if errorForPassword == nil {
		t.Error("Password field cannot be empty")
	}

	//Case => Correct Fields
	userForCorrectField := User{
		ID:       1,
		Name:     "test",
		Email:    "test@hotmail.com",
		Status:   true,
		Password: "test",
	}
	errorForCorrectField := userForCorrectField.ValidateFor(ValidationStatus.CREATE)
	if errorForCorrectField != nil {
		t.Error("It is correct body!")
	}
}

func TestUserUpdateValidation(t *testing.T) {
	//Case => Empty fields (name,password)
	userForEmptyField := User{
		ID:       1,
		Name:     "",
		Email:    "test@hotmail.com",
		Status:   true,
		Password: "",
	}
	errorForEmptyField := userForEmptyField.ValidateFor(ValidationStatus.UPDATE)
	if errorForEmptyField == nil {
		t.Error("the request should be body have a any field for update")
	}

	//Case => Invalid Mail
	userForInvalidEmail := User{
		ID:       1,
		Name:     "test",
		Email:    "test@.com",
		Status:   true,
		Password: "test",
	}
	errorForInvalidEmail := userForInvalidEmail.ValidateFor(ValidationStatus.UPDATE)
	if errorForInvalidEmail != nil {
		t.Error("It's not valid email address")
	}
}

func TestUserToPublic(t *testing.T) {
	userForToPublic := User{
		ID:       1,
		Name:     "",
		Email:    "test@hotmail.com",
		Status:   true,
		Password: "test",
	}
	publicData := userForToPublic.ToPublic()
	manualPublicData := public.User{
		ID:    userForToPublic.ID,
		Name:  userForToPublic.Name,
		Email: userForToPublic.Email,
	}
	if publicData != manualPublicData {
		t.Error("It's not valid email address")
	}
}
