package field

import (
	"github.com/chenjiayao/godeng/constant"
	"github.com/chenjiayao/godeng/inter"
)

type FieldString struct {
	typ   string //type
	key   string //key
	value interface{}
	rule  inter.Rule
}

var _ inter.Field = &FieldString{}

//generate value by rule
func (f *FieldString) Value() interface{} {
	return f.rule.Generate()
}

func (f *FieldString) Key() string {
	return f.key
}

func MakeFieldString(key string, rule inter.Rule) inter.Field {
	return &FieldString{
		typ:  constant.FIELD_TYPE_STRING,
		key:  key,
		rule: rule,
	}
}
