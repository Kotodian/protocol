package charger

import "strings"

func (x *APDU) IsRequest() bool {
	return x.MessageId>>7 == 0
}

func (x *APDU) Action() string {
	action := strings.TrimPrefix(MessageID_name[int32(x.MessageId)], "ID_")
	if x.IsRequest() {
		action = strings.TrimSuffix(action, "Req")
	} else {
		action = strings.TrimSuffix(action, "Conf")
	}
	return action
}

func (x *Tariff) Sum() float64 {
	return x.GetSharp() + x.GetValley() + x.GetFlat() + x.GetPeak()
}
