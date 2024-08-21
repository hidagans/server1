package roaringforties

// See: https://freeslotshub.com/novomatic/roaring-forties/

import (
	"math"

	"github.com/slotopol/server/game"
)

// reels lengths [58, 66, 66, 66, 58], total reshuffles 967136544
// RTP = 81.412(lined) + 7.1905(scatter) = 88.602035%
var Reels89 = game.Reels5x{
	{8, 8, 8, 8, 5, 5, 5, 5, 1, 1, 1, 1, 6, 6, 6, 6, 4, 4, 4, 4, 7, 7, 7, 7, 3, 3, 3, 3, 2, 2, 2, 2, 10, 5, 5, 5, 5, 3, 3, 3, 3, 7, 7, 7, 7, 10, 4, 4, 4, 4, 8, 8, 8, 8, 6, 6, 6, 6},
	{8, 8, 8, 8, 5, 5, 5, 5, 1, 1, 1, 1, 6, 6, 6, 6, 4, 4, 4, 4, 7, 7, 7, 7, 3, 3, 3, 3, 2, 2, 2, 2, 10, 5, 5, 5, 5, 3, 3, 3, 3, 7, 7, 7, 7, 10, 4, 4, 4, 4, 8, 8, 8, 8, 6, 6, 6, 6, 9, 9, 9, 9, 5, 6, 7, 8},
	{8, 8, 8, 8, 5, 5, 5, 5, 1, 1, 1, 1, 6, 6, 6, 6, 4, 4, 4, 4, 7, 7, 7, 7, 3, 3, 3, 3, 2, 2, 2, 2, 10, 5, 5, 5, 5, 3, 3, 3, 3, 7, 7, 7, 7, 10, 4, 4, 4, 4, 8, 8, 8, 8, 6, 6, 6, 6, 9, 9, 9, 9, 5, 6, 7, 8},
	{8, 8, 8, 8, 5, 5, 5, 5, 1, 1, 1, 1, 6, 6, 6, 6, 4, 4, 4, 4, 7, 7, 7, 7, 3, 3, 3, 3, 2, 2, 2, 2, 10, 5, 5, 5, 5, 3, 3, 3, 3, 7, 7, 7, 7, 10, 4, 4, 4, 4, 8, 8, 8, 8, 6, 6, 6, 6, 9, 9, 9, 9, 5, 6, 7, 8},
	{8, 8, 8, 8, 5, 5, 5, 5, 1, 1, 1, 1, 6, 6, 6, 6, 4, 4, 4, 4, 7, 7, 7, 7, 3, 3, 3, 3, 2, 2, 2, 2, 10, 5, 5, 5, 5, 3, 3, 3, 3, 7, 7, 7, 7, 10, 4, 4, 4, 4, 8, 8, 8, 8, 6, 6, 6, 6},
}

// reels lengths [62, 62, 62, 62, 62], total reshuffles 916132832
// RTP = 85.306(lined) + 7.462(scatter) = 92.768326%
var Reels93 = game.Reels5x{
	{8, 8, 8, 8, 5, 5, 5, 5, 1, 1, 1, 1, 6, 6, 6, 6, 4, 4, 4, 4, 7, 7, 7, 7, 3, 3, 3, 3, 2, 2, 2, 2, 10, 5, 5, 5, 5, 3, 3, 3, 3, 7, 7, 7, 7, 10, 4, 4, 4, 4, 8, 8, 8, 8, 6, 6, 6, 6, 5, 6, 7, 8},
	{8, 8, 8, 8, 5, 5, 5, 5, 1, 1, 1, 1, 6, 6, 6, 6, 4, 4, 4, 4, 7, 7, 7, 7, 3, 3, 3, 3, 2, 2, 2, 2, 10, 5, 5, 5, 5, 3, 3, 3, 3, 7, 7, 7, 7, 10, 4, 4, 4, 4, 8, 8, 8, 8, 6, 6, 6, 6, 9, 9, 9, 9},
	{8, 8, 8, 8, 5, 5, 5, 5, 1, 1, 1, 1, 6, 6, 6, 6, 4, 4, 4, 4, 7, 7, 7, 7, 3, 3, 3, 3, 2, 2, 2, 2, 10, 5, 5, 5, 5, 3, 3, 3, 3, 7, 7, 7, 7, 10, 4, 4, 4, 4, 8, 8, 8, 8, 6, 6, 6, 6, 9, 9, 9, 9},
	{8, 8, 8, 8, 5, 5, 5, 5, 1, 1, 1, 1, 6, 6, 6, 6, 4, 4, 4, 4, 7, 7, 7, 7, 3, 3, 3, 3, 2, 2, 2, 2, 10, 5, 5, 5, 5, 3, 3, 3, 3, 7, 7, 7, 7, 10, 4, 4, 4, 4, 8, 8, 8, 8, 6, 6, 6, 6, 9, 9, 9, 9},
	{8, 8, 8, 8, 5, 5, 5, 5, 1, 1, 1, 1, 6, 6, 6, 6, 4, 4, 4, 4, 7, 7, 7, 7, 3, 3, 3, 3, 2, 2, 2, 2, 10, 5, 5, 5, 5, 3, 3, 3, 3, 7, 7, 7, 7, 10, 4, 4, 4, 4, 8, 8, 8, 8, 6, 6, 6, 6, 5, 6, 7, 8},
}

// reels lengths [58, 66, 67, 66, 58], total reshuffles 981790128
// RTP = 86.974(lined) + 7.114(scatter) = 94.088429%
var Reels94 = game.Reels5x{
	{8, 8, 8, 8, 5, 5, 5, 5, 1, 1, 1, 1, 6, 6, 6, 6, 4, 4, 4, 4, 7, 7, 7, 7, 3, 3, 3, 3, 2, 2, 2, 2, 10, 5, 5, 5, 5, 3, 3, 3, 3, 7, 7, 7, 7, 10, 4, 4, 4, 4, 8, 8, 8, 8, 6, 6, 6, 6},
	{8, 8, 8, 8, 5, 5, 5, 5, 1, 1, 1, 1, 6, 6, 6, 6, 4, 4, 4, 4, 7, 7, 7, 7, 3, 3, 3, 3, 2, 2, 2, 2, 10, 5, 5, 5, 5, 3, 3, 3, 3, 7, 7, 7, 7, 10, 4, 4, 4, 4, 8, 8, 8, 8, 6, 6, 6, 6, 9, 9, 9, 9, 5, 6, 7, 8},
	{8, 8, 8, 8, 5, 5, 5, 5, 1, 1, 1, 1, 6, 6, 6, 6, 4, 4, 4, 4, 7, 7, 7, 7, 3, 3, 3, 3, 2, 2, 2, 2, 10, 5, 5, 5, 5, 3, 3, 3, 3, 7, 7, 7, 7, 10, 4, 4, 4, 4, 8, 8, 8, 8, 6, 6, 6, 6, 9, 9, 9, 9, 5, 6, 7, 8, 9},
	{8, 8, 8, 8, 5, 5, 5, 5, 1, 1, 1, 1, 6, 6, 6, 6, 4, 4, 4, 4, 7, 7, 7, 7, 3, 3, 3, 3, 2, 2, 2, 2, 10, 5, 5, 5, 5, 3, 3, 3, 3, 7, 7, 7, 7, 10, 4, 4, 4, 4, 8, 8, 8, 8, 6, 6, 6, 6, 9, 9, 9, 9, 5, 6, 7, 8},
	{8, 8, 8, 8, 5, 5, 5, 5, 1, 1, 1, 1, 6, 6, 6, 6, 4, 4, 4, 4, 7, 7, 7, 7, 3, 3, 3, 3, 2, 2, 2, 2, 10, 5, 5, 5, 5, 3, 3, 3, 3, 7, 7, 7, 7, 10, 4, 4, 4, 4, 8, 8, 8, 8, 6, 6, 6, 6},
}

// reels lengths [58, 62, 62, 62, 58], total reshuffles 801735392
// RTP = 87.193(lined) + 8.2264(scatter) = 95.419380%
var Reels95 = game.Reels5x{
	{8, 8, 8, 8, 5, 5, 5, 5, 1, 1, 1, 1, 6, 6, 6, 6, 4, 4, 4, 4, 7, 7, 7, 7, 3, 3, 3, 3, 2, 2, 2, 2, 10, 5, 5, 5, 5, 3, 3, 3, 3, 7, 7, 7, 7, 10, 4, 4, 4, 4, 8, 8, 8, 8, 6, 6, 6, 6},
	{8, 8, 8, 8, 5, 5, 5, 5, 1, 1, 1, 1, 6, 6, 6, 6, 4, 4, 4, 4, 7, 7, 7, 7, 3, 3, 3, 3, 2, 2, 2, 2, 10, 5, 5, 5, 5, 3, 3, 3, 3, 7, 7, 7, 7, 10, 4, 4, 4, 4, 8, 8, 8, 8, 6, 6, 6, 6, 9, 9, 9, 9},
	{8, 8, 8, 8, 5, 5, 5, 5, 1, 1, 1, 1, 6, 6, 6, 6, 4, 4, 4, 4, 7, 7, 7, 7, 3, 3, 3, 3, 2, 2, 2, 2, 10, 5, 5, 5, 5, 3, 3, 3, 3, 7, 7, 7, 7, 10, 4, 4, 4, 4, 8, 8, 8, 8, 6, 6, 6, 6, 9, 9, 9, 9},
	{8, 8, 8, 8, 5, 5, 5, 5, 1, 1, 1, 1, 6, 6, 6, 6, 4, 4, 4, 4, 7, 7, 7, 7, 3, 3, 3, 3, 2, 2, 2, 2, 10, 5, 5, 5, 5, 3, 3, 3, 3, 7, 7, 7, 7, 10, 4, 4, 4, 4, 8, 8, 8, 8, 6, 6, 6, 6, 9, 9, 9, 9},
	{8, 8, 8, 8, 5, 5, 5, 5, 1, 1, 1, 1, 6, 6, 6, 6, 4, 4, 4, 4, 7, 7, 7, 7, 3, 3, 3, 3, 2, 2, 2, 2, 10, 5, 5, 5, 5, 3, 3, 3, 3, 7, 7, 7, 7, 10, 4, 4, 4, 4, 8, 8, 8, 8, 6, 6, 6, 6},
}

// reels lengths [59, 62, 62, 62, 59], total reshuffles 829619768
// RTP = 88.997(lined) + 8.0221(scatter) = 97.019081%
var Reels97 = game.Reels5x{
	{8, 8, 8, 8, 5, 5, 5, 5, 1, 1, 1, 1, 6, 6, 6, 6, 4, 4, 4, 4, 7, 7, 7, 7, 3, 3, 3, 3, 2, 2, 2, 2, 10, 5, 5, 5, 5, 3, 3, 3, 3, 7, 7, 7, 7, 10, 4, 4, 4, 4, 8, 8, 8, 8, 6, 6, 6, 6, 1},
	{8, 8, 8, 8, 5, 5, 5, 5, 1, 1, 1, 1, 6, 6, 6, 6, 4, 4, 4, 4, 7, 7, 7, 7, 3, 3, 3, 3, 2, 2, 2, 2, 10, 5, 5, 5, 5, 3, 3, 3, 3, 7, 7, 7, 7, 10, 4, 4, 4, 4, 8, 8, 8, 8, 6, 6, 6, 6, 9, 9, 9, 9},
	{8, 8, 8, 8, 5, 5, 5, 5, 1, 1, 1, 1, 6, 6, 6, 6, 4, 4, 4, 4, 7, 7, 7, 7, 3, 3, 3, 3, 2, 2, 2, 2, 10, 5, 5, 5, 5, 3, 3, 3, 3, 7, 7, 7, 7, 10, 4, 4, 4, 4, 8, 8, 8, 8, 6, 6, 6, 6, 9, 9, 9, 9},
	{8, 8, 8, 8, 5, 5, 5, 5, 1, 1, 1, 1, 6, 6, 6, 6, 4, 4, 4, 4, 7, 7, 7, 7, 3, 3, 3, 3, 2, 2, 2, 2, 10, 5, 5, 5, 5, 3, 3, 3, 3, 7, 7, 7, 7, 10, 4, 4, 4, 4, 8, 8, 8, 8, 6, 6, 6, 6, 9, 9, 9, 9},
	{8, 8, 8, 8, 5, 5, 5, 5, 1, 1, 1, 1, 6, 6, 6, 6, 4, 4, 4, 4, 7, 7, 7, 7, 3, 3, 3, 3, 2, 2, 2, 2, 10, 5, 5, 5, 5, 3, 3, 3, 3, 7, 7, 7, 7, 10, 4, 4, 4, 4, 8, 8, 8, 8, 6, 6, 6, 6, 1},
}

// reels lengths [62, 67, 67, 67, 62], total reshuffles 1156132972
// RTP = 94.287(lined) + 6.3129(scatter) = 100.599786%
var Reels101 = game.Reels5x{
	{8, 8, 8, 8, 5, 5, 5, 5, 1, 1, 1, 1, 6, 6, 6, 6, 4, 4, 4, 4, 7, 7, 7, 7, 3, 3, 3, 3, 2, 2, 2, 2, 10, 5, 5, 5, 5, 3, 3, 3, 3, 7, 7, 7, 7, 10, 4, 4, 4, 4, 8, 8, 8, 8, 6, 6, 6, 6, 5, 6, 7, 8},
	{8, 8, 8, 8, 5, 5, 5, 5, 1, 1, 1, 1, 6, 6, 6, 6, 4, 4, 4, 4, 7, 7, 7, 7, 3, 3, 3, 3, 2, 2, 2, 2, 10, 5, 5, 5, 5, 3, 3, 3, 3, 7, 7, 7, 7, 10, 4, 4, 4, 4, 8, 8, 8, 8, 6, 6, 6, 6, 9, 9, 9, 9, 5, 6, 7, 8, 9},
	{8, 8, 8, 8, 5, 5, 5, 5, 1, 1, 1, 1, 6, 6, 6, 6, 4, 4, 4, 4, 7, 7, 7, 7, 3, 3, 3, 3, 2, 2, 2, 2, 10, 5, 5, 5, 5, 3, 3, 3, 3, 7, 7, 7, 7, 10, 4, 4, 4, 4, 8, 8, 8, 8, 6, 6, 6, 6, 9, 9, 9, 9, 5, 6, 7, 8, 9},
	{8, 8, 8, 8, 5, 5, 5, 5, 1, 1, 1, 1, 6, 6, 6, 6, 4, 4, 4, 4, 7, 7, 7, 7, 3, 3, 3, 3, 2, 2, 2, 2, 10, 5, 5, 5, 5, 3, 3, 3, 3, 7, 7, 7, 7, 10, 4, 4, 4, 4, 8, 8, 8, 8, 6, 6, 6, 6, 9, 9, 9, 9, 5, 6, 7, 8, 9},
	{8, 8, 8, 8, 5, 5, 5, 5, 1, 1, 1, 1, 6, 6, 6, 6, 4, 4, 4, 4, 7, 7, 7, 7, 3, 3, 3, 3, 2, 2, 2, 2, 10, 5, 5, 5, 5, 3, 3, 3, 3, 7, 7, 7, 7, 10, 4, 4, 4, 4, 8, 8, 8, 8, 6, 6, 6, 6, 5, 6, 7, 8},
}

// reels lengths [58, 63, 63, 63, 58], total reshuffles 841158108
// RTP = 103.23(lined) + 7.9464(scatter) = 111.172239%
var Reels111 = game.Reels5x{
	{8, 8, 8, 8, 5, 5, 5, 5, 1, 1, 1, 1, 6, 6, 6, 6, 4, 4, 4, 4, 7, 7, 7, 7, 3, 3, 3, 3, 2, 2, 2, 2, 10, 5, 5, 5, 5, 3, 3, 3, 3, 7, 7, 7, 7, 10, 4, 4, 4, 4, 8, 8, 8, 8, 6, 6, 6, 6},
	{8, 8, 8, 8, 5, 5, 5, 5, 1, 1, 1, 1, 6, 6, 6, 6, 4, 4, 4, 4, 7, 7, 7, 7, 3, 3, 3, 3, 2, 2, 2, 2, 10, 5, 5, 5, 5, 3, 3, 3, 3, 7, 7, 7, 7, 10, 4, 4, 4, 4, 8, 8, 8, 8, 6, 6, 6, 6, 9, 9, 9, 9, 9},
	{8, 8, 8, 8, 5, 5, 5, 5, 1, 1, 1, 1, 6, 6, 6, 6, 4, 4, 4, 4, 7, 7, 7, 7, 3, 3, 3, 3, 2, 2, 2, 2, 10, 5, 5, 5, 5, 3, 3, 3, 3, 7, 7, 7, 7, 10, 4, 4, 4, 4, 8, 8, 8, 8, 6, 6, 6, 6, 9, 9, 9, 9, 9},
	{8, 8, 8, 8, 5, 5, 5, 5, 1, 1, 1, 1, 6, 6, 6, 6, 4, 4, 4, 4, 7, 7, 7, 7, 3, 3, 3, 3, 2, 2, 2, 2, 10, 5, 5, 5, 5, 3, 3, 3, 3, 7, 7, 7, 7, 10, 4, 4, 4, 4, 8, 8, 8, 8, 6, 6, 6, 6, 9, 9, 9, 9, 9},
	{8, 8, 8, 8, 5, 5, 5, 5, 1, 1, 1, 1, 6, 6, 6, 6, 4, 4, 4, 4, 7, 7, 7, 7, 3, 3, 3, 3, 2, 2, 2, 2, 10, 5, 5, 5, 5, 3, 3, 3, 3, 7, 7, 7, 7, 10, 4, 4, 4, 4, 8, 8, 8, 8, 6, 6, 6, 6},
}

// Map with available reels.
var ReelsMap = map[float64]*game.Reels5x{
	88.602035:  &Reels89,
	92.768326:  &Reels93,
	94.088429:  &Reels94,
	95.419380:  &Reels95,
	97.019081:  &Reels97,
	100.599786: &Reels101,
	111.172239: &Reels111,
}

func FindReels(mrtp float64) (rtp float64, reels game.Reels) {
	for p, r := range ReelsMap {
		if math.Abs(mrtp-p) < math.Abs(mrtp-rtp) {
			rtp, reels = p, r
		}
	}
	return
}

// Lined payment.
var LinePay = [10][5]float64{
	{0, 4, 60, 200, 1000}, //  1 seven
	{0, 0, 40, 100, 300},  //  2 bell
	{0, 0, 20, 80, 200},   //  3 melon
	{0, 0, 20, 80, 200},   //  4 grapes
	{0, 0, 8, 40, 100},    //  5 plum
	{0, 0, 8, 40, 100},    //  6 orange
	{0, 0, 8, 40, 100},    //  7 lemon
	{0, 0, 8, 40, 100},    //  8 cherry
	{0, 0, 0, 0, 0},       //  9 wild
	{0, 0, 0, 0, 0},       // 10 star
}

// Scatters payment.
var ScatPay = [5]float64{0, 0, 2, 20, 500} // star

const (
	jid = 1 // jackpot ID
)

// Jackpot win combinations.
var Jackpot = [10][5]int{
	{0, 0, 0, 0, 0}, //  1 seven
	{0, 0, 0, 0, 0}, //  2 bell
	{0, 0, 0, 0, 0}, //  3 melon
	{0, 0, 0, 0, 0}, //  4 grapes
	{0, 0, 0, 0, 0}, //  5 plum
	{0, 0, 0, 0, 0}, //  6 orange
	{0, 0, 0, 0, 0}, //  7 lemon
	{0, 0, 0, 0, 0}, //  8 cherry
	{0, 0, 0, 0, 0}, //  9 wild
	{0, 0, 0, 0, 0}, // 10 star
}

type Game struct {
	game.Slot5x4 `yaml:",inline"`
}

func NewGame(rtp float64) *Game {
	return &Game{
		Slot5x4: game.Slot5x4{
			RTP: rtp,
			SBL: game.MakeBitNum(40),
			Bet: 1,
		},
	}
}

const wild, scat = 9, 10

var bl = game.BetLinesNvm40

func (g *Game) Scanner(screen game.Screen, wins *game.Wins) {
	g.ScanLined(screen, wins)
	g.ScanScatters(screen, wins)
}

// Lined symbols calculation.
func (g *Game) ScanLined(screen game.Screen, wins *game.Wins) {
	for li := g.SBL.Next(0); li != 0; li = g.SBL.Next(li) {
		var line = bl.Line(li)

		var syml, numl = screen.Pos(1, line), 1
		for x := 2; x <= 5; x++ {
			var sx = screen.Pos(x, line)
			if sx != syml && sx != wild {
				break
			}
			numl++
		}

		if pay := LinePay[syml-1][numl-1]; pay > 0 {
			*wins = append(*wins, game.WinItem{
				Pay:  g.Bet * pay,
				Mult: 1,
				Sym:  syml,
				Num:  numl,
				Line: li,
				XY:   line.CopyL(numl),
			})
		}
	}
}

// Scatters calculation.
func (g *Game) ScanScatters(screen game.Screen, wins *game.Wins) {
	if count := screen.ScatNum(scat); count >= 3 {
		var pay = ScatPay[count-1]
		*wins = append(*wins, game.WinItem{
			Pay:  g.Bet * float64(g.SBL.Num()) * pay,
			Mult: 1,
			Sym:  scat,
			Num:  count,
			XY:   screen.ScatPos(scat),
		})
	}
}

func (g *Game) Spin(screen game.Screen) {
	var _, reels = FindReels(g.RTP)
	screen.Spin(reels)
}

func (g *Game) SetLines(sbl game.Bitset) error {
	var mask game.Bitset = (1<<len(bl) - 1) << 1
	if sbl == 0 {
		return game.ErrNoLineset
	}
	if sbl&^mask != 0 {
		return game.ErrLinesetOut
	}
	if g.FreeSpins() > 0 {
		return game.ErrNoFeature
	}
	g.SBL = sbl
	return nil
}
