package service

import (
	"Stats/internal/client"
	"Stats/internal/model"
	"Stats/internal/repository"
)

type StatsService struct {
	repo *repository.StatsRepository
}

func NewStatService(repo *repository.StatsRepository) *StatsService {
	return &StatsService{
		repo: repo,
	}
}

func (s *StatsService) CreateStat(mp map[string]string) {
	stat, isExist, _ := s.repo.GetStatByUrl(mp["url"])
	if !isExist {
		s.repo.CreateStat(&model.Stats{
			Url:            mp["url"],
			Clicks:         1,
			Average_length: len(mp["url"]) / 2,
		})
	} else {
		stat.Clicks += 1
		if stat.Clicks > 5 {
			client.Client(stat, mp["hash"])
		} else {
			s.repo.UpdateStat(stat)
		}
	}
}