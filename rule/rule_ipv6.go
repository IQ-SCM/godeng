package rule

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/chenjiayao/godeng/inter"
)

type RuleIPv6 struct {
}

var _ inter.Rule = &RuleIPv6{}

func (rule *RuleIPv6) Generate() interface{} {
	return gofakeit.IPv6Address()
}

func MakeRuleIPv6() *RuleIPv6 {
	return &RuleIPv6{}
}
