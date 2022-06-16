package field

import (
	"github.com/chenjiayao/godeng/constant"
	"github.com/chenjiayao/godeng/inter"
)

type FieldFloat struct {
	typ   string //type
	key   string //key
	value interface{}
	rule  inter.Rule
}

var _ inter.Field = &FieldFloat{}

//generate value by rule
func (f *FieldFloat) Value() interface{} {
	return f.rule.Generate()
}

func (f *FieldFloat) Key() string {
	return f.key
}

func MakeFieldFloat(key string, rule inter.Rule) inter.Field {
	return &FieldFloat{
		typ:  constant.FIELD_TYPE_FLOAT,
		key:  key,
		rule: rule,
	}
}
