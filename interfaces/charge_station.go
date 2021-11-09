package interfaces

type KindChargeStationState int

const (
	KindChargeStationStateOnline  KindChargeStationState = 1
	KindChargeStationStateOffline KindChargeStationState = 2
	KindChargeStationStateBan     KindChargeStationState = 3
)

type ChargeStation interface {
	Core
	// 状态
	State() KindChargeStationState
	// 设置状态
	SetState(KindChargeStationState)
	// 设备
	Evses() []Evse
	// 添加设备
	AddEvses(evses ...Evse)
	// 所有的枪
	Connectors() []Connector

	Registered() bool

	Register()
}

type chargeStation struct {
	id       uint64
	sn       string
	state    KindChargeStationState
	evses    []Evse
	register bool
}

var _ ChargeStation = &chargeStation{}

func NewDefaultChargeStation(sn string, register bool, coreID uint64) ChargeStation {
	return &chargeStation{
		id:       coreID,
		sn:       sn,
		state:    KindChargeStationStateOffline,
		evses:    make([]Evse, 0),
		register: register,
	}
}

// 序列号
func (c *chargeStation) SN() string {
	return c.sn
}

// 状态
func (c *chargeStation) State() KindChargeStationState {
	return c.state
}

// 设置状态
func (c *chargeStation) SetState(state KindChargeStationState) {
	c.state = state
}

// 设备
func (c *chargeStation) Evses() []Evse {
	return c.evses
}

// 添加设备
func (c *chargeStation) AddEvses(evses ...Evse) {
	c.evses = append(c.evses, evses...)
}

// 平台id
func (c *chargeStation) CoreID() uint64 {
	return c.id
}

// 设置平台id
func (c *chargeStation) SetCoreID(coreID uint64) {
	c.id = coreID
}

// 所有的枪(记住是无法修改的)
func (c *chargeStation) Connectors() []Connector {
	if len(c.evses) == 0 {
		return nil
	}
	connectors := make([]Connector, 0)
	for _, evse := range c.evses {
		connectors = append(connectors, evse.Connectors()...)
	}
	return connectors
}

func (c *chargeStation) Registered() bool {
	return c.register
}

func (c *chargeStation) Register() {
	c.register = true
}
