package gonecsvc

import (
		"github.com/covrom/gonec/version"
)

const indexPage = 
`<!doctype html>
<html lang="ru">
<head>	
	<meta charset="utf-8">
	<!--[if IE]><meta http-equiv="X-UA-Compatible" content="IE=edge" /><![endif]-->
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<script src="/gonec/src?name=jquery"></script>
	
	<title>Гонец:Микросервисы</title>

	<style type="text/css">
		#head {
			float: left;
			height: 45px;
			display:flex;
			align-items:center;
		}
		.header {
			color: #7F6C5F;
			font-size: 16px;
			font-family: sans-serif;
		}
		input[type=button] {
			margin: 10px;
			height: 30px;
			border: 1px solid #FF7822;
			font-size: 16px;
			font-family: sans-serif;
			background: #FF7822;
			color: #FFECDF;
			position: static;
			top: 1px;
			border-radius: 5px;
		}
		#wrap, #about {
			margin: 10px;
			position: absolute;
			top: 45px;
			bottom: 25%;
			left: 0;
			right: 0;
			background: #25282c;
		}
		#wrapout {
			margin: 10px;
			position: absolute;
			top: 75%;
			bottom: 0;
			left: 0;
			right: 0;
			background: #1D1F21;
			border: none;
		}
		#code, #output, pre, .lines {
			font-family: Consolas, Roboto Mono, Menlo, monospace;
			font-size: 11pt;
		}			
		#code, #output {
			border-width: 0;
			width: 100%;
			height: 100%;
			margin: 0;
			outline: none;
		}
		#output {
			color: white;
			background: inherit;
		}
		#output .system, #output .loading {
			color: #999;
		}
		#output .stderr, #output .error {
			color: #900;
		}
	
	</style>

</head>
<body bgcolor=#25282c>
	<div id="head" itemprop="name">	
	<a href="https://github.com/covrom/gonec/wiki" style="text-decoration: none">
	<img id="headimg" alt="ГОНЕЦ" src="data:image/png;base64,
	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
	SUVORK5CYII="/>
	</a>
	<input type="button" value="Выполнить" id="run">
	<span class="header">v`+version.Version+`</span></div>
	<div id="wrap">
		<div id="code" name="code">Функция бенч1(бд)
	дтнач = ТекущаяДата()
		
	ОбработатьСерв = Функция (соед)
		база = соед.Данные()
		запр = соед.Получить()
		Сообщить("Сервер получил соединение:",соед)
		Сообщить("Получен запрос:",запр)
		// транзакция с записью - может быть только одна в один момент времени, но она не мешает транзакциям чтения
		тр = база.НачатьТранзакцию(Истина)
		Попытка
			тр.Таблица("testing").Установить(Формат("%s",соед), запр)
			тр.ЗафиксироватьТранзакцию()
		Исключение
			тр.ОтменитьТранзакцию()
			Сообщить("Транзакция отменена!")
		КонецПопытки
	КонецФункции

	серв = Новый Сервер
	Попытка
		серв.Открыть("tcp", "127.0.0.1:9990", 20, ОбработатьСерв, бд)
	Исключение
		Сообщить(ОписаниеОшибки())
		Сообщить("Кажется сервер уже запущен, или тут какая-то другая ошибка, но мы все равно попробуем отправить запрос :)")
	КонецПопытки

	клиенты = []
	гр = Новый ГруппаОжидания

	фобр = Функция (соед)
			д = соед.Данные()
			Сообщить("Устанавливаем соединение:",соед)
			запр={
				"id":соед.Идентификатор(),
				"query":"Запрос по tcp протоколу",
				"num":д[1],
				}
			Сообщить("Отправляем:", запр)
			соед.Отправить(запр)
			д[0].Завершить()
		КонецФункции

	Для н=1 по 20 Цикл
		кли = Новый Клиент
		гр.Добавить(1)
		кли.Открыть("tcp", "127.0.0.1:9990", фобр, [гр, н])
		клиенты += кли
	КонецЦикла

	гр.Ожидать()
	серв.Закрыть()

	Возврат ПрошлоВремениС(дтнач)
КонецФункции

Функция бенч2(бд)
	дтнач = ТекущаяДата()

	ОбработатьHTTP = Функция (вых,вх)
		база = вх.Данные()
		запр = вх.Сообщение()
		Сообщить("Сервер получил запрос:\n", запр)
		тр = база.НачатьТранзакцию(Истина)
		Попытка
			тр.Таблица("testing").Установить(вх.Адрес(), запр)
			тр.ЗафиксироватьТранзакцию()
			вых.Отправить({"Статус":200, "Тело":"Запрос обработан успешно"})
		Исключение
			тр.ОтменитьТранзакцию()
			Сообщить("Транзакция отменена!")
			вых.Отправить({"Статус":501, "Тело":"Проблемы с записью в базу данных"})
		КонецПопытки
	КонецФункции

	серв = Новый Сервер
	Попытка
		серв.Открыть("http", "127.0.0.1:9990", 20, {"/test":ОбработатьHTTP}, бд)
	Исключение
		Сообщить(ОписаниеОшибки())
		Сообщить("Кажется сервер уже запущен, или тут какая-то другая ошибка, но мы все равно попробуем отправить запрос :)")
	КонецПопытки

	гр = Новый ГруппаОжидания

	фобр = Функция (соед)
		д = соед.Данные()
		Сообщить("Устанавливаем соединение:",соед)
		запр={
			"Метод": "GET",
			"Путь": "http://127.0.0.1:9990/test",
			"Заголовки": {
				"ReqId":соед.Идентификатор(),
				},
			"Параметры": {
					"id":соед.Идентификатор(),
					"query":"Запрос по http протоколу",
					"num": д[1],
				},
			}
		Сообщить("Отправляем:", запр)
		Попытка
			Сообщить("Ответ",д[1],":",соед.Запрос(запр))
		Исключение
			Сообщить(ОписаниеОшибки())
		КонецПопытки
		соед.Закрыть()
		д[0].Завершить() // группа ожидания
	КонецФункции

	Для н=1 по 20 Цикл
		кли = Новый Клиент
		
		// это асинхронный вариант:
		гр.Добавить(1)	
		кли.Открыть("http", "127.0.0.1:9990", фобр, [гр, н])
		
		// это синхронный вариант (раскомментируйте, закомментировав выше асинхронный вариант):
		// соед = кли.Соединить("http","")
		// Попытка
		// 	Сообщить("Ответ",н,":",соед.Запрос({
		// 		"Метод": "GET",
		// 		"Путь": "http://127.0.0.1:9990/test",
		// 		"Заголовки": {
		// 			"ReqId":соед.Идентификатор(),
		// 			},
		// 		"Параметры": {
		// 				"id":соед.Идентификатор(),
		// 				"query":"Запрос по http протоколу",
		// 				"num": н,
		// 			},
		// 		}))
		// Исключение
		// 	Сообщить(ОписаниеОшибки())
		// КонецПопытки
		// кли.Закрыть()

	КонецЦикла

	гр.Ожидать()

	серв.Закрыть()

	Возврат ПрошлоВремениС(дтнач)
КонецФункции

// Результаты запоминаем в базе данных
база = Новый ФайловаяБазаДанных
база.Открыть("test.db")

Сообщить("Все завершилось просто идеально: TCP за", бенч1(база), ", HTTP за", бенч2(база), "!")

// транзакция чтения - таких может быть сколько угодно, в разных горутинах
тран = база.НачатьТранзакцию(Ложь)
Попытка
	Сообщить(тран.Таблица("testing").ПолучитьВсе())
	тран.ОтменитьТранзакцию() // при чтении транзакция закрывается только таким способом
Исключение
	тран.ОтменитьТранзакцию()
	Сообщить(ОписаниеОшибки())
КонецПопытки
база.Закрыть()

	</div>
		<script src="/gonec/src?name=ace" type="text/javascript" charset="utf-8"></script>
		<script src="/gonec/src?name=acetheme" type="text/javascript" charset="utf-8"></script>
		<script src="/gonec/src?name=acelang" type="text/javascript" charset="utf-8"></script>
		<script type="text/javascript">
			require("ace/ext/language_tools");
			var editor = ace.edit("code");
			editor.getSession().setMode("ace/mode/gonec");
			editor.setTheme("ace/theme/tomorrow_night");
			editor.setOptions({
				enableBasicAutocompletion: true,
				enableSnippets: true,
				enableLiveAutocompletion: false
			});
			$(document).ready(function() {
				$('#run').click(function(){
					var body = editor.getValue();
					$.ajax('/gonec', {
						type: 'POST',
						data: body,
						processData : false,
						dataType: 'text',
						cache: false,
						beforeSend: function(xhr){
							xhr.overrideMimeType("text/plain");
							xhr.setRequestHeader('Sid', $("#sid").val());
						},
						success: function(data, textStatus, request) {
							$("#output").text(data);
							$("#sid").val(request.getResponseHeader('Sid'));
						},
						error: function(xhr, status, error) {
							$("#output").text(xhr.responseText);
						}
					});
				});
			});
			</script>	
	</div>
	<div id="wrapout">
	<textarea id="output" autocorrect="off" autocomplete="off" autocapitalize="off" spellcheck="false" wrap="off" readonly></textarea>
	</div>
	<input type="hidden" id="sid" name="sid" value="">
</body>
`