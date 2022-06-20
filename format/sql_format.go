package format

import (
	"fmt"

	"github.com/chenjiayao/godeng/inter"
)

type SQLFormat struct {
	Tablename string
}

var _ inter.Format = &SQLFormat{}

func (f *SQLFormat) Format(row inter.Row) []byte {

	ret := fmt.Sprintf("insert into %s (", f.Tablename)
	// insert into table (field1, field2, field3) values (value1, value2, value3)
	for _, field := range row.Items() {
		ret += fmt.Sprintf("%s, ", field.Key())
	}
	ret = ret[:len(ret)-2] + ") values ("
	for _, field := range row.Items() {

		v, ok := field.Value().(string)
		if ok {
			ret += fmt.Sprintf("\"%v\", ", v)

		} else {
			ret += fmt.Sprintf("%v, ", field.Value())
		}
	}
	ret = ret[:len(ret)-2] + ");"
	return []byte(ret)
}

func MakeSQLFormat(tablename string) *SQLFormat {
	return &SQLFormat{
		Tablename: tablename,
	}
}
