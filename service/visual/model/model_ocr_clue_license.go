package model

import (
	"github.com/volcengine/volc-sdk-golang/base"
)

type LicenseBasePos struct {
	X      interface{} `json:"x"`
	Y      interface{} `json:"y"`
	Width  interface{} `json:"width"`
	Height interface{} `json:"height"`
}

type LicenseBase struct {
	Val string          `json:"val"`
	Pos *LicenseBasePos `json:"pos"`
}

type LicenseMain struct {
	HasHead           bool         `json:"has_head"`
	HasTail           bool         `json:"has_tail"`
	IsCopy            bool         `json:"is_copy"`
	CreditCode        *LicenseBase `json:"credit_code"`
	Name              *LicenseBase `json:"name"`
	Type              *LicenseBase `json:"type"`
	Address           *LicenseBase `json:"address"`
	LegalPerson       *LicenseBase `json:"legal_person"`
	RegisteredCapital *LicenseBase `json:"registered_capital"`
	RegistrationDate  *LicenseBase `json:"registration_date"`
	ValidPeriod       *LicenseBase `json:"valid_period"`
	BusinessScope     *LicenseBase `json:"business_scope"`
	ApprovalDate      *LicenseBase `json:"approval_date"`
}

type LicensePos struct {
	X interface{} `json:"x"`
	Y interface{} `json:"y"`
}

type OcrClueLicenseData struct {
	LicenseMain LicenseMain   `json:"license_main"`
	LicensePos  []*LicensePos `json:"license_pos"`
}

type OcrClueLicenseResult struct {
	ResponseMetadata *base.ResponseMetadata `json:",omitempty"`
	RequestId        string                 `json:"request_id"`
	TimeElapsed      string                 `json:"time_elapsed"`
	Code             int                    `json:"code"`
	Message          string                 `json:"message"`
	Status           int                    `json:"status"`
	Data             *OcrClueLicenseData    `json:"data"`
}
