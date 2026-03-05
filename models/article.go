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

func QueryArticleWithId(id string) (Article, error) {
	sql := fmt.Sprintf("WHERE id = %s", id)
	articles, err := QueryArticleWithCon(sql)
	if err != nil {
		return Article{}, err
	}
	if len(articles) == 0 {
		return Article{}, fmt.Errorf("article with id %s not found", id)
	}
	return articles[0], nil
}

func UpdateArticle(id string, article Article) (int64, error) {
	sql := `
		UPDATE article
		SET title = ?, tags = ?, short = ?, content = ?, author = ?
		WHERE id = ?
	`
	return database.ModifyDB(sql, article.Title, article.Tags, article.Short, article.Content, article.Author, id)
}

func DeleteArticle(id string) (int64, error) {
	sql := fmt.Sprintf("DELETE FROM article WHERE id = %s", id)
	return database.ModifyDB(sql)
}
