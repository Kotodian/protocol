package interfaces

type Core interface {
	// 序列号
	SN() string
	// 平台id
	CoreID() uint64
	// 设置平台id
	SetCoreID(coreID uint64)
}
