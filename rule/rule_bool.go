package rule

import (
	"math/rand"
	"time"

	"github.com/chenjiayao/godeng/inter"
)

type RuleBool struct {
}

var _ inter.Rule = &RuleBool{}

func (rule *RuleBool) Generate() interface{} {
	rand.Seed(time.Now().Unix())
	return rand.Intn(2) == 1
}

func MakeRuleBool() *RuleBool {
	return &RuleBool{}
}
