package gonecsvc

const jsAceLang=`ace.define("ace/mode/gonec_highlight_rules",["require","exports","module","ace/lib/oop","ace/mode/text_highlight_rules"], function(require, exports, module) {
	"use strict";
	
	var oop = require("../lib/oop");
	var TextHighlightRules = require("./text_highlight_rules").TextHighlightRules;
	
	var GonecHighlightRules = function() {
	
		var keywords = (
			"Прервать|Цикл|Новый|Иначе|ИначеЕсли|Для|Функция|Если|из|по|Выбор|Когда|Другое|Старт|Параллельно|"+
			 "Возврат|Тогда|каждого|Пока|ИЛИ|И|НЕ|Попытка|ВызватьИсключение|Исключение|Продолжить|Канал|Модуль|"+
			 "КонецЦикла|КонецЕсли|КонецФункции|КонецПопытки|КонецВыбора"
		);
	
		var builtinConstants = ("Истина|Ложь|Неопределено|NULL|ДлительностьНаносекунды|"+
			"ДлительностьМикросекунды|ДлительностьМиллисекунды|ДлительностьСекунды|"+
			"ДлительностьМинуты|ДлительностьЧаса|ДлительностьДня|АргументыЗапуска");
	
		var functions = (
			"Число|Строка|Булево|ЦелоеЧисло|Массив|Структура|Дата|Длительность|"+
			"Импорт|Длина|Диапазон|ТекущаяДата|ПрошлоВремениС|Пауза|Хэш|"+
			"УникальныйИдентификатор|ПолучитьМассивИзПула|ВернутьМассивВПул|СлучайнаяСтрока|НРег|ВРег|"+
			"Формат|КодСимвола|ТипЗнч|Сообщить|СообщитьФ|ОбработатьГорутины|ЗагрузитьИВыполнить|"+
			"ОписаниеОшибки|ПеременнаяОкружения|СтрСодержит|СтрСодержитЛюбой|СтрКоличество|СтрНайти|"+
			"СтрНайтиЛюбой|СтрНайтиПоследний|СтрЗаменить|Окр"
		);
	
		var builtinTypes = ("ГруппаОжидания|Сервер|Клиент");
	
		var keywordMapper = this.createKeywordMapper({
			"keyword": keywords,
			"support.function": functions,
			"support.type": builtinTypes,
			"constant.language": builtinConstants,
			"variable.language": "self"
		}, "identifier");
		
		this.$rules = {
			"start" : [
			{
				token : "comment",
				regex : "\\/\\/.*$"
			},
			{
				token : "comment",
				regex : "\\#.*$"
			},
			{
				token : "string", // single line
				regex : /"(?:[^"\\]|\\.)*?"/
			}, {
				token : "string", // raw
				regex : '`+"`"+`',
				next : "bqstring"
			}, {
				token : "constant.numeric", // hex
				regex : "0[xX][0-9a-fA-F]+\\b" 
			}, {
				token : "constant.numeric", // float
				regex : "[+-]?\\d+(?:(?:\\.\\d*)?(?:[eE][+-]?\\d+)?)?\\b"
			}, {
				token : ["keyword", "text", "entity.name.function"],
				regex : "(Функция)(\\s+)([a-zA-Zа-яА-ЯёЁ_$][a-zA-Zа-яА-Я0-9_$]*)(?![a-zA-Zа-яА-ЯёЁ])"
			}, {
				token : keywordMapper,
				regex : "[a-zA-Zа-яА-Я_$][a-zA-Zа-яА-Я0-9_$]*(?![a-zA-Zа-яА-ЯёЁ])"
			}, {
				token : "keyword.operator",
				regex : "!|\\$|%|&|\\*|\\-\\-|\\-|\\+\\+|\\+|~|==|=|!=|<=|>=|<<=|>>=|>>>=|<>|<|>|!|&&|\\|\\||\\:|\\*=|%=|\\+=|\\-=|&=|\\^="
			}, {
				token : "paren.lparen",
				regex : "[\\[\\(\\{]"
			}, {
				token : "paren.rparen",
				regex : "[\\]\\)\\}]"
			}, {
				token : "text",
				regex : "\\s+|\\w+"
			} ],
			"bqstring" : [
                {
                    token : "string",
                    regex : '`+"`"+`',
                    next : "start"
                }, {
                    defaultToken : "string"
                }
            ]
		};
		
		this.normalizeRules();
	}
	
	oop.inherits(GonecHighlightRules, TextHighlightRules);
	
	exports.GonecHighlightRules = GonecHighlightRules;
	});
		
	ace.define("ace/mode/gonec",["require","exports","module","ace/lib/oop","ace/mode/text","ace/mode/gonec_highlight_rules","ace/mode/folding/gonec","ace/range","ace/worker/worker_client"], function(require, exports, module) {
	"use strict";
	
	var oop = require("../lib/oop");
	var TextMode = require("./text").Mode;
	var GonecHighlightRules = require("./gonec_highlight_rules").GonecHighlightRules;
	var Range = require("../range").Range;
	var WorkerClient = require("../worker/worker_client").WorkerClient;
	
	var Mode = function() {
		this.HighlightRules = GonecHighlightRules;
		
		this.$behaviour = this.$defaultBehaviour;
	};
	oop.inherits(Mode, TextMode);
	
	(function() {
	   
		this.lineCommentStart = "//";
		
		var indentKeywords = {
			"Функция": 1,
			"Тогда": 1,
			"Цикл": 1,
			"Иначе": 1,
			"ИначеЕсли": 1,
			"Когда": 1,
			"Другое": 1,
			"Пока": 1,
			"Попытка": 1,
			"Исключение": 1,
			"КонецЦикла": -1,
			"КонецЕсли": -1,
			"КонецФункции": -1,
			"КонецПопытки": -1,
			"КонецВыбора": -1
		};
		var outdentKeywords = [
			"Иначе",
			"ИначеЕсли",
			"КонецЦикла",
			"КонецЕсли",
			"КонецФункции",
			"КонецПопытки",
			"КонецВыбора"
		];
	
		function getNetIndentLevel(tokens) {
			var level = 0;
			for (var i = 0; i < tokens.length; i++) {
				var token = tokens[i];
				if (token.type == "keyword") {
					if (token.value in indentKeywords) {
						level += indentKeywords[token.value];
					}
				} else if (token.type == "paren.lparen") {
					level += token.value.length;
				} else if (token.type == "paren.rparen") {
					level -= token.value.length;
				}
			}
			if (level < 0) {
				return -1;
			} else if (level > 0) {
				return 1;
			} else {
				return 0;
			}
		}
	
		this.getNextLineIndent = function(state, line, tab) {
			var indent = this.$getIndent(line);
			var level = 0;
	
			var tokenizedLine = this.getTokenizer().getLineTokens(line, state);
			var tokens = tokenizedLine.tokens;
	
			if (state == "start") {
				level = getNetIndentLevel(tokens);
			}
			if (level > 0) {
				return indent + tab;
			} else if (level < 0 && indent.substr(indent.length - tab.length) == tab) {
				if (!this.checkOutdent(state, line, "\n")) {
					return indent.substr(0, indent.length - tab.length);
				}
			}
			return indent;
		};
	
		this.checkOutdent = function(state, line, input) {
			if (input != "\n" && input != "\r" && input != "\r\n")
				return false;
	
			if (line.match(/^\s*[\)\}\]]$/))
				return true;
	
			var tokens = this.getTokenizer().getLineTokens(line.trim(), state).tokens;
	
			if (!tokens || !tokens.length)
				return false;
	
			return (tokens[0].type == "keyword" && outdentKeywords.indexOf(tokens[0].value) != -1);
		};
	
		this.autoOutdent = function(state, session, row) {
			var prevLine = session.getLine(row - 1);
			var prevIndent = this.$getIndent(prevLine).length;
			var prevTokens = this.getTokenizer().getLineTokens(prevLine, "start").tokens;
			var tabLength = session.getTabString().length;
			var expectedIndent = prevIndent + tabLength * getNetIndentLevel(prevTokens);
			var curIndent = this.$getIndent(session.getLine(row)).length;
			if (curIndent <= expectedIndent) {
				return;
			}
			session.outdentRows(new Range(row, 0, row + 2, 0));
		};
	
		this.createWorker = function(session) {
			var worker = new WorkerClient(["ace"], "ace/mode/gonec_worker", "Worker");
			worker.attachToDocument(session.getDocument());
			
			worker.on("annotate", function(e) {
				session.setAnnotations(e.data);
			});
			
			worker.on("terminate", function() {
				session.clearAnnotations();
			});
			
			return worker;
		};
	
		this.$id = "ace/mode/gonec";
	}).call(Mode.prototype);
	
	exports.Mode = Mode;
	});
`