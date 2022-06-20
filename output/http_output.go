package output

import (
	"bytes"
	"net/http"

	"github.com/chenjiayao/godeng/inter"
)

type HttpOutput struct {
	url string
}

var _ inter.Output = &HttpOutput{}

func (o *HttpOutput) Output(content []byte) {
	var buf bytes.Buffer
	buf.Read(content)
	http.Post(o.url, "application/json", &buf)
}

func MakeHTTPOutput(url string) *HttpOutput {
	return &HttpOutput{
		url: url,
	}
}
