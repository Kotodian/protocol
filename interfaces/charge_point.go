package interfaces

type chargePoint struct {
	id       uint64
	sn       string
	state    KindChargeStationState
	register bool
}

func NewChargePoint(sn string, register bool, coreID uint64) ChargeStation {
	return &chargePoint{
		id:       coreID,
		sn:       sn,
		state:    KindChargeStationStateOffline,
		register: register,
	}
}

// 序列号
func (c *chargePoint) SN() string {
	return c.sn
}

func (c *chargePoint) SetSN(sn string) {
	c.sn = sn
}

// 状态
func (c *chargePoint) State() KindChargeStationState {
	return c.state
}

// 设置状态
func (c *chargePoint) SetState(state KindChargeStationState) {
	c.state = state
}

// 设备
func (c *chargePoint) Evses() []Evse {
	return nil
}

// 添加设备
func (c *chargePoint) AddEvses(evses ...Evse) {

}

// 平台id
func (c *chargePoint) CoreID() uint64 {
	return c.id
}

// 设置平台id
func (c *chargePoint) SetCoreID(coreID uint64) {
	c.id = coreID
}

// 所有的枪(记住是无法修改的)
func (c *chargePoint) Connectors() []Connector {
	return nil
}

func (c *chargePoint) Registered() bool {
	return c.register
}

func (c *chargePoint) Register() {
	c.register = true
}
