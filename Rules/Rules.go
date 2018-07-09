package Rules

func CheckStillAlive(neighbors int) bool{
	if neighbors == 2 || neighbors == 3 {
		return true
	} else {
		return false
	}
}

func CheckStillDead(neighbors int) bool{
	if neighbors == 3 {
		return true
	} else {
		return false
	}
}