package main

import (
	"context"
	"testing"

	"github.com/baronwithyou/go-practice/23redis/01chapter/model"
	"github.com/go-redis/redis/v8"
)

func TestChapter1(t *testing.T) {
	conn := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       15,
	})

	if _, err := conn.Ping(context.Background()).Result(); err != nil {
		t.Fatalf("Connect to redis client failed, err: %v\n", err)
	}

	repo := model.NewArticleRepo(conn)
	// repo.PostArticle("baron", "leave", "wanna go")

	// repo.ArticleVote("article:"+articleId, "baron")

	articles := repo.GetArticles(1, "score:")
	t.Log(articles)

	repo.AddRemoveGroups("7", []string{"new-group"}, []string{})

	groupArticles := repo.GetGroupArticles("new-group", "score:", 1)
	t.Log(groupArticles)
}
