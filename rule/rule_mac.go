package rule

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/chenjiayao/godeng/inter"
)

type RuleMac struct {
}

var _ inter.Rule = &RuleMac{}

func (rule *RuleMac) Generate() interface{} {
	return gofakeit.MacAddress()
}

func MakeRuleMac() *RuleMac {
	return &RuleMac{}
}
