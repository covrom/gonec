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
	<img id="headimg" alt="ГОНЕЦ" src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAHAAAAAtCAYAAAB28O9iAAAABHNCSVQICAgIfAhkiAAAAAlwSFlzAAAEZQAABGUBWZCbYAAAABl0RVh0U29mdHdhcmUAd3d3Lmlua3NjYXBlLm9yZ5vuPBoAABE0SURBVHic7Zt7dNTVtcc/+/wm72SGgIgIXEkySjUkAVMUe2uL1kUR0WuLYanIKlYyQV62Vmtr673h1t6r9laFQiERihrQVuS2PpBaK2J7raBSIZCWx/AsBUTJOyGPmd++f0wmM5OZycvcddt7813rt9b8ztl7n33O/p1z9t7nDAxiEIMYxCAGMYhB/H/HQZh4CL7uha8cAtdAyk53T7xsRH5+WvDdmT1h0kDK/3uFGQghXrjGC+8L/FFhLfCfNhz2QvFAyAcQte9paZIZABQWJiD2loGS/feMT2XAg+A8GDDYVuCz4XUCQ4Gf7IDbFKRXAj1lCfGqBPsjFb0UIL2+7SpgX78V/z8ER38ZD8JEA5sUsjqKzgEJYTLPfASbamH9a/BFYH5cYUUvWCSd+SHN7V+i6PEvs/He6igaQVE1gFgqnwd5B2CIu2CibUtj/eFdBwFcOXkeVetDMZpV5939QpDd6S64w9fueNNhtU3EMA7Eha0H6g/veQ4gIyfvZhGZgIpl2f6NNUf2VqaNzb3AsswCVPwY/aTeu+enuN2JTk37NoBCc8OYzGUZx2umYeyzApMVM8Ly+563HVYyyvVi6zlbUKO6te7w3g9c7oJZQb0ysgsuQfSfGqR5uctO+YaKSUSRpBTf422tjunh+jP6qpSM5KbrGryVr4QPS78MeAhmK5QrpAIngXv/ApsuAJcDlgvkH4N9LXA30Lp7TFb7FROvWpVaf3Lxtm3bfBHCbl9xEebjCuBqwCYx8UbgmVjtikhlWnbeeIVExG4BsG17JJacBUgdN/FC9fsnKP4zIJd1Gi87b7oo7Q7Ln47Iovok/81UVbU53fkLnNkTJtlGmwU7q95bWQoYZ3bBcnJz77XOOb6d6Gj5l0/2729Id+d90ekumA3q9vkcZc3Hdp5KdxfkOo9XL1IRS2zx1R+ufAKKrIycfSuNylbx+35Ze6Rqd+DDyn8M+EDRTr0EvQ1IyNCUe2zMLxsO7T6QcknBKFqt+8COWLWGJtsJ7co4IMKAfV5CvfCQwnoCxvu1wgQ3/OIa8F0KZy+G2V749xaYCZx68fLPPfOriZMXgs4/OyT7u8xZNR6Ar6/NYM6KJRh2gV4NtCPMZsOimMYDEJ/5rRG5FTgTqz7B579T0YrwMld2/kxEzqs7tPsXHVK2UFXVBlBP81oRe6qoXu8gdW0Hiw2yJb0tcbIY/eST/fsbABq9e942IntRNc3Hdp4KlO2uQmQYgLHNSwH2jX5V8xtRxoTroUhr+Lszq+AKjGxH1C+Q3nB49wGAcwd2/9WI+VW3RghDnw1ow2+AGoHSHJh+MXzclWYaPOeHr628dsaOqtEXeQBttJyr96Rfvhj17+GOFdW0nTuLsgwYArRiU0TFop9313bt0V21Aldimze61mVk5Y8D8ZrgQKkaV3beA7aQrWion0Jt52+vt1WFBBHSq7076oPFavzVxq9jsKmJaP/grl0gV7iy8x4IPsBoADFWJ63DSDWq6bYxM105eR5nTv4ylUhZiF5f7939euC32F3bUVsSXDl5Hqc7vzTQt9jo8xJ6Cew4AtlZhA1EDNwEFZPS05MFrjmdMOqZ4yljS4CkjurMyM5QzHOLXupN+6qyoOHwroNOd16kCEtvr/NWLnW58ycAqJhEY5u1DYc/POZy5z+YelHhSPCjaFI4W4dQH4WFCezc2Q5gYSWhdoMtDI9svchC9u2oO7Tn0fDSjJyCb7VLe6dcG00SaDO2/VpwCXW6C+ZkZo3P9wMZ2QUzMPp6WK+6OHlFlph97XXePeWAcebklQL/EWs8+uWF9mS8IN5/+edPHUoYN/l4ytjphIzXFZVULKqIUxeF4FITAZsrxa//BXR+yYK21B758BhAWnPCkw7LFwhpVDrjR1f2+EKUfSpS2eHZBkhs/5U+y3pXhAs7af8hL9OVve8+UUklNG6S4c6/C8CoCcWlak8S5a8ROooe8ItjDKrGGM2r91ZuD5FLAoWFQQ/cZLj3PRzRO0HjjUe/vVAALebLwA2M4htSih2L5uym0n1MKb2UMedNRVlP19mH/DUWXx81+ce6I3uXx6s9eXJns8ud3+5XTTJwLiMn/z6DNKhoQX2S/xtUXeZ35ex/1JWdnwvqVENjo/fDj53Zedsy3AWPYHNARSdpQtv3LV/SPzhz8p5UZKeghdimQkW/gPizXDl5xSDnK1Knok0qZqYrJ+9KG4aoamaDXbvUaTmXqm1uixgBzDpnXfsyzcl7T+ByQZ62bW5x5eR5FEar6q/jBWK9i89iDZmHEmAFgY/gm1LOk3GJb185AUtfpsvG3oEmjDWOZ+/u3pCjr0rhxLvnot4LCxM4e9bi6NGWQEWRReFhw0eJjgj63NzEjGYrSyy5on505vOpRxqGB52RINJz8s9PlJSW8P2QsWOTh1rpw6sP7T0BHTNhyhRH5vGPL6zJTD7Fzp3tGTkF33KoY42V0Gj7WlOs2qO7asnNTXQ1mDSAugy7Keg4Dc/NTf+4qqoxul9F1tCcP19YPea8U2zb5mP0VSku05hcd/wz9bDRDwhjxyaF+hlAnw2oIHj4N+A7HUWruJAlUoovJsPsFV9BqADSAB/ofSBPdGn7fdS+jg1L6mPKGCA4sydcjNHJ9d7dvV6ye4OgAWsO76wbSLm9Qd/3QA+PEDSe8qiUsyCu8eas+C7CiwSMVwtMZ/3iZaC/60I5CZHXmfvEkD7r0wf4/FajLfaxARes7E9poX3A5fYCfZqBWswPEL7f8bpUyimNSbh4eRI11hrQOzqaOQhyI+sX7AdgzqrxqP8dwNlFm520tE2NmYnpqss8RgBZGBTDEVkdOzaM4vMwEiULQxsOjshKzsYknLsumbZzKQC0WzYbSyJn19wnhtCWFBi/k2ca2FYa+Ig9ZQnUkNqjIpm0U17S3LX4FUhNgCw/DANON8GhWeCPJ6bXM1A93IzwvY7XSspZGpe4xqwNGY+ttLZO7jQeQMXdexFzE9AYwacUkpSwmaLHU+LqUcwkLeYtDKcwvAtsx+a0lvCmzmdiN/p/Xj28A5xEeAflfdr5SD28rHfhjmLwN92HsasxdjVJ7We5c3lkSOFLONhZP2p4qN1z7beS1F7b49PkiwibXoe0zbDMwCk/7AXeBvanwdHNcE+8fHKvDKilOIDQviWcEuK7tqj9HeAAUEZqwrSYM6piwduoTgO67HsymaTEh6PoAS3hSwi/R5hCZIcE5Vps/qDzuDqKr5ibgLeAz3WpsoAbsfhAi8mN2x+w8MlN3dR/KrwCqX74HbCErqtSIFHw5BZYHYu3dzPwJG5gbOe7cp0WRw1GCBuWnKBVClm/aD7lJfH3hg2L30GYSpQRmc+U0ogQRz0koDxLKJ5sACoIpPUaOsqSMTyrpaF+qYdUhKcJhUy1CE8jvAAE01uuDpr4UDOz2/rYqAP+EPMxujdIZOD7CpeH8b0FlClUdTYPnlfhy10b6F0c6OcMFjYhg1sIL2kJ10kZu2PybFzYCKBFWGSSLuXE9tAqFu1gzk/mofJCWGkqo4amEm5YYTraGVi3IEyWMv4EoPO4FMMHBPKzYznBFAJHXCB8Fe2MPeuxuVzWcARA5zMFm60EZvNndR75sobK2IOg11H0+NDe7M9hOv+RikXXdkdSGhjTO8OKHrkBvgvwFjia4CWB6QAG5gKvh/P3agbKWqoJfO3hOA9lq86jMB6fLsZJJi8Bv9P5nB+3gYrFGyHiQ3gnRkgRnjvbHDQegKzhz0ApcB/C5zD8PqREBN+moPEAZDXbEP7YWWuRH68rQAJJSTPi9qGfmATnAxeEVOBHwd/XgE9hqcKPgZsJLLER6H0mJokltHI5kQM5FMNv1cM0KWdHOLnOI4tWXgYCpw823wIeiCnbU5ZKc3uwEy3Y5p4oGg11EjjatVrKQx2P4gvtltEhhHIUOj5CZWRMGYgX9GLQW4BnY9P0DzaMCJtF9dMgYobfCO8ReGKi116o/IR6HEwlbF3uwBDgNb2b7GCBzuNqDDsIGE+BH1LeGfhHorTU0OwrA0YEaKWE5xbs7EHXuG51tOIRfNHpPomQZcWUoRrMW05l9vKuTsanRXibve9XB/qUC5Wfclrv5BoSeAMoCKsaip/HgFu0hDtRVgOJQAtwl5TzXEyBRS9YeM8EQw4FWcL6hQP6hQ8IRCtBGoAMxLoBeJ7uvPAgVD/D7BXrosoT7W+zbknUMVx/0OdktqzjY72La7F4m+DyGMAU9fAYyv0d76ex+YqsYXsMMVC0Mp3kM8+jzAB8KCVsWPizvujS4Zn+IEZVhTwVtVL0H2KGoLoZuBW1byFgwLZeMI5EmBtVbPMwMc5R+4N+nUbIWqr1bm7Bz94wGcOg03i7cHCT/JS/xBWSbH8PlRlAI6qz2LC477fM0nDQFHNf3U70Ut9/2JoBsgnRWxGZxpwfpaG09Mz4P49+HyfJKvarhz0Qlf04hXI/Pk53K+AvZx9i1HmpiP0zNiyJHYr8rUBIRZq3oCnNQCp26vWINvbIBx9g2wujSjM5MVCq9duAWkQiMCpG1UiEN4A/6Q+4SR7iUEwBgdxhtLfZF5ygjUxmdbz9CLjoU8mLC0mm4v4m7lixBZiJ4RZsTiMRfkAMNhp4bklcD3Ig0P97oUP5HsSN7WoYx1bOZ6+WcXtvxKVljR/RVxVkI34pZ6OUs5Fe3hLoH4JXHmRT4FVvoOsdl/8l9MuAWsw9KA+FFdUBxwFFOEgevyadRUDyR5Vcs5lujOh2J7nceWstY/YMycqN/0WHXyuQmHqHoj2NCBc07Fd0QtjuIcwIR2LyqwQ863TQ8d3S9hImss0+n8/22YBawgMIT3Y0dgyhiBqGSTkXIYyngJdI5Dag9c8vsuyDldwKVKxLSpreVVZGdsElTk19V1W+DgzxGxP3NAGN8NpiBdyhGaxhR0s984XKtIcjqZ/d1UDgVh4gl3ZL20vYkd6o6/XA2WknXoHPvAoLNkNeaQx79cmA6uFBlEc6Xn9BGxOkjBdlYyAAlTL+hMWDQNm7j7H6yBssFkh7P81Zcc/ocU85cwrWpbsnDs/MLnRluAvuFdGdBJygM0bM1IZDe56O33jYVXplRsd5YFCvGwgZ0I/BG8YZ4hO+qneH7uRoCZchTI7ZRlx0LKPxgv4+4hycJrT8iw/uCqmDGCgVWAlUTgqELxHomxNjsxZDEVAu5ayKRSIltL8AC1PhGYHWNcNHlb/pHFYMpILONeqf6xc/ElrYtjuEomrvru49M+FlAku1C3BheFc9rAvIZXEY3RYp45OwHr6In2VACnA+frarh2eBVJR5hAxxkDWR6cCYcLS9jC+hjUCionsol3HHyjjXN3Qv6xc9Ogv8r8LzErjFDvDjzXAZcGgLfInIE4hNXaX0aQbKGj7iQgrjGS+IWeBvhq/9+IKxC990DlsCMU+ofSil9WOGXl3trezRrZZymlEWEtrTsoB/JXC9I7jsNIYlEgJ8q6hBuDes6BLgYeBBQk5YO0pxt2ecQTz9zVqQN3ukC2BEIMsU41G5Lkik8M8EfAgITKoS4DENM57Ca+/Di10b6PMeGO/6YFfMAv/W00fXIboYIjw2G+UVW6yC+sOVS+n6X4nu2n6KDSgziZWUhj3AF6U8ehmUMlYDs4FYN9+qgGvlKd7urR6IHTUTPg1uhE98cBXwKtGO1DlgeRrMLI3hZPX7WmFfcN64cRlt7SlfEGOnGNv/Xs3hquM9c8WHlmI4wRUIORh8CAdkNR/2yFeExRCuxJCNTRsWB+Py3VE2EtsfOH+0/NVULO48hsJTlkqjP+TEWE37qLi/CYCix4eSlDi2515IPesXeruWboaLBPIVhimcAt6bwd9GyDKIQQxiEIMYxCAGMYhBBPDfJ4ZhgAQp7nEAAAAASUVORK5CYII="/>
	</a>
	<input type="button" value="Выполнить" id="run">
	<span class="header">v`+version+`</span></div>
	<div id="wrap">
		<div id="code" name="code">function int(t)
		return t:byte(1)+t:byte(2)*0x100+t:byte(3)*0x10000+t:byte(4)*0x1000000
	end
	    </div>
		<script src="/gonec/src?name=ace" type="text/javascript" charset="utf-8"></script>
		<script src="/gonec/src?name=acetheme" type="text/javascript" charset="utf-8"></script>
		<script src="/gonec/src?name=acelang" type="text/javascript" charset="utf-8"></script>
		<script type="text/javascript">
			require("ace/ext/language_tools");
			var editor = ace.edit("code");
			editor.getSession().setMode("ace/mode/lua");
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