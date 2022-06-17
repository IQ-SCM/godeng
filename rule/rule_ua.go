package rule

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/chenjiayao/godeng/inter"
)

type RuleUA struct {
}

var _ inter.Rule = &RuleUA{}

func (rule *RuleUA) Generate() interface{} {
	return gofakeit.UserAgent()
}

func MakeRuleUA() *RuleUA {
	return &RuleUA{}
}
