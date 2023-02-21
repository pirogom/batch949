package batch949

import (
	"fmt"
	"os/exec"
	"syscall"

	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/transform"
)

var (
	fixedCodePage string
)

func init() {
	fixedCodePage = "949" // chcp 949 (korean)
}

/**
*	SetCodePage
**/
func SetCodePage(cp string) {
	fixedCodePage = cp
}

/**
*	buildCommand
**/
func buildCommand(execCmd string, args ...string) (string, error) {
	batchString := fmt.Sprintf("chcp %s&%s", fixedCodePage, execCmd)
	for _, v := range args {
		batchString += fmt.Sprintf(" \"%s\"", v)
	}
	return batchString, nil
	// utf8 to cp949
	// str949, _, err := transform.String(korean.EUCKR.NewEncoder(), batchString)
	// return str949, err
}

/**
*	Run
**/
func Run(execCmd string, args ...string) (string, error) {
	cmdBuf, cmdErr := buildCommand(execCmd, args...)
	if cmdErr != nil {
		return "", cmdErr
	}

	cmd := exec.Command("cmd.exe", "/C", cmdBuf)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	buf, _ := cmd.CombinedOutput()

	ub, _, ubErr := transform.Bytes(korean.EUCKR.NewDecoder(), buf)

	return string(ub), ubErr
}
