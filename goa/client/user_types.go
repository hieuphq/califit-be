// Code generated by goagen v1.4.0, DO NOT EDIT.
//
// API "califit-be Backend API": Application User Types
//
// Command:
// $ goagen
// --design=github.com/hieuphq/califit-be/goa/design
// --out=$(GOPATH)/src/github.com/hieuphq/califit-be/goa
// --version=v1.4.0

package client

import (
	"github.com/goadesign/goa"
	"unicode/utf8"
)

// loginPayload user type.
type loginPayload struct {
	Email    *string `form:"email,omitempty" json:"email,omitempty" yaml:"email,omitempty" xml:"email,omitempty"`
	Password *string `form:"password,omitempty" json:"password,omitempty" yaml:"password,omitempty" xml:"password,omitempty"`
}

// Validate validates the loginPayload type instance.
func (ut *loginPayload) Validate() (err error) {
	if ut.Email == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "email"))
	}
	if ut.Password == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "password"))
	}
	if ut.Email != nil {
		if err2 := goa.ValidateFormat(goa.FormatEmail, *ut.Email); err2 != nil {
			err = goa.MergeErrors(err, goa.InvalidFormatError(`request.email`, *ut.Email, goa.FormatEmail, err2))
		}
	}
	if ut.Email != nil {
		if utf8.RuneCountInString(*ut.Email) < 6 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.email`, *ut.Email, utf8.RuneCountInString(*ut.Email), 6, true))
		}
	}
	if ut.Email != nil {
		if utf8.RuneCountInString(*ut.Email) > 400 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.email`, *ut.Email, utf8.RuneCountInString(*ut.Email), 400, false))
		}
	}
	if ut.Password != nil {
		if utf8.RuneCountInString(*ut.Password) < 8 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.password`, *ut.Password, utf8.RuneCountInString(*ut.Password), 8, true))
		}
	}
	if ut.Password != nil {
		if utf8.RuneCountInString(*ut.Password) > 100 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.password`, *ut.Password, utf8.RuneCountInString(*ut.Password), 100, false))
		}
	}
	return
}

// Publicize creates LoginPayload from loginPayload
func (ut *loginPayload) Publicize() *LoginPayload {
	var pub LoginPayload
	if ut.Email != nil {
		pub.Email = *ut.Email
	}
	if ut.Password != nil {
		pub.Password = *ut.Password
	}
	return &pub
}

// LoginPayload user type.
type LoginPayload struct {
	Email    string `form:"email" json:"email" yaml:"email" xml:"email"`
	Password string `form:"password" json:"password" yaml:"password" xml:"password"`
}

// Validate validates the LoginPayload type instance.
func (ut *LoginPayload) Validate() (err error) {
	if ut.Email == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "email"))
	}
	if ut.Password == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "password"))
	}
	if err2 := goa.ValidateFormat(goa.FormatEmail, ut.Email); err2 != nil {
		err = goa.MergeErrors(err, goa.InvalidFormatError(`type.email`, ut.Email, goa.FormatEmail, err2))
	}
	if utf8.RuneCountInString(ut.Email) < 6 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`type.email`, ut.Email, utf8.RuneCountInString(ut.Email), 6, true))
	}
	if utf8.RuneCountInString(ut.Email) > 400 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`type.email`, ut.Email, utf8.RuneCountInString(ut.Email), 400, false))
	}
	if utf8.RuneCountInString(ut.Password) < 8 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`type.password`, ut.Password, utf8.RuneCountInString(ut.Password), 8, true))
	}
	if utf8.RuneCountInString(ut.Password) > 100 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`type.password`, ut.Password, utf8.RuneCountInString(ut.Password), 100, false))
	}
	return
}

// registerPayload user type.
type registerPayload struct {
	CompanyAddress *string `form:"company_address,omitempty" json:"company_address,omitempty" yaml:"company_address,omitempty" xml:"company_address,omitempty"`
	CompanyName    *string `form:"company_name,omitempty" json:"company_name,omitempty" yaml:"company_name,omitempty" xml:"company_name,omitempty"`
	ContactName    *string `form:"contact_name,omitempty" json:"contact_name,omitempty" yaml:"contact_name,omitempty" xml:"contact_name,omitempty"`
	Email          *string `form:"email,omitempty" json:"email,omitempty" yaml:"email,omitempty" xml:"email,omitempty"`
	Password       *string `form:"password,omitempty" json:"password,omitempty" yaml:"password,omitempty" xml:"password,omitempty"`
	Phone          *string `form:"phone,omitempty" json:"phone,omitempty" yaml:"phone,omitempty" xml:"phone,omitempty"`
}

// Validate validates the registerPayload type instance.
func (ut *registerPayload) Validate() (err error) {
	if ut.Email == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "email"))
	}
	if ut.Password == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "password"))
	}
	if ut.ContactName == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "contact_name"))
	}
	if ut.Phone == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "phone"))
	}
	if ut.CompanyAddress == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "company_address"))
	}
	if ut.CompanyName == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "company_name"))
	}
	if ut.CompanyAddress != nil {
		if utf8.RuneCountInString(*ut.CompanyAddress) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.company_address`, *ut.CompanyAddress, utf8.RuneCountInString(*ut.CompanyAddress), 1, true))
		}
	}
	if ut.CompanyAddress != nil {
		if utf8.RuneCountInString(*ut.CompanyAddress) > 500 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.company_address`, *ut.CompanyAddress, utf8.RuneCountInString(*ut.CompanyAddress), 500, false))
		}
	}
	if ut.CompanyName != nil {
		if utf8.RuneCountInString(*ut.CompanyName) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.company_name`, *ut.CompanyName, utf8.RuneCountInString(*ut.CompanyName), 1, true))
		}
	}
	if ut.CompanyName != nil {
		if utf8.RuneCountInString(*ut.CompanyName) > 500 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.company_name`, *ut.CompanyName, utf8.RuneCountInString(*ut.CompanyName), 500, false))
		}
	}
	if ut.ContactName != nil {
		if utf8.RuneCountInString(*ut.ContactName) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.contact_name`, *ut.ContactName, utf8.RuneCountInString(*ut.ContactName), 1, true))
		}
	}
	if ut.ContactName != nil {
		if utf8.RuneCountInString(*ut.ContactName) > 200 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.contact_name`, *ut.ContactName, utf8.RuneCountInString(*ut.ContactName), 200, false))
		}
	}
	if ut.Email != nil {
		if err2 := goa.ValidateFormat(goa.FormatEmail, *ut.Email); err2 != nil {
			err = goa.MergeErrors(err, goa.InvalidFormatError(`request.email`, *ut.Email, goa.FormatEmail, err2))
		}
	}
	if ut.Email != nil {
		if utf8.RuneCountInString(*ut.Email) < 6 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.email`, *ut.Email, utf8.RuneCountInString(*ut.Email), 6, true))
		}
	}
	if ut.Email != nil {
		if utf8.RuneCountInString(*ut.Email) > 150 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.email`, *ut.Email, utf8.RuneCountInString(*ut.Email), 150, false))
		}
	}
	if ut.Password != nil {
		if utf8.RuneCountInString(*ut.Password) < 8 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.password`, *ut.Password, utf8.RuneCountInString(*ut.Password), 8, true))
		}
	}
	if ut.Password != nil {
		if utf8.RuneCountInString(*ut.Password) > 100 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.password`, *ut.Password, utf8.RuneCountInString(*ut.Password), 100, false))
		}
	}
	if ut.Phone != nil {
		if utf8.RuneCountInString(*ut.Phone) > 50 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.phone`, *ut.Phone, utf8.RuneCountInString(*ut.Phone), 50, false))
		}
	}
	return
}

// Publicize creates RegisterPayload from registerPayload
func (ut *registerPayload) Publicize() *RegisterPayload {
	var pub RegisterPayload
	if ut.CompanyAddress != nil {
		pub.CompanyAddress = *ut.CompanyAddress
	}
	if ut.CompanyName != nil {
		pub.CompanyName = *ut.CompanyName
	}
	if ut.ContactName != nil {
		pub.ContactName = *ut.ContactName
	}
	if ut.Email != nil {
		pub.Email = *ut.Email
	}
	if ut.Password != nil {
		pub.Password = *ut.Password
	}
	if ut.Phone != nil {
		pub.Phone = *ut.Phone
	}
	return &pub
}

// RegisterPayload user type.
type RegisterPayload struct {
	CompanyAddress string `form:"company_address" json:"company_address" yaml:"company_address" xml:"company_address"`
	CompanyName    string `form:"company_name" json:"company_name" yaml:"company_name" xml:"company_name"`
	ContactName    string `form:"contact_name" json:"contact_name" yaml:"contact_name" xml:"contact_name"`
	Email          string `form:"email" json:"email" yaml:"email" xml:"email"`
	Password       string `form:"password" json:"password" yaml:"password" xml:"password"`
	Phone          string `form:"phone" json:"phone" yaml:"phone" xml:"phone"`
}

// Validate validates the RegisterPayload type instance.
func (ut *RegisterPayload) Validate() (err error) {
	if ut.Email == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "email"))
	}
	if ut.Password == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "password"))
	}
	if ut.ContactName == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "contact_name"))
	}
	if ut.Phone == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "phone"))
	}
	if ut.CompanyAddress == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "company_address"))
	}
	if ut.CompanyName == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "company_name"))
	}
	if utf8.RuneCountInString(ut.CompanyAddress) < 1 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`type.company_address`, ut.CompanyAddress, utf8.RuneCountInString(ut.CompanyAddress), 1, true))
	}
	if utf8.RuneCountInString(ut.CompanyAddress) > 500 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`type.company_address`, ut.CompanyAddress, utf8.RuneCountInString(ut.CompanyAddress), 500, false))
	}
	if utf8.RuneCountInString(ut.CompanyName) < 1 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`type.company_name`, ut.CompanyName, utf8.RuneCountInString(ut.CompanyName), 1, true))
	}
	if utf8.RuneCountInString(ut.CompanyName) > 500 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`type.company_name`, ut.CompanyName, utf8.RuneCountInString(ut.CompanyName), 500, false))
	}
	if utf8.RuneCountInString(ut.ContactName) < 1 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`type.contact_name`, ut.ContactName, utf8.RuneCountInString(ut.ContactName), 1, true))
	}
	if utf8.RuneCountInString(ut.ContactName) > 200 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`type.contact_name`, ut.ContactName, utf8.RuneCountInString(ut.ContactName), 200, false))
	}
	if err2 := goa.ValidateFormat(goa.FormatEmail, ut.Email); err2 != nil {
		err = goa.MergeErrors(err, goa.InvalidFormatError(`type.email`, ut.Email, goa.FormatEmail, err2))
	}
	if utf8.RuneCountInString(ut.Email) < 6 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`type.email`, ut.Email, utf8.RuneCountInString(ut.Email), 6, true))
	}
	if utf8.RuneCountInString(ut.Email) > 150 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`type.email`, ut.Email, utf8.RuneCountInString(ut.Email), 150, false))
	}
	if utf8.RuneCountInString(ut.Password) < 8 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`type.password`, ut.Password, utf8.RuneCountInString(ut.Password), 8, true))
	}
	if utf8.RuneCountInString(ut.Password) > 100 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`type.password`, ut.Password, utf8.RuneCountInString(ut.Password), 100, false))
	}
	if utf8.RuneCountInString(ut.Phone) > 50 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`type.phone`, ut.Phone, utf8.RuneCountInString(ut.Phone), 50, false))
	}
	return
}
