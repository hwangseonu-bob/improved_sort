package main

func QSort(d []int, l, r int) {
	var i, j, p int
	if l < r {
		for {
			p := d[(l+r)/2]

			for i = l; d[i] < p; i++ {
			}
			for j = r-1; d[j] >= p; j-- {
			}

			if i >= j {
				break
			}

			d[i], d[j] = d[j], d[i]
		}
		d[l], d[j] = d[j], p
		//quickSort(d, i+1, r, md)
		//quickSort(d, l, j-1, md)
		QSort(d, l, j-1)
		QSort(d, j+1, r)
	}
}
