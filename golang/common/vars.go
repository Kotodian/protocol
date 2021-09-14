package common

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
	"math"
	"strconv"
	"strings"
)

// KindCurrent 电流（分辨率: 1mA）
type KindCurrent float64

func (c KindCurrent) Value() (driver.Value, error) {
	return int64(float64(c) * 1000), nil
}

func (c *KindCurrent) Scan(input interface{}) error {
	switch input.(type) {
	case int64:
		*c = KindCurrent(input.(int64)) / 1000
	case float64:
		*c = KindCurrent(input.(float64)) / 1000
	case []byte:
		v, _ := strconv.ParseFloat(string(input.([]byte)), 64)
		*c = KindCurrent(v) / 1000
	}
	return nil
}

func (c KindCurrent) Float64() float64 {
	return Round(float64(c), 3)
}

// KindVoltage 电压 （分辨率: 0.1V）
type KindVoltage float64

func (c KindVoltage) Value() (driver.Value, error) {
	return int64(float64(c) * 10), nil
}

func (c *KindVoltage) Scan(input interface{}) error {
	switch input.(type) {
	case int64:
		*c = KindVoltage(input.(int64)) / 10
	case float64:
		*c = KindVoltage(input.(float64)) / 10
	case []byte:
		v, _ := strconv.ParseFloat(string(input.([]byte)), 64)
		*c = KindVoltage(v) / 10
	}
	return nil
}

func (c KindVoltage) Float64() float64 {
	return Round(float64(c), 1)
}

// KindPowerKW 功率 （分辨率: 0.01W）
type KindPowerKW float64

func (c KindPowerKW) Value() (driver.Value, error) {
	return int64(float64(c) * 100000), nil
}

func (c *KindPowerKW) Scan(input interface{}) error {
	switch input.(type) {
	case int64:
		*c = KindPowerKW(input.(int64)) / 100000
	case float64:
		*c = KindPowerKW(input.(float64)) / 100000
	case []byte:
		v, _ := strconv.ParseFloat(string(input.([]byte)), 64)
		*c = KindPowerKW(v) / 100000
	}
	return nil
}

func (c KindPowerKW) Float64() float64 {
	return Round(float64(c), 5)
}

//
type KindUnitPrice float64

func (c KindUnitPrice) Value() (driver.Value, error) {
	// return int64(float64(c) * 1000), nil
	return strconv.ParseInt(fmt.Sprintf("%.f", float64(c)*1000), 10, 64)
}

func (c *KindUnitPrice) Scan(input interface{}) error {
	switch input.(type) {
	case int64:
		*c = KindUnitPrice(input.(int64)) / 1000
	case float64:
		*c = KindUnitPrice(input.(float64)) / 1000
	case []byte:
		v, _ := strconv.ParseFloat(string(input.([]byte)), 64)
		*c = KindUnitPrice(v) / 1000
	}
	return nil
}

func (c KindUnitPrice) Float64() float64 {
	return Round(float64(c), 3)
}

func (c KindUnitPrice) String() string {
	return fmt.Sprintf("%.3f", c.Float64())
}

// KindPrice 费用
type KindPrice float64

func (c KindPrice) Value() (driver.Value, error) {
	// return int64(float64(c) * 100), nil // 丢精度
	// return int64(utils.Round(float64(c)*100, 2)), nil
	//return strconv.ParseInt(fmt.Sprintf("%.f", float64(c)*100), 10, 64)
	// return decimal.NewFromFloat(float64(c)).Mul(decimal.NewFromInt(100)).Value()
	return decimal.NewFromFloat(float64(c)).Mul(decimal.NewFromInt(100)).IntPart(), nil
}

func (c *KindPrice) Scan(input interface{}) error {
	var tmp float64
	switch input.(type) {
	case int64:
		// *c = KindPrice(input.(int64)) / 100
		tmp, _ = decimal.NewFromInt(input.(int64)).Div(decimal.NewFromInt(100)).Float64()
		*c = KindPrice(tmp)
	case float64:
		// *c = KindPrice(input.(float64)) / 100
		tmp, _ = decimal.NewFromFloat(input.(float64)).Div(decimal.NewFromInt(100)).Float64()
		*c = KindPrice(tmp)
	case []byte:
		// v, _ := strconv.ParseFloat(string(input.([]byte)), 64)
		// *c = KindPrice(v) / 100
		dil, err := decimal.NewFromString(string(input.([]byte)))
		if err != nil {
			return err
		}
		tmp, _ = dil.Div(decimal.NewFromInt(100)).Float64()
		*c = KindPrice(tmp)
	case string:
		dil, err := decimal.NewFromString(input.(string))
		if err != nil {
			return err
		}
		tmp, _ = dil.Div(decimal.NewFromInt(100)).Float64()
		*c = KindPrice(tmp)
	}

	// if !ok {
	// 	return fmt.Errorf("lose precision. src[%+v] ret[%+v]", input, tmp)
	// }
	return nil
}

func (c KindPrice) Float64() float64 {
	// return utils.Round(float64(c), 2)
	return float64(c)
}

func (c KindPrice) String() string {
	// return fmt.Sprintf("%.2f", c.Float64())
	return decimal.NewFromFloat(c.Float64()).String()
}

func (c KindPrice) Add(v KindPrice) KindPrice {
	// return KindPrice((utils.Round(float64(c)*100, 2) + utils.Round(float64(v)*100, 2)) / 100)
	tmp, _ := decimal.NewFromFloat(c.Float64()).Add(decimal.NewFromFloat(v.Float64())).Float64()
	return KindPrice(tmp)

}
func (c KindPrice) Sub(v KindPrice) KindPrice {
	// return KindPrice((utils.Round(float64(c)*100, 2) - utils.Round(float64(v)*100, 2)) / 100)
	tmp, _ := decimal.NewFromFloat(c.Float64()).Sub(decimal.NewFromFloat(v.Float64())).Float64()
	return KindPrice(tmp)
}

//
//func (c KindPrice) Int32() int32 {
//	return int32(float64(c) * 1000)
//}

// KindSOC SOC
type KindSOC float64

func (c KindSOC) Value() (driver.Value, error) {
	return int64(float64(c) * 100), nil
}

func (c *KindSOC) Scan(input interface{}) error {
	switch input.(type) {
	case int64:
		*c = KindSOC(input.(int64)) / 100
	case float64:
		*c = KindSOC(input.(float64)) / 100
	case []byte:
		v, _ := strconv.ParseFloat(string(input.([]byte)), 64)
		*c = KindSOC(v) / 100
	}
	return nil
}

func (c KindSOC) Float64() float64 {
	return Round(float64(c), 2)
}

func (c KindSOC) Int32() int32 {
	return int32(float64(c) * 100)
}

// KindElectricity 电量
type KindElectricity float64

func (c KindElectricity) Value() (driver.Value, error) {
	// return int64(float64(c) * 100), nil
	return strconv.ParseInt(fmt.Sprintf("%.f", float64(c)*100), 10, 64)
}

func (c *KindElectricity) Scan(input interface{}) error {
	switch input.(type) {
	case int64:
		*c = KindElectricity(input.(int64)) / 100
	case float64:
		*c = KindElectricity(input.(float64)) / 100
	case []byte:
		v, _ := strconv.ParseFloat(string(input.([]byte)), 64)
		*c = KindElectricity(v) / 100
	}
	return nil
}

func (c KindElectricity) Float64() float64 {
	return Round(float64(c), 2)
}

func (c KindElectricity) Int32() int32 {
	return int32(float64(c) * 100)
}

type KindDashboardType int

type KindDashboardFolderType int

var (
	KindDashboardTypeEvse    KindDashboardType = 0
	KindDashboardTypeStation KindDashboardType = 1
	KindDashboardTypeManager KindDashboardType = 2
	KindDashboardTypeGroup   KindDashboardType = 3

	KindDashboardFolderTypeEvse    KindDashboardFolderType = 1
	KindDashboardFolderTypeManager KindDashboardFolderType = 2
	KindDashboardFolderTypeStation KindDashboardFolderType = 3
	KindDashboardFolderTypeGroup   KindDashboardFolderType = 4
)

//路径
type RootPath []uint64

func (c RootPath) Value() (driver.Value, error) {
	// return strings.Trim(strings.Replace(fmt.Sprint(c), " ", ",", -1), "[]"), nil
	s, _ := json.Marshal(c)
	return strings.Trim(string(s), "[]"), nil
	//return strings.Join(c, ","), nil
}

func (c *RootPath) Scan(input interface{}) error {
	return json.Unmarshal([]byte(fmt.Sprintf("[%s]", string(input.([]byte)))), c)
}

//
func (c RootPath) String() string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint([]uint64(c))), ","), "[]")
}

//type KindUint64s []uint64
//
//func (k *KindUint64s) Scan(input interface{}) error {
//	//_i := string(input.([]byte))
//	v := make([]uint64, 0)
//	for _, s := range strings.Split(string(input.([]byte)), ",") {
//		_i, _ := datasource.ParseUUID(s)
//		v = append(v, _i.Uint64())
//	}
//	*k = v
//	return nil
//}
//
//func (c KindUint64s) Value() (driver.Value, error) {
//	return strings.Trim(strings.Replace(fmt.Sprint(c), " ", ",", -1), "[]"), nil
//}
//func (c KindUint64s) String() string {
//	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint([]uint64(c))), ","), "[]")
//}

type KindStrings []string

func (k *KindStrings) Scan(input interface{}) error {
	*k = strings.Split(string(input.([]byte)), ",")
	return nil
}

type KindMapStringJSON map[string]string

//
//func (v KindMapStringJSON) MarshalJSON() ([]byte, error) {
//	return json.Marshal(v)
//}

func (v *KindMapStringJSON) UnmarshalJSON(data []byte) error {
	//jsonMap := make(map[string]interface{})
	var jsonMap map[string]string
	if len(data) <= 2 {
		return nil
	}
	s := strings.Replace(strings.Trim(string(data), `"`), `\"`, `"`, -1)
	err := json.Unmarshal([]byte(s), &jsonMap)
	if err != nil {
		return err
	}
	*v = KindMapStringJSON(jsonMap)
	return nil
}

//KindMapInterfaceJSON json格式的map对象
type KindMapInterfaceJSON map[string]interface{}

//
//func (v KindMapInterfaceJSON) MarshalJSON() ([]byte, error) {
//	return json.Marshal(v)
//}

func (v *KindMapInterfaceJSON) UnmarshalJSON(data []byte) error {
	//jsonMap := make(map[string]interface{})
	var jsonMap map[string]interface{}
	if len(data) <= 2 {
		return nil
	}
	var err error
	if string(data)[0] == '"' {
		s := strings.Replace(strings.Trim(string(data), `"`), `\"`, `"`, -1)
		err = json.Unmarshal([]byte(s), &jsonMap)
	} else {
		err = json.Unmarshal(data, &jsonMap)
	}

	if err != nil {
		return err
	}
	*v = KindMapInterfaceJSON(jsonMap)
	return nil
}

func (v *KindMapInterfaceJSON) Scan(input interface{}) error {
	//fmt.Println("vvvvvvvvvvvvvvvvv", input, string(input.([]byte)))
	//return json.Unmarshal(input.([]byte), v)
	return v.UnmarshalJSON(input.([]byte))
}

func (v KindMapInterfaceJSON) Value() (driver.Value, error) {
	return json.Marshal(v)
}

type KindGeoPoint struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type Msg struct {
	Title     string `json:"title"`
	Content   string `json:"content"`
	FromEmail string `json:"from_email,omitempty"`
	FromName  string `json:"from_name"`
	ToEmail   string `json:"to_email,omitempty"`
	BCCEmail  string `json:"cc_email,omitempty"`
}

func Round(f float64, n int) float64 {
	pow10n := math.Pow10(n)
	return math.Trunc((f+0.5/pow10n)*pow10n) / pow10n
}
