package main

import "github.com/Sagleft/tgfun"

const (
	eventAboutHygieneTag   = "aboutHygiene"
	eventAboutAITag        = "aboutAI"
	eventAboutAITag2       = "aboutAI2"
	eventAboutAssistantTag = "aboutAssistant"
	eventAboutTitlesTag    = "aboutTitles"
	eventAboutSourceTag    = "aboutSource"
	eventAboutCheckInfoTag = "aboutCheckInfo"
)

func getData(cfg config) tgfun.FunnelData {
	return tgfun.FunnelData{
		Token:     cfg.TelegramBotToken,
		ImageRoot: "img/",
	}
}

func getStartEvent() tgfun.FunnelEvent {
	return tgfun.FunnelEvent{
		Message: tgfun.EventMessage{
			Text: `*Интернет - это дар, который был дан человечеству. Мы можем обмениваться информацией с невиданной скоростью. Но мы стали применять этот дар себе во вред.*

			😕 Однако интернет *теперь* - это нездоровая среда, потому что все хотят выудить ваши данные, которые стали разменной монетой для маркетологов.

			💊 Твоим антибиотиком в этой нездоровой среде станет *интернет-безопасность*`,

			Image: "start.jpg",

			Buttons: []tgfun.MessageButton{
				{
					Text:          "Подробнее ➡️",
					NextMessageID: eventAboutHygieneTag,
				},
			},
		},
	}
}

func getScript() tgfun.FunnelScript {
	return tgfun.FunnelScript{
		"/start": getStartEvent(),

		eventAboutHygieneTag: tgfun.FunnelEvent{
			Message: tgfun.EventMessage{
				Text: `*Информация как еда: одна нам полезна, а другая может навредить.*

				😵‍💫 Непроверенная информация подобна немытым фруктам. Мы много её потребляем, не проверяя. Это отрицательно сказывается на нашем самочувствии.

				Затем переживаем и репостим, распространяя ненадежные сведения среди друзей 🤦🏻‍♂️`,

				Image: "rotted.jpg",

				Buttons: []tgfun.MessageButton{
					{
						Text:          "🤔 Что делать?",
						NextMessageID: eventAboutAITag,
					},
				},
			},
		},

		eventAboutAITag: getEventAboutAI(),

		eventAboutTitlesTag: getTitlesQuest(),
		eventTitlesAnswer1:  getTitlesAnswer1(),
		eventTitlesAnswer2:  getTitlesAnswer2(),
		eventTitlesAnswer3:  getTitlesAnswer3(),
		eventTitlesAnswer4:  getTitlesAnswer4(),

		eventAboutSourceTag: getSourceQuest(),
		eventSourceAnswer1:  getSourceAnswer1(),
		eventSourceAnswer2:  getSourceAnswer2(),
		eventSourceAnswer3:  getSourceAnswer3(),
		eventSourceAnswer4:  getSourceAnswer4(),

		eventAboutCheckInfoTag: getCheckInfoMessage(),
		eventAboutAITag2:       getEventAboutAI2(),
	}
}

func getDownloadUtopiaButton() tgfun.MessageButton {
	return tgfun.MessageButton{
		Text: "🔗 Скачать Utopia",
		URL:  "https://guest.link/utopia",
	}
}

func getEventAboutAI() tgfun.FunnelEvent {
	return tgfun.FunnelEvent{
		Message: tgfun.EventMessage{
			Text: `*От некачественной информации можно защититься. Для этого надо соблюдать правила цифровой гигиены.*
	
			1️⃣ *Используй правильный инструмент* для поиска информации. Тебе нужен персональный ИИ-ассистент, которому можно написать вопрос и моментально получить вдумчивый ответ.
	
			🤨 Поисковые системы уже бесполезны, они находят мусорные сайты с кучей рекламы и минимумом ответов на заданный вопрос.
	
			😎 Попробуй использовать *приложение Utopia* для своего смартфона или компьютера, чтобы получить себе такого ИИ-ассистента бесплатно.`,

			Image: "hands.jpg",

			Buttons: []tgfun.MessageButton{
				getDownloadUtopiaButton(),
				{
					Text:          "🧐 Совет №2 и тест",
					NextMessageID: eventAboutTitlesTag,
				},
			},
			ButtonsIsColumns: true,
		},
	}
}
