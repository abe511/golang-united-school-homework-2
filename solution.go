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
	switch sidesNum {
	case SidesCircle:
		return math.Pi * (sideLen * sideLen)
	case SidesTriangle:
		return (math.Sqrt(3) * 0.25) * (sideLen * sideLen)
	case SidesSquare:
		return sideLen * sideLen
	default:
		return 0
	}
}
