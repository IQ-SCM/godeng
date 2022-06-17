package field

import (
	"github.com/chenjiayao/godeng/constant"
	"github.com/chenjiayao/godeng/inter"
)

type FieldMac struct {
	typ   string //type
	key   string //key
	value interface{}
	rule  inter.Rule
}

var _ inter.Field = &FieldMac{}

//generate value by rule
func (f *FieldMac) Value() interface{} {
	return f.rule.Generate()
}

func (f *FieldMac) Key() string {
	return f.key
}

func MakeFieldMac(key string, rule inter.Rule) inter.Field {
	return &FieldMac{
		typ:  constant.FILED_TYPE_MAC,
		key:  key,
		rule: rule,
	}
}
