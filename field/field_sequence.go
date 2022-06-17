package field

import (
	"github.com/chenjiayao/godeng/constant"
	"github.com/chenjiayao/godeng/inter"
)

type FieldSequence struct {
	typ   string //type
	key   string //key
	value interface{}
	rule  inter.Rule
}

var _ inter.Field = &FieldSequence{}

//generate value by rule
func (f *FieldSequence) Value() interface{} {
	return f.rule.Generate()
}

func (f *FieldSequence) Key() string {
	return f.key
}

func MakeFieldSequence(key string, rule inter.Rule) inter.Field {
	return &FieldSequence{
		typ:  constant.FIELD_TYPE_STRING,
		key:  key,
		rule: rule,
	}
}
