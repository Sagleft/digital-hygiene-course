package main

import "github.com/Sagleft/tgfun"

func getData(cfg config) tgfun.FunnelData {
	return tgfun.FunnelData{
		Token: cfg.TelegramBotToken,
	}
}

func getScript() tgfun.FunnelScript {
	return tgfun.FunnelScript{}
}
