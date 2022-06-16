package field

import (
	"github.com/chenjiayao/godeng/constant"
	"github.com/chenjiayao/godeng/inter"
)

type FieldURL struct {
	typ   string //type
	key   string //key
	value interface{}
	rule  inter.Rule
}

var _ inter.Field = &FieldURL{}

//generate value by rule
func (f *FieldURL) Value() interface{} {
	return f.rule.Generate()
}

func (f *FieldURL) Key() string {
	return f.key
}

func MakeFieldURL(key string, rule inter.Rule) inter.Field {
	return &FieldURL{
		typ:  constant.FILED_TYPE_IPV4,
		key:  key,
		rule: rule,
	}
}
