package controller

import "github.com/gin-gonic/gin"

type EndpointHandler struct{}

func (s *EndpointHandler) Login(c *gin.Context) {
}

func (s *EndpointHandler) Logout(c *gin.Context) {
}

func (s *EndpointHandler) Register(c *gin.Context) {
}

func (s *EndpointHandler) GetUser(c *gin.Context, userId UserId) {
}

func (s *EndpointHandler) GetUsers(c *gin.Context) {
}

func (s *EndpointHandler) GetArticles(c *gin.Context) {
}

func (s *EndpointHandler) GetArticle(c *gin.Context, articleId ArticleId) {
}
