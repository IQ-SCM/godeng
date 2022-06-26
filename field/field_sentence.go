package field

import (
	"github.com/chenjiayao/godeng/constant"
	"github.com/chenjiayao/godeng/inter"
)

type FieldSentence struct {
	typ   string //type
	key   string //key
	value interface{}
	rule  inter.Rule
}

var _ inter.Field = &FieldSentence{}

//generate value by rule
func (f *FieldSentence) Value() interface{} {
	return f.rule.Generate()
}

func (f *FieldSentence) Key() string {
	return f.key
}

func MakeFieldSentence(key string, rule inter.Rule) inter.Field {
	return &FieldSentence{
		typ:  constant.FIELD_TYPE_SENTCNCE,
		key:  key,
		rule: rule,
	}
}
