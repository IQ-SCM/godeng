package field

import (
	"github.com/chenjiayao/godeng/constant"
	"github.com/chenjiayao/godeng/inter"
)

type FieldEmail struct {
	typ   string //type
	key   string //key
	value interface{}
	rule  inter.Rule
}

var _ inter.Field = &FieldEmail{}

//generate value by rule
func (f *FieldEmail) Value() interface{} {
	return f.rule.Generate()
}

func (f *FieldEmail) Key() string {
	return f.key
}

func MakeFieldEmail(key string, rule inter.Rule) inter.Field {
	return &FieldEmail{
		typ:  constant.FILED_TYPE_IPV4,
		key:  key,
		rule: rule,
	}
}
