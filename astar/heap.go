package astar

type aStarParamHeap struct {
	paramArray []*aStarParam
}

func (h *aStarParamHeap) Len() int {
	return len(h.paramArray)
}

func (h *aStarParamHeap) Less(i, j int) bool {
	left, right := h.paramArray[i], h.paramArray[j]
	return left.gValue+left.hValue < right.gValue+right.hValue
}

func (h *aStarParamHeap) Swap(i, j int) {
	h.paramArray[i], h.paramArray[j] = h.paramArray[j], h.paramArray[i]
}

func (h *aStarParamHeap) Push(x interface{}) {
	h.paramArray = append(h.paramArray, x.(*aStarParam))
}

func (h *aStarParamHeap) Pop() (ret interface{}) {
	n := len(h.paramArray)

	h.paramArray, ret = h.paramArray[:n-1], h.paramArray[n-1]

	return
}

func (h *aStarParamHeap) index(grid int32) (int, *aStarParam) {
	for i, v := range h.paramArray {
		if v.grid == grid {
			return i, v
		}
	}
	return -1, nil
}
