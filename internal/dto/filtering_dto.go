package dto

type Filtering struct {
	PerPage   int `json:"perpage"`
	TotalPage int `json:"total_page"`
	TotalData int `json:"total_data"`
}
