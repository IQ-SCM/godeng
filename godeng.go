package godeng

import (
	"context"
	"fmt"
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
	loop      bool
	url       string
	tablename string
	cancel    context.CancelFunc
	ctx       context.Context
}

func validate(o string, f string, count int64, loop bool, sleep int64, url string, tablename string, file string) error {
	if !loop && count <= 0 {
		return fmt.Errorf("count must be greater than 0")
	}

	if sleep < 0 {
		return fmt.Errorf("sleep must be greater than 0")
	}

	if o == "http" && url == "" {
		return fmt.Errorf("url must be set if output is http")
	}

	if o == "http" && f != "json" {
		return fmt.Errorf("only json format is supported for http output")
	}

	if f == "sql" && tablename == "" {
		return fmt.Errorf("tablename must be set if output is sql")
	}

	if o == "file" && file == "" {
		return fmt.Errorf("file must be set if output is file")
	}
	return nil
}

func MakeGoDeng(cfg *Config, o string, f string, count int64, loop bool, sleep int64, url string, tablename string, file string) (*GoDeng, error) {

	if err := validate(o, f, count, loop, sleep, url, tablename, file); err != nil {
		return nil, err
	}

	g := &GoDeng{
		wangChan: make(chan row),
	}
	switch f {
	case "json":
		g.format = format.MakeJSONFormat()
	case "sql":
		g.format = format.MakeSQLFormat(tablename)
	default:
		return nil, fmt.Errorf("unknown format: %s", f)
	}

	switch o {
	case "stdout":
		g.output = output.MakeStdoutOutput()
	case "file":
		g.output = output.MakeFileOutput(file)
	case "http":
		g.output = output.MakeHTTPOutput(url)
		g.format = format.MakeJSONFormat()
	default:
		return nil, fmt.Errorf("unknown output: %s", f)
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
		case constant.FIELD_TYPE_UUID:
			r := rule.MakeRuleUUID()
			fields[idx] = field.MakeFieldUUID(item.key, r)
		case constant.FIELD_TYPE_SENTCNCE:
			r := rule.MakeRuleSentence()
			fields[idx] = field.MakeFieldSentence(item.key, r)
		default:
			return nil, fmt.Errorf("unknown field type: %s", item.typ)
		}
	}
	g.fields = fields
	ctx, cancel := context.WithCancel(context.Background())
	g.ctx = ctx
	g.cancel = cancel

	g.sleep = sleep
	g.count = count
	g.loop = loop
	g.url = url
	g.tablename = tablename
	return g, nil
}

func (g *GoDeng) waitSignal() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGTERM)
	<-ch
	g.cancel()
}

func (g *GoDeng) Start() {
	go g.waitSignal()

	if g.loop {
		go g.runForever()
	} else {
		go g.run()
	}
	g.barking()
}

func (g *GoDeng) runForever() {

	defer close(g.wangChan)
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
	defer close(g.wangChan)

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
		case item, ok := <-g.wangChan:
			if !ok {
				return
			}
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
