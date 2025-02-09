package services

import (
	"context"
	"field-service/common/gcs"
	"field-service/common/util"
	"field-service/domain/dto"
	"field-service/repositories"
)

type FieldService struct {
	repository repositories.IRepositoryRegistry
	gcs        gcs.IGCSClient
}

type IFieldService interface {
	GetAllWithPagination(context.Context, *dto.FieldRequestParam) (*util.PaginationResult, error)
	GetAllWithoutPagination(context.Context) ([]dto.FieldResponse, error)
	GetByUUID(context.Context, string) (*dto.FieldResponse, error)
	Create(context.Context, *dto.FieldRequest) (*dto.FieldResponse, error)
	Update(context.Context, string, *dto.UpdateFieldRequest) (*dto.FieldResponse, error)
	Delete(context.Context, string) error
}

func NewFieldService(repository repositories.IRepositoryRegistry, gcs gcs.IGCSClient) IFieldService {
	return &FieldService{
		repository: repository,
		gcs:        gcs,
	}
}

// Create implements IFieldService.
func (f *FieldService) Create(context.Context, *dto.FieldRequest) (*dto.FieldResponse, error) {
	panic("unimplemented")
}

// Delete implements IFieldService.
func (f *FieldService) Delete(context.Context, string) error {
	panic("unimplemented")
}

// GetAllWithPagination implements IFieldService.
func (f *FieldService) GetAllWithPagination(ctx context.Context, req *dto.FieldRequestParam) (*util.PaginationResult, error) {
	fields, total, err := f.repository.GetField().FindAllWithPagination(ctx, req)
	if err != nil {
		return nil, err
	}

	fieldResults := make([]dto.FieldResponse, 0, len(fields))
	for _, field := range fields {
		fieldResults = append(fieldResults, dto.FieldResponse{
			UUID:         field.UUID,
			Code:         field.Code,
			Name:         field.Name,
			Images:       field.Images,
			PricePerHour: field.PricePerHour,
			CreatedAt:    field.CreatedAt,
			UpdatedAt:    field.UpdatedAt,
		})
	}
	pagination := &util.PaginationParam{
		Page:  req.Page,
		Limit: req.Limit,
		Count: total,
		Data:  fieldResults,
	}
	result := util.GeneratePagination(*pagination)
	return &result, nil

}

// GetAllWithoutPagination implements IFieldService.
func (f *FieldService) GetAllWithoutPagination(ctx context.Context) ([]dto.FieldResponse, error) {
	fields, err := f.repository.GetField().FindAllWithoutPagination(ctx)
	if err != nil {
		return nil, err
	}
	fieldResults := make([]dto.FieldResponse, 0, len(fields))
	for _, field := range fields {
		fieldResults = append(fieldResults, dto.FieldResponse{
			UUID:         field.UUID,
			Name:         field.Name,
			Images:       field.Images,
			PricePerHour: field.PricePerHour,
		})
	}
	return fieldResults, nil
}

// GetByUUID implements IFieldService.
func (f *FieldService) GetByUUID(ctx context.Context, uuid string) (*dto.FieldResponse, error) {
	field, err := f.repository.GetField().FindByUUID(ctx, uuid)
	if err != nil {
		return nil, err
	}

	pricePerHour := float64(field.PricePerHour)
	fieldResult := dto.FieldResponse{
		UUID:         field.UUID,
		Code:         field.Code,
		Name:         field.Name,
		PricePerHour: util.RupiahFormat(&pricePerHour),
		Images:       field.Images,
		CreatedAt:    field.CreatedAt,
		UpdatedAt:    field.UpdatedAt,
	}

	return &fieldResult, nil
}

// Update implements IFieldService.
func (f *FieldService) Update(context.Context, string, *dto.UpdateFieldRequest) (*dto.FieldResponse, error) {
	panic("unimplemented")
}
