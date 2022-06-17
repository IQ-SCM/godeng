package rule

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
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
	value, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", base+rand.Float64()), 64)

	return value
}

func MakeRuleFloat(min, max float64) *RuleFloat {
	return &RuleFloat{
		min: min,
		max: max,
	}
}
