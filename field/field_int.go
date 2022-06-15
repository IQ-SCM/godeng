package field

import (
	"github.com/chenjiayao/godeng/constant"
	"github.com/chenjiayao/godeng/inter"
)

type FieldInt struct {
	typ   string //type
	key   string //key
	value interface{}
	rule  inter.Rule
}

var _ inter.Field = &FieldInt{}

//generate value by rule
func (f *FieldInt) Value() interface{} {
	return f.rule.Generate()
}

func (f *FieldInt) Key() string {
	return f.key
}

func MakeFieldInt(key string, rule inter.Rule) inter.Field {
	return &FieldInt{
		typ:  constant.FIELD_TYPE_INT,
		key:  key,
		rule: rule,
	}
}
