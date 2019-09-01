package main

func ISort(d []int) {
	n := len(d)
	quickSort(d, 0, n, maxDepth(n))
}

// 자료의 크기가 12이하이면 쉘정렬
// 자료의 크기가 12보다 많을 때 Depth가 0이면 힙정렬, 아니면 퀵정렬
// MaxDepth는 2 * ceil(log(n+1))
func quickSort(d []int, l, r int, md int) {
	if r-l > 12 {
		if md == 0 {
			heapSort(d, l, r)
			return
		}
		md--
		if l < r {
			p := pivot(d, l, r)
			quickSort(d, l, p, md)
			quickSort(d, p, r, md)
		}
	} else {
		shellSort(d, l, r)
	}
}

func heapSort(d []int, l, r int) {
	f, lo, hi := l, 0, r-l

	for i := (hi - 1) / 2; i >= 0; i-- {
		shiftDown(d, i, hi, f)
	}

	for i := hi - 1; i >= 0; i-- {
		d[f], d[f+i] = d[f+i], d[f]
		shiftDown(d, lo, i, f)
	}
}

func shellSort(d []int, l, r int) {
	for i := l + 6; i < r; i++ {
		if d[i] < d[i-6] {
			d[i], d[i-6] = d[i-6], d[i]
		}
	}
	insertSort(d, l, r)
}

func insertSort(d []int, l, r int) {
	for i := l + 1; i < r; i++ {
		for j := i; j > l && d[j] < d[j-1]; j-- {
			d[j], d[j-1] = d[j-1], d[j]
		}
	}
}

func shiftDown(d []int, r, hi int, f int) {
	c := 2*r + 1
	if c >= hi {
		return
	}

	if c+1 < hi && d[f+c] < d[f+c+1] {
		c++
	}

	if !(d[f+r] < d[f+c]) {
		return
	}

	d[f+r], d[f+c] = d[f+c], d[f+r]
	shiftDown(d, c, hi, f)
}

func pivot(d []int, l, r int) int {
	p := l
	i, j := l+1, p
	for ; i < r; i++ {
		if d[i] < d[p] {
			j++
			d[i], d[j] = d[j], d[i]
		}
	}
	d[l], d[j] = d[j], d[l]
	return j
}


// 퀵정렬을 힙정렬로 변환하기 위한 임계값 반환
// 2*ceil(lg(n+1))
func maxDepth(n int) int {
	var depth int
	for i := n; i > 0; i >>= 1 {
		depth++
	}
	return depth * 2
}