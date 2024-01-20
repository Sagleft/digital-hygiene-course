package main

import "github.com/Sagleft/tgfun"

const (
	eventTitlesAnswer1 = "titlesAnswer1"
	eventTitlesAnswer2 = "titlesAnswer2"
	eventTitlesAnswer3 = "titlesAnswer3"
	eventTitlesAnswer4 = "titlesAnswer4"

	eventSourceAnswer1 = "sourceAnswer1"
	eventSourceAnswer2 = "sourceAnswer2"
	eventSourceAnswer3 = "sourceAnswer3"
	eventSourceAnswer4 = "sourceAnswer4"
)

// QUEST 1

func getTitlesQuest() tgfun.FunnelEvent {
	return tgfun.FunnelEvent{
		Message: tgfun.EventMessage{
			Text: `*Придирайся к заголовкам.* Они служат для привлечения внимания, однако за ярким заголовком может прятаться сомнительный контент. Не делай выводов и репостов лишь по заголовкам.

			😀 А теперь проверь себя. Найди заголовок, за которым скрывается фейк:

			1️⃣ Огонь поглотил дома. В Исландии разверзлась земная кора.

			2️⃣ Мавзолей Ленина реконструируют из-за появившейся грибковой плесени.

			3️⃣ В московском метро поставят камеры для распознания "праздношатаний."

			4️⃣ Прокуратура Украины подозревает Сталина в том, что тот еще жив.`,
			Buttons: []tgfun.MessageButton{
				{
					Text:          "1️⃣",
					NextMessageID: eventTitlesAnswer1,
				},
				{
					Text:          "2️⃣",
					NextMessageID: eventTitlesAnswer2,
				},
				{
					Text:          "3️⃣",
					NextMessageID: eventTitlesAnswer3,
				},
				{
					Text:          "4️⃣",
					NextMessageID: eventTitlesAnswer4,
				},
			},
			ButtonsIsColumns: false,
		},
	}
}

func getTitlesQuestRetryBtn() tgfun.MessageButton {
	return tgfun.MessageButton{
		Text:          "🔄 Попробовать снова",
		NextMessageID: eventAboutTitlesTag,
	}
}

func getTitlesAnswer1() tgfun.FunnelEvent {
	return tgfun.FunnelEvent{
		Message: tgfun.EventMessage{
			Text: `😏 *Это не фейк*, такое произошло

			Источник: РИА-Новости.
			https://ria.ru/20240116/vulkan-1921462781.html`,
			Buttons: []tgfun.MessageButton{
				getTitlesQuestRetryBtn(),
				getTip3Button(),
			},
			ButtonsIsColumns: true,
		},
	}
}

func getTitlesAnswer2() tgfun.FunnelEvent {
	return tgfun.FunnelEvent{
		Message: tgfun.EventMessage{
			Text: `✅ Верно! Действительно существовала новость с заголовком

			*"Мавзолей Ленина реконструируют из-за появившейся грибковой плесени"*

			Но автором новости был сайт сатирических новостей "Панорама" - источник забавных выдуманных новостей, которые кажутся серьезными, если не присмотреться к деталям.

			Даже отдельные СМИ, не занимавшиеся проверкой фактов, цитировали их новости. Это показало, что даже серьезные издания допускают грубые ошибки.`,
			Buttons: []tgfun.MessageButton{
				getTip3Button(),
			},
			ButtonsIsColumns: true,
		},
	}
}

func getTitlesAnswer3() tgfun.FunnelEvent {
	return tgfun.FunnelEvent{
		Message: tgfun.EventMessage{
			Text: `😏 *Это не фейк*, известное издание сообщает о госзакупке

			Сообщал "Интерфакс":
			https://www.interfax.ru/moscow/752414`,
			Buttons: []tgfun.MessageButton{
				getTitlesQuestRetryBtn(),
				getTip3Button(),
			},
			ButtonsIsColumns: true,
		},
	}
}

func getTitlesAnswer4() tgfun.FunnelEvent {
	return tgfun.FunnelEvent{
		Message: tgfun.EventMessage{
			Text: `😏 *Это не фейк*, но событие странное и политизированное

			Источник: РИА-Новости:
			https://ria.ru/20201223/ukraina-1590418497.html`,
			Buttons: []tgfun.MessageButton{
				getTitlesQuestRetryBtn(),
				getTip3Button(),
			},
			ButtonsIsColumns: true,
		},
	}
}

func getTip3Button() tgfun.MessageButton {
	return tgfun.MessageButton{
		Text:          "🧐 Совет №3 и тест",
		NextMessageID: eventAboutSourceTag,
	}
}

// QUEST 2

func getSourceQuest() tgfun.FunnelEvent {
	return tgfun.FunnelEvent{
		Message: tgfun.EventMessage{
			Text: `3️⃣ *Проверяй источники.* Достаточно подумать, где информация размещена, кто сделал ее репост, есть ли ссылки на первоисточник и куда они ведут, не является ли источник подделкой.

			😀 А теперь проверь себя. Выбери *неверный* пункт, который НЕ поможет тебе проверить статью, которую ты увидел в каком-либо канале?

			1️⃣ Найти ссылку на первоисточник.

			2️⃣ Проверить репутацию, известность первоисточника.

			3️⃣ Показать друзьям и если большинство из них считает, что это достоверно, значит так оно и есть.

			4️⃣ Определить, есть ли в статье скрытая коммерческая или политическая мотивация.`,
			Buttons: []tgfun.MessageButton{
				{
					Text:          "1️⃣",
					NextMessageID: eventSourceAnswer1,
				},
				{
					Text:          "2️⃣",
					NextMessageID: eventSourceAnswer2,
				},
				{
					Text:          "3️⃣",
					NextMessageID: eventSourceAnswer3,
				},
				{
					Text:          "4️⃣",
					NextMessageID: eventSourceAnswer4,
				},
			},
			ButtonsIsColumns: false,
		},
	}
}

func getSourceQuestRetryBtn() tgfun.MessageButton {
	return tgfun.MessageButton{
		Text:          "🔄 Попробовать снова",
		NextMessageID: eventAboutSourceTag,
	}
}

func getSourceAnswer1() tgfun.FunnelEvent {
	return tgfun.FunnelEvent{
		Message: tgfun.EventMessage{
			Text: `😏 *Нет, это полезный пункт*

			Если в посте, статье или другом материале не говорится об источнике информации, это скорее всего фейк.

			Или владелец этого материала не уважает своих читателей и считает их легковерными, что даже источник указывать не надо.`,
			Buttons: []tgfun.MessageButton{
				getSourceQuestRetryBtn(),
				getTip4Button(),
			},
			ButtonsIsColumns: true,
		},
	}
}

func getSourceAnswer2() tgfun.FunnelEvent {
	return tgfun.FunnelEvent{
		Message: tgfun.EventMessage{
			Text: `😏 *Нет, это полезный пункт*

			Даже если у материала указан источник, это может быть фальшивка или сатирическое издание, работа которого связана с выдумыванием фактов.

			К примеру, известный телеканал РЕН-ТВ активно занимался псевдо-документалистикой. Они специально придумывали невероятные лженаучные истории. И всё ради тв-рейтингов и продажи рекламы.`,

			Image: "source_2.jpg",

			Buttons: []tgfun.MessageButton{
				getSourceQuestRetryBtn(),
				getTip4Button(),
			},
			ButtonsIsColumns: true,
		},
	}
}

func getSourceAnswer3() tgfun.FunnelEvent {
	return tgfun.FunnelEvent{
		Message: tgfun.EventMessage{
			Text: `✅ *Правильно*, этот пункт неверен

			Знакомые и друзья могут совершить ошибку при оценке достоверности статьи. Они могут просто не разбираться в теме материала.

			Если отправлять материал на проверку знакомым, то тем, которые разбираются в этой теме и непредвзяты.`,

			Image: "source_3.jpg",

			Buttons: []tgfun.MessageButton{
				getTip4Button(),
			},
			ButtonsIsColumns: true,
		},
	}
}

func getSourceAnswer4() tgfun.FunnelEvent {
	return tgfun.FunnelEvent{
		Message: tgfun.EventMessage{
			Text: `😏 *Нет, это полезный пункт*

			В интернете нас постоянно атакует информация, в которую вплетен чей-то интерес. Интернет полон настоящих медиа-вирусов, призванных уверить тебя в чем-либо, чтобы использовать твоё внимание.
			
			Полезно критически относиться к информации, оценивая, а кто может быть заинтересован в том, чтобы я в "это" поверил?
			
			Так, например, можно отфильтровать заявления реальных ученых от шарлатанов, которые надели медицинский халат и предлагают купить их чудо-средство.`,

			Image: "source_4.jpg",

			Buttons: []tgfun.MessageButton{
				getSourceQuestRetryBtn(),
				getTip4Button(),
			},

			ButtonsIsColumns: true,
		},
	}
}

func getTip4Button() tgfun.MessageButton {
	return tgfun.MessageButton{
		Text:          "🧐 Совет №4",
		NextMessageID: eventAboutCheckInfoTag,
	}
}

func getCheckInfoMessage() tgfun.FunnelEvent {
	return tgfun.FunnelEvent{
		Message: tgfun.EventMessage{
			Text: `4️⃣ *Не переедай.* Даже здоровая еда вредна, если не знать меры. С информацией то же самое. Отпишись от информационного фастфуда.

			😉 Подытожим:
			1. Используй правильные инструменты.
			2. Придирайся к заголовкам.
			3. Смотри, есть ли источник и какой он.
			4. Не потребляй слишком много информации.`,

			Image: "source_5.jpg",

			Buttons: []tgfun.MessageButton{
				{
					Text:          "Подробнее про ИИ-ассистента",
					NextMessageID: eventAboutAITag2,
				},
			},
			ButtonsIsColumns: false,
		},
	}
}

func getEventAboutAI2() tgfun.FunnelEvent {
	return tgfun.FunnelEvent{
		Message: tgfun.EventMessage{
			Text: `😎 *Получи бесплатный доступ к ИИ-ассистенту*, который подскажет как решить любую проблему, найдет для тебя информацию и даже поможет сгенерировать какую-либо идею.

			Приложение *Utopia* доступно бесплатно для ПК и смартфона: Android, Windows, Linux и MacOS.`,

			Image: "source_6.jpg",

			Buttons: []tgfun.MessageButton{
				getDownloadUtopiaButton(),
			},
			ButtonsIsColumns: false,
		},
	}
}
