package controller

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func NewEndpointHandler(db *sql.DB) *EndpointHandler {
	return &EndpointHandler{db: db}
}

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
	db := s.db
	rows, err := db.Query("SELECT id, nickname, age, gendar.name AS gender, image, in_kitakyushu, email, password FROM users INNER JOIN genders ON users.gender_id = gender_id")
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var users []GetUsersResponse

	for rows.Next() {
		var user GetUsersResponse
		err := rows.Scan(&user.Id, &user.Nickname, &user.Age, &user.Gender, &user.Image, &user.InKitakyushu, &user.Email)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		users = append(users, user)
	}

	c.JSON(200, users)
}

func (s *EndpointHandler) GetUser(c *gin.Context, userId UserId) {
	db := s.db
	rows, err := db.Query("SELECT id, nickname, age, gendar.name AS gender, image, in_kitakyushu, email, password FROM users INNER JOIN genders ON users.gender_id = gender_id WHERE id =$2", userId)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var users []GetUsersResponse

	for rows.Next() {
		var user GetUsersResponse
		err := rows.Scan(&user.Id, &user.Nickname, &user.Age, &user.Gender, &user.Image, &user.InKitakyushu, &user.Email)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		users = append(users, user)
	}

	c.JSON(200, users)
}

func (s *EndpointHandler) GetMe(c *gin.Context) {
}

func (s *EndpointHandler) GetArticles(c *gin.Context) {
}

func (s *EndpointHandler) GetArticle(c *gin.Context, articleId ArticleId) {
}
