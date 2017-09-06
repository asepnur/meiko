package user

import (
	"fmt"
	"html"
	"strconv"

	"regexp"

	"github.com/melodiez14/meiko/src/util/helper"
	validator "gopkg.in/asaskevich/govalidator.v4"
)

func (s setStatusUserParams) Validate() (*setStatusUserArgs, error) {
	// Email Validation
	if len(s.Email) < 1 {
		return nil, fmt.Errorf("Error validation: email cant't be empty")
	}
	if len(s.Email) > 45 {
		return nil, fmt.Errorf("Error validation : email too longer")
	}

	if !validator.IsEmail(s.Email) {
		return nil, fmt.Errorf("%s is not an email", s.Email)
	}
	email, err := helper.NormalizeEmail(html.EscapeString(s.Email))
	if err != nil {
		return nil, err
	}

	//Code Validation
	if len(s.Code) < 1 {
		return nil, fmt.Errorf("Error validation : Code can't be empty")
	} else if len(s.Code) != 4 {
		return nil, fmt.Errorf("Error validation : Wrong code")
	}
	val, err := regexp.MatchString(`[0-9]+$`, s.Code)
	if !val || err != nil {
		return nil, fmt.Errorf("Error validation: Wrong code")
	}

	c, err := strconv.ParseInt(s.Code, 10, 16)
	if err != nil {
		return nil, fmt.Errorf("Error validation: Wrong code")
	}

	args := &setStatusUserArgs{
		Email: email,
		Code:  uint16(c),
	}

	return args, nil

}
func (s signUpParams) Validate() (*signUpArgs, error) {

	// Email Validation
	if len(s.Email) < 1 {
		return nil, fmt.Errorf("Error validation: email cant't be empty")
	}
	if len(s.Email) > 45 {
		return nil, fmt.Errorf("Error validation : email too longer")
	}

	if !validator.IsEmail(s.Email) {
		return nil, fmt.Errorf("%s is not an email", s.Email)
	}
	email, err := helper.NormalizeEmail(html.EscapeString(s.Email))
	if err != nil {
		return nil, err
	}

	// Password Validation
	password := html.EscapeString(s.Password)
	if len(password) < 1 {
		return nil, fmt.Errorf("Error validation: password can't be empty")
	}
	if len(password) < 6 {
		return nil, fmt.Errorf("Error validation: password at least consist of 6 characters")
	}
	regexPassword := []string{`[a-z]`, `[A-Z]`, `[0-9]`}
	for _, val := range regexPassword {
		is, _ := regexp.MatchString(val, password)
		if !is {
			return nil, fmt.Errorf("Error validation: password must contains alphanumeric upper and lower case")
		}
	}
	// ID validation
	if len(s.ID) < 1 {
		return nil, fmt.Errorf("Error validation: ID can't be empty")
	}

	if len(s.ID) != 12 {
		return nil, fmt.Errorf(fmt.Sprintf("ID : %s is wrong", s.ID))
	}

	id, err := strconv.ParseInt(s.ID, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("Error validation: ID must be numeric")
	}

	// Name validation
	if len(s.Name) < 1 {
		return nil, fmt.Errorf("Error validation: name cant't be empty")
	}
	if len(s.Name) > 50 {
		return nil, fmt.Errorf("Error validation: name cant't to long")
	}

	v, err := regexp.MatchString(`[A-z]+$`, html.EscapeString(s.Name))
	if !v || err != nil {
		return nil, fmt.Errorf("Error validation: name contains alphabet only")
	}
	// Result
	args := &signUpArgs{
		ID:       id,
		Name:     s.Name,
		Email:    email,
		Password: s.Password,
	}
	return args, nil
}

func (s signInParams) Validate() (*signInArgs, error) {

	// Email Validation
	if len(s.Email) < 1 {
		return nil, fmt.Errorf("Error validation: email cant't be empty")
	}

	email, err := helper.NormalizeEmail(html.EscapeString(s.Email))
	if err != nil {
		return nil, err
	}

	// Password Validation
	password := html.EscapeString(s.Password)
	if len(password) < 6 {
		return nil, fmt.Errorf("Error validation: password at least consist of 6 characters")
	}

	args := &signInArgs{
		Email:    email,
		Password: password,
	}
	return args, nil
}

func (f forgotRequestParams) Validate() (*forgotRequestArgs, error) {

	// Email Validation
	if len(f.Email) < 1 {
		return nil, fmt.Errorf("Error validation: email cant't be empty")
	}

	email, err := helper.NormalizeEmail(html.EscapeString(f.Email))
	if err != nil {
		return nil, err
	}

	args := &forgotRequestArgs{
		Email: email,
	}
	return args, nil
}

func (f forgotConfirmationParams) Validate() (*forgotConfirmationArgs, error) {

	// Email Validation
	if len(f.Email) < 1 {
		return nil, fmt.Errorf("Error validation: email cant't be empty")
	}

	email, err := helper.NormalizeEmail(html.EscapeString(f.Email))
	if err != nil {
		return nil, err
	}

	// Password Validation (Optional Field)
	if len(f.Password) > 0 {
		f.Password = html.EscapeString(f.Password)
		if len(f.Password) < 6 {
			return nil, fmt.Errorf("Error validation: password at least consist of 6 characters")
		}
	}

	// Code Validation
	if len(f.Code) < 1 {
		return nil, fmt.Errorf("Error validation: code cant't be empty")
	} else if len(f.Code) != 4 {
		return nil, fmt.Errorf("Error validation: code must be 4 digits")
	}

	c, err := strconv.ParseInt(f.Code, 10, 16)
	if err != nil {
		return nil, fmt.Errorf("Error validation: code should be numeric")
	}

	args := &forgotConfirmationArgs{
		Email:    email,
		Code:     uint16(c),
		Password: f.Password,
	}

	return args, nil
}
