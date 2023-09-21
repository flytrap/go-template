package repositories

import (
	"github.com/flytrap/go-template/internal/models"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type LogRepository interface {
	Get(id int64) (*models.Log, error)
	List(q string, offset int64, limit int64, ordering string, data interface{}) (int64, error)

	Create(*models.Log) error
	Update(id int64, info map[string]interface{}) (err error)
	Delete(ids []int64) (err error)
}

func NewLoggingRepository(db *gorm.DB) LogRepository {
	return &LogRepositoryImp{Db: db}
}

type LogRepositoryImp struct {
	Db *gorm.DB
}

func (s *LogRepositoryImp) Get(id int64) (data *models.Log, err error) {
	if err = s.Db.First(&data, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (s *LogRepositoryImp) List(q string, offset int64, limit int64, ordering string, data interface{}) (n int64, err error) {
	query := s.Db
	if len(q) > 0 {
		query = query.Where("content like ?", q+"%")
	}
	if len(ordering) == 0 {
		ordering = "id desc"
	}

	if err := query.Offset(int(offset)).Limit(int(limit)).Order(ordering).Find(data).Error; err != nil {
		return 0, err
	}
	query.Count(&n)
	return n, nil
}

func (s *LogRepositoryImp) Create(data *models.Log) (err error) {
	result := s.Db.Create(data)
	return errors.WithStack(result.Error)
}

func (s *LogRepositoryImp) Delete(ids []int64) (err error) {
	result := s.Db.Where("id in ?", ids).Delete(&models.Log{})
	return errors.WithStack(result.Error)
}

func (s *LogRepositoryImp) Update(id int64, info map[string]interface{}) (err error) {
	result := s.Db.Model(&models.Log{}).Where("id = ?", id).Updates(info)
	return errors.WithStack(result.Error)
}
