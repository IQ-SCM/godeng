package rule

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/chenjiayao/godeng/inter"
)

type RuleEmail struct {
}

var _ inter.Rule = &RuleEmail{}

func (rule *RuleEmail) Generate() interface{} {
	return gofakeit.Email()
}

func MakeRuleEmail() *RuleEmail {
	return &RuleEmail{}
}
