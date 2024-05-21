package astar

type Interface interface {
	GetWidth() int32
	GetHeight() int32
	CanMoveIndex(grid int32) bool
	GetPosIndex(x, y int32) int32
	GetIndex2Pos(grid int32) (x, y int32)
}
