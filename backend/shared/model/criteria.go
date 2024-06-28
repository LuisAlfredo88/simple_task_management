package model

type CriteriaFilter struct {
	Filters map[string]interface{} `json:"filters"`
	Limit   int32                  `json:"limit"`
	Skip    int32                  `json:"skip"`
}
