package exercise_2

func CalculateAverage(notes ...float64) (float64, string) {
	var sumatoria float64
	for _, note := range notes {
		if note < 0 {
			return 0.0, "ERROR!. You can't introduce negative notes"
		}
		sumatoria += note
	}

	return sumatoria / float64(len(notes)), ""
}
