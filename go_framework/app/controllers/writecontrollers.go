package controllers

func (t *T) Write() {
	t.Url = "write.html"
}
func (t *T) Preview_post() {
	t.Html_data = map[string]string{
		"main": t.Form["main"],
		//"side": "<p><a href=\"/teapot\">・紅茶でも飲みませんか？</a></p>",
		"side": "<p><a href=\"/teapot\">" + t.Form["side"] + "</a></p>",
	}
	t.Url = "burogu.html"
}
