package main

import "github.com/Sagleft/tgfun"

const (
	readMoreBtnText = "Далее ➡️"

	eventAboutThreatsTag      = "aboutThreats"
	eventAboutActionTag       = "toAction"
	eventAboutInstructionsTag = "instructions1"
	eventAboutHygieneTag      = "hygiene"
	eventAboutSafetyTag       = "safety"
	eventAboutSafety2Tag      = "safety2"
	eventAboutUtopiaTag       = "aboutUtopia"
	eventAboutAssistantTag    = "aboutAssistant"
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
						Text:          "Как перестать потреблять какую попало информацию?",
						NextMessageID: eventAboutSafetyTag,
					},
				},
			},
		},

		eventAboutSafetyTag: tgfun.FunnelEvent{
			Message: tgfun.EventMessage{
				Text: `*От некачественной информации можно защититься. Для этого надо соблюдать правила цифровой гигиены.*

				1️⃣ *Используй правильный инструмент* для поиска информации. Тебе нужен персональный ИИ-ассистент, которому можно написать вопрос и моментально получить вдумчивый ответ.

				🤨 Поисковые системы уже бесполезны, они находят мусорные сайты с кучей рекламы и минимумом ответов на заданный вопрос.

				😎 Попробуй использовать *приложение Utopia* для своего смартфона или компьютера, чтобы получить себе такого ИИ-ассистента бесплатно.`,

				Image: "hands.jpg",

				Buttons: []tgfun.MessageButton{
					getDownloadUtopiaButton(),
					{
						Text:          "Получить совет №2",
						NextMessageID: eventAboutSafety2Tag,
					},
				},
			},
		},

		eventAboutSafety2Tag: tgfun.FunnelEvent{
			Message: tgfun.EventMessage{
				Text: `2️⃣ *Придирайся к заголовкам.* Они служат для привлечения внимания, однако за ярким заголовком может прятаться сомнительный контент. Не делай выводов и репостов лишь по заголовкам.

				3️⃣ *Проверяй источники.* Достаточно подумать, где информация размещена, кто сделал ее репост, есть ли ссылки на первоисточник и куда они ведут, не является ли источник подделкой.

				4️⃣ *Проверяй информацию.* Кто может быть заинтересован в этой информации? Каковы мотивы публикации или репоста? Если хочешь лучше изучить тему, поищи противоположную информацию.

				5️⃣ *Не переедай.* Даже здоровая еда вредна, если не знать меры. С информацией то же самое. Отпишись от информационного фастфуда.`,
				Buttons: []tgfun.MessageButton{},
			},
		},

		/*eventAboutUtopiaTag: tgfun.FunnelEvent{
			Message: tgfun.EventMessage{
				Text: `🆓 Utopia - это мессенджер без отвлекающих факторов, без множества надоедающих оповещений и без цензуры информации.

				🤖 В приложении можно получить доступ к бесплатному ИИ-ассистенту, который поможет решить любой вопрос и найти для тебя информацию. Это сэкономит тебе массу времени, больше не придется ходить по горам сайтов с рекламой, на которых сложно что-то найти.

				🌐 Также в приложении есть браузер Idyll, в котором можно получить доступ к прокси-сервису, чтобы открыть любой сайт. Тебе больше не понадобится VPN.`,
				Buttons: []tgfun.MessageButton{
					getDownloadUtopiaButton(),
					{
						Text:          "🤖 Подробнее про ассистента",
						NextMessageID: eventAboutAssistantTag,
					},
					getProxyInfoButton(),
				},
			},
		},*/
		/*eventAboutAssistantTag: tgfun.FunnelEvent{
			Message: tgfun.EventMessage{
				Text: `Тут будет инфа`,
				Buttons: []tgfun.MessageButton{
					getDownloadUtopiaButton(),
					getProxyInfoButton(),
				},
			},
		},*/
		"placeholder": tgfun.FunnelEvent{
			Message: tgfun.EventMessage{
				Text: ``,
				Buttons: []tgfun.MessageButton{
					{
						Text:          "",
						NextMessageID: "",
					},
				},
			},
		},
	}
}

func getDownloadUtopiaButton() tgfun.MessageButton {
	return tgfun.MessageButton{
		Text: "Скачать Utopia для своей платформы",
		URL:  "https://guest.link/utopia",
	}
}

/*func getProxyInfoButton() tgfun.MessageButton {
	return tgfun.MessageButton{
		Text: "🤔 Как использовать прокси",
		URL:  "https://telegra.ph/Kak-ispolzovat-proksi-servis-v-Utopia-01-15",
	}
}*/
