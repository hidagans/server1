package twomillionbc

import (
	_ "embed"

	"github.com/slotopol/server/game/slot"
)

//go:embed 2millionbc_bon.yaml
var rbon []byte

var ReelsBon = slot.ReadObj[*slot.Reels5x](rbon)

//go:embed 2millionbc_reel.yaml
var reels []byte

var ReelsMap = slot.ReadMap[*slot.Reels5x](reels)

// Lined payment.
var LinePay = [13][5]float64{
	{0, 30, 100, 300, 500}, //  1 girl
	{0, 15, 75, 200, 400},  //  2 lion
	{0, 10, 60, 150, 300},  //  3 bee
	{0, 5, 50, 125, 250},   //  4 stone
	{0, 5, 40, 100, 200},   //  5 wheel
	{0, 2, 30, 90, 150},    //  6 club
	{0, 0, 25, 75, 125},    //  7 chaplet
	{0, 0, 20, 60, 100},    //  8 gold
	{0, 0, 15, 50, 75},     //  9 vase
	{0, 0, 10, 25, 50},     // 10 ruby
	{},                     // 11 fire
	{},                     // 12 acorn
	{0, 0, 40, 100, 200},   // 13 diamond
}

// Scatter freespins table
var ScatFreespin = [5]int{0, 0, 4, 12, 20} // 11 fire

// Bet lines
var BetLines = slot.BetLinesNetEnt5x3[:30]

const (
	acbn = 1 // acorn bonus
	dlbn = 2 // diamond lion bonus
)

type Game struct {
	slot.Slot5x3 `yaml:",inline"`
	// acorns number
	AN int `json:"an" yaml:"an" xml:"an"`
	// acorns bet
	AB float64 `json:"ab" yaml:"ab" xml:"ab"`
}

// Declare conformity with SlotGame interface.
var _ slot.SlotGame = (*Game)(nil)

func NewGame() *Game {
	return &Game{
		Slot5x3: slot.Slot5x3{
			Sel: len(BetLines),
			Bet: 1,
		},
	}
}

const scat, acorn, diamond = 11, 12, 13

func (g *Game) Scanner(screen slot.Screen, wins *slot.Wins) {
	g.ScanLined(screen, wins)
	g.ScanScatters(screen, wins)
}

// Lined symbols calculation.
func (g *Game) ScanLined(screen slot.Screen, wins *slot.Wins) {
	for li := 1; li <= g.Sel; li++ {
		var line = BetLines[li-1]

		var numl slot.Pos = 5
		var syml = screen.Pos(1, line)
		var x slot.Pos
		for x = 2; x <= 5; x++ {
			var sx = screen.Pos(x, line)
			if sx != syml {
				numl = x - 1
				break
			}
		}

		if pay := LinePay[syml-1][numl-1]; pay > 0 {
			*wins = append(*wins, slot.WinItem{
				Pay:  g.Bet * pay,
				Mult: 1,
				Sym:  syml,
				Num:  numl,
				Line: li,
				XY:   line.CopyL(numl),
			})
		}
		if syml == diamond && numl >= 3 {
			*wins = append(*wins, slot.WinItem{
				Mult: 1,
				Sym:  diamond,
				Num:  numl,
				Line: li,
				XY:   line.CopyL(numl),
				BID:  dlbn,
			})
		}
	}
}

// Scatters calculation.
func (g *Game) ScanScatters(screen slot.Screen, wins *slot.Wins) {
	if count := screen.ScatNum(scat); count >= 3 {
		var fs = ScatFreespin[count-1]
		*wins = append(*wins, slot.WinItem{
			Sym:  scat,
			Num:  count,
			XY:   screen.ScatPos(scat),
			Free: fs,
		})
	}

	if screen.At(5, 1) == acorn || screen.At(5, 2) == acorn || screen.At(5, 3) == acorn {
		if (g.AN+1)%3 == 0 {
			*wins = append(*wins, slot.WinItem{
				Mult: 1,
				Sym:  acorn,
				Num:  1,
				BID:  acbn,
			})
		}
	}
}

func (g *Game) Spin(screen slot.Screen, mrtp float64) {
	if g.FSR == 0 {
		var reels, _ = slot.FindReels(ReelsMap, mrtp)
		screen.Spin(reels)
	} else {
		screen.Spin(ReelsBon)
	}
}

func (g *Game) Spawn(screen slot.Screen, wins slot.Wins) {
	for i, wi := range wins {
		switch wi.BID {
		case acbn:
			wins[i].Pay = AcornSpawn(g.AB + g.Bet*float64(g.Sel))
		case dlbn:
			wins[i].Pay = DiamondLionSpawn(g.Bet)
		}
	}
}

func (g *Game) Apply(screen slot.Screen, wins slot.Wins) {
	if screen.At(5, 1) == acorn || screen.At(5, 2) == acorn || screen.At(5, 3) == acorn {
		g.AN++
		g.AN %= 3
		if g.AN > 0 {
			g.AB += g.Bet * float64(g.Sel)
		} else {
			g.AB = 0
		}
	}

	if g.FSR != 0 {
		g.Gain += wins.Gain()
		g.FSN++
	} else {
		g.Gain = wins.Gain()
		g.FSN = 0
	}

	if g.FSR > 0 {
		g.FSR--
	}
	for _, wi := range wins {
		if wi.Free > 0 {
			g.FSR += wi.Free
		}
	}
}

func (g *Game) SetSel(sel int) error {
	return g.SetSelNum(sel, len(BetLines))
}
