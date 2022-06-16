package field

import (
	"github.com/chenjiayao/godeng/constant"
	"github.com/chenjiayao/godeng/inter"
)

type FieldBool struct {
	typ   string //type
	key   string //key
	value interface{}
	rule  inter.Rule
}

var _ inter.Field = &FieldBool{}

//generate value by rule
func (f *FieldBool) Value() interface{} {
	return f.rule.Generate()
}

func (f *FieldBool) Key() string {
	return f.key
}

func MakeFieldBool(key string, rule inter.Rule) inter.Field {
	return &FieldBool{
		typ:  constant.FIELD_TYPE_BOOL,
		key:  key,
		rule: rule,
	}
}
