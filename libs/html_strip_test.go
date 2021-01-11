package libs

import (
	"testing"
)

func TestStripTags(t *testing.T) {
	html := "<p>Форма для ваших вопросов: <a href=\"https://forms.gle/WqCy4DTAwHbvvpM67\" target=\"_blank\">https://forms.gle/WqCy4DTAwHbvvpM67</a>. В одном из выпусков мы на него, возможно, ответим.</p><p>***<br />Темы выпуска<br />02:38 <a href=\"https://reminder.media/super/14-priemov-kotorye-pomogayut-razobratsya-v-sebe-prinyat-reshenie-ili-povysit-effektivnost\" target=\"_blank\">14 приёмов, которые помогают разобраться в себе: от некролога до часа в полной тишине</a><br />15:28 <a href=\"https://twitter.com/Twitter/status/1334542969530183683\" target=\"_blank\">Твиттер предложил людям описать 2020-й одним словом</a><br />25:13 <a href=\"https://theguardian.com/world/2020/dec/08/mount-everest-china-and-nepal-agree-on-new-taller-height\" target=\"_blank\">Эверест «подрос» почти на метр</a><br />38:16 Рубрика «Вечная тема»: романтизация переработок</p><p>Отвечаем на ваши вопросы:<br />53:18 Как работать с мудаками<br />01:01:19 Про то, когда интересы партнеров различаются и кого-то в паре это начинает напрягать</p><p>***<br />Упоминали в выпуске<br />— <a href=\"https://blog.mann-ivanov-ferber.ru/2018/10/25/dnevnik-i-utrennie-stranicy-pismennye-praktiki-kotorye-pomogayut-ne-sojti-s-uma/\" target=\"_blank\">Про утренние страницы и другие письменные практики</a><br />— <a href=\"https://youtu.be/FSBYf_fCJUk\" target=\"_blank\">Вебинар про конфликты на Хабре с Милой Кудряковой</a><br />— <a href=\"https://www.ozon.ru/context/detail/id/147394385/\" target=\"_blank\">Книга «Кровь, пот и пиксели»</a><br />— <a href=\"https://worldexpeditions.com/everest-base-camp-trek-guide\" target=\"_blank\">Про трек к базовому лагерю Эвереста</a></p><p>***<br />Вступайте в наш чат в Телеграме: <a href=\"https://t.me/hobacast\" target=\"_blank\">https://t.me/hobacast</a></p><p>Подписывайтесь на наш Патреон: <a href=\"https://www.patreon.com/hoba\">https://www.patreon.com/hoba</a></p><p>Чтобы поддержать нас и стать частью сообщества Хобы!</p><p>***<br />Участники выпуска<br />— Аня Линская, <a href=\"https://t.me/shel_sneg\" target=\"_blank\">https://t.me/shel_sneg</a><br />— Ваня Звягин, <a href=\"http://anchor.fm/omfg-podcast\" target=\"_blank\">http://anchor.fm/omfg-podcast</a><br />— Далер Алиёров, <a href=\"https://t.me/dalerblog\" target=\"_blank\">https://t.me/dalerblog</a><br />— Адель Мубаракшин, <a href=\"https://t.me/exarg\" target=\"_blank\">https://t.me/exarg</a></p><p>***<br />Спасибо нашим патронам!</p><p>— Евгению Звягину<br />— Анастасии Смирновой<br />— Паше Пастернаку<br />— Роману Далинкевичу<br />— Евгению Васкивскому</p>"
	validHtml := `Форма для ваших вопросов: <a href="https://forms.gle/WqCy4DTAwHbvvpM67" target="_blank">https://forms.gle/WqCy4DTAwHbvvpM67</a>. В одном из выпусков мы на него, возможно, ответим.***
Темы выпуска
02:38 <a href="https://reminder.media/super/14-priemov-kotorye-pomogayut-razobratsya-v-sebe-prinyat-reshenie-ili-povysit-effektivnost" target="_blank">14 приёмов, которые помогают разобраться в себе: от некролога до часа в полной тишине</a>
15:28 <a href="https://twitter.com/Twitter/status/1334542969530183683" target="_blank">Твиттер предложил людям описать 2020-й одним словом</a>
25:13 <a href="https://theguardian.com/world/2020/dec/08/mount-everest-china-and-nepal-agree-on-new-taller-height" target="_blank">Эверест «подрос» почти на метр</a>
38:16 Рубрика «Вечная тема»: романтизация переработокОтвечаем на ваши вопросы:
53:18 Как работать с мудаками
01:01:19 Про то, когда интересы партнеров различаются и кого-то в паре это начинает напрягать***
Упоминали в выпуске
— <a href="https://blog.mann-ivanov-ferber.ru/2018/10/25/dnevnik-i-utrennie-stranicy-pismennye-praktiki-kotorye-pomogayut-ne-sojti-s-uma/" target="_blank">Про утренние страницы и другие письменные практики</a>
— <a href="https://youtu.be/FSBYf_fCJUk" target="_blank">Вебинар про конфликты на Хабре с Милой Кудряковой</a>
— <a href="https://www.ozon.ru/context/detail/id/147394385/" target="_blank">Книга «Кровь, пот и пиксели»</a>
— <a href="https://worldexpeditions.com/everest-base-camp-trek-guide" target="_blank">Про трек к базовому лагерю Эвереста</a>***
Вступайте в наш чат в Телеграме: <a href="https://t.me/hobacast" target="_blank">https://t.me/hobacast</a>Подписывайтесь на наш Патреон: <a href="https://www.patreon.com/hoba">https://www.patreon.com/hoba</a>Чтобы поддержать нас и стать частью сообщества Хобы!***
Участники выпуска
— Аня Линская, <a href="https://t.me/shel_sneg" target="_blank">https://t.me/shel_sneg</a>
— Ваня Звягин, <a href="http://anchor.fm/omfg-podcast" target="_blank">http://anchor.fm/omfg-podcast</a>
— Далер Алиёров, <a href="https://t.me/dalerblog" target="_blank">https://t.me/dalerblog</a>
— Адель Мубаракшин, <a href="https://t.me/exarg" target="_blank">https://t.me/exarg</a>***
Спасибо нашим патронам!— Евгению Звягину
— Анастасии Смирновой
— Паше Пастернаку
— Роману Далинкевичу
— Евгению Васкивскому`

	newHtml := ValidateHTML(html)

	if newHtml != validHtml {
		t.Errorf("newHtml != validHtml, newHtml=%s", newHtml)
	}
}

func TestTagsShielding(t *testing.T) {
	t.Run("case #1, ToHtml: false", func(t *testing.T) {
		{
			a := TagsShielding("<br />", false)
			if a != "\n" {
				t.Errorf("not valid tags")
			}
		}

		{
			a := TagsShielding("<br />", true)
			if a != "<br />" {
				t.Errorf("not valid tags")
			}
		}
	})

	t.Run("case #2, ToHtml: true", func(t *testing.T) {
		{
			a := TagsShielding("<a ", false)
			if a != "[a] " {
				t.Errorf("not valid tags")
			}
		}
		{
			a := TagsShielding("[a] ", true)
			if a != "<a " {
				t.Errorf("not valid tags")
			}
		}
	})
}
