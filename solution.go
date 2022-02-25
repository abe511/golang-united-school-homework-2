package square

import "math"

type Sides byte

func CalcSquare(sideLen float64, sidesNum Sides) float64 {

	const (
		SidesCircle Sides  = iota
		_
		_
		SidesTriangle
		SidesSquare
	)
	
	if sidesNum == SidesCircle {
		return math.Pi * (sideLen * sideLen)
	} else if sidesNum == SidesTriangle {
		return (math.Sqrt(3) / 4) * (sideLen * sideLen)
	} else if sidesNum == SidesSquare {
		return sideLen * sideLen
	}
	return 0
}