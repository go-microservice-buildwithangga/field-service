package repositories

import (
	"context"
	error2 "field-service/common/error"
	errConst "field-service/constants/error"
	errConstField "field-service/constants/error/field"
	"field-service/domain/dto"
	"field-service/domain/models"
	"fmt"

	"gorm.io/gorm"
)

type FieldRepository struct {
	db *gorm.DB
}

type IFieldRepository interface {
	FindAllWithPagination(context.Context, *dto.FieldRequestParam) ([]models.Field, int64, error)
	FindAllWithoutPagination(context.Context) ([]models.Field, error)
	FindByUUID(context.Context, string) (*models.Field, error)
	Create(context.Context, string, *models.Field) error
}

func NewFieldRepository(db *gorm.DB) IFieldRepository {
	return &FieldRepository{
		db: db,
	}
}

// Create implements IFieldRepository.
func (f *FieldRepository) Create(context.Context, string, *models.Field) error {
	panic("unimplemented")
}

// FindAllWithPagination implements IFieldRepository.
func (f *FieldRepository) FindAllWithPagination(ctx context.Context, param *dto.FieldRequestParam) ([]models.Field, int64, error) {
	var (
		fields []models.Field
		sort   string
	)
	if param.SortColumn != nil {
		sort = fmt.Sprintf("%s %s", *param.SortColumn, *param.SortOrder)
	} else {
		sort = "created_at desc"
	}
	limit := param.Limit
	offset := (param.Page - 1) * limit
	err := f.db.WithContext(ctx).Limit(limit).Offset(offset).Order(sort).Find(&fields).Error
	if err != nil {
		return nil, 0, error2.WrapError(errConst.ErrSQLError)
	}
	var total int64
	err = f.db.WithContext(ctx).Model(&fields).Count(&total).Error
	if err != nil {
		return nil, 0, error2.WrapError(errConst.ErrSQLError)
	}
	return fields, total, nil

}

// FindAllWithoutPagination implements IFieldRepository.
func (f *FieldRepository) FindAllWithoutPagination(ctx context.Context) ([]models.Field, error) {
	var fields []models.Field
	err := f.db.WithContext(ctx).Find(&fields).Error
	if err != nil {
		return nil, error2.WrapError(errConst.ErrSQLError)
	}
	return fields, nil
}

// FindByUUID implements IFieldRepository.
func (f *FieldRepository) FindByUUID(ctx context.Context, UUID string) (*models.Field, error) {
	var field models.Field
	err := f.db.WithContext(ctx).Where("uuid = ?", UUID).First(&field).Error
	if err != nil {
		if err.Error() == "record not found" {
			return nil, error2.WrapError(errConstField.ErrFieldNotFound)
		}
		return nil, error2.WrapError(errConst.ErrSQLError)
	}
	return &field, nil
}
