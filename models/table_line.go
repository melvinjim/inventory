package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type Lines struct {
	ID                 uuid.UUID `json:"id" db:"id"`
	PhoneLine          string    `json:"phone_line" db:"phone_line"`
	Carrier            string    `json:"carrier" db:"carrier"`
	AssociatedTo       string    `json:"associated_to" db:"associated_to"`
	AssociatedToDevice string    `json:"associated_to_device" db:"associated_to_device"`
	EndContractDate    string    `json:"end_contract_date" db:"end_contract_date"`
	UpgradeEligibility string    `json:"upgrade_eligibility" db:"upgrade_eligibility"`
	IMEI               string    `json:"imei" db:"imei"`
	CreatedAt          time.Time `form:"created_at" db:"created_at"`
	UpdatedAt          time.Time `form:"updated_at" db:"updated_at"`
}
