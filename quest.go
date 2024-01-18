package main

import "github.com/Sagleft/tgfun"

const (
	eventTitlesAnswer1 = "titlesAnswer1"
	eventTitlesAnswer2 = "titlesAnswer2"
	eventTitlesAnswer3 = "titlesAnswer3"
	eventTitlesAnswer4 = "titlesAnswer4"
)

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
