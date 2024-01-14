package models

import (
	"aitu-news/aitu-news/aitu-news/pkg/drivers"
)

// Article структура представляет собой модель для статьи
type Article struct {
	ID      int
	Title   string
	Content string
	ForWho  string
}

func GetArticles() ([]Article, error) {
	rows, err := drivers.DB.Query("SELECT id, title, content FROM articles")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []Article

	for rows.Next() {
		var article Article
		if err := rows.Scan(&article.ID, &article.Title, &article.Content); err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}

	return articles, nil
}

// GetLatestArticles возвращает последние статьи из базы данных
func GetLatestArticles() ([]Article, error) {
	rows, err := drivers.DB.Query("SELECT id, title, content FROM articles ORDER BY id DESC LIMIT 10")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []Article

	for rows.Next() {
		var article Article
		if err := rows.Scan(&article.ID, &article.Title, &article.Content); err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}

	return articles, nil
}

func AddArticle(article Article) error {
	_, err := drivers.DB.Exec("INSERT INTO articles (title, content, forwho) VALUES ($1, $2, $3)",
		article.Title, article.Content, article.ForWho)
	return err
}

func DeleteArticleByID(id int) error {
	_, err := drivers.DB.Exec("DELETE FROM articles WHERE id = $1", id)
	return err
}

func GetArticlesByCategory(category string) ([]Article, error) {
	rows, err := drivers.DB.Query("SELECT id, title, content, forwho FROM articles WHERE forwho = $1", category)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []Article

	for rows.Next() {
		var article Article
		if err := rows.Scan(&article.ID, &article.Title, &article.Content, &article.ForWho); err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}

	return articles, nil
}
