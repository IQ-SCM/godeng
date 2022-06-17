package rule

import (
	"time"

	"github.com/chenjiayao/godeng/inter"
)

type RuleTimestamp struct {
}

var _ inter.Rule = &RuleTimestamp{}

func (rule *RuleTimestamp) Generate() interface{} {
	return time.Now().Unix()
}

func MakeRuleTimestamp() *RuleTimestamp {
	return &RuleTimestamp{}
}
