package handlers

import (
	"aitu-news/aitu-news/aitu-news/pkg/models"
	"aitu-news/aitu-news/aitu-news/ui"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Получение последних статей из базы данных
	articles, err := models.GetLatestArticles()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Рендеринг HTML-шаблона
	ui.RenderTemplate(w, "home.html", articles)
}

func CategoriesHandler(w http.ResponseWriter, r *http.Request) {

	ui.RenderTemplate(w, "category.html", nil)
}

func CategoryArticlesHandler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered in CategoryArticlesHandler:", r)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}()

	vars := mux.Vars(r)
	category := vars["category"]

	fmt.Println("Category:", category) // Добавьте логирование

	articles, err := models.GetArticlesByCategory(category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Articles []models.Article
		Category string
	}{
		Articles: articles,
		Category: category,
	}

	ui.RenderTemplate(w, "category_articles.html", data)
}

func ContactsHandler(w http.ResponseWriter, r *http.Request) {
	// Получение списка контактов из базы данных
	contacts, err := models.GetContacts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Рендеринг HTML-шаблона с передачей списка контактов
	ui.RenderTemplate(w, "contacts.html", contacts)
}

func AddArticleHandler(w http.ResponseWriter, r *http.Request) {
	// Если это GET запрос, отображаем форму для добавления статьи
	if r.Method == http.MethodGet {
		ui.RenderTemplate(w, "add_article.html", nil)
		return
	}

	// Если это POST запрос, обрабатываем данные из формы
	if r.Method == http.MethodPost {
		// Получение данных из формы
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Создание новой статьи с данными из формы
		newArticle := models.Article{
			Title:   r.Form.Get("title"),
			Content: r.Form.Get("content"),
			ForWho:  r.Form.Get("forwho"),
			// Другие поля статьи, если они есть
		}

		// Сохранение статьи в базе данных
		err = models.AddArticle(newArticle)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Рендеринг HTML-шаблона для отображения успешного добавления статьи
		ui.RenderTemplate(w, "add_article.html", nil)
		return
	}

	// Если метод запроса не GET и не POST, возвращаем ошибку "Method Not Allowed"
	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
}

func PostArticleHandler(w http.ResponseWriter, r *http.Request) {
	// Получение данных формы из POST-запроса
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Создание объекта Article с данными из формы
	article := models.Article{
		Title:   r.Form.Get("title"),
		Content: r.Form.Get("content"),
		ForWho:  r.Form.Get("forwho"),
	}

	// Добавление статьи в базу данных
	err = models.AddArticle(article)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Перенаправление на главную страницу после успешного добавления
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func DeleteArticleHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		// Рендеринг HTML-шаблона для отображения формы удаления статьи
		ui.RenderTemplate(w, "delete_article.html", nil)
		return
	}

	if r.Method == http.MethodPost {
		// Получение ID статьи из данных формы
		idStr := r.FormValue("id")
		if idStr == "" {
			http.Error(w, "ID статьи пуст", http.StatusBadRequest)
			return
		}

		// Преобразование значения поля в число
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Неверный ID статьи. Пожалуйста, введите числовой ID.", http.StatusBadRequest)
			return
		}

		// Удаление статьи из базы данных
		err = models.DeleteArticleByID(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Перенаправление на главную страницу после успешного удаления
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// Если запрос не GET и не POST, возвращаем ошибку Method Not Allowed
	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
}
