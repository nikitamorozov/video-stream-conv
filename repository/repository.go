package repository

import (
	"github.com/nikitamorozov/video-stream-conv/models"
)

type QueueRepository interface {
	Queue(name string, job models.Job) error
}
