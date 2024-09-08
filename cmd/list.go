package cmd

import (
	"fmt"
	"sort"
	"strings"

	cfg "github.com/slotopol/server/config"
	"github.com/slotopol/server/config/links"
	"github.com/slotopol/server/util"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var listflags *pflag.FlagSet

var (
	fAll, fRTP bool
)

const listShort = "List of available games released on server"
const listLong = ``
const listExmp = `Get the list of all available games:
  %[1]s list --all
Get the list of available 'NetExt' and 'BetSoft' games only:
  %[1]s list --netent --betsoft
Get the list of available 'Play'n GO' games with RTP list for each:
  %[1]s list --playngo --rtp`

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:     "list",
	Short:   listShort,
	Long:    listLong,
	Example: fmt.Sprintf(listExmp, cfg.AppName),
	Run: func(cmd *cobra.Command, args []string) {
		var num, alg int
		var prov = map[string]int{}
		for _, gi := range links.GameList {
			if prv, _ := listflags.GetBool(util.ToID(gi.Provider)); prv || fAll {
				num += len(gi.Aliases)
			}
		}

		var list = make([]string, num)
		var i int
		for _, gi := range links.GameList {
			if prv, _ := listflags.GetBool(util.ToID(gi.Provider)); prv || fAll {
				prov[gi.Provider] += len(gi.Aliases)
				alg++
				for _, ga := range gi.Aliases {
					var rtpinfo string
					if fRTP && len(gi.RtpList) > 0 {
						var rtpstr = make([]string, len(gi.RtpList))
						for i, rtp := range gi.RtpList {
							rtpstr[i] = fmt.Sprintf("%.2f", rtp)
						}
						rtpinfo = ", RTP: " + strings.Join(rtpstr, ", ")
					}
					var info = fmt.Sprintf("'%s' %s %dx%d videoslot%s", ga.Name, gi.Provider, gi.ScrnX, gi.ScrnY, rtpinfo)
					list[i] = info
					i++
				}
			}
		}

		if is, _ := listflags.GetBool("name"); is {
			sort.Strings(list)
			fmt.Println()
			for _, s := range list {
				fmt.Println(s)
			}
		}
		if is, _ := listflags.GetBool("stat"); is {
			fmt.Println()
			fmt.Printf("total: %d games, %d algorithms, %d providers\n", num, alg, len(prov))
			for p, n := range prov {
				fmt.Printf("%s: %d games\n", p, n)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listflags = listCmd.Flags()
	listflags.BoolP("name", "n", true, "list of provided games names")
	listflags.BoolP("stat", "s", true, "summary statistics of provided games")
	listflags.BoolVar(&fAll, "all", false, "list games of all available providers")
	listflags.BoolVar(&fRTP, "rtp", false, "RTP (Return to Player) percents list of available reels for each game")
	listflags.Bool("megajack", false, "include games of 'Megajack' provider")
	listflags.Bool("novomatic", false, "include games of 'Novomatic' provider")
	listflags.Bool("netent", false, "include games of 'NetExt' provider")
	listflags.Bool("betsoft", false, "include games of 'BetSoft' provider")
	listflags.Bool("playtech", false, "include games of 'Playtech' provider")
	listflags.Bool("playngo", false, "include games of 'Play'n GO' provider")
	listCmd.MarkFlagsOneRequired("all", "megajack", "novomatic", "netent", "betsoft", "playtech", "playngo")
}
