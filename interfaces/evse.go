package interfaces

type Evse interface {
	Core
	// 所有的枪
	Connectors() []Connector
	// 设置枪
	AddConnectors(...Connector)
	// 设备序号
	Num() uint32
}

func NewDefaultEvse(num uint32) Evse {
	return &evse{
		num:        num,
		register:   false,
		id:         0,
		connectors: make([]Connector, 0),
	}
}

var _ Evse = &evse{}

type evse struct {
	id         uint64
	num        uint32
	connectors []Connector
	register   bool
}

// 所有的枪
func (e *evse) Connectors() []Connector {
	return e.connectors
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

// 序号
func (e *evse) Num() uint32 {
	return e.num
}
