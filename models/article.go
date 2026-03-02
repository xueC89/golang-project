package models

import (
	"fmt"
	"frontend-backend/database"
)

type Article struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	Tags       string `json:"tags"`
	Short      string `json:"short"`
	Content    string `json:"content"`
	Author     string `json:"author"`
	Createtime string `json:"createtime"`
}

func AddArticle(article Article) (int64, error) {
	i, err := insertArticle(article)
	if err != nil {
		return 0, err
	}
	return i, nil
}

func insertArticle(article Article) (int64, error) {
	sql := `
		INSERT INTO article (title, tags, short, content, author)
		VALUES (?, ?, ?, ?, ?)
	`
	return database.ModifyDB(sql, article.Title, article.Tags, article.Short, article.Content, article.Author)
}

func QueryArticleWithCon(sql string) ([]Article, error) {
	sql = "SELECT id, title, tags, short, content, author, createtime FROM article " + sql
	rows, err := database.QueryDB(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []Article
	for rows.Next() {
		var article Article
		if err := rows.Scan(&article.ID, &article.Title, &article.Tags, &article.Short, &article.Content, &article.Author, &article.Createtime); err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return articles, nil
}

func QueryArticleWithPage(page, size int) ([]Article, error) {
	// sql := fmt.Sprintf("LIMIT %d OFFSET %d", size, (page-1)*size)
	sql := fmt.Sprintf("LIMIT %d, %d", (page-1)*size, size)
	return QueryArticleWithCon(sql)
}
