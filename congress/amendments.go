package congress

import "github.com/mylez/go-sunlight/internal"

type AmendmentsResult struct {
	Result
	Results []*struct {
		AmendmentId       string `json:"amendment_id"`
		AmendmentType     string `json:"amendment_type"`
		AmendsAmendmentId string `json:"amends_amendment_id"`
		AmendsBillId      string `json:"amends_bill_id"`
		Chamber           string `json:"chamber"`
		Congress          int    `json:"congress"`
		Description       string `json:"description"`
		IntroducedOn      string `json:"introduced_on"`
		LastActionAt      string `json:"last_action_at"`
		Number            int    `json:"number"`
		Purpose           string `json:"purpose"`
		SponsorId         string `json:"sponsor_id"`
		SponsorType       string `json:"sponsor_type"`
	} `json:"results"`
}

func GetAmendments(criteria map[string]string) (*AmendmentsResult, error) {
	v := AmendmentsResult{}
	err := internal.GetURL(&v, congressRoot, criteria, "amendments")
	if err != nil {
		return nil, err
	}
	return &v, nil
}
