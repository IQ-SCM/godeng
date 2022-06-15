package godeng

import (
	"github.com/chenjiayao/godeng/inter"
)

type Row struct {
	fields []inter.Field
}

var _ inter.Row = &Row{}

func (r *Row) Fields() []inter.Field {
	return r.fields
}

func MakeRow(fields []inter.Field) *Row {

	r := &Row{
		fields: fields,
	}
	r.generate()
	return r
}

func (r *Row) generate() {
	for _, field := range r.fields {
		field.Value()
	}
}
