// package main

// import (
// 	"net/http"
// )

// func indexHandler(w http.ResponseWriter, r *http.Request) {
// 	// Устанавливаем заголовки CORS
// 	w.Header().Set("Access-Control-Allow-Origin", "*") // Разрешаем доступ из любого источника
// 	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
// 	w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

// 	// Отправляем HTML-страницу с воспроизведением видео
// 	http.ServeFile(w, r, "index.html")
// 	http.ServeFile(w, r, "main.js")
// 	http.ServeFile(w, r, "style.css")
// 	http.ServeFile(w, r, "1.jpg")

// }

// func main() {
// 	// Обработчик для вашей HTML-страницы
// 	http.HandleFunc("/tv", indexHandler)

// 	// Запускаем веб-сервер на порту 8080
// 	err := http.ListenAndServe(":8080", nil)
// 	if err != nil {
// 		panic(err)
// 	}
// }

package main

import (
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// Отправляем HTML-страницу с воспроизведением видео
	http.ServeFile(w, r, "index.html")
}

func main() {
	// Настройка маршрутизации статических файлов
	http.Handle("/", http.FileServer(http.Dir(".")))

	// Устанавливаем заголовки CORS глобально
	cors := func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*") // Разрешаем доступ из любого источника
			w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")
			if r.Method == "OPTIONS" {
				return
			}
			h.ServeHTTP(w, r)
		})
	}

	// Обработчик для вашей HTML-страницы
	http.Handle("/tv", cors(http.HandlerFunc(indexHandler)))

	// Запускаем веб-сервер на порту 8080
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
