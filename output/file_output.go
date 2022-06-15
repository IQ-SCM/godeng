package output

import "github.com/chenjiayao/godeng/inter"

type FileOutput struct {
	filepath string
}

var _ inter.Output = &FileOutput{}

func (o *FileOutput) Output(content []byte) {

}
