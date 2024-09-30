package redroo

// See: https://freeslotshub.com/aristocrat/big-red/

import (
	"math/rand/v2"

	slot "github.com/slotopol/server/game/slot"
)

// Lined payment.
var LinePay = [13][5]float64{
	{},                     //  1 wild
	{},                     //  2 scatter
	{0, 50, 150, 200, 250}, //  3 redroo
	{0, 20, 80, 150, 200},  //  4 girl
	{0, 20, 80, 150, 200},  //  5 boy
	{0, 10, 40, 100, 150},  //  6 dog
	{0, 10, 40, 100, 150},  //  7 parrot
	{0, 0, 10, 50, 140},    //  8 ace
	{0, 0, 10, 50, 140},    //  9 king
	{0, 0, 5, 40, 120},     // 10 queen
	{0, 0, 5, 40, 120},     // 11 jack
	{0, 0, 5, 20, 100},     // 12 ten
	{0, 2, 5, 20, 100},     // 13 nine
}

// Scatters payment.
var ScatPay = [5]float64{0, 0, 80, 400, 800} // scatter

// Scatter freespins table on regular games
var ScatFreespinReg = [5]int{0, 0, 8, 15, 20} // scatter

// Scatter freespins table on bonus games
var ScatFreespinBon = [5]int{0, 5, 8, 15, 20} // scatter

type Game struct {
	slot.Slot5x4 `yaml:",inline"`
	// free spin number
	FS int `json:"fs,omitempty" yaml:"fs,omitempty" xml:"fs,omitempty"`
	// wild multipliers
	MW [3]float64 `json:"mw" yaml:"mw" xml:"mw"`
}

func NewGame() *Game {
	return &Game{
		Slot5x4: slot.Slot5x4{
			Sel: slot.MakeBitNum(40, 1),
			Bet: 1,
		},
		FS: 0,
		MW: [3]float64{1, 1, 1},
	}
}

const wild, scat = 1, 2

const prob2x = 0.5 // probability of 2x multiplier for wild at free games

func (g *Game) Scanner(screen slot.Screen, wins *slot.Wins) {
	g.ScanLined(screen, wins)
	g.ScanScatters(screen, wins)
}

// Lined symbols calculation.
func (g *Game) ScanLined(screen slot.Screen, wins *slot.Wins) {
	var line slot.Linex
loop1:
	for line[0] = 1; line[0] <= 4; line[0]++ {
	loop2:
		for line[1] = 1; line[1] <= 4; line[1]++ {
		loop3:
			for line[2] = 1; line[2] <= 4; line[2]++ {
			loop4:
				for line[3] = 1; line[3] <= 4; line[3]++ {
				loop5:
					for line[4] = 1; line[4] <= 4; line[4]++ {
						var numl slot.Pos = 5
						var syml slot.Sym
						var mw float64 = 1 // mult wild
						var x slot.Pos
						for x = 1; x <= 5; x++ {
							var sx = screen.Pos(x, line)
							if sx == wild {
								mw *= g.MW[x-2]
							} else if syml == 0 && sx != scat {
								syml = sx
							} else if sx != syml {
								numl = x - 1
								break
							}
						}

						if numl >= 2 && syml > 0 {
							// var li = (int(line[0])-1)*256 + (int(line[1])-1)*64 + (int(line[2])-1)*16 + (int(line[line[4]])-1)*4 + int(line[5])
							*wins = append(*wins, slot.WinItem{
								Pay:  g.Bet * LinePay[syml-1][numl-1],
								Mult: mw,
								Sym:  syml,
								Num:  numl,
								Line: 1024,
								XY:   line.CopyL(numl),
							})
							switch numl {
							case 2:
								continue loop2
							case 3:
								continue loop3
							case 4:
								continue loop4
							case 5:
								continue loop5
							}
						}
						switch numl + 1 {
						case 1:
							continue loop1
						case 2:
							continue loop2
						case 3:
							continue loop3
						case 4:
							continue loop4
						case 5:
							continue loop5
						}
					}
				}
			}
		}
	}
}

// Scatters calculation.
func (g *Game) ScanScatters(screen slot.Screen, wins *slot.Wins) {
	if count := screen.ScatNum(scat); count >= 2 {
		var pay = ScatPay[count-1]
		var fs int
		if g.FS > 0 {
			fs = ScatFreespinBon[count-1]
		} else {
			fs = ScatFreespinReg[count-1]
		}
		if pay >= 0 || fs >= 0 {
			*wins = append(*wins, slot.WinItem{
				Pay:  g.Bet * pay,
				Mult: 1,
				Sym:  scat,
				Num:  count,
				XY:   screen.ScatPos(scat),
				Free: fs,
			})
		}
	}
}

func (g *Game) Spin(screen slot.Screen, mrtp float64) {
	var _, reels = slot.FindReels(ReelsMap, mrtp)
	screen.Spin(reels)
}

func (g *Game) Prepare() {
	if g.FS > 0 {
		for x := range g.MW {
			if rand.Float64() < prob2x {
				g.MW[x] = 2
			} else {
				g.MW[x] = 3
			}
		}
	}
}

func (g *Game) Apply(screen slot.Screen, wins slot.Wins) {
	if g.FS > 0 {
		g.Gain += wins.Gain()
	} else {
		g.Gain = wins.Gain()
	}

	if g.FS > 0 {
		g.FS--
	}
	for _, wi := range wins {
		if wi.Free > 0 {
			g.FS += wi.Free
		}
	}
}

func (g *Game) FreeSpins() int {
	return g.FS
}

func (g *Game) SetSel(sel slot.Bitset) error {
	return slot.ErrNoFeature
}
