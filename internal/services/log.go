package services

import (
	"context"

	"github.com/flytrap/go-template/internal/models"
	"github.com/flytrap/go-template/internal/repositories"
	"github.com/mitchellh/mapstructure"
)

type LogService interface {
	List(q string, page int64, size int64, ordering string, data interface{}) (n int64, err error)

	Update(ctx context.Context, info map[string]interface{}) error
	Create(ctx context.Context, info map[string]interface{}) error
	Delete(ctx context.Context, ids []int64) (err error)
}

func NewLogService(repo repositories.LogRepository) LogService {
	return &LogServiceImp{repo: repo}
}

type LogServiceImp struct {
	repo repositories.LogRepository
}

func (s *LogServiceImp) List(q string, page int64, size int64, ordering string, data interface{}) (n int64, err error) {
	n, err = s.repo.List(q, (page-1)*size, size, ordering, data)
	return
}

func (s *LogServiceImp) Create(ctx context.Context, info map[string]interface{}) error {
	data := models.Log{}
	err := mapstructure.Decode(info, &data)
	if err != nil {
		return err
	}
	err = s.repo.Create(&data)
	return err
}

func (s *LogServiceImp) Update(ctx context.Context, info map[string]interface{}) error {
	id := info["id"].(int64)
	return s.repo.Update(id, info)
}

func (s *LogServiceImp) Delete(ctx context.Context, ids []int64) (err error) {
	return s.repo.Delete(ids)
}
