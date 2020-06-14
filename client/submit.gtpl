<!DOCTYPE html>
<html>
	<head>
		<meta charset="utf-8">
		<title>Send answer</title>
		<style>
			body {
				font-family: sans-serif;
			}
			h1 {
				background: #ddd;
			}
			#sidebar {
				float: left;
			}
		</style>
	</head>
	<body>
		<form action="/send" method="post">
            <h1>{{.Current_word}}</h1>
            Your answer:<input type="text" name="answer">
            <input type="submit" value="Send!">
        </form>
	</body>
</html>
