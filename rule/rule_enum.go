package rule

import (
	"math/rand"
	"time"

	"github.com/chenjiayao/godeng/inter"
)

type RuleEnum struct {
	enums []interface{}
}

var _ inter.Rule = &RuleEnum{}

func (rule *RuleEnum) Generate() interface{} {
	rand.Seed(time.Now().UnixNano())
	idx := rand.Int63n(int64(len(rule.enums)))
	return rule.enums[idx]
}

func MakeRuleEnum(enums []interface{}) *RuleEnum {
	return &RuleEnum{
		enums: enums,
	}
}
