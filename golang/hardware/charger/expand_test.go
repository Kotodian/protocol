package charger

import "testing"

func TestAPDU_Action(t *testing.T) {
	apdu := &APDU{
		MessageId: MessageID_ID_DataTransferConf,
	}

	t.Log(apdu.Action())
}
