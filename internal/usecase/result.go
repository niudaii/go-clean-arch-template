package usecase

import "go-clean-template/internal/entity"

type resultUsecase struct {
	resultRepository entity.ResultRepository
}

func NewResultUsecase(resultRepository entity.ResultRepository) entity.ResultUsecase {
	return &resultUsecase{
		resultRepository: resultRepository,
	}
}

func (a resultUsecase) FindList(searchResult *entity.SearchResult) (list []entity.Result, total int64, err error) {
	f := &entity.ResultFilter{}
	f.PageInfo = searchResult.PageInfo
	f.TaskUUID = searchResult.TaskUUID
	return a.resultRepository.FindList(f)
}
