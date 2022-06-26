package rule

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/chenjiayao/godeng/inter"
)

type RuleUUID struct {
}

var _ inter.Rule = &RuleUUID{}

func (rule *RuleUUID) Generate() interface{} {
	return gofakeit.URL()
}

func MakeRuleUUID() *RuleUUID {
	return &RuleUUID{}
}
