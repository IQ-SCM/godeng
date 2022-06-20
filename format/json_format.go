package format

import (
	"encoding/json"

	"github.com/chenjiayao/godeng/inter"
)

type JSONFormat struct {
}

var _ inter.Format = &JSONFormat{}

func (f *JSONFormat) Format(row inter.Row) []byte {

	m := make(map[string]interface{})
	for _, pair := range row.Items() {
		m[pair.Key()] = pair.Value()
	}
	b, _ := json.Marshal(m)
	return b
}

func MakeJSONFormat() *JSONFormat {
	return &JSONFormat{}
}
