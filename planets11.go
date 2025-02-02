package planets11

import (
	"math"

	"github.com/RulezKT/floatsfile"
)

const (
	MOON_FILE       = "moon.bin"
	SUN_FILE        = "sun.bin"
	MER_FILE        = "mercury.bin"
	VEN_FILE        = "venus.bin"
	MAR_FILE        = "mars.bin"
	JUP_FILE        = "jupyter.bin"
	SAT_FILE        = "saturn.bin"
	URA_FILE        = "uranus.bin"
	NEP_FILE        = "neptune.bin"
	PLU_FILE        = "pluto.bin"
	EARTH_FILE      = "earth.bin"
	EARTH_BARY_FILE = "earth_bary.bin"
)

type Position struct {
	X float64
	Y float64
	Z float64
}

type Pl11 struct {
	moon   []float64
	sun    []float64
	mer    []float64
	ven    []float64
	mar    []float64
	jup    []float64
	sat    []float64
	ura    []float64
	nep    []float64
	plu    []float64
	earth  []float64
	earthb []float64
}

func (pl *Pl11) Load(dir string) {

	pl.earth = floatsfile.LoadBinary(dir+EARTH_FILE, 752801)
	pl.moon = floatsfile.LoadBinary(dir+MOON_FILE, 752801)
	pl.sun = floatsfile.LoadBinary(dir+SUN_FILE, 160685)
	pl.mer = floatsfile.LoadBinary(dir+MER_FILE, 403788)
	pl.ven = floatsfile.LoadBinary(dir+VEN_FILE, 146848)
	pl.mar = floatsfile.LoadBinary(dir+MAR_FILE, 80325)
	pl.jup = floatsfile.LoadBinary(dir+JUP_FILE, 59670)
	pl.sat = floatsfile.LoadBinary(dir+SAT_FILE, 52785)
	pl.ura = floatsfile.LoadBinary(dir+URA_FILE, 45900)
	pl.nep = floatsfile.LoadBinary(dir+NEP_FILE, 45900)
	pl.plu = floatsfile.LoadBinary(dir+PLU_FILE, 45900)
	pl.earthb = floatsfile.LoadBinary(dir+EARTH_BARY_FILE, 188149)
}

func (pl *Pl11) Calc(seconds int64) []Position {
	return []Position{
		// earth
		calcPos11(seconds, -3_157_963_200, 41, 345_600, 12, pl.earth),

		// moon
		calcPos11(seconds, -3_157_963_200, 41, 345_600, 12, pl.moon),

		// sun
		calcPos11(seconds, -3_157_444_800, 35, 1_382_400, 10, pl.sun),

		// mercury
		calcPos11(seconds, -3_155_716_800, 44, 691_200, 13, pl.mer),

		// venus
		calcPos11(seconds, -3_156_062_400, 32, 1_382_400, 9, pl.ven),

		// mars
		calcPos11(seconds, -3_156_753_600, 35, 2_764_800, 10, pl.mar),

		// jupyter
		calcPos11(seconds, -3_156_753_600, 26, 2_764_800, 7, pl.jup),

		// saturn
		calcPos11(seconds, -3_156_753_600, 23, 2_764_800, 6, pl.sat),

		// uranus
		calcPos11(seconds, -3_156_753_600, 20, 2_764_800, 5, pl.ura),

		// neptune
		calcPos11(seconds, -3_156_753_600, 20, 2_764_800, 5, pl.nep),

		// pluto
		calcPos11(seconds, -3_156_753_600, 20, 2_764_800, 5, pl.plu),

		// earth barycenter
		calcPos11(seconds, -3_156_062_400, 41, 1_382_400, 12, pl.earthb),
	}

}

func calcPos11(seconds int64, startTime int64, rsize int, intlen int, order int, arrPtr []float64) Position {

	deg := order + 1

	offset := math.Floor((float64(seconds)-float64(startTime))/float64(intlen)) * float64(rsize)
	data := arrPtr[int64(offset) : int64(offset)+int64(rsize)]
	tau := (float64(seconds) - data[0]) / data[1]

	return Position{
		X: chebyshev(order, tau, data[2:2+deg]),
		Y: chebyshev(order, tau, data[2+deg:2+2*deg]),
		Z: chebyshev(order, tau, data[2+2*deg:2+3*deg]),
	}
}

func chebyshev(order int, x float64, data []float64) float64 {

	// Evaluate a Chebyshev polynomial
	var bk float64
	two_x := 2 * x
	bkp2 := data[order]
	bkp1 := two_x*bkp2 + data[order-1]

	for n := order - 2; n > 0; n-- {
		bk = data[n] + two_x*bkp1 - bkp2
		bkp2 = bkp1
		bkp1 = bk
	}
	return data[0] + x*bkp1 - bkp2
}
