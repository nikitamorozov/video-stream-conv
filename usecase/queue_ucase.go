package usecase

import (
	"github.com/nikitamorozov/video-stream-conv/models"
	"github.com/nikitamorozov/video-stream-conv/repository"
)

type QueueUseCases interface {
	Queue(name string, job models.Job) error
}

type queueUseCases struct {
	repo repository.QueueRepository
}

func (uc queueUseCases) Queue(name string, job models.Job) error {
	err := uc.repo.Queue(name, job)
	if err != nil {
		return err
	}

	return nil
}

func NewQueueUseCases(repo repository.QueueRepository) QueueUseCases {
	return &queueUseCases{
		repo: repo,
	}
}
