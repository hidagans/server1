package happysanta50

import (
	"github.com/slotopol/server/game/slot"
)

// reels lengths [52, 55, 56, 55, 52], total reshuffles 458057600
// RTP = 81.515(lined) + 6.1288(scatter) = 87.643964%
var Reels876 = slot.Reels5x{
	{10, 9, 9, 9, 9, 9, 9, 9, 2, 10, 10, 10, 10, 10, 10, 10, 7, 7, 7, 7, 7, 7, 7, 2, 5, 5, 5, 5, 5, 9, 4, 4, 4, 4, 4, 6, 6, 6, 6, 6, 8, 8, 8, 8, 8, 8, 8, 3, 3, 3, 3, 3},
	{10, 7, 7, 7, 7, 7, 7, 7, 9, 4, 4, 4, 4, 4, 2, 9, 9, 9, 9, 9, 9, 9, 1, 1, 1, 1, 3, 3, 3, 3, 3, 10, 10, 10, 10, 10, 10, 10, 8, 8, 8, 8, 8, 8, 8, 6, 6, 6, 6, 6, 5, 5, 5, 5, 5},
	{10, 8, 8, 8, 8, 8, 8, 8, 1, 1, 1, 1, 9, 2, 7, 7, 7, 7, 7, 7, 7, 5, 5, 5, 5, 5, 2, 6, 6, 6, 6, 6, 3, 3, 3, 3, 3, 9, 9, 9, 9, 9, 9, 9, 4, 4, 4, 4, 4, 10, 10, 10, 10, 10, 10, 10},
	{10, 7, 7, 7, 7, 7, 7, 7, 9, 4, 4, 4, 4, 4, 2, 9, 9, 9, 9, 9, 9, 9, 1, 1, 1, 1, 3, 3, 3, 3, 3, 10, 10, 10, 10, 10, 10, 10, 8, 8, 8, 8, 8, 8, 8, 6, 6, 6, 6, 6, 5, 5, 5, 5, 5},
	{10, 9, 9, 9, 9, 9, 9, 9, 2, 10, 10, 10, 10, 10, 10, 10, 7, 7, 7, 7, 7, 7, 7, 2, 5, 5, 5, 5, 5, 9, 4, 4, 4, 4, 4, 6, 6, 6, 6, 6, 8, 8, 8, 8, 8, 8, 8, 3, 3, 3, 3, 3},
}

// reels lengths [52, 55, 56, 55, 52], total reshuffles 458057600
// RTP = 82.203(lined) + 6.1288(scatter) = 88.332215%
var Reels883 = slot.Reels5x{
	{10, 10, 10, 10, 10, 10, 10, 9, 9, 9, 9, 9, 9, 9, 5, 5, 5, 5, 5, 5, 3, 3, 3, 3, 3, 8, 8, 8, 8, 8, 8, 8, 2, 7, 7, 7, 7, 7, 7, 7, 6, 6, 6, 6, 6, 6, 4, 4, 4, 4, 4, 2},
	{3, 3, 3, 3, 3, 7, 7, 7, 7, 7, 7, 7, 9, 9, 9, 9, 9, 9, 9, 4, 4, 4, 4, 4, 5, 5, 5, 5, 5, 5, 8, 8, 8, 8, 8, 8, 8, 10, 10, 10, 10, 10, 10, 10, 6, 6, 6, 6, 6, 6, 1, 1, 1, 1, 2},
	{8, 8, 8, 8, 8, 8, 8, 7, 7, 7, 7, 7, 7, 7, 2, 5, 5, 5, 5, 5, 5, 4, 4, 4, 4, 4, 3, 3, 3, 3, 3, 6, 6, 6, 6, 6, 6, 10, 10, 10, 10, 10, 10, 10, 9, 9, 9, 9, 9, 9, 9, 1, 1, 1, 1, 2},
	{3, 3, 3, 3, 3, 7, 7, 7, 7, 7, 7, 7, 9, 9, 9, 9, 9, 9, 9, 4, 4, 4, 4, 4, 5, 5, 5, 5, 5, 5, 8, 8, 8, 8, 8, 8, 8, 10, 10, 10, 10, 10, 10, 10, 6, 6, 6, 6, 6, 6, 1, 1, 1, 1, 2},
	{10, 10, 10, 10, 10, 10, 10, 9, 9, 9, 9, 9, 9, 9, 5, 5, 5, 5, 5, 5, 3, 3, 3, 3, 3, 8, 8, 8, 8, 8, 8, 8, 2, 7, 7, 7, 7, 7, 7, 7, 6, 6, 6, 6, 6, 6, 4, 4, 4, 4, 4, 2},
}

// reels lengths [51, 55, 56, 55, 51], total reshuffles 440609400
// RTP = 82.986(lined) + 6.3107(scatter) = 89.296878%
var Reels892 = slot.Reels5x{
	{8, 8, 8, 8, 8, 8, 8, 9, 9, 9, 9, 9, 9, 5, 5, 5, 5, 5, 5, 2, 10, 10, 10, 10, 10, 10, 10, 4, 4, 4, 4, 4, 3, 3, 3, 3, 3, 6, 6, 6, 6, 6, 6, 7, 7, 7, 7, 7, 7, 7, 2},
	{8, 8, 8, 8, 8, 8, 8, 9, 9, 9, 9, 9, 9, 9, 10, 10, 10, 10, 10, 10, 10, 4, 4, 4, 4, 4, 5, 5, 5, 5, 5, 5, 6, 6, 6, 6, 6, 6, 1, 1, 1, 1, 2, 3, 3, 3, 3, 3, 7, 7, 7, 7, 7, 7, 7},
	{7, 7, 7, 7, 7, 7, 7, 2, 4, 4, 4, 4, 4, 1, 1, 1, 1, 9, 9, 9, 9, 9, 9, 9, 5, 5, 5, 5, 5, 5, 8, 8, 8, 8, 8, 8, 8, 3, 3, 3, 3, 3, 6, 6, 6, 6, 6, 6, 10, 10, 10, 10, 10, 10, 10, 2},
	{8, 8, 8, 8, 8, 8, 8, 9, 9, 9, 9, 9, 9, 9, 10, 10, 10, 10, 10, 10, 10, 4, 4, 4, 4, 4, 5, 5, 5, 5, 5, 5, 6, 6, 6, 6, 6, 6, 1, 1, 1, 1, 2, 3, 3, 3, 3, 3, 7, 7, 7, 7, 7, 7, 7},
	{8, 8, 8, 8, 8, 8, 8, 9, 9, 9, 9, 9, 9, 5, 5, 5, 5, 5, 5, 2, 10, 10, 10, 10, 10, 10, 10, 4, 4, 4, 4, 4, 3, 3, 3, 3, 3, 6, 6, 6, 6, 6, 6, 7, 7, 7, 7, 7, 7, 7, 2},
}

// reels lengths [51, 54, 55, 54, 51], total reshuffles 417148380
// RTP = 84.14(lined) + 6.5282(scatter) = 90.668061%
var Reels906 = slot.Reels5x{
	{10, 10, 10, 10, 10, 10, 10, 3, 3, 3, 3, 3, 6, 6, 6, 6, 6, 6, 9, 9, 9, 9, 9, 9, 9, 5, 5, 5, 5, 5, 2, 8, 8, 8, 8, 8, 8, 8, 4, 4, 4, 4, 4, 7, 7, 7, 7, 7, 7, 7, 2},
	{6, 6, 6, 6, 6, 6, 5, 5, 5, 5, 5, 1, 1, 1, 1, 3, 3, 3, 3, 3, 2, 8, 8, 8, 8, 8, 8, 8, 10, 10, 10, 10, 10, 10, 10, 7, 7, 7, 7, 7, 7, 7, 4, 4, 4, 4, 4, 9, 9, 9, 9, 9, 9, 9},
	{1, 1, 1, 1, 10, 10, 10, 10, 10, 10, 10, 2, 9, 9, 9, 9, 9, 9, 9, 6, 6, 6, 6, 6, 6, 5, 5, 5, 5, 5, 7, 7, 7, 7, 7, 7, 7, 3, 3, 3, 3, 3, 8, 8, 8, 8, 8, 8, 8, 4, 4, 4, 4, 4, 2},
	{6, 6, 6, 6, 6, 6, 5, 5, 5, 5, 5, 1, 1, 1, 1, 3, 3, 3, 3, 3, 2, 8, 8, 8, 8, 8, 8, 8, 10, 10, 10, 10, 10, 10, 10, 7, 7, 7, 7, 7, 7, 7, 4, 4, 4, 4, 4, 9, 9, 9, 9, 9, 9, 9},
	{10, 10, 10, 10, 10, 10, 10, 3, 3, 3, 3, 3, 6, 6, 6, 6, 6, 6, 9, 9, 9, 9, 9, 9, 9, 5, 5, 5, 5, 5, 2, 8, 8, 8, 8, 8, 8, 8, 4, 4, 4, 4, 4, 7, 7, 7, 7, 7, 7, 7, 2},
}

// reels lengths [51, 54, 55, 54, 51], total reshuffles 417148380
// RTP = 85.361(lined) + 6.5282(scatter) = 91.889572%
var Reels918 = slot.Reels5x{
	{8, 8, 8, 8, 8, 8, 8, 9, 9, 9, 9, 9, 9, 5, 5, 5, 5, 5, 5, 2, 10, 10, 10, 10, 10, 10, 10, 4, 4, 4, 4, 4, 3, 3, 3, 3, 3, 6, 6, 6, 6, 6, 6, 7, 7, 7, 7, 7, 7, 7, 2},
	{4, 4, 4, 4, 4, 5, 5, 5, 5, 5, 5, 7, 7, 7, 7, 7, 7, 7, 1, 1, 1, 1, 3, 3, 3, 3, 3, 2, 10, 10, 10, 10, 10, 10, 10, 6, 6, 6, 6, 6, 6, 8, 8, 8, 8, 8, 8, 8, 9, 9, 9, 9, 9, 9},
	{6, 6, 6, 6, 6, 6, 5, 5, 5, 5, 5, 5, 2, 1, 1, 1, 1, 10, 10, 10, 10, 10, 10, 10, 9, 9, 9, 9, 9, 9, 4, 4, 4, 4, 4, 2, 7, 7, 7, 7, 7, 7, 7, 8, 8, 8, 8, 8, 8, 8, 3, 3, 3, 3, 3},
	{4, 4, 4, 4, 4, 5, 5, 5, 5, 5, 5, 7, 7, 7, 7, 7, 7, 7, 1, 1, 1, 1, 3, 3, 3, 3, 3, 2, 10, 10, 10, 10, 10, 10, 10, 6, 6, 6, 6, 6, 6, 8, 8, 8, 8, 8, 8, 8, 9, 9, 9, 9, 9, 9},
	{8, 8, 8, 8, 8, 8, 8, 9, 9, 9, 9, 9, 9, 5, 5, 5, 5, 5, 5, 2, 10, 10, 10, 10, 10, 10, 10, 4, 4, 4, 4, 4, 3, 3, 3, 3, 3, 6, 6, 6, 6, 6, 6, 7, 7, 7, 7, 7, 7, 7, 2},
}

// reels lengths [50, 53, 54, 53, 50], total reshuffles 379215000
// RTP = 86.185(lined) + 6.9633(scatter) = 93.148401%
var Reels931 = slot.Reels5x{
	{3, 3, 3, 3, 3, 7, 7, 7, 7, 7, 7, 4, 4, 4, 4, 4, 6, 6, 6, 6, 6, 6, 2, 8, 8, 8, 8, 8, 8, 8, 10, 10, 10, 10, 10, 10, 10, 9, 9, 9, 9, 9, 9, 9, 2, 5, 5, 5, 5, 5},
	{3, 3, 3, 3, 3, 4, 4, 4, 4, 4, 1, 1, 1, 1, 10, 10, 10, 10, 10, 10, 10, 2, 8, 8, 8, 8, 8, 8, 8, 6, 6, 6, 6, 6, 6, 5, 5, 5, 5, 5, 7, 7, 7, 7, 7, 7, 9, 9, 9, 9, 9, 9, 9},
	{1, 1, 1, 1, 8, 8, 8, 8, 8, 8, 8, 2, 3, 3, 3, 3, 3, 4, 4, 4, 4, 4, 6, 6, 6, 6, 6, 6, 5, 5, 5, 5, 5, 7, 7, 7, 7, 7, 7, 2, 10, 10, 10, 10, 10, 10, 10, 9, 9, 9, 9, 9, 9, 9},
	{3, 3, 3, 3, 3, 4, 4, 4, 4, 4, 1, 1, 1, 1, 10, 10, 10, 10, 10, 10, 10, 2, 8, 8, 8, 8, 8, 8, 8, 6, 6, 6, 6, 6, 6, 5, 5, 5, 5, 5, 7, 7, 7, 7, 7, 7, 9, 9, 9, 9, 9, 9, 9},
	{3, 3, 3, 3, 3, 7, 7, 7, 7, 7, 7, 4, 4, 4, 4, 4, 6, 6, 6, 6, 6, 6, 2, 8, 8, 8, 8, 8, 8, 8, 10, 10, 10, 10, 10, 10, 10, 9, 9, 9, 9, 9, 9, 9, 2, 5, 5, 5, 5, 5},
}

// reels lengths [50, 53, 54, 53, 50], total reshuffles 379215000
// RTP = 87.369(lined) + 6.9633(scatter) = 94.332391%
var Reels943 = slot.Reels5x{
	{4, 4, 4, 4, 4, 7, 7, 7, 7, 7, 7, 7, 5, 5, 5, 5, 5, 9, 9, 9, 9, 9, 9, 6, 6, 6, 6, 6, 6, 2, 10, 10, 10, 10, 10, 10, 10, 2, 8, 8, 8, 8, 8, 8, 8, 3, 3, 3, 3, 3},
	{10, 10, 10, 10, 10, 10, 10, 4, 4, 4, 4, 4, 1, 1, 1, 1, 2, 9, 9, 9, 9, 9, 9, 8, 8, 8, 8, 8, 8, 8, 3, 3, 3, 3, 3, 5, 5, 5, 5, 5, 7, 7, 7, 7, 7, 7, 7, 6, 6, 6, 6, 6, 6},
	{8, 8, 8, 8, 8, 8, 8, 1, 1, 1, 1, 2, 3, 3, 3, 3, 3, 6, 6, 6, 6, 6, 6, 7, 7, 7, 7, 7, 7, 7, 10, 10, 10, 10, 10, 10, 10, 5, 5, 5, 5, 5, 2, 4, 4, 4, 4, 4, 9, 9, 9, 9, 9, 9},
	{10, 10, 10, 10, 10, 10, 10, 4, 4, 4, 4, 4, 1, 1, 1, 1, 2, 9, 9, 9, 9, 9, 9, 8, 8, 8, 8, 8, 8, 8, 3, 3, 3, 3, 3, 5, 5, 5, 5, 5, 7, 7, 7, 7, 7, 7, 7, 6, 6, 6, 6, 6, 6},
	{4, 4, 4, 4, 4, 7, 7, 7, 7, 7, 7, 7, 5, 5, 5, 5, 5, 9, 9, 9, 9, 9, 9, 6, 6, 6, 6, 6, 6, 2, 10, 10, 10, 10, 10, 10, 10, 2, 8, 8, 8, 8, 8, 8, 8, 3, 3, 3, 3, 3},
}

// reels lengths [50, 53, 54, 53, 50], total reshuffles 379215000
// RTP = 88.68(lined) + 6.9633(scatter) = 95.643745%
var Reels956 = slot.Reels5x{
	{4, 4, 4, 4, 4, 6, 6, 6, 6, 6, 6, 9, 9, 9, 9, 9, 9, 2, 5, 5, 5, 5, 5, 5, 7, 7, 7, 7, 7, 7, 7, 3, 3, 3, 3, 3, 8, 8, 8, 8, 8, 8, 8, 2, 10, 10, 10, 10, 10, 10},
	{4, 4, 4, 4, 4, 5, 5, 5, 5, 5, 5, 8, 8, 8, 8, 8, 8, 8, 10, 10, 10, 10, 10, 10, 1, 1, 1, 1, 6, 6, 6, 6, 6, 6, 9, 9, 9, 9, 9, 9, 2, 3, 3, 3, 3, 3, 7, 7, 7, 7, 7, 7, 7},
	{5, 5, 5, 5, 5, 5, 6, 6, 6, 6, 6, 6, 2, 8, 8, 8, 8, 8, 8, 8, 10, 10, 10, 10, 10, 10, 3, 3, 3, 3, 3, 9, 9, 9, 9, 9, 9, 1, 1, 1, 1, 7, 7, 7, 7, 7, 7, 7, 2, 4, 4, 4, 4, 4},
	{4, 4, 4, 4, 4, 5, 5, 5, 5, 5, 5, 8, 8, 8, 8, 8, 8, 8, 10, 10, 10, 10, 10, 10, 1, 1, 1, 1, 6, 6, 6, 6, 6, 6, 9, 9, 9, 9, 9, 9, 2, 3, 3, 3, 3, 3, 7, 7, 7, 7, 7, 7, 7},
	{4, 4, 4, 4, 4, 6, 6, 6, 6, 6, 6, 9, 9, 9, 9, 9, 9, 2, 5, 5, 5, 5, 5, 5, 7, 7, 7, 7, 7, 7, 7, 3, 3, 3, 3, 3, 8, 8, 8, 8, 8, 8, 8, 2, 10, 10, 10, 10, 10, 10},
}

// reels lengths [49, 52, 53, 52, 49], total reshuffles 344092112
// RTP = 89.63(lined) + 7.4381(scatter) = 97.067837%
var Reels970 = slot.Reels5x{
	{5, 5, 5, 5, 5, 5, 10, 10, 10, 10, 10, 10, 10, 2, 8, 8, 8, 8, 8, 8, 6, 6, 6, 6, 6, 6, 3, 3, 3, 3, 3, 9, 9, 9, 9, 9, 9, 4, 4, 4, 4, 4, 2, 7, 7, 7, 7, 7, 7},
	{3, 3, 3, 3, 3, 8, 8, 8, 8, 8, 8, 6, 6, 6, 6, 6, 6, 7, 7, 7, 7, 7, 7, 2, 10, 10, 10, 10, 10, 10, 10, 9, 9, 9, 9, 9, 9, 5, 5, 5, 5, 5, 5, 4, 4, 4, 4, 4, 1, 1, 1, 1},
	{8, 8, 8, 8, 8, 8, 9, 9, 9, 9, 9, 9, 7, 7, 7, 7, 7, 7, 10, 10, 10, 10, 10, 10, 10, 1, 1, 1, 1, 2, 4, 4, 4, 4, 4, 3, 3, 3, 3, 3, 2, 6, 6, 6, 6, 6, 6, 5, 5, 5, 5, 5, 5},
	{3, 3, 3, 3, 3, 8, 8, 8, 8, 8, 8, 6, 6, 6, 6, 6, 6, 7, 7, 7, 7, 7, 7, 2, 10, 10, 10, 10, 10, 10, 10, 9, 9, 9, 9, 9, 9, 5, 5, 5, 5, 5, 5, 4, 4, 4, 4, 4, 1, 1, 1, 1},
	{5, 5, 5, 5, 5, 5, 10, 10, 10, 10, 10, 10, 10, 2, 8, 8, 8, 8, 8, 8, 6, 6, 6, 6, 6, 6, 3, 3, 3, 3, 3, 9, 9, 9, 9, 9, 9, 4, 4, 4, 4, 4, 2, 7, 7, 7, 7, 7, 7},
}

// reels lengths [48, 51, 52, 51, 48], total reshuffles 311620608
// RTP = 90.299(lined) + 7.9573(scatter) = 98.256580%
var Reels982 = slot.Reels5x{
	{4, 4, 4, 4, 4, 5, 5, 5, 5, 5, 2, 7, 7, 7, 7, 7, 7, 6, 6, 6, 6, 6, 2, 10, 10, 10, 10, 10, 10, 10, 8, 8, 8, 8, 8, 8, 9, 9, 9, 9, 9, 9, 9, 3, 3, 3, 3, 3},
	{1, 1, 1, 1, 6, 6, 6, 6, 6, 5, 5, 5, 5, 5, 4, 4, 4, 4, 4, 2, 3, 3, 3, 3, 3, 9, 9, 9, 9, 9, 9, 9, 10, 10, 10, 10, 10, 10, 10, 7, 7, 7, 7, 7, 7, 8, 8, 8, 8, 8, 8},
	{10, 10, 10, 10, 10, 10, 10, 3, 3, 3, 3, 3, 5, 5, 5, 5, 5, 6, 6, 6, 6, 6, 2, 1, 1, 1, 1, 9, 9, 9, 9, 9, 9, 9, 7, 7, 7, 7, 7, 7, 2, 8, 8, 8, 8, 8, 8, 4, 4, 4, 4, 4},
	{1, 1, 1, 1, 6, 6, 6, 6, 6, 5, 5, 5, 5, 5, 4, 4, 4, 4, 4, 2, 3, 3, 3, 3, 3, 9, 9, 9, 9, 9, 9, 9, 10, 10, 10, 10, 10, 10, 10, 7, 7, 7, 7, 7, 7, 8, 8, 8, 8, 8, 8},
	{4, 4, 4, 4, 4, 5, 5, 5, 5, 5, 2, 7, 7, 7, 7, 7, 7, 6, 6, 6, 6, 6, 2, 10, 10, 10, 10, 10, 10, 10, 8, 8, 8, 8, 8, 8, 9, 9, 9, 9, 9, 9, 9, 3, 3, 3, 3, 3},
}

// reels lengths [49, 52, 53, 52, 49], total reshuffles 344092112
// RTP = 92.073(lined) + 7.4381(scatter) = 99.511302%
var Reels995 = slot.Reels5x{
	{9, 9, 9, 9, 9, 9, 9, 4, 4, 4, 4, 4, 4, 8, 8, 8, 8, 8, 8, 6, 6, 6, 6, 6, 2, 3, 3, 3, 3, 3, 5, 5, 5, 5, 5, 2, 7, 7, 7, 7, 7, 7, 10, 10, 10, 10, 10, 10, 10},
	{10, 10, 10, 10, 10, 10, 10, 7, 7, 7, 7, 7, 7, 3, 3, 3, 3, 3, 6, 6, 6, 6, 6, 5, 5, 5, 5, 5, 4, 4, 4, 4, 4, 4, 2, 9, 9, 9, 9, 9, 9, 9, 1, 1, 1, 1, 8, 8, 8, 8, 8, 8},
	{6, 6, 6, 6, 6, 4, 4, 4, 4, 4, 4, 5, 5, 5, 5, 5, 1, 1, 1, 1, 3, 3, 3, 3, 3, 8, 8, 8, 8, 8, 8, 2, 10, 10, 10, 10, 10, 10, 10, 9, 9, 9, 9, 9, 9, 9, 7, 7, 7, 7, 7, 7, 2},
	{10, 10, 10, 10, 10, 10, 10, 7, 7, 7, 7, 7, 7, 3, 3, 3, 3, 3, 6, 6, 6, 6, 6, 5, 5, 5, 5, 5, 4, 4, 4, 4, 4, 4, 2, 9, 9, 9, 9, 9, 9, 9, 1, 1, 1, 1, 8, 8, 8, 8, 8, 8},
	{9, 9, 9, 9, 9, 9, 9, 4, 4, 4, 4, 4, 4, 8, 8, 8, 8, 8, 8, 6, 6, 6, 6, 6, 2, 3, 3, 3, 3, 3, 5, 5, 5, 5, 5, 2, 7, 7, 7, 7, 7, 7, 10, 10, 10, 10, 10, 10, 10},
}

// reels lengths [48, 51, 52, 51, 48], total reshuffles 311620608
// RTP = 91.818(lined) + 7.9573(scatter) = 99.775308%
var Reels997 = slot.Reels5x{
	{4, 4, 4, 4, 4, 8, 8, 8, 8, 8, 8, 2, 6, 6, 6, 6, 6, 3, 3, 3, 3, 3, 2, 5, 5, 5, 5, 5, 5, 7, 7, 7, 7, 7, 7, 9, 9, 9, 9, 9, 9, 10, 10, 10, 10, 10, 10, 10},
	{7, 7, 7, 7, 7, 7, 9, 9, 9, 9, 9, 9, 1, 1, 1, 1, 4, 4, 4, 4, 4, 6, 6, 6, 6, 6, 2, 5, 5, 5, 5, 5, 5, 10, 10, 10, 10, 10, 10, 10, 3, 3, 3, 3, 3, 8, 8, 8, 8, 8, 8},
	{10, 10, 10, 10, 10, 10, 10, 3, 3, 3, 3, 3, 1, 1, 1, 1, 4, 4, 4, 4, 4, 2, 5, 5, 5, 5, 5, 5, 8, 8, 8, 8, 8, 8, 9, 9, 9, 9, 9, 9, 2, 7, 7, 7, 7, 7, 7, 6, 6, 6, 6, 6},
	{7, 7, 7, 7, 7, 7, 9, 9, 9, 9, 9, 9, 1, 1, 1, 1, 4, 4, 4, 4, 4, 6, 6, 6, 6, 6, 2, 5, 5, 5, 5, 5, 5, 10, 10, 10, 10, 10, 10, 10, 3, 3, 3, 3, 3, 8, 8, 8, 8, 8, 8},
	{4, 4, 4, 4, 4, 8, 8, 8, 8, 8, 8, 2, 6, 6, 6, 6, 6, 3, 3, 3, 3, 3, 2, 5, 5, 5, 5, 5, 5, 7, 7, 7, 7, 7, 7, 9, 9, 9, 9, 9, 9, 10, 10, 10, 10, 10, 10, 10},
}

// reels lengths [49, 52, 53, 52, 49], total reshuffles 344092112
// RTP = 93.277(lined) + 7.4381(scatter) = 100.714730%
var Reels100 = slot.Reels5x{
	{9, 9, 9, 9, 9, 9, 3, 3, 3, 3, 3, 6, 6, 6, 6, 6, 6, 10, 10, 10, 10, 10, 10, 10, 2, 4, 4, 4, 4, 4, 4, 8, 8, 8, 8, 8, 8, 5, 5, 5, 5, 5, 7, 7, 7, 7, 7, 7, 2},
	{4, 4, 4, 4, 4, 4, 8, 8, 8, 8, 8, 8, 2, 7, 7, 7, 7, 7, 7, 5, 5, 5, 5, 5, 9, 9, 9, 9, 9, 9, 3, 3, 3, 3, 3, 6, 6, 6, 6, 6, 6, 10, 10, 10, 10, 10, 10, 10, 1, 1, 1, 1},
	{4, 4, 4, 4, 4, 4, 3, 3, 3, 3, 3, 8, 8, 8, 8, 8, 8, 6, 6, 6, 6, 6, 6, 7, 7, 7, 7, 7, 7, 10, 10, 10, 10, 10, 10, 10, 2, 1, 1, 1, 1, 5, 5, 5, 5, 5, 2, 9, 9, 9, 9, 9, 9},
	{4, 4, 4, 4, 4, 4, 8, 8, 8, 8, 8, 8, 2, 7, 7, 7, 7, 7, 7, 5, 5, 5, 5, 5, 9, 9, 9, 9, 9, 9, 3, 3, 3, 3, 3, 6, 6, 6, 6, 6, 6, 10, 10, 10, 10, 10, 10, 10, 1, 1, 1, 1},
	{9, 9, 9, 9, 9, 9, 3, 3, 3, 3, 3, 6, 6, 6, 6, 6, 6, 10, 10, 10, 10, 10, 10, 10, 2, 4, 4, 4, 4, 4, 4, 8, 8, 8, 8, 8, 8, 5, 5, 5, 5, 5, 7, 7, 7, 7, 7, 7, 2},
}

// reels lengths [48, 51, 52, 51, 48], total reshuffles 311620608
// RTP = 93.109(lined) + 7.9573(scatter) = 101.065874%
var Reels101 = slot.Reels5x{
	{9, 9, 9, 9, 9, 9, 5, 5, 5, 5, 5, 5, 10, 10, 10, 10, 10, 10, 3, 3, 3, 3, 3, 7, 7, 7, 7, 7, 7, 2, 6, 6, 6, 6, 6, 6, 8, 8, 8, 8, 8, 8, 4, 4, 4, 4, 4, 2},
	{5, 5, 5, 5, 5, 5, 10, 10, 10, 10, 10, 10, 4, 4, 4, 4, 4, 1, 1, 1, 1, 2, 3, 3, 3, 3, 3, 8, 8, 8, 8, 8, 8, 7, 7, 7, 7, 7, 7, 6, 6, 6, 6, 6, 6, 9, 9, 9, 9, 9, 9},
	{8, 8, 8, 8, 8, 8, 4, 4, 4, 4, 4, 2, 7, 7, 7, 7, 7, 7, 5, 5, 5, 5, 5, 5, 6, 6, 6, 6, 6, 6, 10, 10, 10, 10, 10, 10, 1, 1, 1, 1, 2, 3, 3, 3, 3, 3, 9, 9, 9, 9, 9, 9},
	{5, 5, 5, 5, 5, 5, 10, 10, 10, 10, 10, 10, 4, 4, 4, 4, 4, 1, 1, 1, 1, 2, 3, 3, 3, 3, 3, 8, 8, 8, 8, 8, 8, 7, 7, 7, 7, 7, 7, 6, 6, 6, 6, 6, 6, 9, 9, 9, 9, 9, 9},
	{9, 9, 9, 9, 9, 9, 5, 5, 5, 5, 5, 5, 10, 10, 10, 10, 10, 10, 3, 3, 3, 3, 3, 7, 7, 7, 7, 7, 7, 2, 6, 6, 6, 6, 6, 6, 8, 8, 8, 8, 8, 8, 4, 4, 4, 4, 4, 2},
}

// reels lengths [49, 52, 53, 52, 49], total reshuffles 344092112
// RTP = 94.687(lined) + 7.4381(scatter) = 102.124789%
var Reels102 = slot.Reels5x{
	{9, 9, 9, 9, 9, 9, 5, 5, 5, 5, 5, 5, 10, 10, 10, 10, 10, 10, 7, 7, 7, 7, 7, 7, 8, 8, 8, 8, 8, 8, 6, 6, 6, 6, 6, 6, 2, 4, 4, 4, 4, 4, 4, 3, 3, 3, 3, 3, 2},
	{3, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 8, 8, 8, 8, 8, 8, 9, 9, 9, 9, 9, 9, 1, 1, 1, 1, 5, 5, 5, 5, 5, 5, 2, 10, 10, 10, 10, 10, 10, 7, 7, 7, 7, 7, 7, 6, 6, 6, 6, 6, 6},
	{3, 3, 3, 3, 3, 1, 1, 1, 1, 8, 8, 8, 8, 8, 8, 2, 7, 7, 7, 7, 7, 7, 6, 6, 6, 6, 6, 6, 4, 4, 4, 4, 4, 4, 10, 10, 10, 10, 10, 10, 2, 5, 5, 5, 5, 5, 5, 9, 9, 9, 9, 9, 9},
	{3, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 8, 8, 8, 8, 8, 8, 9, 9, 9, 9, 9, 9, 1, 1, 1, 1, 5, 5, 5, 5, 5, 5, 2, 10, 10, 10, 10, 10, 10, 7, 7, 7, 7, 7, 7, 6, 6, 6, 6, 6, 6},
	{9, 9, 9, 9, 9, 9, 5, 5, 5, 5, 5, 5, 10, 10, 10, 10, 10, 10, 7, 7, 7, 7, 7, 7, 8, 8, 8, 8, 8, 8, 6, 6, 6, 6, 6, 6, 2, 4, 4, 4, 4, 4, 4, 3, 3, 3, 3, 3, 2},
}

// reels lengths [46, 49, 50, 49, 46], total reshuffles 254025800
// RTP = 97.315(lined) + 9.151(scatter) = 106.465502%
var Reels106 = slot.Reels5x{
	{10, 10, 10, 10, 10, 10, 2, 4, 4, 4, 4, 4, 8, 8, 8, 8, 8, 8, 7, 7, 7, 7, 7, 7, 2, 3, 3, 3, 3, 3, 5, 5, 5, 5, 5, 6, 6, 6, 6, 6, 9, 9, 9, 9, 9, 9},
	{8, 8, 8, 8, 8, 8, 1, 1, 1, 1, 9, 9, 9, 9, 9, 9, 2, 5, 5, 5, 5, 5, 4, 4, 4, 4, 4, 7, 7, 7, 7, 7, 7, 3, 3, 3, 3, 3, 6, 6, 6, 6, 6, 10, 10, 10, 10, 10, 10},
	{7, 7, 7, 7, 7, 7, 9, 9, 9, 9, 9, 9, 8, 8, 8, 8, 8, 8, 2, 1, 1, 1, 1, 3, 3, 3, 3, 3, 2, 5, 5, 5, 5, 5, 6, 6, 6, 6, 6, 10, 10, 10, 10, 10, 10, 4, 4, 4, 4, 4},
	{8, 8, 8, 8, 8, 8, 1, 1, 1, 1, 9, 9, 9, 9, 9, 9, 2, 5, 5, 5, 5, 5, 4, 4, 4, 4, 4, 7, 7, 7, 7, 7, 7, 3, 3, 3, 3, 3, 6, 6, 6, 6, 6, 10, 10, 10, 10, 10, 10},
	{10, 10, 10, 10, 10, 10, 2, 4, 4, 4, 4, 4, 8, 8, 8, 8, 8, 8, 7, 7, 7, 7, 7, 7, 2, 3, 3, 3, 3, 3, 5, 5, 5, 5, 5, 6, 6, 6, 6, 6, 9, 9, 9, 9, 9, 9},
}

// reels lengths [47, 50, 51, 50, 47], total reshuffles 281647500
// RTP = 99.541(lined) + 8.5262(scatter) = 108.067564%
var Reels108 = slot.Reels5x{
	{9, 9, 9, 9, 9, 9, 7, 7, 7, 7, 7, 7, 2, 5, 5, 5, 5, 5, 6, 6, 6, 6, 6, 2, 8, 8, 8, 8, 8, 8, 4, 4, 4, 4, 4, 4, 3, 3, 3, 3, 3, 10, 10, 10, 10, 10, 10},
	{4, 4, 4, 4, 4, 4, 3, 3, 3, 3, 3, 8, 8, 8, 8, 8, 8, 7, 7, 7, 7, 7, 7, 9, 9, 9, 9, 9, 9, 2, 10, 10, 10, 10, 10, 10, 1, 1, 1, 1, 5, 5, 5, 5, 5, 6, 6, 6, 6, 6},
	{6, 6, 6, 6, 6, 1, 1, 1, 1, 9, 9, 9, 9, 9, 9, 2, 5, 5, 5, 5, 5, 8, 8, 8, 8, 8, 8, 3, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 7, 7, 7, 7, 7, 7, 2, 10, 10, 10, 10, 10, 10},
	{4, 4, 4, 4, 4, 4, 3, 3, 3, 3, 3, 8, 8, 8, 8, 8, 8, 7, 7, 7, 7, 7, 7, 9, 9, 9, 9, 9, 9, 2, 10, 10, 10, 10, 10, 10, 1, 1, 1, 1, 5, 5, 5, 5, 5, 6, 6, 6, 6, 6},
	{9, 9, 9, 9, 9, 9, 7, 7, 7, 7, 7, 7, 2, 5, 5, 5, 5, 5, 6, 6, 6, 6, 6, 2, 8, 8, 8, 8, 8, 8, 4, 4, 4, 4, 4, 4, 3, 3, 3, 3, 3, 10, 10, 10, 10, 10, 10},
}

// reels lengths [47, 50, 51, 50, 47], total reshuffles 281647500
// RTP = 102.11(lined) + 8.5262(scatter) = 110.638032%
var Reels110 = slot.Reels5x{
	{10, 10, 10, 10, 10, 10, 6, 6, 6, 6, 6, 6, 7, 7, 7, 7, 7, 4, 4, 4, 4, 4, 4, 3, 3, 3, 3, 3, 2, 5, 5, 5, 5, 5, 5, 9, 9, 9, 9, 9, 9, 2, 8, 8, 8, 8, 8},
	{6, 6, 6, 6, 6, 6, 7, 7, 7, 7, 7, 10, 10, 10, 10, 10, 10, 4, 4, 4, 4, 4, 4, 5, 5, 5, 5, 5, 5, 3, 3, 3, 3, 3, 9, 9, 9, 9, 9, 9, 8, 8, 8, 8, 8, 1, 1, 1, 1, 2},
	{4, 4, 4, 4, 4, 4, 7, 7, 7, 7, 7, 5, 5, 5, 5, 5, 5, 9, 9, 9, 9, 9, 9, 1, 1, 1, 1, 6, 6, 6, 6, 6, 6, 3, 3, 3, 3, 3, 2, 10, 10, 10, 10, 10, 10, 8, 8, 8, 8, 8, 2},
	{6, 6, 6, 6, 6, 6, 7, 7, 7, 7, 7, 10, 10, 10, 10, 10, 10, 4, 4, 4, 4, 4, 4, 5, 5, 5, 5, 5, 5, 3, 3, 3, 3, 3, 9, 9, 9, 9, 9, 9, 8, 8, 8, 8, 8, 1, 1, 1, 1, 2},
	{10, 10, 10, 10, 10, 10, 6, 6, 6, 6, 6, 6, 7, 7, 7, 7, 7, 4, 4, 4, 4, 4, 4, 3, 3, 3, 3, 3, 2, 5, 5, 5, 5, 5, 5, 9, 9, 9, 9, 9, 9, 2, 8, 8, 8, 8, 8},
}

// Map with available reels.
var ReelsMap = map[float64]*slot.Reels5x{
	87.643964:  &Reels876,
	88.332215:  &Reels883,
	89.296878:  &Reels892,
	90.668061:  &Reels906,
	91.889572:  &Reels918,
	93.148401:  &Reels931,
	94.332391:  &Reels943,
	95.643745:  &Reels956,
	97.067837:  &Reels970,
	98.256580:  &Reels982,
	99.511302:  &Reels995,
	99.775308:  &Reels997,
	100.714730: &Reels100,
	101.065874: &Reels101,
	102.124789: &Reels102,
	106.465502: &Reels106,
	108.067564: &Reels108,
	110.638032: &Reels110,
}