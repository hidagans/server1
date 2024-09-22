package powerstars

import (
	"math"

	slot "github.com/slotopol/server/game/slot"
)

// reels lengths [33, 33, 33, 33, 33], total reshuffles 39135393
// RTP[---] = 69.062907%
// RTP[*--] = 379.939305%
// RTP[-*-] = 565.086545%
// RTP[--*] = 379.939305%
// RTP[**-] = 3333.277680%
// RTP[-**] = 3261.763642%
// RTP[*-*] = 1598.352673%
// RTP[***] = 9366.391185%
var Reels = slot.Reels5x{
	{2, 7, 8, 6, 4, 3, 7, 6, 5, 4, 8, 6, 7, 4, 1, 8, 6, 7, 5, 1, 8, 3, 2, 5, 6, 2, 3, 5, 4, 3, 8, 7, 5},
	{5, 1, 8, 6, 4, 7, 3, 6, 8, 7, 5, 3, 7, 4, 5, 3, 8, 4, 6, 5, 3, 7, 8, 6, 2, 8, 1, 5, 6, 2, 4, 7, 2},
	{5, 6, 8, 4, 7, 1, 8, 6, 4, 1, 3, 8, 2, 5, 4, 3, 8, 2, 3, 6, 7, 5, 2, 6, 7, 5, 4, 7, 5, 8, 6, 3, 7},
	{6, 8, 2, 6, 3, 8, 7, 3, 1, 5, 7, 1, 3, 8, 2, 5, 8, 4, 6, 5, 4, 8, 2, 4, 7, 3, 6, 7, 5, 6, 7, 5, 4},
	{6, 4, 7, 6, 5, 7, 3, 5, 2, 8, 4, 7, 2, 5, 6, 7, 3, 8, 5, 1, 2, 8, 3, 6, 4, 8, 7, 4, 5, 3, 6, 8, 1},
}

// Map with wild chances.
var ChanceMap = map[float64]float64{
	// free spins: q = 0.036141, 1/q = 27.669, rtpfs = 470.021964%
	// RTP = 69.063(sym) + q*470.02(fg) = 86.049978%
	86.049978: 1 / 82.,
	// free spins: q = 0.039995, 1/q = 25.003, rtpfs = 473.142683%
	// RTP = 69.063(sym) + q*473.14(fg) = 87.986326%
	87.986326: 1 / 74.,
	// free spins: q = 0.044111, 1/q = 22.67, rtpfs = 476.496465%
	// RTP = 69.063(sym) + q*476.5(fg) = 90.081711%
	90.081711: 1 / 67.,
	// free spins: q = 0.046146, 1/q = 21.67, rtpfs = 478.162925%
	// RTP = 69.063(sym) + q*478.16(fg) = 91.128401%
	91.128401: 1 / 64.,
	// free spins: q = 0.047611, 1/q = 21.004, rtpfs = 479.365362%
	// RTP = 69.063(sym) + q*479.37(fg) = 91.885902%
	91.885902: 1 / 62.,
	// free spins: q = 0.051714, 1/q = 19.337, rtpfs = 482.749000%
	// RTP = 69.063(sym) + q*482.75(fg) = 94.027604%
	94.027604: 1 / 57.,
	// free spins: q = 0.05356, 1/q = 18.671, rtpfs = 484.278763%
	// RTP = 69.063(sym) + q*484.28(fg) = 95.000746%
	95.000746: 1 / 55.,
	// free spins: q = 0.055542, 1/q = 18.004, rtpfs = 485.926793%
	// RTP = 69.063(sym) + q*485.93(fg) = 96.052493%
	96.052493: 1 / 53.,
	// free spins: q = 0.058808, 1/q = 17.004, rtpfs = 488.652432%
	// RTP = 69.063(sym) + q*488.65(fg) = 97.799579%
	97.799579: 1 / 50.,
	// free spins: q = 0.062481, 1/q = 16.005, rtpfs = 491.735591%
	// RTP = 69.063(sym) + q*491.74(fg) = 99.787205%
	99.787205: 1 / 47.,
	// free spins: q = 0.083289, 1/q = 12.006, rtpfs = 509.549567%
	// RTP = 69.063(sym) + q*509.55(fg) = 111.502592%
	111.502592: 1 / 35.,
}

func FindChance(mrtp float64) (rtp float64, chance float64) {
	for p, c := range ChanceMap {
		if math.Abs(mrtp-p) < math.Abs(mrtp-rtp) {
			rtp, chance = p, c
		}
	}
	return
}

// Returns the probability of getting at least one star on the 3 reels,
// including several stars at once.
func AnyStarProb(b float64) float64 {
	return (b*b + (b-1)*b + (b-1)*(b-1)) / b / b / b
}