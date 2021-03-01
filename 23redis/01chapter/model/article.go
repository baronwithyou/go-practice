package model

import (
	"context"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

type Article interface {
	ArticleVote(string, string)
	PostArticle(string, string, string) string
	GetArticles(int64, string) []map[string]string
	AddRemoveGroups(string, []string, []string)
	GetGroupArticles(string, string, int64) []map[string]string
	Reset()
}

const (
	OneWeekInSeconds       = 7 * 86400
	VoteScore              = 432
	ArticlesPerPage  int64 = 25
)

type ArticleRepo struct {
	Conn *redis.Client
	Ctx  context.Context
}

func NewArticleRepo(conn *redis.Client) *ArticleRepo {
	return &ArticleRepo{Conn: conn, Ctx: context.Background()}
}

func (a *ArticleRepo) ArticleVote(article, user string) {
	// 判断该文章是否已经过期
	cutoff := time.Now().Unix() - OneWeekInSeconds
	if a.Conn.ZScore(a.Ctx, "time:", article).Val() < float64(cutoff) {
		return
	}

	// 判断该用户是否已经投过票
	if a.Conn.SAdd(a.Ctx, "voted:"+article, user).Val() == 1 {
		// 上升排名
		a.Conn.ZIncrBy(a.Ctx, "score:", float64(VoteScore), article)
		// 文章信息中的投票数目上升
		a.Conn.HIncrBy(a.Ctx, article, "votes", 1)
	}
}

func (a *ArticleRepo) PostArticle(user, title, link string) string {
	articleId := strconv.Itoa(int(a.Conn.Incr(a.Ctx, "article:").Val()))

	// 自己给自己投票，并且给投票set设置一个过期时间
	voted := "voted:" + articleId
	a.Conn.SAdd(a.Ctx, voted, user)
	a.Conn.Expire(a.Ctx, voted, time.Second*OneWeekInSeconds)

	now := time.Now().Unix()

	a.Conn.HSet(a.Ctx, "article:"+articleId, map[string]interface{}{
		"user":  user,
		"title": title,
		"link":  link,
		"time":  now,
		"votes": 1,
	})

	// 初始化排名分数
	a.Conn.ZAdd(a.Ctx, "score:", &redis.Z{Score: float64(now + VoteScore), Member: articleId})
	// 初始化文章时间
	a.Conn.ZAdd(a.Ctx, "time:", &redis.Z{Score: float64(now), Member: articleId})

	return ""
}

func (a *ArticleRepo) GetArticles(page int64, order string) []map[string]string {
	if order == "" {
		order = "score:"
	}

	start := (page - 1) * ArticlesPerPage
	end := start + ArticlesPerPage - 1
	articles := []map[string]string{}

	ids := a.Conn.ZRevRange(a.Ctx, order, start, end).Val()
	for _, id := range ids {
		data := a.Conn.HGetAll(a.Ctx, "article:"+id).Val()
		data["id"] = id
		articles = append(articles, data)
	}

	return articles
}

func (a *ArticleRepo) AddRemoveGroups(articleId string, toAdd []string, toRemove []string) {

}

func (a *ArticleRepo) GetGroupArticles(group, order string, page int64) []map[string]string {
	return nil
}

func (a *ArticleRepo) Reset() {

}
