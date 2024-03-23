package controller

import (
	"log"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
)

type mockHandler struct {
	ctx *gin.Context
}

func (s *mockHandler) Login(c *gin.Context) {
	log.Println("Call Login")
	s.ctx = c
}

func (s *mockHandler) Logout(c *gin.Context) {
	log.Println("Call Logout")
	s.ctx = c
}

func (s *mockHandler) Register(c *gin.Context) {
	log.Println("Call Register")
	s.ctx = c
}

func (s *mockHandler) GetUsers(c *gin.Context) {
	log.Println("Call GetUsers")
	s.ctx = c
}

func (s *mockHandler) GetUser(c *gin.Context, userId UserId) {
	log.Println("Call GetUser")
	s.ctx = c
}

func (s *mockHandler) GetArticles(c *gin.Context) {
	log.Println("Call GetArticles")
	s.ctx = c
}

func (s *mockHandler) GetArticle(c *gin.Context, articleId ArticleId) {
	log.Println("Call GetArticle")
	s.ctx = c
}

func TestStartServer(t *testing.T) {
	testHandler := &mockHandler{}

	go StartServer(testHandler)

	tests := []struct {
		name   string
		method string
		path   string
	}{
		{
			name:   "Login",
			method: http.MethodPost,
			path:   "/auth/v1/login",
		},
		{
			name:   "Logout",
			method: http.MethodPost,
			path:   "/auth/v1/logout",
		},
		{
			name:   "Register",
			method: http.MethodPost,
			path:   "/auth/v1/register",
		},
		{
			name:   "GetUser",
			method: http.MethodGet,
			path:   "/api/v1/users/1",
		},
		{
			name:   "GetUsers",
			method: http.MethodGet,
			path:   "/api/v1/users",
		},
		{
			name:   "GetArticle",
			method: http.MethodGet,
			path:   "/api/v1/articles/1",
		},
		{
			name:   "GetArticles",
			method: http.MethodGet,
			path:   "/api/v1/articles",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(tt.method, "http://localhost:8080"+tt.path, nil)
			if err != nil {
				t.Fatal(err)
			}

			client := http.Client{}
			_, err = client.Do(req)
			if err != nil {
				t.Fatal(err)
			}

			calledPath := testHandler.ctx.Request.URL.Path

			log.Println("***********************")
			log.Println(calledPath)
			log.Println(tt.path)
			log.Println("***********************")

			if calledPath != tt.path {
				t.Errorf("Handler not called for %s", tt.name)
			}
		})
	}

}
