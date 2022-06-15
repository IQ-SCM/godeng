package output

import "github.com/chenjiayao/godeng/inter"

type HttpOutput struct {
	url string
}

var _ inter.Output = &HttpOutput{}

func (o *HttpOutput) Output(content []byte) {

}
