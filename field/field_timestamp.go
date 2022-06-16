package field

import (
	"github.com/chenjiayao/godeng/constant"
	"github.com/chenjiayao/godeng/inter"
)

type FieldTimestamp struct {
	typ   string //type
	key   string //key
	value interface{}
	rule  inter.Rule
}

var _ inter.Field = &FieldTimestamp{}

//generate value by rule
func (f *FieldTimestamp) Value() interface{} {
	return f.rule.Generate()
}

func (f *FieldTimestamp) Key() string {
	return f.key
}

func MakeFieldTimestamp(key string, rule inter.Rule) inter.Field {
	return &FieldTimestamp{
		typ:  constant.FILED_TYPE_IPV4,
		key:  key,
		rule: rule,
	}
}
