package validators

import (
	"errors"
	uc "final-project/deliveries/controllers/user"
	"regexp"
)

func ValidateCreateUser(req uc.RequestCreateUser) error {
	namePattern, _ := regexp.Compile(`^([A-Za-z]+ ?[A-Za-z]*){4,29}[A-Z-a-z]$`)
	if len(req.Name) > 30 || !namePattern.MatchString(req.Name) {
		return errors.New("input nama tidak sesuai (alfabet; tanpa simbol; boleh ada spasi di antara kata; total karakter: minimal 4, maksimal 30;)\ncontoh: \"Yusuf Nur Wahid\"")
	}

	emailPattern, _ := regexp.Compile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !emailPattern.MatchString(req.Email) {
		return errors.New("input email tidak sesuai (username: lowercase alfanumerik, simbol yang diperbolehkan hanya '_' dan '-'; pemisah: simbol '@'; domain name: lowercase alfanumerik, simbol yang diperbolehkan hanya '-'; pemisah: simbol '.'; domain: lowercase alfabet, minimal 2 karakter, maksimal 4 karakter;)\ncontoh: \"yusuf@mail.com\"")
	}

	passwordPattern, _ := regexp.Compile(`^[a-zA-Z0-9!@#$&()\\\-\x60.+,/\"]{5,8}$`)
	if !passwordPattern.MatchString(req.Password) {
		return errors.New("input password tidak sesuai (alfanumerik; boleh menggunakan simbol; total karakter: minimal 5, maksimal 8;)")
	}
	return nil
}
