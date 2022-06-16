package rule

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/chenjiayao/godeng/inter"
)

type RuleIPv4 struct {
}

var _ inter.Rule = &RuleIPv4{}

func (rule *RuleIPv4) Generate() interface{} {
	return gofakeit.IPv4Address()
}

func MakeRuleIPv4() *RuleIPv4 {
	return &RuleIPv4{}
}
