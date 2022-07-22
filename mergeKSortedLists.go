func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	res := &ListNode{}
	head := res
	var minVal, index int

	A := make(map[int]*ListNode, len(lists))
	for idx, lst := range lists {
		if lst != nil {
			A[idx] = lst
		}
	}

	keys := make([]int, 0, len(A))
	for k := range A {
		keys = append(keys, k)
	}

	for {
		if len(keys) <= 0 {
			break
		}
		minVal = A[keys[0]].Val
		index = keys[0]
		for idx, v := range A {
			if minVal > v.Val {
				minVal = v.Val
				index = idx
			}
		}

		n := &ListNode{Val: minVal}
		if res.Next == nil {
			res.Next = n
			res = res.Next
		}
		if A[index].Next == nil {
			delete(A, index)
			keys = make([]int, 0, len(A))
			for k := range A {
				keys = append(keys, k)
			}
		} else {
			A[index] = A[index].Next
		}
	}
	return head.Next
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	res := &ListNode{}
	head := res
	var minVal int

	index := 0
	for len(lists) > 0 {
		if lists[index] == nil {
			lists = append(lists[:index], lists[index+1:]...)
			continue
		}
		minVal = lists[index].Val
		for idx, v := range lists {
			if v == nil {
				continue
			}
			if minVal > v.Val {
				minVal = v.Val
				index = idx
			}
		}
		if res.Next == nil {
			res.Next = &ListNode{Val: minVal}
			res = res.Next
		}
		if lists[index].Next != nil {
			lists[index] = lists[index].Next
		} else {
			lists = append(lists[:index], lists[index+1:]...)
			index = 0
		}
	}
	return head.Next
}
