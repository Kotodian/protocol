package coregw

import (
	pCharger "github.com/Kotodian/protocol/golang/hardware/charger"
)

type KindOrderPayState int32

const (
	KindOrderPayStateFail    KindOrderPayState = -10 //支付失败
	KindOrderPayStateInit    KindOrderPayState = 0   //初始化状态,刷卡鉴权卡订单或其他后付费的订单使用
	KindOrderPayStatePending KindOrderPayState = 10  //待支付状态，先付费创建订单
	KindOrderPayStateSuccess KindOrderPayState = 100 //已支付，钱包卡, 白名单
)

type KindOrderRefundState int32

const (
	KindOrderRefundStateFail    KindOrderRefundState = -10 //退款失败
	KindOrderRefundStateInit    KindOrderRefundState = 0   //初始化状态
	KindOrderRefundStatePending KindOrderRefundState = 10  //待处理状态
	KindOrderRefundStateSuccess KindOrderRefundState = 100 //已退款
)

type KindOrderBackPayState int32

const (
	KindOrderBackPayStateInit    KindOrderBackPayState = 0   //初始化状态
	KindOrderBackPayStatePending KindOrderBackPayState = 10  //待补缴
	KindOrderBackPayStateSuccess KindOrderBackPayState = 100 //已退款
)

type KindOrderSettlementState int32

const (
	KindOrderSettlementStateFail    KindOrderSettlementState = -10 //结算失败
	KindOrderSettlementStateInit    KindOrderSettlementState = 0   //初始化，重试结算的订单条件, state=100 and settlement_state=0
	KindOrderSettlementStatePending KindOrderSettlementState = 10  //待结算
	KindOrderSettlementStateSuccess KindOrderSettlementState = 100 //结算成功
)

type KindOrderState int32

const (
	KindOrderStateUnknownSuspend              KindOrderState = -199 //枪头未知原因导致订单挂起
	KindOrderStateConnectorDiffSessionSuspend KindOrderState = -152 //枪头的订单会话不一致导致挂起
	KindOrderStateConnectorUnavailableSuspend KindOrderState = -151 //枪头不可用导致订单挂起
	KindOrderStateConnectorAvailableSuspend   KindOrderState = -150 //枪头空闲可用导致订单挂起
	KindOrderStateConnectorFaultSuspend       KindOrderState = -101 //枪头故障导致订单挂起
	KindOrderStateConnectorOfflineSuspend     KindOrderState = -100 //枪头离线导致订单挂起
	KindOrderStateCostConflict                KindOrderState = -60  //费用冲突
	//KindOrderStateSettlementErrorSuspend      KindOrderState = -50  //结算错误，订单挂起
	KindOrderStateRemoteStopTimeout      KindOrderState = -30 //结束充电超时
	KindOrderStateConnectTimeout         KindOrderState = -20 //连接超时（3天内退款）
	KindOrderStateRemoteStartTimeout     KindOrderState = -10 //启动充电超时（马上退款）
	KindOrderStateRemoteStartConfTimeout KindOrderState = -7  //启动充电回复超时（原是马上退款，先修改成3天内退款）
	KindOrderStateRemoteStartFail        KindOrderState = -6  //启动充电失败，后台导致的（马上退款）
	KindOrderStateRemoteStartReject      KindOrderState = -5  //设备拒绝启动充电（马上退款）
	KindOrderStatePayTimeout             KindOrderState = -3  //支付超时
	KindOrderStatePayFail                KindOrderState = -1  //支付失败
	KindOrderStatePending                KindOrderState = 0   //订单等待处理
	KindOrderStateRemoteStarted          KindOrderState = 10  //启动充电成功，等待连接
	KindOrderStateCharging               KindOrderState = 20  //连接成功，充电中
	KindOrderStateRemoteStopped          KindOrderState = 30  //结束充电请求成功
	KindOrderStateSuspendedEvse          KindOrderState = 40  // 暂停充电,已连接未输出
	KindOrderStateStopped                KindOrderState = 50  //结束充电成功，等待计费结果
	KindOrderStateFinish                 KindOrderState = 100 //已完成
)

func (e KindOrderState) IsInProcessing() bool {
	// return e >= KindOrderStatePendding && e < KindOrderStateFinish || e == KindOrderStateRemoteStartTimeout
	return e >= KindOrderStatePending && e < KindOrderStateStopped
}

func (e KindOrderState) IsStopped() bool {
	return e >= KindOrderStateStopped
}

func (e KindOrderState) IsFinish() bool {
	return e == KindOrderStateFinish
}

func (e KindOrderState) IsStarting() bool {
	return e == KindOrderStateRemoteStarted || e == KindOrderStatePending
}

func (e KindOrderState) IsFailure() bool {
	return e < KindOrderStatePending
}

func (e KindOrderState) IsCharging() bool {
	return e == KindOrderStateCharging
}

func (e KindOrderState) IsStoping() bool {
	return e == KindOrderStateRemoteStopped
}

type KindConnectorState pCharger.ChargerStatus

func (cs KindConnectorState) Int32() int32 {
	return int32(cs)
}

func (cs KindConnectorState) CanCharge() bool {
	_cs := pCharger.ChargerStatus(cs)
	if _cs != pCharger.ChargerStatus_CHS_Available &&
		_cs != pCharger.ChargerStatus_CHS_Preparing {
		return false
	}
	return true
}

func (cs KindConnectorState) IsReleaseOrderState() bool {
	return pCharger.ChargerStatus(cs) == pCharger.ChargerStatus_CHS_Available
}

func (cs KindConnectorState) IsChargingState() bool {
	_cs := pCharger.ChargerStatus(cs)
	return _cs == pCharger.ChargerStatus_CHS_Charging ||
		_cs == pCharger.ChargerStatus_CHS_SuspendedEV ||
		_cs == pCharger.ChargerStatus_CHS_SuspendedEVSE
}

func (cs KindConnectorState) Eq(s pCharger.ChargerStatus) bool {
	return pCharger.ChargerStatus(cs) == s
}

//是否能预约充电
func (cs *KindConnectorState) CanReserve() bool {
	return false
}

func (cs KindConnectorState) Desc() string {
	return pCharger.ChargerStatus_name[int32(cs)]
}
