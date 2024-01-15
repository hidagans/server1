package spi

import (
	"encoding/json"
	"encoding/xml"
	"log"
	"net/http"
	"strings"

	cfg "github.com/slotopol/server/config"
	"xorm.io/xorm"

	"github.com/gin-gonic/gin"
	"github.com/slotopol/server/game"
)

func SpiGameJoin(c *gin.Context) {
	var err error
	var ok bool
	var arg struct {
		XMLName xml.Name `json:"-" yaml:"-" xml:"arg"`
		UID     uint64   `json:"uid" yaml:"uid" xml:"uid,attr" form:"uid"`
		RID     uint64   `json:"rid" yaml:"rid" xml:"rid,attr" form:"rid"`
		Alias   string   `json:"alias" yaml:"alias" xml:"alias" form:"alias"`
	}
	var ret struct {
		XMLName xml.Name    `json:"-" yaml:"-" xml:"ret"`
		GID     uint64      `json:"gid" yaml:"gid" xml:"gid,attr"`
		Screen  game.Screen `json:"screen" yaml:"screen" xml:"screen"`
	}

	if err = c.Bind(&arg); err != nil {
		Ret400(c, SEC_game_join_nobind, err)
		return
	}
	if arg.UID == 0 {
		Ret400(c, SEC_game_join_nouid, ErrNoUID)
		return
	}
	if arg.RID == 0 {
		Ret400(c, SEC_game_join_norid, ErrNoRID)
		return
	}
	if arg.Alias == "" {
		Ret400(c, SEC_game_join_nodata, ErrNoData)
		return
	}

	var user *User
	if user, ok = Users.Get(arg.UID); !ok {
		Ret400(c, SEC_game_join_nouser, ErrNoUser)
		return
	}

	var room *Room
	if room, ok = Rooms.Get(arg.RID); !ok {
		Ret500(c, SEC_game_join_noroom, ErrNoRoom)
		return
	}
	_ = room

	var props *Props
	if props, ok = user.props.Get(arg.RID); !ok {
		Ret403(c, SEC_game_join_noprops, ErrNoProps)
		return
	}
	_ = props

	var alias = strings.ToLower(arg.Alias)
	var gname string
	if gname, ok = cfg.GameAliases[alias]; !ok {
		Ret400(c, SEC_game_join_noalias, ErrNoAliase)
		return
	}

	var maker = cfg.GameFactory[gname]
	var slotgame = maker("96")
	if slotgame == nil {
		Ret400(c, SEC_game_join_noreels, ErrNoReels)
		return
	}

	var og = OpenGame{
		UID:   arg.UID,
		RID:   arg.RID,
		Alias: alias,
		game:  slotgame.(game.SlotGame),
	}
	if _, err = cfg.XormStorage.Insert(&og); err != nil {
		Ret500(c, SEC_game_join_insert, err)
		return
	}

	OpenGames.Set(og.GID, og)
	user.games.Set(og.GID, og)

	var scrn = og.game.NewScreen()
	og.game.Spin(scrn)

	ret.GID = og.GID
	ret.Screen = scrn

	RetOk(c, ret)
}

func SpiGamePart(c *gin.Context) {
	var err error
	var ok bool
	var arg struct {
		XMLName xml.Name `json:"-" yaml:"-" xml:"arg"`
		GID     uint64   `json:"gid" yaml:"gid" xml:"gid,attr" form:"gid"`
	}

	if err = c.Bind(&arg); err != nil {
		Ret400(c, SEC_game_part_nobind, err)
		return
	}
	if arg.GID == 0 {
		Ret400(c, SEC_game_part_nogid, ErrNoGID)
		return
	}

	var og OpenGame
	if og, ok = OpenGames.Get(arg.GID); !ok {
		Ret400(c, SEC_game_part_notopened, ErrNotOpened)
		return
	}

	var user *User
	if user, ok = Users.Get(og.UID); !ok {
		Ret500(c, SEC_game_part_nouser, ErrNoUser)
		return
	}

	OpenGames.Delete(arg.GID)
	user.games.Delete(arg.GID)

	c.Status(http.StatusOK)
}

func SpiGameGetBet(c *gin.Context) {
	var err error
	var ok bool
	var arg struct {
		XMLName xml.Name `json:"-" yaml:"-" xml:"arg"`
		GID     uint64   `json:"gid" yaml:"gid" xml:"gid,attr" form:"gid"`
	}
	var ret struct {
		XMLName xml.Name `json:"-" yaml:"-" xml:"ret"`
		Bet     int      `json:"bet" yaml:"bet" xml:"bet"`
	}

	if err = c.Bind(&arg); err != nil {
		Ret400(c, SEC_game_getbet_nobind, err)
		return
	}
	if arg.GID == 0 {
		Ret400(c, SEC_game_getbet_nogid, ErrNoGID)
		return
	}

	var og OpenGame
	if og, ok = OpenGames.Get(arg.GID); !ok {
		Ret400(c, SEC_game_getbet_notopened, ErrNotOpened)
		return
	}

	ret.Bet = og.game.GetBet()

	RetOk(c, ret)
}

func SpiGameSetBet(c *gin.Context) {
	var err error
	var ok bool
	var arg struct {
		XMLName xml.Name `json:"-" yaml:"-" xml:"arg"`
		GID     uint64   `json:"gid" yaml:"gid" xml:"gid,attr"`
		Bet     int      `json:"bet" yaml:"bet" xml:"bet"`
	}

	if err = c.Bind(&arg); err != nil {
		Ret400(c, SEC_game_setbet_nobind, err)
		return
	}
	if arg.GID == 0 {
		Ret400(c, SEC_game_setbet_nogid, ErrNoGID)
		return
	}
	if arg.Bet == 0 {
		Ret400(c, SEC_game_setbet_nodata, ErrNoData)
		return
	}

	var og OpenGame
	if og, ok = OpenGames.Get(arg.GID); !ok {
		Ret400(c, SEC_game_setbet_notopened, ErrNotOpened)
		return
	}

	if err = og.game.SetBet(arg.Bet); err != nil {
		Ret400(c, SEC_game_setbet_badbet, err)
		return
	}

	c.Status(http.StatusOK)
}

func SpiGameGetSbl(c *gin.Context) {
	var err error
	var ok bool
	var arg struct {
		XMLName xml.Name `json:"-" yaml:"-" xml:"arg"`
		GID     uint64   `json:"gid" yaml:"gid" xml:"gid,attr" form:"gid"`
	}
	var ret struct {
		XMLName xml.Name `json:"-" yaml:"-" xml:"ret"`
		SBL     game.SBL `json:"sbl" yaml:"sbl" xml:"sbl"`
	}

	if err = c.Bind(&arg); err != nil {
		Ret400(c, SEC_game_getsbl_nobind, err)
		return
	}
	if arg.GID == 0 {
		Ret400(c, SEC_game_getsbl_nogid, ErrNoGID)
		return
	}

	var og OpenGame
	if og, ok = OpenGames.Get(arg.GID); !ok {
		Ret400(c, SEC_game_getsbl_notopened, ErrNotOpened)
		return
	}

	ret.SBL = og.game.GetLines()

	RetOk(c, ret)
}

func SpiGameSetSbl(c *gin.Context) {
	var err error
	var ok bool
	var arg struct {
		XMLName xml.Name `json:"-" yaml:"-" xml:"arg"`
		GID     uint64   `json:"gid" yaml:"gid" xml:"gid,attr"`
		SBL     game.SBL `json:"sbl" yaml:"sbl" xml:"sbl"`
	}

	if err = c.Bind(&arg); err != nil {
		Ret400(c, SEC_game_setsbl_nobind, err)
		return
	}
	if arg.GID == 0 {
		Ret400(c, SEC_game_setsbl_nogid, ErrNoGID)
		return
	}
	if arg.SBL == 0 {
		Ret400(c, SEC_game_setsbl_nodata, ErrNoData)
		return
	}

	var og OpenGame
	if og, ok = OpenGames.Get(arg.GID); !ok {
		Ret400(c, SEC_game_setsbl_notopened, ErrNotOpened)
		return
	}

	if err = og.game.SetLines(arg.SBL); err != nil {
		Ret400(c, SEC_game_setsbl_badlines, err)
		return
	}

	c.Status(http.StatusOK)
}

func SpiGameSpin(c *gin.Context) {
	var err error
	var ok bool
	var arg struct {
		XMLName xml.Name `json:"-" yaml:"-" xml:"arg"`
		GID     uint64   `json:"gid" yaml:"gid" xml:"gid,attr" form:"gid"`
	}
	var ret struct {
		XMLName xml.Name       `json:"-" yaml:"-" xml:"ret"`
		SID     uint64         `json:"sid" yaml:"sid" xml:"sid,attr" form:"sid"`
		Screen  game.Screen    `json:"screen" yaml:"screen" xml:"screen"`
		Wins    []game.WinItem `json:"wins" yaml:"wins" xml:"wins"`
		Gain    int            `json:"gain" yaml:"gain" xml:"gain"`
		FS      int            `json:"fs,omitempty" yaml:"fs,omitempty" xml:"fs,omitempty"`
		Wallet  int            `json:"wallet" yaml:"wallet" xml:"wallet"`
	}

	if err = c.Bind(&arg); err != nil {
		Ret400(c, SEC_game_spin_nobind, err)
		return
	}
	if arg.GID == 0 {
		Ret400(c, SEC_game_spin_nogid, ErrNoGID)
		return
	}

	var og OpenGame
	if og, ok = OpenGames.Get(arg.GID); !ok {
		Ret400(c, SEC_game_spin_notopened, ErrNotOpened)
		return
	}

	var user *User
	if user, ok = Users.Get(og.UID); !ok {
		Ret500(c, SEC_game_spin_nouser, ErrNoUser)
		return
	}

	var (
		fs       = og.game.FreeSpins()
		bet      = og.game.GetBet()
		sbl      = og.game.GetLines()
		totalbet int
		totalwin int
	)
	if fs == 0 {
		totalbet = bet * sbl.Num()
	}

	var props *Props
	if props, ok = user.props.Get(og.RID); !ok {
		Ret403(c, SEC_game_spin_noprops, ErrNoProps)
		return
	}
	if props.Wallet < totalbet {
		Ret403(c, SEC_game_spin_nomoney, ErrNoMoney)
		return
	}

	var room *Room
	if room, ok = Rooms.Get(og.RID); !ok {
		Ret500(c, SEC_game_spin_noroom, ErrNoRoom)
		return
	}

	// get game screen object
	var scrn = og.game.NewScreen()

	// spin until gain less than bank value
	room.mux.RLock()
	var bank = room.Bank
	room.mux.RUnlock()
	var ws game.WinScan
	var n = 0
	for {
		og.game.Spin(scrn)
		og.game.Scanner(scrn, &ws)
		og.game.Spawn(scrn, &ws)
		totalwin = ws.SumPay()
		if bank+float64(totalbet-totalwin) >= 0 || (bank < 0 && totalbet > totalwin) {
			break
		}
		if n >= cfg.Cfg.MaxSpinAttempts {
			Ret500(c, SEC_game_spin_badbank, ErrBadBank)
			return
		}
		n++
	}

	// write gain and total bet as transaction
	if _, err = cfg.XormStorage.Transaction(func(session *xorm.Session) (ret interface{}, err error) {
		defer func() {
			if err != nil {
				session.Rollback()
			}
		}()

		const sql1 = `UPDATE room SET bank=bank+? WHERE rid=?`
		if ret, err = session.Exec(sql1, totalbet-totalwin, room.RID); err != nil {
			Ret500(c, SEC_game_spin_sqlbank, err)
			return
		}

		const sql2 = `UPDATE props SET wallet=wallet+? WHERE uid=? AND rid=?`
		if ret, err = session.Exec(sql2, totalwin-totalbet, props.UID, props.RID); err != nil {
			Ret500(c, SEC_game_spin_sqlbalance, err)
			return
		}

		return
	}); err != nil {
		return
	}

	// make changes to memory data
	room.mux.Lock()
	room.Bank += float64(totalbet - totalwin)
	room.mux.Unlock()

	props.Wallet += totalwin - totalbet

	og.game.Apply(scrn, &ws)

	// write spin result to log and get spin ID
	var sl = Spinlog{
		GID:    arg.GID,
		Gain:   totalwin,
		Wallet: props.Wallet,
	}
	var b []byte
	b, _ = json.Marshal(scrn)
	sl.Screen = string(b)
	b, _ = json.Marshal(og.game)
	sl.Game = string(b)
	b, _ = json.Marshal(ws.Wins)
	sl.Wins = string(b)
	if _, err = cfg.XormSpinlog.Insert(&sl); err != nil {
		log.Printf("can not write to spin log: %s", err.Error())
	}

	// prepare result
	ret.SID = sl.SID
	ret.Screen = scrn
	ret.Wins = ws.Wins
	ret.Wallet = props.Wallet
	ret.FS = fs
	ret.Gain = totalwin

	RetOk(c, ret)
}