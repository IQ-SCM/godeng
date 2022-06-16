package rule

import (
	"time"

	"github.com/chenjiayao/godeng/inter"
)

type RuleTimestamp struct {
	enums []interface{}
}

var _ inter.Rule = &RuleTimestamp{}

func (rule *RuleTimestamp) Generate() interface{} {
	return time.Now().Unix()
}

func MakeRuleTimestamp(enums []interface{}) *RuleTimestamp {
	return &RuleTimestamp{
		enums: enums,
	}
}
