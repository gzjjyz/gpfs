package astar

var aStarMapMgr = make(map[int32]*aStart)

func InitAStarMap(id int32, conf Interface) {
	aStarMapMgr[id] = &aStart{
		mapData: conf,
	}
}

func FindRoads(id, x1, y1, x2, y2 int32) []int32 {
	aStarMap, ok := aStarMapMgr[id]
	if !ok {
		return nil
	}
	start := aStarMap.mapData.GetPosIndex(x1, y1)
	end := aStarMap.mapData.GetPosIndex(x2, y2)

	return aStarMap.findRoads(start, end)
}

func Grid2Pos(id int32, grid int32) (x, y int32) {
	aStarMap, ok := aStarMapMgr[id]
	if !ok {
		return
	}
	return aStarMap.mapData.GetIndex2Pos(grid)
}
