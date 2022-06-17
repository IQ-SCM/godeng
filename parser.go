package godeng

import (
	"errors"
	"log"

	"github.com/chenjiayao/godeng/constant"
	"github.com/chenjiayao/godeng/validator"
	"github.com/spf13/viper"
)

func Parser(cfgFile string) (*Config, error) {
	viper.SetConfigFile(cfgFile)
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	fieldsVal := viper.Get("fields")
	if fieldsVal == nil {
		return nil, errors.New("fields is empty")
	}
	inters, ok := fieldsVal.([]interface{})
	if !ok {
		return nil, errors.New("fields is not array")
	}

	fields := make([]map[string]interface{}, len(inters))
	for idx, inter := range inters {
		fields[idx] = inter.(map[string]interface{})
	}
	return makeFields(fields)
}

func makeFields(fields []map[string]interface{}) (*Config, error) {
	cfg := &Config{}
	cfgItems := make([]*ConfigItem, len(fields))

	for idx, field := range fields {

		fieldType, ok := field["type"].(string)
		if !ok {
			return nil, errors.New("field type is not string")
		}
		var cfgItem *ConfigItem
		var err error

		switch fieldType {
		case constant.FIELD_TYPE_INT:
			cfgItem, err = parserInt(field)
		case constant.FIELD_TYPE_FLOAT:
			cfgItem, err = parserFloat(field)
		case constant.FIELD_TYPE_BOOL:
			cfgItem, err = parserBool(field)
		case constant.FIELD_TYPE_STRING:
			cfgItem, err = parserString(field)
		case constant.FILED_TYPE_IPV4:
			cfgItem, err = parserIPv4(field)
		case constant.FILED_TYPE_IPV6:
			cfgItem, err = parserIPv6(field)
		case constant.FIELD_TYPE_URL:
			cfgItem, err = parserURL(field)
		case constant.FILELD_TYPE_EMAIL:
			cfgItem, err = parserEmail(field)
		case constant.FILED_TYPE_MAC:
			cfgItem, err = parserMac(field)
		case constant.FIELD_TYPE_ENUM:
			cfgItem, err = parserEnum(field)
		case constant.FIELD_TYPE_DATETIME:
			cfgItem, err = parserDatetime(field)
		case constant.FIELD_TYPE_TIMESTAMP:
			cfgItem, err = parserTimestamp(field)
		case constant.FIELD_TYPE_SEQUENCE:
			cfgItem, err = parserSequence(field)
		default:
			log.Printf("unknown field type: %s", fieldType)
		}
		if err != nil {
			return nil, err
		}
		cfgItem.typ = fieldType
		cfgItems[idx] = cfgItem
	}

	cfg.items = cfgItems
	return cfg, nil
}

func parserInt(field map[string]interface{}) (*ConfigItem, error) {
	if err := validator.ValidateInt(field); err != nil {
		return nil, err
	}

	max := float64(0)
	min := float64(0)
	_, ok := field["max"]
	if ok {
		max = field["max"].(float64)
	}
	_, ok = field["min"]
	if ok {
		min = field["min"].(float64)
	}

	cfgItem := &ConfigItem{
		key: field["key"].(string),
		max: max,
		min: min,
	}
	return cfgItem, nil
}

func parserFloat(field map[string]interface{}) (*ConfigItem, error) {
	if err := validator.ValidateFloat(field); err != nil {
		return nil, err
	}

	max := float64(100)
	min := float64(0)
	_, ok := field["max"]
	if ok {
		max = field["max"].(float64)
	}
	_, ok = field["min"]
	if ok {
		min = field["min"].(float64)
	}

	cfgItem := &ConfigItem{
		key: field["key"].(string),
		max: max,
		min: min,
	}
	return cfgItem, nil
}

func parserString(field map[string]interface{}) (*ConfigItem, error) {
	if err := validator.ValidateString(field); err != nil {
		return nil, err
	}

	len := int64(0)
	_, ok := field["len"]
	if ok {
		len = field["len"].(int64)
	}

	cfgItem := &ConfigItem{
		key: field["key"].(string),
		len: len,
	}
	return cfgItem, nil
}

func parserBool(field map[string]interface{}) (*ConfigItem, error) {
	if err := validator.ValidateBool(field); err != nil {
		return nil, err
	}

	cfgItem := &ConfigItem{
		key: field["key"].(string),
	}
	return cfgItem, nil
}

func parserIPv4(field map[string]interface{}) (*ConfigItem, error) {
	if err := validator.ValidateIPv4(field); err != nil {
		return nil, err
	}

	cfgItem := &ConfigItem{
		key: field["key"].(string),
	}
	return cfgItem, nil
}

func parserIPv6(field map[string]interface{}) (*ConfigItem, error) {
	if err := validator.ValidateIPv6(field); err != nil {
		return nil, err
	}

	cfgItem := &ConfigItem{
		key: field["key"].(string),
	}
	return cfgItem, nil
}

func parserURL(field map[string]interface{}) (*ConfigItem, error) {
	if err := validator.ValidateURL(field); err != nil {
		return nil, err
	}

	cfgItem := &ConfigItem{
		key: field["key"].(string),
	}
	return cfgItem, nil
}

func parserTimestamp(field map[string]interface{}) (*ConfigItem, error) {
	if err := validator.ValidateTimestamp(field); err != nil {
		return nil, err
	}

	cfgItem := &ConfigItem{
		key: field["key"].(string),
	}
	return cfgItem, nil
}

func parserDatetime(field map[string]interface{}) (*ConfigItem, error) {
	if err := validator.ValidateDatetime(field); err != nil {
		return nil, err
	}

	cfgItem := &ConfigItem{
		key: field["key"].(string),
	}
	return cfgItem, nil
}

func parserMac(field map[string]interface{}) (*ConfigItem, error) {
	if err := validator.ValidateMac(field); err != nil {
		return nil, err
	}

	cfgItem := &ConfigItem{
		key: field["key"].(string),
	}
	return cfgItem, nil
}

func parserEmail(field map[string]interface{}) (*ConfigItem, error) {
	if err := validator.ValidateEmail(field); err != nil {
		return nil, err
	}

	cfgItem := &ConfigItem{
		key: field["key"].(string),
	}
	return cfgItem, nil
}

func parserEnum(field map[string]interface{}) (*ConfigItem, error) {
	if err := validator.ValidateEnum(field); err != nil {
		return nil, err
	}

	cfgItem := &ConfigItem{
		key:   field["key"].(string),
		enums: field["enum"].([]interface{}),
	}
	return cfgItem, nil
}

func parserSequence(field map[string]interface{}) (*ConfigItem, error) {
	if err := validator.ValidateSequence(field); err != nil {
		return nil, err
	}

	cfgItem := &ConfigItem{
		key:   field["key"].(string),
		begin: field["begin"].(float64),
		step:  field["step"].(float64),
	}
	return cfgItem, nil
}
