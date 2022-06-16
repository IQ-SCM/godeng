package field

import (
	"github.com/chenjiayao/godeng/constant"
	"github.com/chenjiayao/godeng/inter"
)

type FieldIPv6 struct {
	typ   string //type
	key   string //key
	value interface{}
	rule  inter.Rule
}

var _ inter.Field = &FieldIPv6{}

//generate value by rule
func (f *FieldIPv6) Value() interface{} {
	return f.rule.Generate()
}

func (f *FieldIPv6) Key() string {
	return f.key
}

func MakeFieldIPv6(key string, rule inter.Rule) inter.Field {
	return &FieldIPv6{
		typ:  constant.FILED_TYPE_IPV4,
		key:  key,
		rule: rule,
	}
}
