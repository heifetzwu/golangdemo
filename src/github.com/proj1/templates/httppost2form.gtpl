{{if .Success}}
	<h1>Thanks for your message!</h1>
{{else}}
	<h1>Contact</h1>
	<form method="POST">
		<label>Email:</label><br />
		<input type="text" name="email"><br />
		<label>Subject:</label><br />
		<input type="text" name="subject"><br />
		<label>Message:</label><br />
		<textarea name="message"></textarea><br />
		<input type="submit">
	</form>
{{end}}