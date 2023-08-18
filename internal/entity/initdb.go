package entity

type SourceRepository interface {
	InitTableData() error
}

type InitDBUsecase interface {
	InitTableData() error
}
