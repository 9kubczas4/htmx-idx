package main

import (
	"fmt"
	"net/http"
	"text/template"
)

var tmpl = template.Must(template.New("").Parse(`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>htmx and Go App</title>
		<script src="https://unpkg.com/htmx.org@1.9.10" integrity="sha384-D1Kt99CQMDuVetoL1lrYwg5t+9QdHe7NLX/SoJYkXDFfX37iInKRy5xLSi8nO7UC" crossorigin="anonymous"></script>
</head>
<body>
    <h1>htmx and Go App</h1>
    <div id="content" hx-get="/api/data" hx-trigger="load">
        Loading...
    </div>
</body>
</html>
`))

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, nil)
	})

	http.HandleFunc("/api/data", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World - loaded via htmx!")
	})

	fmt.Println("Server is running on http://localhost:9000")
	http.ListenAndServe(":9000", nil)
}
