package service

import (
	"Lynks/stats/internal/model"
	"Lynks/stats/internal/repository"
	
)

type StatsService struct {
	repo *repository.StatsRepository
}

func NewStatService(repo *repository.StatsRepository) *StatsService {
	return &StatsService{
		repo: repo,
	}
}

func (s *StatsService) CreateStat(url string) {
	stat, isExist, _ := s.repo.GetStatByUrl(url)
	if !isExist {
		s.repo.CreateStat(&model.Stats{
			Url: url,
			Clicks: 1,
			Average_length: len(url) / 2,
		})
	} else {
		stat.Clicks += 1
		s.repo.UpdateStat(stat)
	}
}
