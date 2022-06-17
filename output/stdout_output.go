package output

import (
	"os"

	"github.com/chenjiayao/godeng/inter"
)

type StdoutOutput struct {
}

var _ inter.Output = &StdoutOutput{}

func (o *StdoutOutput) Output(content []byte) {
	content = append(content, '\n')
	os.Stdout.Write(content)
}

func MakeStdoutOutput() *StdoutOutput {
	return &StdoutOutput{}
}
