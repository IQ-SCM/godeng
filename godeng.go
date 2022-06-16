package godeng

import (
	"log"

	"github.com/chenjiayao/godeng/constant"
	"github.com/chenjiayao/godeng/field"
	"github.com/chenjiayao/godeng/format"
	"github.com/chenjiayao/godeng/inter"
	"github.com/chenjiayao/godeng/output"
	"github.com/chenjiayao/godeng/rule"
)

type GoDeng struct {
	output   inter.Output
	wangChan chan row
	format   inter.Format
	fields   []inter.Field
	config   Config
}

func makeGoDeng(cfg *Config) *GoDeng {

	g := &GoDeng{
		wangChan: make(chan row),
	}
	switch cfg.f {
	case "json":
		g.format = &format.JSONFormat{}
	default:
		log.Println("unknow format")
	}

	switch cfg.o {
	case "stdout":
		g.output = &output.StdoutOutput{}
	}

	fields := make([]inter.Field, len(cfg.items))
	for idx, item := range cfg.items {
		switch item.typ {
		case constant.FIELD_TYPE_INT:
			r := rule.MakeRuleInt(item.min, item.max)
			fields[idx] = field.MakeFieldInt(item.key, r)
		case constant.FIELD_TYPE_STRING:
			r := rule.MakeRuleString(item.len)
			fields[idx] = field.MakeFieldString(item.key, r)
		default:
			log.Println("unknow field type")
		}
	}
	g.fields = fields
	return g
}

func (g *GoDeng) generateRow() row {
	kvs := make([]rowItem, len(g.fields))
	for idx, field := range g.fields {
		kv := rowItem{
			key: field.Key(),
			val: field.Value(),
		}
		kvs[idx] = kv
	}
	return row{
		kvs: kvs,
	}
}

func (g *GoDeng) barking() {
	for item := range g.wangChan {
		b := g.format.Format(item)
		g.output.Output(b)
	}
}

type rowItem struct {
	key string
	val interface{}
}

type row struct {
	kvs []rowItem
}
