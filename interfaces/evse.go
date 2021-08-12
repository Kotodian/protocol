package interfaces

type KindEvseState int

const (
	KindEvseStateOnline  KindEvseState = 1
	KindEvseStateOffline KindEvseState = 2
	KindEvseStateBan     KindEvseState = 3
)

type Evse interface {
	// 序列号
	SN() string
	// 所有的枪
	Connectors() []Connector
	// 设置枪
	AddConnectors(...Connector)
	// 状态
	State() KindEvseState
	// 设置状态
	SetState(KindEvseState)
	// 在平台的ID
	CoreID() uint64
	// 设置平台ID
	SetCoreID(uint64)
}

func NewDefaultEvse(sn string) Evse {
	return &evse{
		sn:         sn,
		state:      KindEvseStateOffline,
		register:   false,
		id:         0,
		connectors: make([]Connector, 0),
	}
}

var _ Evse = &evse{}

type evse struct {
	id         uint64
	sn         string
	connectors []Connector
	state      KindEvseState
	register   bool
}

// 序列号
func (e *evse) SN() string {
	return e.sn
}

// 所有的枪
func (e *evse) Connectors() []Connector {
	return e.connectors
}

// 状态
func (e *evse) State() KindEvseState {
	return e.state
}

// 设置状态
func (e *evse) SetState(state KindEvseState) {
	e.state = state
}

// 在平台的ID
func (e *evse) CoreID() uint64 {
	return e.id
}

// 设置平台ID
func (e *evse) SetCoreID(coreID uint64) {
	e.id = coreID
}

// 是否需要注册
func (e *evse) NeedRegister() bool {
	return e.register
}

// 注册
func (e *evse) Register() {
	e.register = true
}

// 设置枪
func (e *evse) AddConnectors(connectors ...Connector) {
	e.connectors = append(e.connectors, connectors...)
}
