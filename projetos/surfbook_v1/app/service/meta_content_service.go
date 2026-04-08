package service

import (
	"context"
	"time"

	"github.com/YagoSchramm/intensivo-surfbook_v1/model"
	"github.com/YagoSchramm/intensivo-surfbook_v1/repository"
	"github.com/google/uuid"
)

type MetaContentService struct {
	repo *repository.MetaContentRepository
}

func NewMetaContentService(r *repository.MetaContentRepository) *MetaContentService {
	return &MetaContentService{repo: r}
}

func (srv *MetaContentService) Create(ctx context.Context, input model.CreateMetaContentDTO) (*model.MetaContentEntity, error) {
	id := uuid.New()
	now := time.Now()
	mc := model.MetaContentEntity{
		MetaContentID: id,
		NotebookID:    input.NotebookID,
		UserID:        input.UserID,
		Icon:          input.Icon,
		Name:          input.Name,
		CreatedAt:     now,
		UpdatedAt:     now,
		DeletedAt:     nil,
	}
	err := srv.repo.Create(ctx, mc)
	if err != nil {
		return nil, err
	}
	return &mc, nil
}
func (srv *MetaContentService) ListMetaContentById(ctx context.Context, input model.ListMetaContentByIdDTO) (*model.MetaContentEntity, error) {
	return srv.repo.ListMetaContentById(ctx, input)
}
func (srv *MetaContentService) ListMetaContentFromUser(ctx context.Context, input model.ListMetaContentFromUserDTO) (*[]model.MetaContentEntity, error) {
	return srv.repo.ListMetaContentByUserId(ctx, input)
}
func (srv *MetaContentService) SoftDelete(ctx context.Context, input model.DeleteMetaContentDTO) error {
	return srv.repo.Delete(ctx, input)
}
func (srv *MetaContentService) Update(ctx context.Context, input model.UpdateMetaContentDTO) (*model.MetaContentEntity, error) {
	err := srv.repo.Update(ctx, input)
	if err != nil {
		return nil, err
	}
	in := model.ListMetaContentByIdDTO{
		MetaContentID: input.MetaContentID,
		UserID:        input.UserID,
	}
	mc, err := srv.repo.ListMetaContentById(ctx, in)
	if err != nil {
		return nil, err
	}
	return mc, err
}
