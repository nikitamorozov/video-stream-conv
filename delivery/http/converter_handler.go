package http

import (
	"github.com/labstack/echo"
	"github.com/nikitamorozov/video-stream-conv/common"
	"github.com/nikitamorozov/video-stream-conv/models/response"
	"github.com/nikitamorozov/video-stream-conv/tools"
	"github.com/nikitamorozov/video-stream-conv/usecase"
	"io"
	"net/http"
	"os"
)

type ConverterHandler struct {
	converterUseCases usecase.ConverterUseCases
	domain            string
	filePostfix       string
}

func (handler *ConverterHandler) Convert(c echo.Context) error {
	hash := tools.HashGenerator()
	destFileName := "video/" + hash
	sourceFileName := "video/source_" + hash + ".mp4"

	// Source
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(sourceFileName)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	resp := response.ConvertResponse{
		Dest: handler.domain + hash + handler.filePostfix,
	}

	go handler.converterUseCases.ConvertVideo(sourceFileName, destFileName)

	return c.JSON(http.StatusOK, resp)
}

func NewConverterHttpHandler(e *echo.Echo, converterUseCases usecase.ConverterUseCases, domain string, filePostfix string) {
	handler := ConverterHandler{
		converterUseCases: converterUseCases,
		domain:            domain,
		filePostfix:       filePostfix,
	}

	e.GET(common.API_VER_1_0+"/convert", handler.Convert)
}
