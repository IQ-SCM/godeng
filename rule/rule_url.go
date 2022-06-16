package rule

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/chenjiayao/godeng/inter"
)

type RuleURL struct {
}

var _ inter.Rule = &RuleURL{}

func (rule *RuleURL) Generate() interface{} {
	return gofakeit.URL()
}

func MakeRuleURL() *RuleURL {
	return &RuleURL{}
}
