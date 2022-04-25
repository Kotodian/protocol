package charger

import (
	"strings"
)

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

type ChargingExpand interface {
	SetConnector(id string)
	SetEvse(id string)
	SetTariff(tariff *Tariff)
	SetChargingProfile(chargingProfile *CurrentChargingProfile)
	SetToken(token *IdToken)
}

func (x *StartTransactionReq) SetTariff(tariff *Tariff) {
	x.Tariff = tariff
}

func (x *StartTransactionReq) SetChargingProfile(chargingProfile *CurrentChargingProfile) {
	x.Profile = chargingProfile
}

func (x *StartTransactionReq) SetConnector(id string) {
	x.ConnectorId = id
}

func (x *StartTransactionReq) SetEvse(id string) {
	x.EvseId = id
}

func (x *StartTransactionReq) SetToken(token *IdToken) {
	x.IdData = token
}

func (x *StopTransactionReq) SetTariff(tariff *Tariff) {
	x.Tariff = tariff
}

func (x *StopTransactionReq) SetChargingProfile(chargingProfile *CurrentChargingProfile) {
	x.Profile = chargingProfile
}

func (x *StopTransactionReq) SetConnector(id string) {
	x.ConnectorId = id
}

func (x *StopTransactionReq) SetEvse(id string) {
	x.EvseId = id
}

func (x *StopTransactionReq) SetToken(token *IdToken) {
	x.IdData = token
}

func (x *ChargingInfoReq) SetTariff(tariff *Tariff) {
	x.Tariff = tariff
}

func (x *ChargingInfoReq) SetChargingProfile(chargingProfile *CurrentChargingProfile) {
	x.Profile = chargingProfile
}

func (x *ChargingInfoReq) SetConnector(id string) {
	x.ConnectorId = id
}

func (x *ChargingInfoReq) SetEvse(id string) {
	x.EvseId = id
}

func (x *ChargingInfoReq) SetToken(token *IdToken) {
	x.IdData = token
}

func (x *TransactionReq) SetTariff(tariff *Tariff) {
	x.Tariff = tariff
}

func (x *TransactionReq) SetConnector(id string) {
	x.ConnectorId = id
}

func (x *TransactionReq) SetEvse(id string) {
	x.EvseId = id
}

func (x *TransactionReq) SetToken(token *IdToken) {
	x.IdData = token
}

func (x *TransactionReq) SetChargingProfile(chargingProfile *CurrentChargingProfile) {
	x.Profile = chargingProfile
}
