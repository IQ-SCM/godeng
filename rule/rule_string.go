package rule

import (
	"math/rand"

	"github.com/chenjiayao/godeng/inter"
)

type RuleString struct {
	len int64
}

var letterRunes = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

var _ inter.Rule = &RuleString{}

func MakeRuleString(len int64) inter.Rule {
	return &RuleString{
		len: len,
	}
}

func (rule *RuleString) Generate() interface{} {
	b := make([]rune, rule.len)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
