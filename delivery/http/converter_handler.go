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
}

func (handler *ConverterHandler) Convert(c echo.Context) error {
	destFileName := tools.HashGenerator()
	sourceFileName := "source-" + destFileName

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
		Dest: destFileName,
	}

	go handler.converterUseCases.ConvertVideo(sourceFileName, destFileName)

	return c.JSON(http.StatusOK, resp)
}

func NewConverterHttpHandler(e *echo.Echo, converterUseCases usecase.ConverterUseCases) {
	handler := ConverterHandler{
		converterUseCases: converterUseCases,
	}

	e.GET(common.API_VER_1_0+"/convert", handler.Convert)
}
