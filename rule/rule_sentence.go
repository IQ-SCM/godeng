package rule

import (
	"math/rand"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/chenjiayao/godeng/inter"
)

type RuleSentence struct {
}

var _ inter.Rule = &RuleSentence{}

func (rule *RuleSentence) Generate() interface{} {
	rand.Seed(time.Now().UnixNano())
	count := rand.Intn(10)
	return gofakeit.Sentence(count + 10)
}

func MakeRuleSentence() *RuleSentence {
	return &RuleSentence{}
}
