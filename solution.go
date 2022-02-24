package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

type Sides int

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