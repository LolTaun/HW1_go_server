package pkg

import (
	"fmt"
	// "regexp"
	"strings"

)

func PhoneNormalize(phone string) (normalizedPhone string, err error) {
	// 8 (900) 123-12-21 -> +7 900 123-12-21
	// 89001231221 -> +7 900 123-12-21
	// +7 900 123-12-21 -> +7 900 123-12-21
	// 8 900 123-12-21 -> +7 900 123-12-21
	// 8 900 123 12 21 -> +7 900 123-12-21
	eW := NewEWrapper("PhoneNormalize()")
	normalizedPhone = removeNonDigits(phone)
	if normalizedPhone[0] == '8' {
		normalizedPhone = "+7" + normalizedPhone[1:]
	}

	if normalizedPhone[:2] != "+7" {
		err = eW.WrapError(fmt.Errorf("wrong phone number format in phone (length error): %s", phone), "normalizedPhone[:2] != \"+7\"")
		return phone, err
	}

	if len(normalizedPhone) != 12 {
		err = eW.WrapError(fmt.Errorf("wrong phone number format in phone (length error): %s", phone), "len(normalizedPhone) != 12")
		return phone, err
	}

	normalizedPhone = normalizedPhone[:2] + " (" + normalizedPhone[2:5] + ") " + normalizedPhone[5:8] + "-" + normalizedPhone[8:10] + "-" + normalizedPhone[10:12]

	return normalizedPhone, nil
}

func removeNonDigits(s string) string {
    return strings.Map(func(r rune) rune {
        if r >= '0' && r <= '9' {
            return r
        }
        return -1
    }, s)
}