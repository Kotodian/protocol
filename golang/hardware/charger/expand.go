package charger

func (x *APDU) IsRequest() bool {
	return x.MessageId>>7 == 0
}
