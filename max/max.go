package max

import ()

func MaximumDistance(xx, yy []int) (out int) {
	for ii := 0; ii < len(xx); ii++ {
		for jj := ii; jj < len(yy); jj++ {
			if yy[jj] >= xx[ii] && jj-ii > out {
				out = jj - ii
			}
		}
	}
	return
}
