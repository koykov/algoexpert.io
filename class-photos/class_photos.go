package class_photos

import "sort"

func ClassPhotos(redShirtHeights []int, blueShirtHeights []int) (r bool) {
	sort.Ints(redShirtHeights)
	sort.Ints(blueShirtHeights)
	r = true
	for i := 0; i < len(redShirtHeights); i++ {
		if blueShirtHeights[i] <= redShirtHeights[i] {
			r = false
			break
		}
	}
	if r {
		return
	}
	r = true
	for i := 0; i < len(redShirtHeights); i++ {
		if redShirtHeights[i] <= blueShirtHeights[i] {
			r = false
			break
		}
	}
	return
}
