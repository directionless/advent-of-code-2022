package day09

// From https://gist.github.com/ota42y/db4ff0298d9c945cd261

type Point struct {
	x float64
	y float64
}

func leastSquaresMethod(points *[]Point) (a float64, b float64) {
	// http://ja.wikipedia.org/wiki/%E6%9C%80%E5%B0%8F%E4%BA%8C%E4%B9%97%E6%B3%95

	n := float64(len(*points))

	sumX := 0.0
	sumY := 0.0
	sumXY := 0.0
	sumXX := 0.0

	for _, p := range *points {
		sumX += p.x
		sumY += p.y
		sumXY += p.x * p.y
		sumXX += p.x * p.x
	}

	base := (n*sumXX - sumX*sumX)
	a = (n*sumXY - sumX*sumY) / base
	b = (sumXX*sumY - sumXY*sumX) / base

	return a, b
}
