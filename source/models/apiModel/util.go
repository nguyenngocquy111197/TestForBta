package apiModel

import (
	"errors"
	"regexp"
	"strings"
)

func checkNumber(number string) bool {
	regexNumber := regexp.MustCompile("^[0-9]*$")

	return regexNumber.MatchString(number)
}

func CheckPhoneNumber(phonenumber string) error {
	//check phonenumber
	result := false

	if len(phonenumber) == 10 && strings.HasPrefix(phonenumber, "0") {
		result = checkNumber(phonenumber)

	}
	if !result {
		return errors.New("Phone invalid")
	}
	return nil
}
