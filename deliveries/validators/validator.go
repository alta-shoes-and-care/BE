package validators

import (
	"errors"
	"mime/multipart"
	"net/http"
	"regexp"
	"strings"
)

func ValidateCreateUser(name, email, password string) error {
	namePattern, _ := regexp.Compile(`^([A-Za-z]+ ?[A-Za-z]*){3,29}[A-Z-a-z]$`)
	if len(name) > 30 || !namePattern.MatchString(name) {
		return errors.New("input nama tidak sesuai (alfabet; tanpa simbol; boleh ada spasi di antara kata; total karakter: minimal 4, maksimal 30;)\ncontoh: \"Yusuf Nur Wahid\"")
	}

	emailPattern, _ := regexp.Compile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !emailPattern.MatchString(email) {
		return errors.New("input email tidak sesuai (username: lowercase alfanumerik, simbol yang diperbolehkan hanya '_' dan '-'; pemisah: simbol '@'; domain name: lowercase alfanumerik, simbol yang diperbolehkan hanya '-'; pemisah: simbol '.'; domain: lowercase alfabet, minimal 2 karakter, maksimal 4 karakter;)\ncontoh: \"yusuf@mail.com\"")
	}

	passwordPattern, _ := regexp.Compile(`^[a-zA-Z0-9!@#$&()\\\-\x60.+,/\"]{5,8}$`)
	if !passwordPattern.MatchString(password) {
		return errors.New("input password tidak sesuai (alfanumerik; boleh menggunakan simbol; total karakter: minimal 5, maksimal 8;)")
	}
	return nil
}

func ValidateUpdateUser(name, email, password string) error {
	namePattern, _ := regexp.Compile(`^([A-Za-z]+ ?[A-Za-z]*){3,29}[A-Z-a-z]$`)
	if len(name) != 0 && (len(name) > 30 || !namePattern.MatchString(name)) {
		return errors.New("input nama tidak sesuai (alfabet; tanpa simbol; boleh ada spasi di antara kata; total karakter: minimal 4, maksimal 30;)\ncontoh: \"Yusuf Nur Wahid\"")
	}

	emailPattern, _ := regexp.Compile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if len(email) != 0 && !emailPattern.MatchString(email) {
		return errors.New("input email tidak sesuai (username: lowercase alfanumerik, simbol yang diperbolehkan hanya '_' dan '-'; pemisah: simbol '@'; domain name: lowercase alfanumerik, simbol yang diperbolehkan hanya '-'; pemisah: simbol '.'; domain: lowercase alfabet, minimal 2 karakter, maksimal 4 karakter;)\ncontoh: \"yusuf@mail.com\"")
	}

	passwordPattern, _ := regexp.Compile(`^[a-zA-Z0-9!@#$&()\\\-\x60.+,/\"]{5,8}$`)
	if len(password) != 0 && !passwordPattern.MatchString(password) {
		return errors.New("input password tidak sesuai (alfanumerik; boleh menggunakan simbol; total karakter: minimal 5, maksimal 8;)")
	}
	return nil
}

func ValidateCreateService(title, description string) error {
	titlePattern, _ := regexp.Compile(`^[a-zA-Z0-9 ]*$`)
	descPattern, _ := regexp.Compile(`^[a-zA-Z0-9\W ]*$`)
	if len(title) > 30 || !titlePattern.MatchString(title) {
		return errors.New("input title tidak sesuai (alfanumerik; tanpa simbol; boleh ada spasi di antara kata; total karakter: maksimal 30;)\ncontoh: \"Service 1\"")
	}

	if len(description) > 320 || !descPattern.MatchString(description) {
		return errors.New("input description tidak sesuai (alfanumerik; boleh simbol; boleh ada spasi di antara kata; total karakter: maksimal 320;)\ncontoh: \"Layanan regular yang disukai orang.\"")
	}
	return nil
}

func ValidateUpdateServiceData(title, description string) error {
	titlePattern, _ := regexp.Compile(`^[a-zA-Z0-9 ]*$`)
	descPattern, _ := regexp.Compile(`^[a-zA-Z0-9\W ]*$`)
	if len(title) > 30 || !titlePattern.MatchString(title) {
		return errors.New("input title tidak sesuai (alfanumerik; tanpa simbol; boleh ada spasi di antara kata; total karakter: maksimal 30;)\ncontoh: \"Service 1\"")
	}

	if len(description) > 320 || !descPattern.MatchString(description) {
		return errors.New("input description tidak sesuai (alfanumerik; boleh simbol; boleh ada spasi di antara kata; total karakter: maksimal 320;)\ncontoh: \"Layanan regular yang disukai orang.\"")
	}
	return nil
}

func ValidateServiceImage(file *multipart.FileHeader) (int, error) {
	if file.Size > 512000 {
		return http.StatusRequestEntityTooLarge, errors.New("ukuran gambar melebihi 500kB")
	}

	if !strings.HasSuffix(file.Filename, ".jpg") || !strings.HasSuffix(file.Filename, ".jpeg") || !strings.HasSuffix(file.Filename, ".png") {
		return http.StatusBadRequest, errors.New("hanya menerima file dengan ekstensi jpg, jpeg, dan png")
	}
	return http.StatusOK, nil
}
