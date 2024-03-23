package controller

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

type EndpointHandler struct {
	db *sql.DB
}

func (s *EndpointHandler) Login(c *gin.Context) {
}

func (s *EndpointHandler) Logout(c *gin.Context) {
}

func (s *EndpointHandler) Register(c *gin.Context) {
}

func (s *EndpointHandler) GetUsers(c *gin.Context) {
}

func (s *EndpointHandler) GetUser(c *gin.Context, userId UserId) {
}

func (s *EndpointHandler) GetMe(c *gin.Context) {
}

func (s *EndpointHandler) GetArticles(c *gin.Context) {
	db := s.db
	rows, err := db.Query("SELECT id, title, discription, image, categories.name AS category FROM articles INNER JOIN categories ON articles.category_id = categories.id")
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var articles []GetArticleResponse

	for rows.Next() {
		var article GetArticleResponse
		err := rows.Scan(&article.Id, &article.Title, &article.Description, &article.Image, &article.Category)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		articles = append(articles, article)
	}

	c.JSON(200, articles)
}

func (s *EndpointHandler) GetArticle(c *gin.Context, articleId ArticleId) {
	db := s.db
	rows, err := db.Query("SELECT id, title, discription, image, categories.name AS category FROM articles INNER JOIN categories ON articles.category_id = categories.id WHERE articles.id = $1", articleId)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var article GetArticleResponse

	if rows.Next() {
		err := rows.Scan(&article.Id, &article.Title, &article.Description, &article.Image, &article.Category)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(200, article)
}
