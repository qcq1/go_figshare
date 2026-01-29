package go_figshare

type TotalArticle struct {
	Views     int64 `json:"views"`
	Downloads int64 `json:"downloads"`
	Shares    int64 `json:"shares"`
	Cites     int64 `json:"cites"`
}
