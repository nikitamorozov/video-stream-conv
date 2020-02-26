package usecase

import (
	"os/exec"
)

type ConverterUseCases interface {
	ConvertVideo(source string, dest string)
}

type converterUseCases struct {
}

func (uc converterUseCases) ConvertVideo(source string, dest string) {
	exec.Command("/bin/bash", "converter.sh", source, dest).Run()
}

func NewConverterUseCases() ConverterUseCases {
	return &converterUseCases{}
}
