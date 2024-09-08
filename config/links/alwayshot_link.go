//go:build !prod || full || novomatic

package links

import (
	"context"

	"github.com/slotopol/server/game/alwayshot"
	"github.com/spf13/pflag"
)

func init() {
	var gi = GameInfo{
		Aliases: []GameAlias{
			{"alwayshot", "Always Hot"},
		},
		Provider: "Novomatic",
		ScrnX:    3,
		ScrnY:    3,
		RtpList:  MakeRtpList(alwayshot.ReelsMap),
	}
	GameList = append(GameList, gi)

	for _, ga := range gi.Aliases {
		ScanIters = append(ScanIters, func(flags *pflag.FlagSet, ctx context.Context) {
			if is, _ := flags.GetBool(ga.ID); is {
				var rn, _ = flags.GetString("reels")
				alwayshot.CalcStat(ctx, rn)
			}
		})
		GameFactory[ga.ID] = func() any {
			return alwayshot.NewGame()
		}
	}
}
