package usecase

import (
	"bytes"
	"fmt"
	"github.com/labstack/gommon/log"
	"os/exec"
	"strings"
)

type ConverterUseCases interface {
	ConvertVideo(source string, dest string)
}

type converterUseCases struct {
}

func (uc converterUseCases) ConvertVideo(source string, dest string) {
	cmd := exec.Command("bash", fmt.Sprintf("convert-hls.sh %s %s", source, dest))
	cmd.Stdin = strings.NewReader("")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Output: %s", out.String())
}

func NewConverterUseCases() ConverterUseCases {
	return &converterUseCases{}
}
