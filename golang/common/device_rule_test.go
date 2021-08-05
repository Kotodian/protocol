package common

import (
	"testing"
)

func TestDeviceRuleMatch(t *testing.T) {
	s := "N010216905180001"
	dr := NewDeviceRule(`N\d{3}\d{2}(\d{4})\d{3}\d{3}`, "u")
	if s, ok := dr.Match(s); !ok {
		t.Error("根据规则引擎匹配失败")
	} else if s != "6905" {
		t.Error("无法根据规则引擎获取标示代码")
	}
}
