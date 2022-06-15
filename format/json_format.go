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
	for _, field := range row.Fields() {
		m[field.Key()] = field.Value()
	}
	b, _ := json.Marshal(m)
	return b
}
