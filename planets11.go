package planets11

import "github.com/RulezKT/floatsfile"

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
	pl.earth = floatsfile.LoadBinary(dir+EARTH_FILE, 752801)
	pl.earthb = floatsfile.LoadBinary(dir+EARTH_BARY_FILE, 188149)
}

func (pl *Pl11) Calc(seconds int64) []Position {
	return []Position{}

}
