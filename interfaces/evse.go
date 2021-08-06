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
	// 连接在哪个pod上
	Controllee() string
	// 所有的枪
	Connectors() []Connector
	// 状态
	State() KindEvseState
	// 设置状态
	SetState(KindEvseState)
	// 在平台的ID
	CoreID() uint64
	// 设置平台ID
	SetCoreID(uint64)
	// 是否需要注册
	NeedRegister() bool
	// 注册
	Register()
}
