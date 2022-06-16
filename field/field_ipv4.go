package field

import (
	"github.com/chenjiayao/godeng/constant"
	"github.com/chenjiayao/godeng/inter"
)

type FieldIPv4 struct {
	typ   string //type
	key   string //key
	value interface{}
	rule  inter.Rule
}

var _ inter.Field = &FieldIPv4{}

//generate value by rule
func (f *FieldIPv4) Value() interface{} {
	return f.rule.Generate()
}

func (f *FieldIPv4) Key() string {
	return f.key
}

func MakeFieldIPv4(key string, rule inter.Rule) inter.Field {
	return &FieldIPv4{
		typ:  constant.FILED_TYPE_IPV4,
		key:  key,
		rule: rule,
	}
}
