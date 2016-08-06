package openstates

import (
	"github.com/mylez/go-sunlight/internal"
)

type Bill struct {
	Timestamps
	Sources

	BillId     string   `json:"bill_id"`
	Id         string   `json:"id"`
	AllIds     []string `json:"all_ids"`
	Companions []string `json:"companions"`

	AlternateTitles []string `json:"alternate_titles"`
	Title           string   `json:"title"`

	Country string `json:"country"`
	Level   string `json:"level"`
	Session string `json:"session"`
	State   string `json:"state"`

	Actions []struct {
		Action string   `json:"action"`
		Actor  string   `json:"actor"`
		Types  []string `json:"type"`
	} `json:"actions"`
	Sponsors []struct {
		LegId string `json:"leg_id"`
		Name  string `json:"name"`
		Type  string `json:"type"`
	} `json:"sponsors"`
	Chamber string `json:"chamber"`

	Subjects []string `json:"subjects"`
	Types    []string `json:"type"`
}

func GetBill(state, session, billId string) (*Bill, error) {
	t := &Bill{}
	err := internal.GetURL(t, openstatesRoot, map[string]string{}, "bills", state, session, billId)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func GetBills(criteria map[string]string) (*[]Bill, error) {
	t := &[]Bill{}
	err := internal.GetURL(t, openstatesRoot, criteria, "bills")
	if err != nil {
		return nil, err
	}
	return t, nil
}

func GetBillById(bigId string) (*Bill, error) {
	t := &Bill{}
	err := internal.GetURL(t, openstatesRoot, map[string]string{}, "bills", bigId)
	if err != nil {
		return nil, err
	}
	return t, nil
}
