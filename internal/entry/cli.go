package entry

import (
	"fmt"
	"log"

	"github.com/anhoder/foxful-cli/util"
	neteaseutil "github.com/go-musicfox/netease-music/util"
	"github.com/gookit/gcli/v2"

	"github.com/go-musicfox/go-musicfox/internal/commands"
	"github.com/go-musicfox/go-musicfox/internal/configs"
	"github.com/go-musicfox/go-musicfox/internal/types"
	"github.com/go-musicfox/go-musicfox/internal/web"
	"github.com/go-musicfox/go-musicfox/utils"
)

func runCLI() {
	log.SetOutput(utils.LogWriter())

	var app = gcli.NewApp()
	app.Name = types.AppName
	app.Version = types.AppVersion
	if types.BuildTags != "" {
		app.Version += " [" + types.BuildTags + "]"
	}
	app.Description = types.AppDescription
	app.GOptsBinder = func(gf *gcli.Flags) {
		gf.BoolOpt(&commands.GlobalOptions.PProfMode, "pprof", "p", false, "enable PProf mode")
	}

	// 加载config
	utils.LoadIniConfig()

	util.PrimaryColor = configs.ConfigRegistry.Main.PrimaryColor
	var (
		logo         = util.GetAlphaAscii(app.Name)
		randomColor  = util.GetPrimaryColor()
		logoColorful = util.SetFgStyle(logo, randomColor)
	)

	gcli.AppHelpTemplate = fmt.Sprintf(types.AppHelpTemplate, logoColorful)
	app.Logo.Text = logoColorful

	// 更新netease配置
	neteaseutil.UNMSwitch = configs.ConfigRegistry.UNM.Enable
	neteaseutil.Sources = configs.ConfigRegistry.UNM.Sources
	neteaseutil.SearchLimit = configs.ConfigRegistry.UNM.SearchLimit
	neteaseutil.EnableLocalVip = configs.ConfigRegistry.UNM.EnableLocalVip
	neteaseutil.UnlockSoundEffects = configs.ConfigRegistry.UNM.UnlockSoundEffects

	var playerCommand = commands.NewPlayerCommand()
	go web.Headler()
	app.Add(playerCommand)
	app.Add(commands.NewConfigCommand())
	app.DefaultCommand(playerCommand.Name)

	app.Run()
}
