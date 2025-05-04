package simple

func SwitchCase(day int) int {
	switch day {
	case 1:
		return 1
	case 2:
		return 2
	default:
		return day + 1
	}
}
