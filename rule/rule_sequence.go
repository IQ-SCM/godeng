package rule

import (
	"github.com/chenjiayao/godeng/inter"
)

type RuleSequence struct {
	begin   float64
	step    float64
	current float64
}

var _ inter.Rule = &RuleIPv6{}

func (rule *RuleSequence) Generate() interface{} {
	defer func() {
		rule.current += rule.step
	}()
	return rule.current
}

func MakeRuleSequence(begin, step float64) *RuleSequence {
	return &RuleSequence{
		begin:   begin,
		step:    step,
		current: begin,
	}
}
