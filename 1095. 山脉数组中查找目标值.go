package main

type MountainArray struct {
 }

func (this *MountainArray) get(index int) int {
	return 0
}
func (this *MountainArray) length() int {
	return 0
}


func findInMountainArray(target int, mountainArr *MountainArray) int {
	n := mountainArr.length() - 1
	l := 0
	r := n
	for r - l > 2 {
		c := (r - l) / 3
		m1 := l + c
		m2 := r - c
		a1 := mountainArr.get(m1)
		a2 := mountainArr.get(m2)
		if a1 == a2 {
			l = m1
			r = m2
		} else if a1 < a2 {
			l = m1
		} else {
			r = m2
		}
	}
	peak := l
	tmp := mountainArr.get(peak)
	for i := l; i <= r; i++ {
		if mountainArr.get(i) > tmp {
			peak = i
			tmp = mountainArr.get(i)
		}
	}

	l = 0
	r = peak
	for l <= r {
		x := (l + r) >> 1
		tmp := mountainArr.get(x)
		if tmp == target {
			return x
		} else if tmp < target {
			l = x + 1
		} else {
			r = x - 1
		}
	}

	l = peak + 1
	r = n
	for l <= r {
		x := (l + r) >> 1
		tmp := mountainArr.get(x)
		if tmp == target {
			return x
		} else if tmp > target {
			l = x + 1
		} else {
			r = x - 1
		}
	}

	return -1
}
