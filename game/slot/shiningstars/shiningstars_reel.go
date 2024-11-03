package shiningstars

import (
	"github.com/slotopol/server/game/slot"
)

// reels lengths [53, 52, 54, 52, 53], total reshuffles 410158944
// RTP = 76.095(lined) + 10.283(scatter) = 86.377743%
var Reels863 = slot.Reels5x{
	{8, 10, 10, 6, 4, 4, 5, 5, 5, 3, 10, 10, 10, 11, 11, 11, 7, 10, 10, 10, 3, 9, 9, 2, 6, 6, 6, 8, 8, 8, 9, 9, 9, 11, 11, 7, 7, 7, 9, 9, 9, 11, 11, 11, 7, 7, 7, 8, 8, 8, 5, 2, 8},
	{7, 5, 5, 5, 9, 9, 9, 4, 4, 6, 6, 6, 11, 11, 11, 8, 8, 8, 7, 7, 1, 10, 10, 8, 8, 8, 7, 10, 10, 10, 9, 8, 7, 7, 7, 5, 8, 9, 11, 11, 11, 2, 10, 10, 10, 9, 9, 9, 2, 11, 11, 6},
	{11, 9, 11, 11, 11, 9, 7, 7, 7, 2, 4, 9, 8, 8, 8, 9, 10, 10, 10, 6, 6, 6, 5, 10, 10, 10, 1, 11, 7, 7, 5, 5, 5, 2, 8, 8, 8, 4, 6, 10, 10, 8, 3, 7, 9, 9, 9, 3, 11, 11, 11, 9, 7, 8},
	{7, 5, 5, 5, 9, 9, 9, 4, 4, 6, 6, 6, 11, 11, 11, 8, 8, 8, 7, 7, 1, 10, 10, 8, 8, 8, 7, 10, 10, 10, 9, 8, 7, 7, 7, 5, 8, 9, 11, 11, 11, 2, 10, 10, 10, 9, 9, 9, 2, 11, 11, 6},
	{8, 10, 10, 6, 4, 4, 5, 5, 5, 3, 10, 10, 10, 11, 11, 11, 7, 10, 10, 10, 3, 9, 9, 2, 6, 6, 6, 8, 8, 8, 9, 9, 9, 11, 11, 7, 7, 7, 9, 9, 9, 11, 11, 11, 7, 7, 7, 8, 8, 8, 5, 2, 8},
}

// reels lengths [49, 52, 54, 52, 49], total reshuffles 350584416
// RTP = 76.288(lined) + 11.515(scatter) = 87.803869%
var Reels878 = slot.Reels5x{
	{3, 10, 7, 8, 4, 10, 9, 9, 9, 11, 6, 2, 7, 6, 8, 7, 5, 5, 5, 10, 10, 10, 3, 11, 11, 11, 9, 8, 8, 8, 4, 7, 10, 6, 9, 9, 9, 8, 11, 11, 11, 6, 5, 2, 7, 7, 7, 8, 10},
	{7, 5, 5, 5, 9, 9, 9, 4, 4, 6, 6, 6, 11, 11, 11, 8, 8, 8, 7, 7, 1, 10, 10, 8, 8, 8, 7, 10, 10, 10, 9, 8, 7, 7, 7, 5, 8, 9, 11, 11, 11, 2, 10, 10, 10, 9, 9, 9, 2, 11, 11, 6},
	{11, 9, 11, 11, 11, 9, 7, 7, 7, 2, 4, 9, 8, 8, 8, 9, 10, 10, 10, 6, 6, 6, 5, 10, 10, 10, 1, 11, 7, 7, 5, 5, 5, 2, 8, 8, 8, 4, 6, 10, 10, 8, 3, 7, 9, 9, 9, 3, 11, 11, 11, 9, 7, 8},
	{7, 5, 5, 5, 9, 9, 9, 4, 4, 6, 6, 6, 11, 11, 11, 8, 8, 8, 7, 7, 1, 10, 10, 8, 8, 8, 7, 10, 10, 10, 9, 8, 7, 7, 7, 5, 8, 9, 11, 11, 11, 2, 10, 10, 10, 9, 9, 9, 2, 11, 11, 6},
	{3, 10, 7, 8, 4, 10, 9, 9, 9, 11, 6, 2, 7, 6, 8, 7, 5, 5, 5, 10, 10, 10, 3, 11, 11, 11, 9, 8, 8, 8, 4, 7, 10, 6, 9, 9, 9, 8, 11, 11, 11, 6, 5, 2, 7, 7, 7, 8, 10},
}

// reels lengths [44, 52, 54, 52, 44], total reshuffles 282686976
// RTP = 76.368(lined) + 13.5(scatter) = 89.867755%
var Reels898 = slot.Reels5x{
	{7, 7, 7, 8, 8, 8, 11, 11, 11, 2, 6, 6, 6, 10, 10, 10, 8, 3, 9, 9, 9, 5, 5, 5, 3, 7, 5, 4, 9, 8, 9, 8, 9, 7, 10, 10, 10, 2, 4, 11, 11, 11, 7, 6},
	{7, 5, 5, 5, 9, 9, 9, 4, 4, 6, 6, 6, 11, 11, 11, 8, 8, 8, 7, 7, 1, 10, 10, 8, 8, 8, 7, 10, 10, 10, 9, 8, 7, 7, 7, 5, 8, 9, 11, 11, 11, 2, 10, 10, 10, 9, 9, 9, 2, 11, 11, 6},
	{11, 9, 11, 11, 11, 9, 7, 7, 7, 2, 4, 9, 8, 8, 8, 9, 10, 10, 10, 6, 6, 6, 5, 10, 10, 10, 1, 11, 7, 7, 5, 5, 5, 2, 8, 8, 8, 4, 6, 10, 10, 8, 3, 7, 9, 9, 9, 3, 11, 11, 11, 9, 7, 8},
	{7, 5, 5, 5, 9, 9, 9, 4, 4, 6, 6, 6, 11, 11, 11, 8, 8, 8, 7, 7, 1, 10, 10, 8, 8, 8, 7, 10, 10, 10, 9, 8, 7, 7, 7, 5, 8, 9, 11, 11, 11, 2, 10, 10, 10, 9, 9, 9, 2, 11, 11, 6},
	{7, 7, 7, 8, 8, 8, 11, 11, 11, 2, 6, 6, 6, 10, 10, 10, 8, 3, 9, 9, 9, 5, 5, 5, 3, 7, 5, 4, 9, 8, 9, 8, 9, 7, 10, 10, 10, 2, 4, 11, 11, 11, 7, 6},
}

// reels lengths [44, 52, 49, 52, 44], total reshuffles 256512256
// RTP = 77.562(lined) + 14.473(scatter) = 92.035571%
var Reels920 = slot.Reels5x{
	{7, 7, 7, 8, 8, 8, 11, 11, 11, 2, 6, 6, 6, 10, 10, 10, 8, 3, 9, 9, 9, 5, 5, 5, 3, 7, 5, 4, 9, 8, 9, 8, 9, 7, 10, 10, 10, 2, 4, 11, 11, 11, 7, 6},
	{7, 5, 5, 5, 9, 9, 9, 4, 4, 6, 6, 6, 11, 11, 11, 8, 8, 8, 7, 7, 1, 10, 10, 8, 8, 8, 7, 10, 10, 10, 9, 8, 7, 7, 7, 5, 8, 9, 11, 11, 11, 2, 10, 10, 10, 9, 9, 9, 2, 11, 11, 6},
	{5, 7, 8, 10, 6, 11, 11, 11, 10, 8, 9, 9, 9, 10, 8, 5, 7, 7, 7, 1, 5, 6, 10, 10, 10, 8, 8, 8, 3, 6, 9, 7, 11, 3, 9, 10, 9, 2, 4, 11, 4, 6, 8, 9, 11, 5, 11, 7, 2},
	{7, 5, 5, 5, 9, 9, 9, 4, 4, 6, 6, 6, 11, 11, 11, 8, 8, 8, 7, 7, 1, 10, 10, 8, 8, 8, 7, 10, 10, 10, 9, 8, 7, 7, 7, 5, 8, 9, 11, 11, 11, 2, 10, 10, 10, 9, 9, 9, 2, 11, 11, 6},
	{7, 7, 7, 8, 8, 8, 11, 11, 11, 2, 6, 6, 6, 10, 10, 10, 8, 3, 9, 9, 9, 5, 5, 5, 3, 7, 5, 4, 9, 8, 9, 8, 9, 7, 10, 10, 10, 2, 4, 11, 11, 11, 7, 6},
}

// reels lengths [53, 47, 49, 47, 53], total reshuffles 304048969
// RTP = 80.453(lined) + 12.069(scatter) = 92.522682%
var Reels925 = slot.Reels5x{
	{8, 10, 10, 6, 4, 4, 5, 5, 5, 3, 10, 10, 10, 11, 11, 11, 7, 10, 10, 10, 3, 9, 9, 2, 6, 6, 6, 8, 8, 8, 9, 9, 9, 11, 11, 7, 7, 7, 9, 9, 9, 11, 11, 11, 7, 7, 7, 8, 8, 8, 5, 2, 8},
	{9, 10, 10, 8, 9, 2, 5, 5, 5, 7, 7, 10, 8, 8, 11, 11, 11, 2, 6, 4, 11, 7, 7, 7, 10, 10, 10, 11, 11, 11, 8, 9, 9, 4, 6, 6, 6, 10, 1, 7, 9, 9, 9, 8, 8, 8, 5},
	{5, 7, 8, 10, 6, 11, 11, 11, 10, 8, 9, 9, 9, 10, 8, 5, 7, 7, 7, 1, 5, 6, 10, 10, 10, 8, 8, 8, 3, 6, 9, 7, 11, 3, 9, 10, 9, 2, 4, 11, 4, 6, 8, 9, 11, 5, 11, 7, 2},
	{9, 10, 10, 8, 9, 2, 5, 5, 5, 7, 7, 10, 8, 8, 11, 11, 11, 2, 6, 4, 11, 7, 7, 7, 10, 10, 10, 11, 11, 11, 8, 9, 9, 4, 6, 6, 6, 10, 1, 7, 9, 9, 9, 8, 8, 8, 5},
	{8, 10, 10, 6, 4, 4, 5, 5, 5, 3, 10, 10, 10, 11, 11, 11, 7, 10, 10, 10, 3, 9, 9, 2, 6, 6, 6, 8, 8, 8, 9, 9, 9, 11, 11, 7, 7, 7, 9, 9, 9, 11, 11, 11, 7, 7, 7, 8, 8, 8, 5, 2, 8},
}

// reels lengths [49, 47, 49, 47, 49], total reshuffles 259886641
// RTP = 80.697(lined) + 13.475(scatter) = 94.171658%
var Reels941 = slot.Reels5x{
	{3, 10, 7, 8, 4, 10, 9, 9, 9, 11, 6, 2, 7, 6, 8, 7, 5, 5, 5, 10, 10, 10, 3, 11, 11, 11, 9, 8, 8, 8, 4, 7, 10, 6, 9, 9, 9, 8, 11, 11, 11, 6, 5, 2, 7, 7, 7, 8, 10},
	{9, 10, 10, 8, 9, 2, 5, 5, 5, 7, 7, 10, 8, 8, 11, 11, 11, 2, 6, 4, 11, 7, 7, 7, 10, 10, 10, 11, 11, 11, 8, 9, 9, 4, 6, 6, 6, 10, 1, 7, 9, 9, 9, 8, 8, 8, 5},
	{5, 7, 8, 10, 6, 11, 11, 11, 10, 8, 9, 9, 9, 10, 8, 5, 7, 7, 7, 1, 5, 6, 10, 10, 10, 8, 8, 8, 3, 6, 9, 7, 11, 3, 9, 10, 9, 2, 4, 11, 4, 6, 8, 9, 11, 5, 11, 7, 2},
	{9, 10, 10, 8, 9, 2, 5, 5, 5, 7, 7, 10, 8, 8, 11, 11, 11, 2, 6, 4, 11, 7, 7, 7, 10, 10, 10, 11, 11, 11, 8, 9, 9, 4, 6, 6, 6, 10, 1, 7, 9, 9, 9, 8, 8, 8, 5},
	{3, 10, 7, 8, 4, 10, 9, 9, 9, 11, 6, 2, 7, 6, 8, 7, 5, 5, 5, 10, 10, 10, 3, 11, 11, 11, 9, 8, 8, 8, 4, 7, 10, 6, 9, 9, 9, 8, 11, 11, 11, 6, 5, 2, 7, 7, 7, 8, 10},
}

// reels lengths [44, 47, 49, 52, 44], total reshuffles 231847616
// RTP = 80.374(lined) + 15.088(scatter) = 95.462772%
var Reels954 = slot.Reels5x{
	{7, 7, 7, 8, 8, 8, 11, 11, 11, 2, 6, 6, 6, 10, 10, 10, 8, 3, 9, 9, 9, 5, 5, 5, 3, 7, 5, 4, 9, 8, 9, 8, 9, 7, 10, 10, 10, 2, 4, 11, 11, 11, 7, 6},
	{9, 10, 10, 8, 9, 2, 5, 5, 5, 7, 7, 10, 8, 8, 11, 11, 11, 2, 6, 4, 11, 7, 7, 7, 10, 10, 10, 11, 11, 11, 8, 9, 9, 4, 6, 6, 6, 10, 1, 7, 9, 9, 9, 8, 8, 8, 5},
	{5, 7, 8, 10, 6, 11, 11, 11, 10, 8, 9, 9, 9, 10, 8, 5, 7, 7, 7, 1, 5, 6, 10, 10, 10, 8, 8, 8, 3, 6, 9, 7, 11, 3, 9, 10, 9, 2, 4, 11, 4, 6, 8, 9, 11, 5, 11, 7, 2},
	{7, 5, 5, 5, 9, 9, 9, 4, 4, 6, 6, 6, 11, 11, 11, 8, 8, 8, 7, 7, 1, 10, 10, 8, 8, 8, 7, 10, 10, 10, 9, 8, 7, 7, 7, 5, 8, 9, 11, 11, 11, 2, 10, 10, 10, 9, 9, 9, 2, 11, 11, 6},
	{7, 7, 7, 8, 8, 8, 11, 11, 11, 2, 6, 6, 6, 10, 10, 10, 8, 3, 9, 9, 9, 5, 5, 5, 3, 7, 5, 4, 9, 8, 9, 8, 9, 7, 10, 10, 10, 2, 4, 11, 11, 11, 7, 6},
}

// reels lengths [44, 47, 49, 47, 44], total reshuffles 209554576
// RTP = 81.285(lined) + 15.735(scatter) = 97.020072%
var Reels970 = slot.Reels5x{
	{7, 7, 7, 8, 8, 8, 11, 11, 11, 2, 6, 6, 6, 10, 10, 10, 8, 3, 9, 9, 9, 5, 5, 5, 3, 7, 5, 4, 9, 8, 9, 8, 9, 7, 10, 10, 10, 2, 4, 11, 11, 11, 7, 6},
	{9, 10, 10, 8, 9, 2, 5, 5, 5, 7, 7, 10, 8, 8, 11, 11, 11, 2, 6, 4, 11, 7, 7, 7, 10, 10, 10, 11, 11, 11, 8, 9, 9, 4, 6, 6, 6, 10, 1, 7, 9, 9, 9, 8, 8, 8, 5},
	{5, 7, 8, 10, 6, 11, 11, 11, 10, 8, 9, 9, 9, 10, 8, 5, 7, 7, 7, 1, 5, 6, 10, 10, 10, 8, 8, 8, 3, 6, 9, 7, 11, 3, 9, 10, 9, 2, 4, 11, 4, 6, 8, 9, 11, 5, 11, 7, 2},
	{9, 10, 10, 8, 9, 2, 5, 5, 5, 7, 7, 10, 8, 8, 11, 11, 11, 2, 6, 4, 11, 7, 7, 7, 10, 10, 10, 11, 11, 11, 8, 9, 9, 4, 6, 6, 6, 10, 1, 7, 9, 9, 9, 8, 8, 8, 5},
	{7, 7, 7, 8, 8, 8, 11, 11, 11, 2, 6, 6, 6, 10, 10, 10, 8, 3, 9, 9, 9, 5, 5, 5, 3, 7, 5, 4, 9, 8, 9, 8, 9, 7, 10, 10, 10, 2, 4, 11, 11, 11, 7, 6},
}

// reels lengths [49, 43, 49, 43, 49], total reshuffles 217533001
// RTP = 83.885(lined) + 14.621(scatter) = 98.506479%
var Reels985 = slot.Reels5x{
	{3, 10, 7, 8, 4, 10, 9, 9, 9, 11, 6, 2, 7, 6, 8, 7, 5, 5, 5, 10, 10, 10, 3, 11, 11, 11, 9, 8, 8, 8, 4, 7, 10, 6, 9, 9, 9, 8, 11, 11, 11, 6, 5, 2, 7, 7, 7, 8, 10},
	{10, 8, 8, 8, 8, 11, 2, 7, 7, 7, 6, 1, 5, 4, 11, 6, 10, 10, 10, 9, 11, 2, 8, 9, 8, 9, 5, 6, 4, 11, 11, 11, 7, 7, 5, 5, 9, 9, 9, 10, 10, 6, 7},
	{5, 7, 8, 10, 6, 11, 11, 11, 10, 8, 9, 9, 9, 10, 8, 5, 7, 7, 7, 1, 5, 6, 10, 10, 10, 8, 8, 8, 3, 6, 9, 7, 11, 3, 9, 10, 9, 2, 4, 11, 4, 6, 8, 9, 11, 5, 11, 7, 2},
	{10, 8, 8, 8, 8, 11, 2, 7, 7, 7, 6, 1, 5, 4, 11, 6, 10, 10, 10, 9, 11, 2, 8, 9, 8, 9, 5, 6, 4, 11, 11, 11, 7, 7, 5, 5, 9, 9, 9, 10, 10, 6, 7},
	{3, 10, 7, 8, 4, 10, 9, 9, 9, 11, 6, 2, 7, 6, 8, 7, 5, 5, 5, 10, 10, 10, 3, 11, 11, 11, 9, 8, 8, 8, 4, 7, 10, 6, 9, 9, 9, 8, 11, 11, 11, 6, 5, 2, 7, 7, 7, 8, 10},
}

// reels lengths [44, 43, 49, 43, 49], total reshuffles 195335756
// RTP = 84.266(lined) + 15.774(scatter) = 100.039862%
var Reels100 = slot.Reels5x{
	{7, 7, 7, 11, 10, 8, 8, 8, 9, 11, 11, 11, 6, 5, 2, 6, 10, 8, 5, 3, 10, 9, 9, 9, 3, 7, 5, 5, 4, 8, 9, 8, 9, 7, 10, 10, 10, 2, 6, 4, 11, 11, 7, 6},
	{10, 8, 8, 8, 8, 11, 11, 11, 2, 7, 7, 7, 5, 6, 1, 5, 4, 6, 10, 10, 10, 11, 2, 8, 9, 8, 9, 5, 6, 4, 11, 9, 10, 11, 7, 7, 5, 9, 9, 9, 10, 6, 7},
	{5, 7, 8, 10, 6, 11, 11, 11, 10, 8, 9, 9, 9, 10, 8, 5, 7, 7, 7, 1, 5, 6, 10, 10, 10, 8, 8, 8, 3, 6, 9, 7, 11, 3, 9, 10, 9, 2, 4, 11, 4, 6, 8, 9, 11, 5, 11, 7, 2},
	{10, 8, 8, 8, 8, 11, 11, 11, 2, 7, 7, 7, 5, 6, 1, 5, 4, 6, 10, 10, 10, 11, 2, 8, 9, 8, 9, 5, 6, 4, 11, 9, 10, 11, 7, 7, 5, 9, 9, 9, 10, 6, 7},
	{3, 10, 7, 8, 4, 10, 9, 9, 9, 11, 6, 2, 7, 6, 8, 7, 5, 5, 5, 10, 10, 10, 3, 11, 11, 11, 9, 8, 8, 8, 4, 7, 10, 6, 9, 9, 9, 8, 11, 11, 11, 6, 5, 2, 7, 7, 7, 8, 10},
}

// reels lengths [44, 43, 49, 43, 44], total reshuffles 175403536
// RTP = 84.175(lined) + 17.016(scatter) = 101.190799%
var Reels101 = slot.Reels5x{
	{7, 7, 7, 11, 10, 8, 8, 8, 9, 11, 11, 11, 6, 5, 2, 6, 10, 8, 5, 3, 10, 9, 9, 9, 3, 7, 5, 5, 4, 8, 9, 8, 9, 7, 10, 10, 10, 2, 6, 4, 11, 11, 7, 6},
	{10, 8, 8, 8, 8, 11, 11, 11, 2, 7, 7, 7, 5, 6, 1, 5, 4, 6, 10, 10, 10, 11, 2, 8, 9, 8, 9, 5, 6, 4, 11, 9, 10, 11, 7, 7, 5, 9, 9, 9, 10, 6, 7},
	{5, 7, 8, 10, 6, 11, 11, 11, 10, 8, 9, 9, 9, 10, 8, 5, 7, 7, 7, 1, 5, 6, 10, 10, 10, 8, 8, 8, 3, 6, 9, 7, 11, 3, 9, 10, 9, 2, 4, 11, 4, 6, 8, 9, 11, 5, 11, 7, 2},
	{10, 8, 8, 8, 8, 11, 11, 11, 2, 7, 7, 7, 5, 6, 1, 5, 4, 6, 10, 10, 10, 11, 2, 8, 9, 8, 9, 5, 6, 4, 11, 9, 10, 11, 7, 7, 5, 9, 9, 9, 10, 6, 7},
	{7, 7, 7, 11, 10, 8, 8, 8, 9, 11, 11, 11, 6, 5, 2, 6, 10, 8, 5, 3, 10, 9, 9, 9, 3, 7, 5, 5, 4, 8, 9, 8, 9, 7, 10, 10, 10, 2, 6, 4, 11, 11, 7, 6},
}

// reels lengths [44, 43, 45, 43, 44], total reshuffles 161084880
// RTP = 87.959(lined) + 18.061(scatter) = 106.020478%
var Reels106 = slot.Reels5x{
	{7, 7, 7, 11, 10, 8, 8, 8, 9, 11, 11, 11, 6, 5, 2, 6, 10, 8, 5, 3, 10, 9, 9, 9, 3, 7, 5, 5, 4, 8, 9, 8, 9, 7, 10, 10, 10, 2, 6, 4, 11, 11, 7, 6},
	{10, 8, 8, 8, 8, 11, 11, 11, 2, 7, 7, 7, 5, 6, 1, 5, 4, 6, 10, 10, 10, 11, 2, 8, 9, 8, 9, 5, 6, 4, 11, 9, 10, 11, 7, 7, 5, 9, 9, 9, 10, 6, 7},
	{2, 7, 7, 7, 9, 9, 9, 8, 3, 7, 10, 10, 5, 11, 5, 7, 10, 1, 11, 9, 8, 8, 8, 11, 11, 11, 4, 7, 4, 6, 5, 6, 9, 5, 3, 10, 10, 10, 9, 2, 8, 6, 8, 11, 6},
	{10, 8, 8, 8, 8, 11, 11, 11, 2, 7, 7, 7, 5, 6, 1, 5, 4, 6, 10, 10, 10, 11, 2, 8, 9, 8, 9, 5, 6, 4, 11, 9, 10, 11, 7, 7, 5, 9, 9, 9, 10, 6, 7},
	{7, 7, 7, 11, 10, 8, 8, 8, 9, 11, 11, 11, 6, 5, 2, 6, 10, 8, 5, 3, 10, 9, 9, 9, 3, 7, 5, 5, 4, 8, 9, 8, 9, 7, 10, 10, 10, 2, 6, 4, 11, 11, 7, 6},
}

// reels lengths [39, 43, 45, 43, 39], total reshuffles 126554805
// RTP = 89.08(lined) + 21.517(scatter) = 110.597176%
var Reels110 = slot.Reels5x{
	{10, 10, 10, 8, 8, 8, 5, 5, 5, 2, 4, 9, 11, 11, 11, 8, 7, 7, 7, 3, 6, 6, 6, 5, 11, 7, 3, 6, 11, 7, 9, 9, 9, 10, 10, 2, 4, 9, 8},
	{10, 8, 8, 8, 8, 11, 11, 11, 2, 7, 7, 7, 5, 6, 1, 5, 4, 6, 10, 10, 10, 11, 2, 8, 9, 8, 9, 5, 6, 4, 11, 9, 10, 11, 7, 7, 5, 9, 9, 9, 10, 6, 7},
	{2, 7, 7, 7, 9, 9, 9, 8, 3, 7, 10, 10, 5, 11, 5, 7, 10, 1, 11, 9, 8, 8, 8, 11, 11, 11, 4, 7, 4, 6, 5, 6, 9, 5, 3, 10, 10, 10, 9, 2, 8, 6, 8, 11, 6},
	{10, 8, 8, 8, 8, 11, 11, 11, 2, 7, 7, 7, 5, 6, 1, 5, 4, 6, 10, 10, 10, 11, 2, 8, 9, 8, 9, 5, 6, 4, 11, 9, 10, 11, 7, 7, 5, 9, 9, 9, 10, 6, 7},
	{10, 10, 10, 8, 8, 8, 5, 5, 5, 2, 4, 9, 11, 11, 11, 8, 7, 7, 7, 3, 6, 6, 6, 5, 11, 7, 3, 6, 11, 7, 9, 9, 9, 10, 10, 2, 4, 9, 8},
}

// Map with available reels.
var ReelsMap = map[float64]*slot.Reels5x{
	86.377743:  &Reels863,
	87.803869:  &Reels878,
	89.867755:  &Reels898,
	92.035571:  &Reels920,
	92.522682:  &Reels925,
	94.171658:  &Reels941,
	95.462772:  &Reels954,
	97.020072:  &Reels970,
	98.506479:  &Reels985,
	100.039862: &Reels100,
	101.190799: &Reels101,
	106.020478: &Reels106,
	110.597176: &Reels110,
}