package field

import (
	"github.com/chenjiayao/godeng/constant"
	"github.com/chenjiayao/godeng/inter"
)

type FieldEnum struct {
	typ   string //type
	key   string //key
	value interface{}
	rule  inter.Rule
}

var _ inter.Field = &FieldEnum{}

//generate value by rule
func (f *FieldEnum) Value() interface{} {
	return f.rule.Generate()
}

func (f *FieldEnum) Key() string {
	return f.key
}

func MakeFieldEnum(key string, rule inter.Rule) inter.Field {
	return &FieldEnum{
		typ:  constant.FILED_TYPE_IPV4,
		key:  key,
		rule: rule,
	}
}
