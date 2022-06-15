package rule

import (
	"math/rand"
	"time"

	"github.com/chenjiayao/godeng/inter"
)

type RuleInt struct {
	max int64
	min int64
}

var _ inter.Rule = &RuleInt{}

func (rule *RuleInt) Generate() interface{} {
	rand.Seed(time.Now().UnixNano())
	return rand.Int63n(rule.max-rule.min) + rule.min
}

func MakeRuleInt(min, max int64) *RuleInt {
	return &RuleInt{
		min: min,
		max: max,
	}
}
