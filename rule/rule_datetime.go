package rule

import (
	"time"

	"github.com/chenjiayao/godeng/inter"
)

type RuleDatetime struct {
}

var _ inter.Rule = &RuleDatetime{}

func (rule *RuleDatetime) Generate() interface{} {
	return time.Now().Local().Format("2006-01-02 15:04:05")
}

func MakeRuleDatetime() *RuleDatetime {
	return &RuleDatetime{}
}
