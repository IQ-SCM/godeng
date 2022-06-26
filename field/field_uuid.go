package field

import (
	"github.com/chenjiayao/godeng/constant"
	"github.com/chenjiayao/godeng/inter"
)

type FieldUUID struct {
	typ   string //type
	key   string //key
	value interface{}
	rule  inter.Rule
}

var _ inter.Field = &FieldUUID{}

//generate value by rule
func (f *FieldUUID) Value() interface{} {
	return f.rule.Generate()
}

func (f *FieldUUID) Key() string {
	return f.key
}

func MakeFieldUUID(key string, rule inter.Rule) inter.Field {
	return &FieldUUID{
		typ:  constant.FILED_TYPE_IPV4,
		key:  key,
		rule: rule,
	}
}
