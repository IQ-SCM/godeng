package output

import (
	"os"

	"github.com/chenjiayao/godeng/inter"
)

type FileOutput struct {
	filepath string
}

var _ inter.Output = &FileOutput{}

func (o *FileOutput) Output(content []byte) {
	file, err := os.OpenFile(o.filepath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.Write(append(content, '\n'))
}

func MakeFileOutput(filepath string) *FileOutput {
	return &FileOutput{
		filepath: filepath,
	}
}
