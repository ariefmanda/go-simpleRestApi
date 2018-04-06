package main
import (
  "encoding/json"
  "io"
  "net/http"
  "time"
)

type Todo struct {
    Name      string
    Completed bool
    Due       time.Time
}

type Todos []Todo

func todo(w http.ResponseWriter, r *http.Request)  {
  todos := Todos{
    Todo{Name: "Write presentation"},
    Todo{Name: "Host meetup"},
  }
  json.NewEncoder(w).Encode(todos)
}

func hello(w http.ResponseWriter, r *http.Request) {
  io.WriteString(w, "Hello world!")
}

func main() {
  mux := http.NewServeMux()
  mux.HandleFunc("/", hello)
  mux.HandleFunc("/todo", todo)
  http.ListenAndServe(":8000", mux)
}