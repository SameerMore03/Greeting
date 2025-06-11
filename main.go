package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

var tmpl = template.Must(template.New("form").Parse(`
<!DOCTYPE html>
<html>
<head>
    <title>Greeting</title>
    <style>
        body {
            font-family: 'Segoe UI', sans-serif;
            background-color: #e6f2f0;
            color: #004d40;
            text-align: center;
            padding-top: 50px;
        }
        input {
            padding: 10px;
            font-size: 16px;
            border: 1px solid #00796b;
            border-radius: 5px;
        }
        button {
            padding: 10px 20px;
            background-color: #00796b;
            color: white;
            border: none;
            border-radius: 5px;
            font-size: 16px;
            margin-left: 10px;
            cursor: pointer;
        }
        button:hover {
            background-color: #004d40;
        }
        h1 {
            color: #00695c;
        }
        h2 {
            color: #004d40;
        }
    </style>
</head>
<body>
    <h1>Welcome to the App</h1>
    {{if .Message}}
        <h2>{{.Message}}</h2>
		<form method="GET" action="/">
            <button type="submit">Go Back</button>
        </form>
    {{else}}
        <form method="POST" action="/">
            <label for="name">What's your name?</label><br><br>
            <input type="text" id="name" name="name" required />
            <button type="submit">Submit</button>
        </form>
    {{end}}
</body>
</html>
`))

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		name := strings.TrimSpace(strings.ToLower(r.FormValue("name")))

		var message string
		switch name {
		case "sidaa", "sida":
			message = "Hastess Kayy, Tithee Yeun Marin 😂"
		case "siddhi":
			message = "Tujhyaa Nanachiii Tangg, Ashi Haak Marto Ka Me Tula? "
		case "sarkaar", "sarkar", "sarrkar", "sarrkaar":
			message = "Namskaaaaaaaaar Sarkaaaaaaaar"
		default:
			message = "Tumhi amche sarkaar nahit, he greeting tumchyasathi nahi. 🙅‍♂️"
		}

		tmpl.Execute(w, map[string]string{"Message": message})
	} else {
		tmpl.Execute(w, nil)
	}
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
