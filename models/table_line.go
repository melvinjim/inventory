package models

import (
	"time"

	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
	"github.com/gofrs/uuid"
)

type Lines struct {
	ID                 uuid.UUID `form:"id" db:"id"`
	PhoneLine          string    `form:"phone_line" db:"phone_line"`
	Carrier            string    `form:"carrier" db:"carrier"`
	AssociatedTo       string    `form:"associated_to" db:"associated_to"`
	AssociatedToDevice string    `form:"associated_to_device" db:"associated_to_device"`
	EndContractDate    string    `form:"end_contract_date" db:"end_contract_date"`
	UpgradeEligibility string    `form:"upgrade_eligibility" db:"upgrade_eligibility"`
	IMEI               string    `form:"imei" db:"imei"`
	CreatedAt          time.Time `form:"created_at" db:"created_at"`
	UpdatedAt          time.Time `form:"updated_at" db:"updated_at"`
}

func (l *Lines) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(&validators.StringIsPresent{
		Field:   l.PhoneLine,
		Name:    "number",
		Message: "This field is required",
	}), nil
}

func (l *Lines) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {

	return validate.NewErrors(), nil
}
