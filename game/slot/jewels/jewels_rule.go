package jewels

import (
	"github.com/slotopol/server/game/slot"
)

// Lined payment.
var LinePay = [7][5]float64{
	{0, 0, 20, 200, 2000}, // 1 crown
	{0, 0, 15, 100, 500},  // 2 gold
	{0, 0, 15, 100, 500},  // 3 money
	{0, 0, 10, 50, 200},   // 4 ruby
	{0, 0, 10, 50, 200},   // 5 sapphire
	{0, 0, 5, 25, 100},    // 6 emerald
	{0, 0, 5, 25, 100},    // 7 amethyst
}

const (
	jid = 1 // jackpot ID
)

// Jackpot win combinations.
var Jackpot = [7][5]int{
	{0, 0, 0, 0, 0}, //  1 crown
	{0, 0, 0, 0, 0}, //  2 gold
	{0, 0, 0, 0, 0}, //  3 money
	{0, 0, 0, 0, 0}, //  4 ruby
	{0, 0, 0, 0, 0}, //  5 sapphire
	{0, 0, 0, 0, 0}, //  6 emerald
	{0, 0, 0, 0, 0}, //  7 amethyst
}

// Bet lines
var BetLines = slot.BetLinesNvm10

type Game struct {
	slot.Slot5x3 `yaml:",inline"`
}

func NewGame() *Game {
	return &Game{
		Slot5x3: slot.Slot5x3{
			Sel: 5,
			Bet: 1,
		},
	}
}

func (g *Game) Scanner(screen slot.Screen, wins *slot.Wins) {
	g.ScanLined(screen, wins)
}

// Lined symbols calculation.
func (g *Game) ScanLined(screen slot.Screen, wins *slot.Wins) {
	for li := 1; li <= g.Sel; li++ {
		var line = BetLines[li-1]

		var numl slot.Pos = 1
		var syml = screen.Pos(3, line)
		var xy slot.Linex
		xy.Set(3, line.At(3))
		if screen.Pos(2, line) == syml {
			xy.Set(2, line.At(2))
			numl++
			if screen.Pos(1, line) == syml {
				xy.Set(1, line.At(1))
				numl++
			}
		}
		if screen.Pos(4, line) == syml {
			xy.Set(4, line.At(4))
			numl++
			if screen.Pos(5, line) == syml {
				xy.Set(5, line.At(5))
				numl++
			}
		}

		if numl >= 3 {
			*wins = append(*wins, slot.WinItem{
				Pay:  g.Bet * LinePay[syml-1][numl-1],
				Mult: 1,
				Sym:  syml,
				Num:  numl,
				Line: li,
				XY:   xy,
			})
		}
	}
}

func (g *Game) Spin(screen slot.Screen, mrtp float64) {
	var reels, _ = slot.FindReels(ReelsMap, mrtp)
	screen.Spin(reels)
}

func (g *Game) SetSel(sel int) error {
	return g.SetSelNum(sel, len(BetLines))
}