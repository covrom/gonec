package main

const jsAceLang=`ace.define("ace/mode/gonec_highlight_rules",["require","exports","module","ace/lib/oop","ace/mode/text_highlight_rules"], function(require, exports, module) {
	"use strict";
	
	var oop = require("../lib/oop");
	var TextHighlightRules = require("./text_highlight_rules").TextHighlightRules;
	
	var GonecHighlightRules = function() {
	
		var keywords = (
			"прервать|цикл|новый|иначе|иначеесли|для|функция|если|из|по|выбор|когда|другое|старт|параллельно|"+
			 "возврат|тогда|каждого|пока|или|и|не|попытка|вызватьисключение|исключение|продолжить|канал|модуль|"+
			 "конеццикла|конецесли|конецфункции|конецпопытки|конецвыбора|?"
		);
	
		var builtinConstants = ("истина|ложь|неопределено|null|длительностьнаносекунды|"+
			"длительностьмикросекунды|длительностьмиллисекунды|длительностьсекунды|"+
			"длительностьминуты|длительностьчаса|длительностьдня");
	
		var functions = (
			"число|строка|булево|целоечисло|массив|структура|дата|длительность|"+
			"импорт|длина|диапазон|текущаядата|прошловременис|пауза|хэш|"+
			"уникальныйидентификатор|получитьмассивизпула|вернутьмассиввпул|случайнаястрока|нрег|врег|"+
			"формат|кодсимвола|типзнч|сообщить|сообщитьф|обработатьгорутины"
		);
	
		var builtinTypes = ("группаожидания|сервер|клиент");
	
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
				regex : "(функция)(\\s+)([a-zA-Zа-яА-Я_$][a-zA-Zа-яА-Я0-9_$]*)\\b"
			}, {
				token : keywordMapper,
				regex : "[a-zA-Zа-яА-Я_$][a-zA-Zа-яА-Я0-9_$]*\\b"
			}, {
				token : "keyword.operator",
				regex : "!|\\$|%|&|\\*|\\-\\-|\\-|\\+\\+|\\+|~|==|=|!=|<=|>=|<<=|>>=|>>>=|<>|<|>|!|&&|\\|\\||\\?\\:|\\*=|%=|\\+=|\\-=|&=|\\^="
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
	
	ace.define("ace/mode/folding/gonec",["require","exports","module","ace/lib/oop","ace/mode/folding/fold_mode","ace/range","ace/token_iterator"], function(require, exports, module) {
	"use strict";
	
	var oop = require("../../lib/oop");
	var BaseFoldMode = require("./fold_mode").FoldMode;
	var Range = require("../../range").Range;
	var TokenIterator = require("../../token_iterator").TokenIterator;
	
	
	var FoldMode = exports.FoldMode = function() {};
	
	oop.inherits(FoldMode, BaseFoldMode);
	
	(function() {
	
		this.foldingStartMarker = /\b(function|then|do|repeat)\b|{\s*$|(\[=*\[)/;
		this.foldingStopMarker = /\bend\b|^\s*}|\]=*\]/;
	
		this.getFoldWidget = function(session, foldStyle, row) {
			var line = session.getLine(row);
			var isStart = this.foldingStartMarker.test(line);
			var isEnd = this.foldingStopMarker.test(line);
	
			if (isStart && !isEnd) {
				var match = line.match(this.foldingStartMarker);
				if (match[1] == "then" && /\belseif\b/.test(line))
					return;
				if (match[1]) {
					if (session.getTokenAt(row, match.index + 1).type === "keyword")
						return "start";
				} else if (match[2]) {
					var type = session.bgTokenizer.getState(row) || "";
					if (type[0] == "bracketedComment" || type[0] == "bracketedString")
						return "start";
				} else {
					return "start";
				}
			}
			if (foldStyle != "markbeginend" || !isEnd || isStart && isEnd)
				return "";
	
			var match = line.match(this.foldingStopMarker);
			if (match[0] === "end") {
				if (session.getTokenAt(row, match.index + 1).type === "keyword")
					return "end";
			} else if (match[0][0] === "]") {
				var type = session.bgTokenizer.getState(row - 1) || "";
				if (type[0] == "bracketedComment" || type[0] == "bracketedString")
					return "end";
			} else
				return "end";
		};
	
		this.getFoldWidgetRange = function(session, foldStyle, row) {
			var line = session.doc.getLine(row);
			var match = this.foldingStartMarker.exec(line);
			if (match) {
				if (match[1])
					return this.luaBlock(session, row, match.index + 1);
	
				if (match[2])
					return session.getCommentFoldRange(row, match.index + 1);
	
				return this.openingBracketBlock(session, "{", row, match.index);
			}
	
			var match = this.foldingStopMarker.exec(line);
			if (match) {
				if (match[0] === "end") {
					if (session.getTokenAt(row, match.index + 1).type === "keyword")
						return this.luaBlock(session, row, match.index + 1);
				}
	
				if (match[0][0] === "]")
					return session.getCommentFoldRange(row, match.index + 1);
	
				return this.closingBracketBlock(session, "}", row, match.index + match[0].length);
			}
		};
	
		this.luaBlock = function(session, row, column) {
			var stream = new TokenIterator(session, row, column);
			var indentKeywords = {
				"function": 1,
				"do": 1,
				"then": 1,
				"elseif": -1,
				"end": -1,
				"repeat": 1,
				"until": -1
			};
	
			var token = stream.getCurrentToken();
			if (!token || token.type != "keyword")
				return;
	
			var val = token.value;
			var stack = [val];
			var dir = indentKeywords[val];
	
			if (!dir)
				return;
	
			var startColumn = dir === -1 ? stream.getCurrentTokenColumn() : session.getLine(row).length;
			var startRow = row;
	
			stream.step = dir === -1 ? stream.stepBackward : stream.stepForward;
			while(token = stream.step()) {
				if (token.type !== "keyword")
					continue;
				var level = dir * indentKeywords[token.value];
	
				if (level > 0) {
					stack.unshift(token.value);
				} else if (level <= 0) {
					stack.shift();
					if (!stack.length && token.value != "elseif")
						break;
					if (level === 0)
						stack.unshift(token.value);
				}
			}
	
			var row = stream.getCurrentTokenRow();
			if (dir === -1)
				return new Range(row, session.getLine(row).length, startRow, startColumn);
			else
				return new Range(startRow, startColumn, row, stream.getCurrentTokenColumn());
		};
	
	}).call(FoldMode.prototype);
	
	});
	
	ace.define("ace/mode/lua",["require","exports","module","ace/lib/oop","ace/mode/text","ace/mode/gonec_highlight_rules","ace/mode/folding/gonec","ace/range","ace/worker/worker_client"], function(require, exports, module) {
	"use strict";
	
	var oop = require("../lib/oop");
	var TextMode = require("./text").Mode;
	var GonecHighlightRules = require("./gonec_highlight_rules").GonecHighlightRules;
	var LuaFoldMode = require("./folding/gonec").FoldMode;
	var Range = require("../range").Range;
	var WorkerClient = require("../worker/worker_client").WorkerClient;
	
	var Mode = function() {
		this.HighlightRules = GonecHighlightRules;
		
		this.foldingRules = new LuaFoldMode();
		this.$behaviour = this.$defaultBehaviour;
	};
	oop.inherits(Mode, TextMode);
	
	(function() {
	   
		this.lineCommentStart = "--";
		this.blockComment = {start: "--[", end: "]--"};
		
		var indentKeywords = {
			"function": 1,
			"then": 1,
			"do": 1,
			"else": 1,
			"elseif": 1,
			"repeat": 1,
			"end": -1,
			"until": -1
		};
		var outdentKeywords = [
			"else",
			"elseif",
			"end",
			"until"
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
			var worker = new WorkerClient(["ace"], "ace/mode/lua_worker", "Worker");
			worker.attachToDocument(session.getDocument());
			
			worker.on("annotate", function(e) {
				session.setAnnotations(e.data);
			});
			
			worker.on("terminate", function() {
				session.clearAnnotations();
			});
			
			return worker;
		};
	
		this.$id = "ace/mode/lua";
	}).call(Mode.prototype);
	
	exports.Mode = Mode;
	});
`