package interfaces

type Core interface {
	CoreID() uint64
	SetCoreID(coreID uint64)
}
