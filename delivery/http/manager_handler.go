package http

import (
	"github.com/labstack/echo"
	"github.com/nikitamorozov/video-stream-conv/common"
	"github.com/nikitamorozov/video-stream-conv/models"
	"github.com/nikitamorozov/video-stream-conv/models/response"
	"github.com/nikitamorozov/video-stream-conv/tools"
	"github.com/nikitamorozov/video-stream-conv/usecase"
	"io"
	"net/http"
	"os"
)

type ManagerHandler struct {
	queue         string
	queueUseCases usecase.QueueUseCases
	domain        string
	filePostfix   string
}

func (handler *ManagerHandler) Convert(c echo.Context) error {
	hash := tools.HashGenerator()
	destFileName := "video/" + hash
	sourceFileName := "video/source_" + hash + ".mp4"

	// Source
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(sourceFileName)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	resp := response.ConvertResponse{
		Dest: handler.domain + hash + handler.filePostfix,
	}

	job := &models.Job{
		Source: sourceFileName,
		Dest:   destFileName,
	}

	go handler.queueUseCases.Queue(handler.queue, *job)

	return c.JSON(http.StatusOK, resp)
}

func NewManagerHttpHandler(e *echo.Echo, queue string, queueUseCases usecase.QueueUseCases, domain string, filePostfix string) {
	handler := ManagerHandler{
		queue:         queue,
		queueUseCases: queueUseCases,
		domain:        domain,
		filePostfix:   filePostfix,
	}

	e.POST(common.API_VER_1_0+"/convert", handler.Convert)
}
