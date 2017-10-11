package main

const indexPage = 
`<!doctype html>
<html lang="ru">
<head>	
	<meta charset="utf-8">
	<!--[if IE]><meta http-equiv="X-UA-Compatible" content="IE=edge" /><![endif]-->
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<script src="https://ajax.googleapis.com/ajax/libs/jquery/1.8.2/jquery.min.js"></script>
	
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
			background: #FFECDF;
		}
		#wrapout {
			margin: 10px;
			position: absolute;
			top: 75%;
			bottom: 0;
			left: 0;
			right: 0;
			background: #7F6C5F;
			border: none;
		}
		#code, #output, pre, .lines {
			font-family: Consolas, Roboto Mono, Menlo, monospace;
			font-size: 11pt;
		}			
		#code, #output {
			border-width: 0;
			background: inherit;
			width: 100%;
			height: 100%;
			margin: 0;
			outline: none;
		}
		#code {
			color: black;
		}
		#output {
			color: white;
		}
		#output .system, #output .loading {
			color: #999;
		}
		#output .stderr, #output .error {
			color: #900;
		}

		div.numberedtextarea-wrapper {
			position: relative;
			height: 100%;
			width: 100%;
		}
		
		div.numberedtextarea-wrapper textarea {
			display: block;
			-webkit-box-sizing: border-box;
			-moz-box-sizing: border-box;
			box-sizing: border-box;
		}
		
		div.numberedtextarea-line-numbers {
			position: absolute;
			top: 0;
			left: 0;
			right: 0;
			bottom: 0;
			width: 50px;
			border-right: 1px solid rgba(0, 0, 0, 0.15);
			color: rgba(0, 0, 0, 0.15);
			overflow: hidden;
		}
		
		div.numberedtextarea-number {
			padding-right: 6px;
			text-align: right;
			font-family: Consolas, Roboto Mono, Menlo, monospace;
			font-size: 11pt;
		}
		
	</style>

	<script type='text/javascript'>
		(function ($) {
		
		   $.fn.numberedtextarea = function(options) {
		
			   var settings = $.extend({
				   color:          null,        // Font color
				   borderColor:    null,        // Border color
				   class:          null,        // Add class to the 'numberedtextarea-wrapper'
				   allowTabChar:   false,       // If true Tab key creates indentation
			   }, options);
		
			   this.each(function() {
				   if(this.nodeName.toLowerCase() !== "textarea") {
					   console.log('This is not a <textarea>, so no way Jose...');
					   return false;
				   }
				   
				   addWrapper(this, settings);
				   addLineNumbers(this, settings);
				   
				   if(settings.allowTabChar) {
					   $(this).allowTabChar();
				   }
			   });
			   
			   return this;
		   };
		   
		   $.fn.allowTabChar = function() {
			   if (this.jquery) {
				   this.each(function() {
					   if (this.nodeType == 1) {
						   var nodeName = this.nodeName.toLowerCase();
						   if (nodeName == "textarea" || (nodeName == "input" && this.type == "text")) {
							   allowTabChar(this);
						   }
					   }
				   })
			   }
			   return this;
		   }
		   
		   function addWrapper(element, settings) {
			   var wrapper = $('<div class="numberedtextarea-wrapper"></div>').insertAfter(element);
			   $(element).detach().appendTo(wrapper);
		   }
		   
		   function addLineNumbers(element, settings) {
			   element = $(element);
			   
			   var wrapper = element.parents('.numberedtextarea-wrapper');
			   
			   // Get textarea styles to implement it on line numbers div
			   var paddingLeft = parseFloat(element.css('padding-left'));
			   var paddingTop = parseFloat(element.css('padding-top'));
			   var paddingBottom = parseFloat(element.css('padding-bottom'));
			   
			   var lineNumbers = $('<div class="numberedtextarea-line-numbers"></div>').insertAfter(element);
			   
			   element.css({
				   paddingLeft: paddingLeft + lineNumbers.width() + 'px'
			   }).on('input propertychange change keyup paste', function() {
				   renderLineNumbers(element, settings);
			   }).on('scroll', function() {
				   scrollLineNumbers(element, settings);
			   }); 
			   
			   lineNumbers.css({
				   paddingLeft: paddingLeft + 'px',
				   paddingTop: paddingTop + 'px',
				   lineHeight: element.css('line-height'),
				   fontFamily: element.css('font-family'),
				   width: lineNumbers.width() - paddingLeft + 'px',
			   });
			   
			   element.trigger('change');
		   }
		   
		   function renderLineNumbers(element, settings) {
			   element = $(element);
			   
			   var linesDiv = element.parent().find('.numberedtextarea-line-numbers');
			   var count = element.val().split("\n").length;
			   var paddingBottom = parseFloat(element.css('padding-bottom'));
			   
			   linesDiv.find('.numberedtextarea-number').remove();
			   
			   for(i = 1; i<=count; i++) {
				   var line = $('<div class="numberedtextarea-number numberedtextarea-number-' + i + '">' + i + '</div>').appendTo(linesDiv);
				   
				   if(i === count) {
					   line.css('margin-bottom', paddingBottom + 'px');
				   }
			   }
		   }
		   
		   function scrollLineNumbers(element, settings) {
			   element = $(element);
			   var linesDiv = element.parent().find('.numberedtextarea-line-numbers');
			   linesDiv.scrollTop(element.scrollTop());
		   }
		   
		   function pasteIntoInput(el, text) {
			   el.focus();
			   if (typeof el.selectionStart == "number") {
				   var val = el.value;
				   var selStart = el.selectionStart;
				   el.value = val.slice(0, selStart) + text + val.slice(el.selectionEnd);
				   el.selectionEnd = el.selectionStart = selStart + text.length;
			   } else if (typeof document.selection != "undefined") {
				   var textRange = document.selection.createRange();
				   textRange.text = text;
				   textRange.collapse(false);
				   textRange.select();
			   }
		   }
	   
		   function allowTabChar(el) {
			   $(el).keydown(function(e) {
				   if (e.which == 9) {
					   pasteIntoInput(this, "\t");
					   return false;
				   }
			   });
	   
			   // For Opera, which only allows suppression of keypress events, not keydown
			   $(el).keypress(function(e) {
				   if (e.which == 9) {
					   return false;
				   }
			   });
		   }
		
	   }(jQuery));
	</script>

	<script type='text/javascript'>
		$(document).ready(function() {
			$('#code').numberedtextarea({
				allowTabChar: true,  // If true Tab key creates indentation
			  });
			$('#code').attr('wrap', 'off');
			$('#run').click(function(){
				var body = $("textarea#code").val();
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
</head>
<body bgcolor=#CCBDB3>
	<div id="head" itemprop="name">	
	<a href="https://github.com/covrom/gonec/wiki" style="text-decoration: none">
	<img id="headimg" alt="Интерпретатор ГОНЕЦ" src="data:image/png;base64,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"/>
	</a>
	<input type="button" value="Выполнить" id="run">
	<span class="header">v`+version+`</span></div>
	<div id="wrap">
		<textarea itemprop="description" id="code" name="code" autocorrect="off" autocomplete="off" autocapitalize="off" spellcheck="false" wrap="off"></textarea>
	</div>
	<div id="wrapout">
	<textarea id="output" autocorrect="off" autocomplete="off" autocapitalize="off" spellcheck="false" wrap="off" readonly></textarea>
	</div>
	<input type="hidden" id="sid" name="sid" value="">
</body>
`