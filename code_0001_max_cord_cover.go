package main

func maxCordCover(points []int, k int) int {
	var l, r, res int
	for l < len(points) {
		for r < len(points) && points[r]-points[l] <= k {
			r++
		}
		res = max(res, r-l)
		l++
	}
	return res
}
