package validator

import (
	"fmt"
)

func ValidateInt(cfgMap map[string]interface{}) error {
	_, ok := cfgMap["key"]
	if !ok {
		return fmt.Errorf("key is empty")
	}
	_, minExist := cfgMap["min"]
	_, maxExist := cfgMap["max"]
	if minExist && maxExist {
		min, minOk := cfgMap["min"].(float64)
		max, maxOk := cfgMap["max"].(float64)

		if !minOk || !maxOk {
			return fmt.Errorf("min or max is not number")
		}

		if min > max {
			return fmt.Errorf("min is greater than max")
		}
	}
	return nil
}

func ValidateFloat(cfgMap map[string]interface{}) error {
	_, ok := cfgMap["key"]
	if !ok {
		return fmt.Errorf("key is empty")
	}
	_, minExist := cfgMap["min"]
	_, maxExist := cfgMap["max"]
	if minExist && maxExist {
		min, minOk := cfgMap["min"].(float64)
		max, maxOk := cfgMap["max"].(float64)

		if !minOk || !maxOk {
			return fmt.Errorf("min or max is not number")
		}

		if min > max {
			return fmt.Errorf("min is greater than max")
		}
	}
	return nil
}

func ValidateString(cfgMap map[string]interface{}) error {
	_, ok := cfgMap["key"]
	if !ok {
		return fmt.Errorf("key is empty")
	}

	_, lenExist := cfgMap["len"]
	if lenExist {
		len, lenOk := cfgMap["len"].(int64)
		if !lenOk {
			return fmt.Errorf("len is not number")
		}
		if len < 0 {
			return fmt.Errorf("len is less than 0")
		}
	}
	return nil
}

func ValidateBool(cfgMap map[string]interface{}) error {
	_, ok := cfgMap["key"]
	if !ok {
		return fmt.Errorf("key is empty")
	}

	return nil
}

func ValidateIPv4(cfgMap map[string]interface{}) error {
	_, ok := cfgMap["key"]
	if !ok {
		return fmt.Errorf("key is empty")
	}

	return nil
}

func ValidateIPv6(cfgMap map[string]interface{}) error {
	_, ok := cfgMap["key"]
	if !ok {
		return fmt.Errorf("key is empty")
	}

	return nil
}

func ValidateURL(cfgMap map[string]interface{}) error {
	_, ok := cfgMap["key"]
	if !ok {
		return fmt.Errorf("key is empty")
	}
	return nil
}

func ValidateTimestamp(cfgMap map[string]interface{}) error {
	_, ok := cfgMap["key"]
	if !ok {
		return fmt.Errorf("key is empty")
	}
	return nil
}
func ValidateDatetime(cfgMap map[string]interface{}) error {
	_, ok := cfgMap["key"]
	if !ok {
		return fmt.Errorf("key is empty")
	}
	return nil
}

func ValidateEmail(cfgMap map[string]interface{}) error {
	_, ok := cfgMap["key"]
	if !ok {
		return fmt.Errorf("key is empty")
	}
	return nil
}

func ValidateEnum(cfgMap map[string]interface{}) error {
	_, ok := cfgMap["key"]
	if !ok {
		return fmt.Errorf("key is empty")
	}

	_, enumExist := cfgMap["enum"]
	if !enumExist {
		return fmt.Errorf("enum is nil")
	}
	enumVal, ok := cfgMap["enum"].([]interface{})
	if !ok {
		return fmt.Errorf("enum is not array")
	}
	if len(enumVal) == 0 {
		return fmt.Errorf("enum is empty")
	}
	return nil
}
