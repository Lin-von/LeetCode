package main


func gcd(a, b int) int {
	if a % b == 0 {
		return b
	}

	return gcd(b, a % b)
}

func canMeasureWater(x int, y int, z int) bool {
	if z == 0 {
		return true
	}

	if z > x + y {
		return false
	}

	if x == 0 && y == 0 {
		return false
	}

	if x == 0 {
		return z % y == 0
	}

	if y == 0 {
		return x % y == 0
	}

	return z % gcd(x, y) == 0

}
