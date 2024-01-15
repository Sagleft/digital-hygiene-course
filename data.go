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
	eventAboutSafety3Tag      = "safety3"
)

func getData(cfg config) tgfun.FunnelData {
	return tgfun.FunnelData{
		Token: cfg.TelegramBotToken,
	}
}

func getStartEvent() tgfun.FunnelEvent {
	return tgfun.FunnelEvent{
		Message: tgfun.EventMessage{
			Text: `Интернет - это дар, который был дан человечеству. Мы можем обмениваться информацией с невиданной скоростью. Но мы стали применять этот дар себе во вред.

			Сеть интернет сейчас - это нездоровая среда. Интернет-безопасность - это здоровое нахождение в этой нездоровой среде. Точно также как в пандемию нам приходилось защищать себя от угроз, в сети интернет надо защищать себя от чужого влияния.`,
			Buttons: []tgfun.MessageButton{
				{
					Text:          readMoreBtnText,
					NextMessageID: eventAboutThreatsTag,
				},
			},
		},
	}
}

func getScript() tgfun.FunnelScript {
	return tgfun.FunnelScript{
		"/start": getStartEvent(),
		eventAboutThreatsTag: tgfun.FunnelEvent{
			Message: tgfun.EventMessage{
				Text: `Каждый день, контактируя с интернетом, мы становимся чьими-то клиентами. Наше внимание монетизируют, продают, используют, манипулируют нами. Это игра, в которой ты один выступаешь против армии маркетологов, которые ищут способы проникнуть в твой разум и заставить тебя сделать то, что выгодно им.

				Поэтому возникает вопрос, как оставаться свободным в интернете, чтобы не терять самого себя? Чтобы не вовлекаться в навязанное мнение, чтобы не попадаться на уловку fake news, чтобы не забывать, а какая цель у тебя была изначально. Ведь нас так легко отвлечь, а затем подменить ценности и понятия.`,
				Buttons: []tgfun.MessageButton{
					{
						Text:          "Что с этим делать?",
						NextMessageID: eventAboutActionTag,
					},
				},
			},
		},
		eventAboutActionTag: tgfun.FunnelEvent{
			Message: tgfun.EventMessage{
				Text: `Для этого надо научиться выстраивать границы. Надо управлять своим вниманием и направлять его на свои мысли.`,
				Buttons: []tgfun.MessageButton{
					{
						Text:          "Что конкретно можно сделать?",
						NextMessageID: eventAboutInstructionsTag,
					},
				},
			},
		},
		eventAboutInstructionsTag: tgfun.FunnelEvent{
			Message: tgfun.EventMessage{
				Text: `Есть несколько простых практик:

				1. Начать стоит с малого: достаточно не брать телефон в руку сразу как только проснешься. Сначала идем в ванную, а уже затем можно брать телефон. Если так хочется посмотреть утром время, то на ночь надо выключать в телефоне WIFI, чтобы первые минуты дня не начинались с накопившихся уведомлений.
				
				2. Перед сном 6 минут читай какую-нибудь книгу в бумажном виде. Это позволит на 68% снизить напряженность и поможет легче заснуть. Ведь постоянные уведомления и отвлечение в течение дня раскачивают нашу нервную систему. Всего 6 минут чтения перед сном успокоит тебя и утром ты проснешься с ясным сознанием.
				
				Простой попробуй, это не сложно.
				
				Это только первый шаг: снизить свою зависимость от устройств, чтобы применять их по назначению как инструмент, а не как дешевый способ для развлечений.`,
				Buttons: []tgfun.MessageButton{
					{
						Text:          "Узнать про цифровую гигиену",
						NextMessageID: eventAboutHygieneTag,
					},
				},
			},
		},
		eventAboutHygieneTag: tgfun.FunnelEvent{
			Message: tgfun.EventMessage{
				Text: `Шаг 2: цифровая гигиена.

				Информация сравнима с едой: одна нам полезна, а другая может навредить. Например, если не мыть фрукты перед употреблением, легко получить расстройство желудка. Большинство людей соблюдает гигиену в отношении еды, но к информации мы относимся очень небрежно: мы поглощаем ее тоннами, немытую и непроверенную и это сказывается на нашем самочувствии.
				
				С продуктами всё понятно: мы их моем, подвергаем термической обработке, проверяем на испорченность. Но если дело касается информации, мы реагируем на посты и без причины переживаем, не проверяем достоверность, делаем репосты и травим испорченной информацией друзей.`,
				Buttons: []tgfun.MessageButton{
					{
						Text:          "Как защитить себя?",
						NextMessageID: eventAboutSafetyTag,
					},
				},
			},
		},
		eventAboutSafetyTag: tgfun.FunnelEvent{
			Message: tgfun.EventMessage{
				Text: `Хорошая новость в том, что от некачественной информации можно защититься. Для этого надо соблюдать правила цифровой гигиены.

				1. Придирайся к заголовкам. Они служат для привлечения внимания, однако за ярким заголовком может прятаться сомнительный контент. Не делай выводов и репостов лишь по заголовкам.
				
				2. Проверяй источники. Достаточно подумать, где информация размещена, кто сделал ее репост, есть ли ссылки на первоисточник и куда они ведут, не является ли источник подделкой.`,
				Buttons: []tgfun.MessageButton{
					{
						Text:          readMoreBtnText,
						NextMessageID: eventAboutSafety2Tag,
					},
				},
			},
		},
		eventAboutSafety2Tag: tgfun.FunnelEvent{
			Message: tgfun.EventMessage{
				Text: `3. Проверяй информацию. Кто может быть заинтересован в этой информации? Каковы мотивы публикации или репоста? Если хочешь лучше изучить тему, поищи противоположную информацию.
				
				4. Не переедай. Даже здоровая еда вредна, если не знать меры. С информацией то же самое. Отпишись от информационного фастфуда.`,
				Buttons: []tgfun.MessageButton{
					{
						Text:          readMoreBtnText,
						NextMessageID: eventAboutSafety3Tag,
					},
				},
			},
		},
		eventAboutSafety3Tag: tgfun.FunnelEvent{
			Message: tgfun.EventMessage{
				Text: `5. Используй источники свободные от цензуры. Например, мессенджер Utopia, который доступен бесплатно для смартфона и компьютера.`,
				/*Buttons: []tgfun.MessageButton{
					{
						Text:          "",
						NextMessageID: "",
					},
				},*/
			},
		},
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
