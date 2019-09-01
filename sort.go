package main

func QSort(d []int, l, r int) {
	p := l
	i, j := l+1, p
	if l < r {
		for ; i <= r; i++ {
			if d[i] < d[p] {
				j++
				d[i], d[j] = d[j], d[i]
			}
		}
		d[l], d[j] = d[j], d[l]
		p = j

		QSort(d, l, p-1)
		QSort(d, p+1, r)
	}
}
