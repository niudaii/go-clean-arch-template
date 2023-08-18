package entity

import "github.com/lib/pq"

type Result struct {
	BasicModel
	UUID     string         `json:"uuid"`
	TaskUUID string         `json:"taskUUID"`
	Inputs   pq.StringArray `json:"inputs" gorm:"type:text[]"`
	Result   int            `json:"result"`
}

func (Result) TableName() string {
	return "result"
}

type ResultRepository interface {
	Create(result Result) (err error)
	FindList(f *ResultFilter) (list []Result, total int64, err error)
}

type ResultFilter struct {
	PageInfo
	Result
}

type ResultUsecase interface {
	FindList(searchResult *SearchResult) (list []Result, total int64, err error)
}

type SearchResult struct {
	PageInfo
	Result
}
