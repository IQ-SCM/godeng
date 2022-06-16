package rule

import (
	"math/rand"
	"time"

	"github.com/chenjiayao/godeng/inter"
)

type RuleString struct {
	len int64
}

const (
	DEFAULT_LEN = 20
)

var letterRunes = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

var _ inter.Rule = &RuleString{}

func MakeRuleString(len int64) inter.Rule {
	return &RuleString{
		len: len,
	}
}

func (rule *RuleString) Generate() interface{} {

	if rule.len == 0 {
		rand.Seed(time.Now().UnixNano())
		rule.len = rand.Int63n(DEFAULT_LEN)
	}

	b := make([]rune, rule.len)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
