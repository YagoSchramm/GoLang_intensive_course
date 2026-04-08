package service

import "github.com/YagoSchramm/intensivo-surfbook_v1/repository"

type MetaContentService struct {
	repo *repository.MetaContentRepository
}

func NewMetaContentService(r *repository.MetaContentRepository) *MetaContentService {
	return &MetaContentService{repo: r}
}

func
