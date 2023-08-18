package usecase

import (
	"go-clean-template/internal/entity"
)

type initDBUsecase struct {
	sourceRepository entity.SourceRepository
}

func NewInitdbUsecase(sourceRepository entity.SourceRepository) entity.InitDBUsecase {
	return &initDBUsecase{
		sourceRepository: sourceRepository,
	}
}

func (i initDBUsecase) InitTableData() error {
	return i.sourceRepository.InitTableData()
}
