package common

import (
	"bytes"
	"regexp"
)

type DeviceRule struct {
	Regexp *regexp.Regexp
	prefix string
}

// NewDeviceRule 规则匹配
func NewDeviceRule(s string, p string) *DeviceRule {
	return &DeviceRule{
		Regexp: regexp.MustCompile(s),
		prefix: p,
	}
}

// 匹配并返回匹配的结果
func (dr DeviceRule) Match(s string) (ret string, b bool) {
	bs := dr.Regexp.FindSubmatch([]byte(s))
	if len(bs) == 0 {
		return "", false
	} else if len(bs) == 1 {
		return string(bs[0]), true
	}
	return string(bytes.Join(bs[1:], []byte(""))), true
}

func (dr DeviceRule) GetPrefix() string {
	return dr.prefix
}

//-------------------------------------------------------------------
type KindRule int

const (
	KindRuleCl           KindRule = 0
	KindRuleSzunit       KindRule = 1
	KindRuleSinexcel     KindRule = 2
	KindRuleYunKuaiChong KindRule = 3
)

func (kr KindRule) Int() int {
	return int(kr)
}

func (kr KindRule) String() string {
	switch kr {
	case 0:
		return "cl"
	case 1:
		return "scmd"
	case 2:
		return "sinexcel"
	case 3:
		return "yunkuaichong"
	default:
		return "unknown"
	}
}

var Rules map[KindRule]*DeviceRule

func init() {
	Rules = map[KindRule]*DeviceRule{
		KindRuleSzunit:       NewDeviceRule(`N\d{3}\d{2}(\d{4})\d{3}\d{3}`, "u"),
		KindRuleCl:           NewDeviceRule(`\d+`, "cl"),
		KindRuleSinexcel:     NewDeviceRule(`\d+`, "s"),
		KindRuleYunKuaiChong: NewDeviceRule(`\d+`, "y"),
	}
}
