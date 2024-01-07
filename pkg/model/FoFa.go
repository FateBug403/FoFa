package model

type FoFa struct {
	Error   bool       `json:"error"`
	Size    int64      `json:"size"`
	Page    int64      `json:"page"`
	Model   string     `json:"demo"`
	Query   string     `json:"query"`
	Results [][]string `json:"results"`
}
