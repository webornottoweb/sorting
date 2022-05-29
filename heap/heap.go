package heap

// Sort sorts received array slice
func Sort(src []int) []int {
	// first we need to create heap
	// heap is a binary tree which fits to the rule - root node is greater then any of its leaves
	// we can represent a heap (or a tree) with an array
	//   n
	//  / \
	// l   r
	// n has an address of a[i]
	// left node l has an address a[2i+1]
	// right node r has an address a[2i+2]
	res := heapify(src)

	qSort := res

	for i := 0; i < len(src); i++ {
		res[len(res)-1-i], qSort[0] = qSort[0], res[len(res)-1-i] // switch first element with tailed non sorted elements part
		qSort = heapifyNode(0, qSort[:len(qSort)-1])
	}

	return res
}

func heapify(src []int) []int {
	for i := len(src) / 2; i >= 0; i-- {
		src = heapifyNode(i, src)
	}

	return src
}

func heapifyNode(idx int, src []int) []int {
	for {
		if idx > len(src)-1 {
			break
		}

		var selected bool
		var maxChild, maxChildId int

		if len(src) > idx*2+2 { // has two chilren
			selected = true
			maxChild, maxChildId = src[idx*2+2], idx*2+2

			if src[idx*2+1] > src[idx*2+2] {
				maxChild, maxChildId = src[idx*2+1], idx*2+1
			}
		} else if len(src) > idx*2+1 { // has one child
			selected = true
			maxChild, maxChildId = src[idx*2+1], idx*2+1
		}

		// switch with current node if it is greater then max child
		if selected && maxChild > src[idx] {
			src[idx], src[maxChildId] = src[maxChildId], src[idx]
			idx = maxChildId
			continue
		}

		break
	}

	return src
}
