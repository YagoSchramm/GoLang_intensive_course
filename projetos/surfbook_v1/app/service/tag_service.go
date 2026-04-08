package service

import (
	"context"
	"time"

	"github.com/YagoSchramm/intensivo-surfbook_v1/model"
	"github.com/YagoSchramm/intensivo-surfbook_v1/repository"
	"github.com/google/uuid"
)

type TagService struct {
	repo *repository.TagRepository
}

func NewTagService(r *repository.TagRepository) *TagService {
	return &TagService{repo: r}
}

func (srv *TagService) Create(ctx context.Context, input model.CreateTagDTO) (*model.TagEntity, error) {
	id := uuid.New()
	now := time.Now()
	tag := model.TagEntity{
		TagID:     id,
		Name:      input.Name,
		Color:     input.Color,
		UserID:    input.UserID,
		DeletedAt: nil,
		CreatedAt: now,
		UpdatedAt: now,
	}
	err := srv.repo.Create(ctx, &tag)
	if err != nil {
		return nil, err
	}
	return &tag, nil
}

func (srv *TagService) ListFromUser(ctx context.Context, input model.ListTagsFromUserDTO) ([]*model.TagEntity, error) {
	return srv.repo.ListTags(ctx, input)
}

func (srv *TagService) FindByUserTagId(ctx context.Context, input model.FindTagFromUserDTO) (*model.TagEntity, error) {
	return srv.repo.FindByUserIdTagId(ctx, input)
}

func (srv *TagService) SoftDelete(ctx context.Context, input model.DeleteTagDTO) error {
	return srv.repo.Delete(ctx, input)
}

func (srv *TagService) Update(ctx context.Context, input model.UpdateTagDTO) (*model.TagEntity, error) {
	err := srv.repo.Update(ctx, input)
	if err != nil {
		return nil, err
	}
	find := model.FindTagFromUserDTO{
		UserID: input.UserID,
		TagID:  input.TagID,
	}
	tag, err := srv.repo.FindByUserIdTagId(ctx, find)
	if err != nil {
		return nil, err
	}
	return tag, nil
}
