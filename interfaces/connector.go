package interfaces

type KindConnectorChargingState int
type KindConnectorState int

const (
	KindConnectorStateUnavailable   KindConnectorState = 0 // 不可用
	KindConnectorStateAvailable     KindConnectorState = 1 // 空闲可用
	KindConnectorStateOccupied      KindConnectorState = 2 // 占用
	KindConnectorStateReserved      KindConnectorState = 3 // 预约
	KindConnectorStateFaulted       KindConnectorState = 4 // 故障
	KindConnectorStatePreparing     KindConnectorState = 5 // 准备中
	KindConnectorStateCharging      KindConnectorState = 6 // 充电中
	KindConnectorStateSuspendedEV   KindConnectorState = 7 // 车端挂起，不输入电能
	KindConnectorStateSuspendedEVSE KindConnectorState = 8 // 桩端挂起，不输出电能
	KindConnectorStateFinishing     KindConnectorState = 9 // 结束中
)
const (
	KindConnectorChargingStateIdle          KindConnectorChargingState = 0 //空闲
	KindConnectorChargingStateEVConnected   KindConnectorChargingState = 1 //插枪
	KindConnectorChargingStateCharging      KindConnectorChargingState = 2 //充电中
	KindConnectorChargingStateSuspendedEV   KindConnectorChargingState = 3 //车端挂起，不输入电能
	KindConnectorChargingStateSuspendedEVSE KindConnectorChargingState = 4 //桩端挂起，不输出电能
)

var _ Connector = &connector{}

type Connectors []Connector

func ConnectorNumbers(conns Connectors) []int32 {
	data := make([]int32, 0)
	for _, conn := range conns {
		data = append(data, int32(conn.Num()))
	}
	return data
}

type Connector interface {
	Core
	// 序号
	Num() uint32
	// 枪的状态
	State() KindConnectorState
	// 设置枪的状态
	SetState(KindConnectorState)
	// 枪的充电状态
	ChargingState() KindConnectorChargingState
	// 设置枪的充电状态
	SetChargingState(KindConnectorChargingState)
}

type connector struct {
	id            uint64
	num           uint32
	sn            string
	state         KindConnectorState
	chargingState KindConnectorChargingState
}

func NewDefaultConnector(sn string, num uint32) Connector {
	return &connector{
		sn:  sn,
		num: num,
	}
}

// 枪的id
func (c *connector) CoreID() uint64 {
	return c.id
}

// 设置枪的ID
func (c *connector) SetCoreID(coreID uint64) {
	c.id = coreID
}

// 枪属于哪个桩
func (c *connector) SN() string {
	return c.sn
}

// 序号
func (c *connector) Num() uint32 {
	return c.num
}

// 枪的状态
func (c *connector) State() KindConnectorState {
	return c.state
}

// 设置枪的状态
func (c *connector) SetState(state KindConnectorState) {
	c.state = state
}

// 枪的充电状态
func (c *connector) ChargingState() KindConnectorChargingState {
	return c.chargingState
}

// 设置枪的充电状态
func (c *connector) SetChargingState(chargingState KindConnectorChargingState) {
	c.chargingState = chargingState
}

//
//type ConnectorFSM struct {
//	c Connector
//	fsm.MachineAbs
//}
//
//func NewConnectorFSM(c Connector) fsm.IMachine {
//	return &ConnectorFSM{
//		c: c,
//	}
//}
//
//func (c *ConnectorFSM) SetState(ctx context.Context, s fsm.State) error {
//	c.c.SetState(s.(KindConnectorState))
//	return nil
//}
//
//func (c *ConnectorFSM) GetState() fsm.State {
//	return c.c.State()
//}
//
//func (c *ConnectorFSM) OnInitWithMachine(f *fsm.FSM) {
//
//}
//
//type ConnectorChargeFSM struct {
//	c Connector
//	fsm.MachineAbs
//}
//
//func NewConnectorChargeFSM(c Connector) fsm.IMachine {
//	return &ConnectorChargeFSM{
//		c: c,
//	}
//}
//
//func (c *ConnectorChargeFSM) SetState(ctx context.Context, s fsm.State) error {
//	c.c.SetChargingState(s.(KindConnectorChargingState))
//	return nil
//}
//
//func (c *ConnectorChargeFSM) GetState() fsm.State {
//	return c.c.ChargingState()
//}
//
//func (c *ConnectorChargeFSM) OnInitWithMachine(f *fsm.FSM) {
//}
//
//var ConnectorChargingStateFSM *fsm.FSM
//var ConnectorStateFSM *fsm.FSM
//
//func init() {
//	{
//		ConnectorStateFSM = fsm.NewFSM()
//		//Special 设置特殊状态
//		ConnectorStateFSM.Special(KindConnectorStateFaulted)
//		ConnectorStateFSM.Special(KindConnectorStateUnavailable)
//
//		ConnectorStateFSM.Special(KindConnectorStateAvailable)
//		ConnectorStateFSM.SetStateFuncs(KindConnectorStateAvailable,
//			nil,
//			//onEnterFunc
//			func(object fsm.IMachine, ctx context.Context, args ...interface{}) (err error) {
//				obj := object.(*ConnectorFSM)
//				obj.c.SetChargingState(KindConnectorChargingStateIdle)
//				return nil
//			})
//
//		ConnectorStateFSM.Special(KindConnectorStateReserved)
//
//		//从充电完成跳转到占用，就不更新状态了
//		ConnectorStateFSM.From(KindConnectorStateAvailable).To(KindConnectorStateOccupied).
//			Then(func(object fsm.IMachine, ctx context.Context, state fsm.State, state2 fsm.State, i ...interface{}) error {
//				obj := object.(*ConnectorFSM)
//				if obj.c.ChargingState() == KindConnectorChargingStateIdle {
//					obj.c.SetState(KindConnectorStateReserved)
//					obj.SetSkip(true)
//				}
//				return nil
//			})
//
//		ConnectorStateFSM.Special(KindConnectorStateOccupied)
//		ConnectorStateFSM.SetStateFuncs(KindConnectorStateOccupied,
//			nil,
//			func(object fsm.IMachine, ctx context.Context, args ...interface{}) (err error) {
//				obj := object.(*ConnectorFSM)
//				if obj.GetIgnore() == true || obj.GetSkip() {
//					return nil
//				}
//				//obj.connector.State = models.KindConnectorStateEVConnected
//				switch obj.c.ChargingState() {
//				case KindConnectorChargingStateCharging:
//					obj.c.SetState(KindConnectorStateCharging)
//					obj.SetSkip(true)
//				case KindConnectorChargingStateEVConnected, KindConnectorChargingStateIdle:
//					if obj.c.State() != KindConnectorStateReserved {
//						obj.c.SetState(KindConnectorStateEVConnected)
//						obj.SetSkip(true)
//					} else {
//						obj.SetIgnore(true)
//					}
//				case KindConnectorChargingStateSuspendedEVSE, KindConnectorChargingStateSuspendedEV:
//					//obj.connector.State = models.KindConnectorStateCharging
//					obj.SetIgnore(true)
//				}
//				return nil
//			})
//		//从充电完成跳转到占用，就不更新状态了
//		ConnectorStateFSM.From(KindConnectorStateFinished).To(KindConnectorStateOccupied).
//			Then(func(object fsm.IMachine, ctx context.Context, state fsm.State, state2 fsm.State, i ...interface{}) error {
//				obj := object.(*ConnectorFSM)
//				obj.SetSkip(true)
//				obj.c.SetState(KindConnectorStateFinished)
//				return nil
//			})
//	}
//	{
//		ConnectorChargingStateFSM = fsm.NewFSM()
//		{
//			ConnectorChargingStateFSM.Special(KindConnectorChargingStateIdle)
//			ConnectorChargingStateFSM.SetStateFuncs(KindConnectorChargingStateIdle,
//				nil,
//				func(object fsm.IMachine, ctx context.Context, args ...interface{}) (err error) {
//					obj := object.(*ConnectorChargeFSM)
//					if obj.GetIgnore() {
//						return
//					} else if obj.c.State() == KindConnectorStateOccupied ||
//						obj.c.State() == KindConnectorStateCharging ||
//						obj.c.State() == KindConnectorStateSuspendedEV ||
//						obj.c.State() == KindConnectorStateSuspendedEVSE {
//						obj.c.SetState(KindConnectorStateEVConnected)
//						obj.c.SetChargingState(KindConnectorChargingStateEVConnected)
//					}
//					return
//				})
//			ConnectorChargingStateFSM.Special(KindConnectorChargingStateEVConnected)
//			ConnectorChargingStateFSM.SetStateFuncs(KindConnectorChargingStateEVConnected,
//				nil,
//				func(object fsm.IMachine, ctx context.Context, args ...interface{}) (err error) {
//					obj := object.(*ConnectorChargeFSM)
//					if obj.GetIgnore() {
//						return
//					}
//					if obj.c.State() == KindConnectorStateAvailable ||
//						(obj.c.State() != KindConnectorStateReserved && obj.c.ChargingState() == KindConnectorChargingStateIdle) {
//						obj.c.SetState(KindConnectorStateEVConnected)
//					}
//					return
//				})
//
//			ConnectorChargingStateFSM.Special(KindConnectorChargingStateSuspendedEV)
//			ConnectorChargingStateFSM.SetStateFuncs(KindConnectorChargingStateSuspendedEV,
//				nil,
//				func(object fsm.IMachine, ctx context.Context, args ...interface{}) (err error) {
//					obj := object.(*ConnectorChargeFSM)
//					obj.c.SetState(KindConnectorStateSuspendedEV)
//					return nil
//				})
//
//			ConnectorChargingStateFSM.Special(KindConnectorChargingStateSuspendedEVSE)
//			ConnectorChargingStateFSM.SetStateFuncs(KindConnectorChargingStateSuspendedEVSE,
//				nil,
//				func(object fsm.IMachine, ctx context.Context, args ...interface{}) (err error) {
//					obj := object.(*ConnectorChargeFSM)
//					obj.c.SetState(KindConnectorStateSuspendedEVSE)
//					return nil
//				})
//
//			ConnectorChargingStateFSM.Special(KindConnectorChargingStateCharging)
//			ConnectorChargingStateFSM.SetStateFuncs(KindConnectorChargingStateCharging,
//				nil,
//				func(object fsm.IMachine, ctx context.Context, args ...interface{}) (err error) {
//					obj := object.(*ConnectorChargeFSM)
//					obj.c.SetState(KindConnectorStateCharging)
//					return nil
//				})
//		}
//	}
//}
