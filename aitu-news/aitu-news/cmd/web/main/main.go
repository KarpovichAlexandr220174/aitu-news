// main.go

package main

import (
	"aitu-news/aitu-news/aitu-news/cmd/web/handlers"
	"aitu-news/aitu-news/aitu-news/pkg/drivers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func main() {

	router := mux.NewRouter()
	handler := cors.Default().Handler(router)

	// Инициализация базы данных
	drivers.InitDB("user=aitu_news_user dbname=aitu_news_db password=0000 sslmode=disable")

	// Обработка маршрутов
	router.HandleFunc("/", handlers.HomeHandler)
	router.HandleFunc("/home", handlers.HomeHandler)
	router.HandleFunc("/contacts", handlers.ContactsHandler)
	router.HandleFunc("/add", handlers.AddArticleHandler)
	router.HandleFunc("/add_article", handlers.PostArticleHandler)
	router.HandleFunc("/delete_article", handlers.DeleteArticleHandler)

	// Передача управления роутеру для обработки маршрута "/category"
	router.HandleFunc("/category", handlers.CategoriesHandler)
	router.HandleFunc("/category/{category}", handlers.CategoryArticlesHandler)

	// Использование роутера вместо nil
	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		log.Fatal(err)
	}
}
