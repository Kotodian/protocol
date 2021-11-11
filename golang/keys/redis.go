package keys

import (
	"strings"
)

func Equipment(equipmentID string) string {
	return strings.Join([]string{"ac", "online", equipmentID}, ":")
}

func Connector(sn, evse, connector string) string {
	return strings.Join([]string{"sn", sn, "evse", evse, "cs", connector}, ":")
}

func LockConnector(sn, evse, connector string) string {
	return "lock:" + Connector(sn, evse, connector)
}

func OrderCalculate(orderID string) string {
	return strings.Join([]string{"order", "calculate", orderID}, ":")
}
