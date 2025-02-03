package planets11

import (
	"math"
	"path/filepath"

	"github.com/RulezKT/floatsfile"
	"github.com/RulezKT/types"
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

	EARTH_LENGTH      = 752801
	MOON_LENGTH       = 752801
	SUN_LENGTH        = 160685
	MER_LENGTH        = 403788
	VEN_LENGTH        = 146848
	MAR_LENGTH        = 80325
	JUP_LENGTH        = 59670
	SAT_LENGTH        = 52785
	URA_LENGTH        = 45900
	NEP_LENGTH        = 45900
	PLU_LENGTH        = 45900
	EARTH_BARY_LENGTH = 188149
)

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

	pl.earth = floatsfile.LoadBinary(filepath.Join(dir, EARTH_FILE), EARTH_LENGTH)
	pl.moon = floatsfile.LoadBinary(filepath.Join(dir, MOON_FILE), MOON_LENGTH)
	pl.sun = floatsfile.LoadBinary(filepath.Join(dir, SUN_FILE), SUN_LENGTH)
	pl.mer = floatsfile.LoadBinary(filepath.Join(dir, MER_FILE), MER_LENGTH)
	pl.ven = floatsfile.LoadBinary(filepath.Join(dir, VEN_FILE), VEN_LENGTH)
	pl.mar = floatsfile.LoadBinary(filepath.Join(dir, MAR_FILE), MAR_LENGTH)
	pl.jup = floatsfile.LoadBinary(filepath.Join(dir, JUP_FILE), JUP_LENGTH)
	pl.sat = floatsfile.LoadBinary(filepath.Join(dir, SAT_FILE), SAT_LENGTH)
	pl.ura = floatsfile.LoadBinary(filepath.Join(dir, URA_FILE), URA_LENGTH)
	pl.nep = floatsfile.LoadBinary(filepath.Join(dir, NEP_FILE), NEP_LENGTH)
	pl.plu = floatsfile.LoadBinary(filepath.Join(dir, PLU_FILE), PLU_LENGTH)
	pl.earthb = floatsfile.LoadBinary(filepath.Join(dir, EARTH_BARY_FILE), EARTH_BARY_LENGTH)
}

func (pl *Pl11) Calc(seconds float64) []types.Position {
	return []types.Position{
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

func calcPos11(seconds float64, startTime int64, rsize int, intlen int, order int, arrPtr []float64) types.Position {

	deg := order + 1

	offset := math.Floor((seconds-float64(startTime))/float64(intlen)) * float64(rsize)
	data := arrPtr[int64(offset) : int64(offset)+int64(rsize)]
	tau := (seconds - data[0]) / data[1]

	return types.Position{
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
