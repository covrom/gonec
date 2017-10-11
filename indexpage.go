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
			font-size: 20px;
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
	<img id="headimg" alt="Интерпретатор ГОНЕЦ" src="data:image/png;base64,
	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"/>
	</a>
	<input type="button" value="Выполнить" id="run">
	<span class="header"><b>v`+version+`</b></span></div>
	<div id="wrap">
		<textarea itemprop="description" id="code" name="code" autocorrect="off" autocomplete="off" autocapitalize="off" spellcheck="false" wrap="off"></textarea>
	</div>
	<div id="wrapout">
	<textarea id="output" autocorrect="off" autocomplete="off" autocapitalize="off" spellcheck="false" wrap="off" readonly></textarea>
	</div>
	<input type="hidden" id="sid" name="sid" value="">
</body>
`