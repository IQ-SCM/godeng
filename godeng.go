package godeng

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/chenjiayao/godeng/constant"
	"github.com/chenjiayao/godeng/field"
	"github.com/chenjiayao/godeng/format"
	"github.com/chenjiayao/godeng/inter"
	"github.com/chenjiayao/godeng/output"
	"github.com/chenjiayao/godeng/rule"
)

type GoDeng struct {
	output    inter.Output
	wangChan  chan row
	format    inter.Format
	fields    []inter.Field
	config    Config
	count     int64
	sleep     int64
	forever   bool
	url       string
	tablename string
	cancel    context.CancelFunc
	ctx       context.Context
}

func MakeGoDeng(cfg *Config, o string, f string, count int64, forever bool, sleep int64, url string, tablename string) *GoDeng {

	g := &GoDeng{
		wangChan: make(chan row, 20480),
	}
	switch f {
	case "json":
		g.format = &format.JSONFormat{}
	default:
		log.Println("unknow format")
	}

	switch o {
	case "stdout":
		g.output = &output.StdoutOutput{}
	}

	fields := make([]inter.Field, len(cfg.items))
	for idx, item := range cfg.items {
		switch item.typ {
		case constant.FIELD_TYPE_INT:
			r := rule.MakeRuleInt(int64(item.min), int64(item.max))
			fields[idx] = field.MakeFieldInt(item.key, r)
		case constant.FIELD_TYPE_STRING:
			r := rule.MakeRuleString(item.len)
			fields[idx] = field.MakeFieldString(item.key, r)
		case constant.FIELD_TYPE_DATETIME:
			r := rule.MakeRuleDatetime()
			fields[idx] = field.MakeFieldDatetime(item.key, r)
		case constant.FIELD_TYPE_FLOAT:
			r := rule.MakeRuleFloat(float64(item.min), float64(item.max))
			fields[idx] = field.MakeFieldFloat(item.key, r)
		case constant.FILED_TYPE_MAC:
			r := rule.MakeRuleMac()
			fields[idx] = field.MakeFieldMac(item.key, r)
		case constant.FIELD_TYPE_BOOL:
			r := rule.MakeRuleBool()
			fields[idx] = field.MakeFieldBool(item.key, r)
		case constant.FILED_TYPE_IPV4:
			r := rule.MakeRuleIPv4()
			fields[idx] = field.MakeFieldIPv4(item.key, r)
		case constant.FILED_TYPE_IPV6:
			r := rule.MakeRuleIPv6()
			fields[idx] = field.MakeFieldIPv6(item.key, r)
		case constant.FIELD_TYPE_ENUM:
			r := rule.MakeRuleEnum(item.enums)
			fields[idx] = field.MakeFieldEnum(item.key, r)
		case constant.FIELD_TYPE_URL:
			r := rule.MakeRuleURL()
			fields[idx] = field.MakeFieldURL(item.key, r)
		case constant.FILELD_TYPE_EMAIL:
			r := rule.MakeRuleEmail()
			fields[idx] = field.MakeFieldEmail(item.key, r)
		case constant.FIELD_TYPE_SEQUENCE:
			r := rule.MakeRuleSequence(item.begin, item.step)
			fields[idx] = field.MakeFieldSequence(item.key, r)
		case constant.FIELD_TYPE_TIMESTAMP:
			r := rule.MakeRuleTimestamp()
			fields[idx] = field.MakeFieldTimestamp(item.key, r)
		case constant.FIELD_TYPE_UA:
			r := rule.MakeRuleUA()
			fields[idx] = field.MakeFieldUA(item.key, r)
		default:
			log.Println("unknow field type:", item.typ)
		}
	}
	g.fields = fields
	ctx, cancel := context.WithCancel(context.Background())
	g.ctx = ctx
	g.cancel = cancel

	g.sleep = sleep
	g.count = count
	g.forever = forever
	g.url = url
	g.tablename = tablename
	return g
}

func (g *GoDeng) waitSignal() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGTERM)
	<-ch
	g.cancel()
}

func (g *GoDeng) Start() {
	go g.waitSignal()
	go g.barking()

	if g.forever {
		g.runForever()
	} else {
		g.run()
	}
}

func (g *GoDeng) runForever() {

	for {
		select {
		case <-g.ctx.Done():
			return
		default:
			g.wangChan <- g.generateRow()
		}
	}
}

func (g *GoDeng) run() {

	for i := 1; i <= int(g.count); i++ {
		g.wangChan <- g.generateRow()
	}
	g.cancel()
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

	for {
		select {
		case <-g.ctx.Done():
			return
		case item := <-g.wangChan:
			if g.sleep > 0 {
				time.Sleep(time.Duration(g.sleep) * time.Second)
			}
			b := g.format.Format(item)
			g.output.Output(b)
		}
	}
}

type rowItem struct {
	key string
	val interface{}
}

var _ inter.Pair = &rowItem{}

func (item rowItem) Key() string {
	return item.key
}

func (item rowItem) Value() interface{} {
	return item.val
}

//---------------------splite---------------------/
type row struct {
	kvs []rowItem
}

var _ inter.Row = row{}

func (r row) Items() []inter.Pair {
	ret := make([]inter.Pair, len(r.kvs))
	for idx, kv := range r.kvs {
		ret[idx] = kv
	}
	return ret
}
