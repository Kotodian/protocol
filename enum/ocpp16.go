package enum

type OCPP16ConnectorStatus int

const (
	OCPP16ConnectorStatusUnavailable   OCPP16ConnectorStatus = 0
	OCPP16ConnectorStatusAvailable     OCPP16ConnectorStatus = 1
	OCPP16ConnectorStatusPreparing     OCPP16ConnectorStatus = 2
	OCPP16ConnectorStatusCharging      OCPP16ConnectorStatus = 3
	OCPP16ConnectorStatusSuspendedEV   OCPP16ConnectorStatus = 4
	OCPP16ConnectorStatusSuspendedEVSE OCPP16ConnectorStatus = 5
	OCPP16ConnectorStatusFinishing     OCPP16ConnectorStatus = 6
	OCPP16ConnectorStatusFaulted       OCPP16ConnectorStatus = 7
	OCPP16ConnectorStatusReserved      OCPP16ConnectorStatus = 8
)
