package panda

// See: https://demo.agtsoftware.com/games/agt/panda

import (
	_ "embed"

	"github.com/slotopol/server/game/slot"
)

//go:embed panda_reel.yaml
var reels []byte

var ReelsMap = slot.ReadMap[*slot.Reels3x](reels)

// Lined payment.
var LinePay = [9][3]float64{
	{0, 0, 500}, // 1 wild
	{0, 0, 160}, // 2 bonsai
	{0, 0, 80},  // 3 fish
	{0, 0, 40},  // 4 fan
	{0, 0, 20},  // 5 lamp
	{0, 0, 20},  // 6 pot
	{0, 0, 20},  // 7 flower
	{0, 0, 10},  // 8 button
	{},          // 9 scatter
}

// Bet lines
var BetLines = slot.BetLinesAgt3x3[:27]

type Game struct {
	slot.Slot3x3 `yaml:",inline"`
}

// Declare conformity with SlotGame interface.
var _ slot.SlotGame = (*Game)(nil)

func NewGame() *Game {
	return &Game{
		Slot3x3: slot.Slot3x3{
			Sel: len(BetLines),
			Bet: 1,
		},
	}
}

const wild, scat = 1, 9

func (g *Game) Scanner(screen slot.Screen, wins *slot.Wins) {
	for li := 1; li <= g.Sel; li++ {
		var line = BetLines[li-1]

		var numw, numl slot.Pos = 0, 3
		var syml slot.Sym
		var x slot.Pos
		for x = 1; x <= 3; x++ {
			var sx = screen.Pos(x, line)
			if sx == wild {
				if syml == 0 {
					numw = x
				}
			} else if syml == 0 && sx != scat {
				syml = sx
			} else if sx != syml {
				numl = x - 1
				break
			}
		}

		if numw == 3 {
			var pay = LinePay[wild-1][numw-1]
			*wins = append(*wins, slot.WinItem{
				Pay:  g.Bet * pay,
				Mult: 1,
				Sym:  wild,
				Num:  numw,
				Line: li,
				XY:   line.CopyL(numw),
			})
		} else if numl == 3 {
			var pay = LinePay[syml-1][numl-1]
			*wins = append(*wins, slot.WinItem{
				Pay:  g.Bet * pay,
				Mult: 1,
				Sym:  syml,
				Num:  numl,
				Line: li,
				XY:   line.CopyL(numl),
			})
		}
	}

	if count := screen.ScatNum(scat); count > 0 {
		*wins = append(*wins, slot.WinItem{
			Sym:  scat,
			Num:  count,
			XY:   screen.ScatPos(scat),
			Free: int(count),
		})
	}
}

func (g *Game) Spin(screen slot.Screen, mrtp float64) {
	var reels, _ = slot.FindReels(ReelsMap, mrtp)
	screen.Spin(reels)
}

func (g *Game) SetSel(sel int) error {
	return slot.ErrNoFeature
}
