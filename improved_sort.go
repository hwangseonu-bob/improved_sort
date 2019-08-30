package main

func ISort(d []int) {
	n := len(d)
	quickSort(d, 0, n, maxDepth(n))
}

func quickSort(d []int, l, r int, md int) {
	if r-l > 12 {
		if md == 0 {
			heapSort(d, l, r)
			return
		}
		md--
		i, j := pivot(d, l, r)
		quickSort(d, i+1, r, md)
		quickSort(d, l, j-1, md)
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

func pivot(d []int, l, r int) (int, int) {
	var i, j int
	for {
		p := d[(l+r)/2]

		for i = l; d[i] < p; i++ {
		}
		for j = r - 1; d[j] >= p; j-- {
		}

		if i >= j {
			break
		}

		d[i], d[j] = d[j], d[i]
	}
	return i, j
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

// 퀵정렬을 힙정렬로 변환하기 위한 임계값 반환
// 2*ceil(lg(n+1))
func maxDepth(n int) int {
	var depth int
	for i := n; i > 0; i >>= 1 {
		depth++
	}
	return depth * 2
}
