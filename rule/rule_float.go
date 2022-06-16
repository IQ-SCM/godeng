package rule

import (
	"math"
	"math/rand"
	"time"

	"github.com/chenjiayao/godeng/inter"
)

type RuleFloat struct {
	max float64
	min float64
}

var _ inter.Rule = &RuleInt{}

func (rule *RuleFloat) Generate() interface{} {
	rand.Seed(time.Now().UnixNano())

	//not corrent, but it works
	minInt := math.Ceil(rule.min)
	maxInt := math.Ceil(rule.max)

	base := float64(rand.Int63n(int64(maxInt-minInt)) + int64(minInt))
	return base + rand.Float64()
}

func MakeRuleFloat(min, max float64) *RuleFloat {
	return &RuleFloat{
		min: min,
		max: max,
	}
}
