package main
 
import (
	"log"
	"net/http"
)
 
// Обработчик главной страницы.
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Привет из Snippetbox"))
}
 
// Обработчик для отображения содержимого заметки.
func showSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Отображение заметки..."))
}
 
// Обработчик для создания новой заметки.
func createSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Форма для создания новой заметки..."))
}
 
func main() {
	// Регистрируем два новых обработчика и соответствующие URL-шаблоны в
	// маршрутизаторе servemux
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)
 
	log.Println("Запуск веб-сервера на http://127.0.0.1:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

package main
 
...
 
// Обработчик для создания новой заметки.
func createSnippet(w http.ResponseWriter, r *http.Request) {
    // Используем r.Method для проверки, использует ли запрос метод POST или нет. Обратите внимание,
    // что http.MethodPost является строкой и содержит текст "POST".
    if r.Method != http.MethodPost {
        // Если это не так, то вызывается метод w.WriteHeader() для возвращения статус-кода 405
        // и вызывается метод w.Write() для возвращения тела-ответа с текстом "Метод запрещен".
        // Затем мы завершаем работу функции вызвав "return", чтобы
        // последующий код не выполнялся.
        w.WriteHeader(405)
        w.Write([]byte("GET-Метод запрещен!"))
        return
    }
 
    w.Write([]byte("Создание новой заметки..."))
}

package main
 
import (
    "fmt" // новый импорт
    "log"
    "net/http"
    "strconv" // новый импорт
)
 
...
 

func showSnippet(w http.ResponseWriter, r *http.Request) {

    id, err := strconv.Atoi(r.URL.Query().Get("id"))
    if err != nil || id < 1 {
        http.NotFound(w, r)
        return
    }
 

    fmt.Fprintf(w, "Отображение выбранной заметки с ID %d...", id)
}
