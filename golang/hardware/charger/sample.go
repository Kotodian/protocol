package charger

type Sample interface {
	SetElectricity(value float64)
	SetPower(value float64)
	SetVoltageA(value float64)
	SetVoltageB(value float64)
	SetVoltageC(value float64)
	SetCurrentA(value float64)
	SetCurrentB(value float64)
	SetCurrentC(value float64)
	SetSOC(value uint32)
}

func (x *StartTransactionReq) SetElectricity(value float64) {
	x.MeterStart = value
}

func (x *StartTransactionReq) SetPower(value float64) {
	x.Power = value
}

func (x *StartTransactionReq) SetVoltageA(value float64) {
	x.VoltageA = value
}

func (x *StartTransactionReq) SetVoltageB(value float64) {
	x.VoltageB = value
}

func (x *StartTransactionReq) SetVoltageC(value float64) {
	x.VoltageC = value
}

func (x *StartTransactionReq) SetCurrentA(value float64) {
	x.CurrentA = value
}

func (x *StartTransactionReq) SetCurrentB(value float64) {
	x.CurrentB = value
}

func (x *StartTransactionReq) SetCurrentC(value float64) {
	x.CurrentC = value
}

func (x *StartTransactionReq) SetSOC(soc uint32) {
	x.Soc = soc
}

func (x *StopTransactionReq) SetElectricity(value float64) {
	x.MeterStop = value
}

func (x *StopTransactionReq) SetPower(value float64) {
	x.Power = value
}

func (x *StopTransactionReq) SetVoltageA(value float64) {
	x.VoltageA = value
}

func (x *StopTransactionReq) SetVoltageB(value float64) {
	x.VoltageB = value
}

func (x *StopTransactionReq) SetVoltageC(value float64) {
	x.VoltageC = value
}

func (x *StopTransactionReq) SetCurrentA(value float64) {
	x.CurrentA = value
}

func (x *StopTransactionReq) SetCurrentB(value float64) {
	x.CurrentB = value
}

func (x *StopTransactionReq) SetCurrentC(value float64) {
	x.CurrentC = value
}

func (x *StopTransactionReq) SetSOC(soc uint32) {
	x.Soc = soc
}

func (x *ChargingInfoReq) SetElectricity(value float64) {
	x.Electricity = value
}

func (x *ChargingInfoReq) SetPower(value float64) {
	x.Power = value
}

func (x *ChargingInfoReq) SetVoltageA(value float64) {
	x.VoltageA = value
}

func (x *ChargingInfoReq) SetVoltageB(value float64) {
	x.VoltageB = value
}

func (x *ChargingInfoReq) SetVoltageC(value float64) {
	x.VoltageC = value
}

func (x *ChargingInfoReq) SetCurrentA(value float64) {
	x.CurrentA = value
}

func (x *ChargingInfoReq) SetCurrentB(value float64) {
	x.CurrentB = value
}

func (x *ChargingInfoReq) SetCurrentC(value float64) {
	x.CurrentC = value
}

func (x *ChargingInfoReq) SetSOC(soc uint32) {
	x.Soc = soc
}

func (x *MeterValuesReq) SetElectricity(value float64) {
	x.Electricity = value
}

func (x *MeterValuesReq) SetPower(value float64) {
	x.Power = value
}

func (x *MeterValuesReq) SetCurrentA(value float64) {
	x.CurrentA = value
}

func (x *MeterValuesReq) SetCurrentB(value float64) {
	x.CurrentB = value
}

func (x *MeterValuesReq) SetCurrentC(value float64) {
	x.CurrentC = value
}

func (x *MeterValuesReq) SetVoltageA(value float64) {
	x.VoltageA = value
}

func (x *MeterValuesReq) SetVoltageB(value float64) {
	x.VoltageB = value
}

func (x *MeterValuesReq) SetVoltageC(value float64) {
	x.VoltageC = value
}

func (x *MeterValuesReq) SetSOC(soc uint32) {
	x.Soc = soc
}
