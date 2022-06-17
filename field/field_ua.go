package field

import (
	"github.com/chenjiayao/godeng/constant"
	"github.com/chenjiayao/godeng/inter"
)

type FieldUA struct {
	typ   string //type
	key   string //key
	value interface{}
	rule  inter.Rule
}

var _ inter.Field = &FieldUA{}

//generate value by rule
func (f *FieldUA) Value() interface{} {
	return f.rule.Generate()
}

func (f *FieldUA) Key() string {
	return f.key
}

func MakeFieldUA(key string, rule inter.Rule) inter.Field {
	return &FieldUA{
		typ:  constant.FIELD_TYPE_UA,
		key:  key,
		rule: rule,
	}
}
