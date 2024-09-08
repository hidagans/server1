//go:build !prod || full || megajack

package links

import (
	"context"

	"github.com/slotopol/server/game/slotopol"
	"github.com/spf13/pflag"
)

func init() {
	var gi = GameInfo{
		Aliases: []GameAlias{
			{"slotopol", "Slotopol"},
		},
		Provider: "Megajack",
		ScrnX:    5,
		ScrnY:    3,
		RtpList:  MakeRtpList(slotopol.ReelsMap),
	}
	GameList = append(GameList, gi)

	for _, ga := range gi.Aliases {
		ScanIters = append(ScanIters, func(flags *pflag.FlagSet, ctx context.Context) {
			if is, _ := flags.GetBool(ga.ID); is {
				var rn, _ = flags.GetString("reels")
				slotopol.CalcStat(ctx, rn)
			}
		})
		GameFactory[ga.ID] = func() any {
			return slotopol.NewGame()
		}
	}
}
