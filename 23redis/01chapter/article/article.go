package article

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

	a.Conn.HSet(a.Ctx, "article:"+articleId, []map[string]string{
		"user":  user,
		"title": title,
		"link":  link,
	})
	return ""
}

func (a *ArticleRepo) GetArticles(page int64, order string) []map[string]string {

	return nil
}

func (a *ArticleRepo) AddRemoveGroups(articleId string, toAdd []string, toRemove []string) {

}

func (a *ArticleRepo) GetGroupArticles(group, order string, page int64) []map[string]string {
	return nil
}

func (a *ArticleRepo) Reset() {

}
