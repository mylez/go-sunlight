package congress

import "github.com/mylez/go-sunlight/internal"

type VotesResult struct {
	Result
	Results []*struct {
		AmendmentId string `json:"amendment_id"`
		BillId      string `json:"bill_id"`
		Chamber     string `json:"chamber"`
		Congress    int    `json:"congress"`
		Number      int    `json:"number"`
		Question    string `json:"question"`
		Required    string `json:"required"`
		Result      string `json:"result"`
		RollId      string `json:"roll_id"`
		RollType    string `json:"roll_type"`
		Source      string `json:"source"`
		Url         string `json:"url"`
		VoteType    string `json:"vote_type"`
		VotedAt     string `json:"voted_at"`
		Year        int    `json:"year"`
	}
}

func GetVotes(criteria map[string]string) (*VotesResult, error) {
	v := VotesResult{}
	err := internal.GetURL(&v, congressRoot, criteria, "votes")
	if err != nil {
		return nil, err
	}
	return &v, nil
}
