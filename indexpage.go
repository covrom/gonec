package main

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
	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"/>
	</a>
	<input type="button" value="Выполнить" id="run">
	<span class="header">v`+version+`</span></div>
	<div id="wrap">
		<div id="code" name="code">Функция ОбработатьСерв(соед)
	Сообщить("Сервер получил соединение:",соед)
	Сообщить("Получен запрос:",соед.Получить())
КонецФункции

серв = Новый Сервер
Попытка
	серв.Открыть("tcp", "127.0.0.1:9990", 1000, ОбработатьСерв, 0)
Исключение
	Сообщить(ОписаниеОшибки())
	Сообщить("Кажется сервер уже запущен, или тут какая-то другая ошибка, но мы все равно попробуем отправить запрос :)")
КонецПопытки

клиенты = []
Для н=1 по 10 Цикл

	кли = Новый Клиент

	фобр = Функция (соед)
			Сообщить("Устанавливаем соединение:",соед)
			запр={
				"id":соед.Идентификатор(),
				"query":"Запрос по tcp протоколу",
				"num":соед.Параметры(),
			}
			Сообщить("Отправляем:", запр)
			соед.Отправить(запр)
		КонецФункции

	кли.Открыть("tcp", "127.0.0.1:9990", фобр, н)

	клиенты += кли

КонецЦикла

Завершено = Ложь
Пока НЕ Завершено Цикл
	Завершено = Истина
	Для н=1 по 10 Цикл
		Если клиенты[н-1].Работает() Тогда
			Завершено = Ложь
			Прервать
		КонецЕсли
	КонецЦикла
	ОбработатьГорутины()
КонецЦикла

Сообщить("Все завершилось просто идеально!")
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