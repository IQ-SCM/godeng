package field

import (
	"github.com/chenjiayao/godeng/constant"
	"github.com/chenjiayao/godeng/inter"
)

type FieldDatetime struct {
	typ   string //type
	key   string //key
	value interface{}
	rule  inter.Rule
}

var _ inter.Field = &FieldDatetime{}

//generate value by rule
func (f *FieldDatetime) Value() interface{} {
	return f.rule.Generate()
}

func (f *FieldDatetime) Key() string {
	return f.key
}

func MakeFieldDatetime(key string, rule inter.Rule) inter.Field {
	return &FieldDatetime{
		typ:  constant.FILED_TYPE_IPV4,
		key:  key,
		rule: rule,
	}
}
