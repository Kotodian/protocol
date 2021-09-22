package interfaces

type KindOrderState int

const (
	KindOrderStateTimeout         KindOrderState = iota - 1 // 订单启动超时
	KindOrderStateEnded                                     // 订单正常结束
	KindOrderStateUnexpectedEnded                           // 订单意外结束
	KindOrderStateStarted                                   // 订单启动了 桩还没上报
	KindOrderStateCharging                                  // 订单正在充电中
	KindOrderStateChargingSuspend                           // 订单在充电 但是枪未输出
	KindOrderStateFaultedSuspend                            // 订单在充电中 但是枪上报NotifyEvent
	KindOrderStateOfflineSuspend                            // 订单在充电中 但是桩突然离线
)

type Transaction interface {
	// 订单的充电事件ID
	TransactionID() string
	// 属于哪个桩
	SN() string
	// 属于哪个枪
	ConnectorNum() uint32
	// 启动电量
	StartElectricity() int32
	SetStartElectricity(electricity int32)
	// 计费模板
	TraffID() uint64
	// 远程启动的ID
	RemoteStartID() int32
	// CoreID 平台ID
	CoreID() uint64
}

var _ Transaction = &transaction{}

type transaction struct {
	id              uint64
	transactionID   string
	sn              string
	connectorNum    uint32
	startElectricty int32
	traiffID        uint64
	remoteStartID   int32
}

func NewDefaultTransaction(transactionID string, sn string, connectorNum uint32, remoteStartID ...int32) Transaction {
	var _remoteStartID int32
	if len(remoteStartID) > 0 {
		_remoteStartID = remoteStartID[0]
	}
	return &transaction{
		sn:            sn,
		transactionID: transactionID,
		connectorNum:  connectorNum,
		remoteStartID: _remoteStartID,
	}
}

// 订单的充电事件ID
func (t *transaction) TransactionID() string {
	return t.transactionID
}

// 属于哪个桩
func (t *transaction) SN() string {
	return t.sn
}

// 属于哪个枪
func (t *transaction) ConnectorNum() uint32 {
	return t.connectorNum
}

// 启动电量
func (t *transaction) StartElectricity() int32 {
	return t.startElectricty
}

// 计费模板
func (t *transaction) TraffID() uint64 {
	return t.traiffID
}

// 远程启动的ID
func (t *transaction) RemoteStartID() int32 {
	return t.remoteStartID
}

// CoreID 平台ID
func (t *transaction) CoreID() uint64 {
	return t.id
}

func (t *transaction) SetStartElectricity(electricity int32) {
	t.startElectricty = electricity
}
