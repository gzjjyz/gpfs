package astar

import (
	"container/heap"
	"math"
)

const MaxAStartRoads = 5000
const MaxFindCount = 100000

var (
	neighbors  = make([]int32, 0, 8)
	roadsSlice = make([]int32, 0, MaxAStartRoads)
)

type aStart struct {
	mapData Interface
}

type aStarParam struct {
	grid   int32
	parent *aStarParam
	gValue int32
	hValue int32
}

func (s *aStart) findRoads(start, end int32) []int32 {
	if start == end {
		return nil
	}

	mapData := s.mapData

	if !mapData.CanMoveIndex(start) {
		return nil
	}
	if !mapData.CanMoveIndex(end) {
		return nil
	}

	roadsSlice = roadsSlice[:0]

	closeList := make(map[int32]struct{})

	openList := &aStarParamHeap{}
	heap.Init(openList)

	closeList[start] = struct{}{}

	heap.Push(openList, &aStarParam{grid: start})

	width := mapData.GetWidth()
	var count int32 = 0
	for openList.Len() > 0 {
		if count >= MaxFindCount {
			break
		}
		count++
		current := heap.Pop(openList).(*aStarParam)
		if current.grid == end {
			tail := current
			for {
				if tail.parent == nil {
					break
				}
				roadsSlice = append(roadsSlice, tail.grid)
				tail = tail.parent
			}
			break
		}
		closeList[current.grid] = struct{}{}

		for _, neighbor := range s.findNeighbor(current.grid) {
			if !mapData.CanMoveIndex(neighbor) {
				continue
			}

			if _, exists := closeList[neighbor]; exists {
				continue
			}

			index, param := openList.index(neighbor)

			g := s.calculateGValue(current.grid, neighbor, width)
			if index > 0 {
				if current.gValue+g < param.gValue {
					param.parent = current
					param.gValue = current.gValue + g
					heap.Fix(openList, index)
				}
			} else {
				heap.Push(openList, &aStarParam{
					grid: neighbor, gValue: g, hValue: s.calculateHValue(neighbor, end), parent: current,
				})
			}
		}
	}
	roadsSlice = append(roadsSlice, start)
	return roadsSlice
}

func (s *aStart) findNeighbor(point int32) []int32 {
	x, y := s.mapData.GetIndex2Pos(point)
	width, height := s.mapData.GetWidth(), s.mapData.GetHeight()
	neighbors = neighbors[:0]
	for i := int32(-1); i <= 1; i++ {
		for j := -int32(1); j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			nx := x + i
			ny := y + j
			if nx < 0 || ny < 0 || nx >= width || ny >= height {
				continue
			}
			neighbors = append(neighbors, s.mapData.GetPosIndex(nx, ny))
		}
	}
	return neighbors
}

func (s *aStart) calculateGValue(pointA, pointB int32, width int32) int32 {
	return 10
	//if pointA-width == pointB || pointA+width == pointB || pointA+1 == pointB || pointA-1 == pointB {
	//	return 10
	//}
	//return 15
}

func (s *aStart) calculateHValue(pointA, pointB int32) int32 {
	colA, rowA := s.mapData.GetIndex2Pos(pointA)
	colB, rowB := s.mapData.GetIndex2Pos(pointB)
	return int32((math.Abs(float64(rowB-rowA)) + math.Abs(float64(colB-colA))) * 10)
}
