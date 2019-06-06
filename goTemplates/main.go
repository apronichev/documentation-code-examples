package gcompl

type PageData struct {
	Items []Data
	Title string
	Name  string
}

type Data struct {
	Greeting string
	Info     string
}