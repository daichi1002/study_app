package usecase

import (
	"backend/domain/model"
	"backend/domain/repository"
	"backend/util"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

var logger = util.NewLogger()

type ArticleUsecase struct {
	articleRepository repository.ArticleRepository
}

func NewArticleUsecase(repo repository.ArticleRepository) *ArticleUsecase {
	return &ArticleUsecase{
		articleRepository: repo,
	}
}

func (u *ArticleUsecase) GetArticles(c *gin.Context) {
	articles, err := u.articleRepository.ListArticles()

	if err != nil {
		logger.Error(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, articles)
}

func (u *ArticleUsecase) CreateArticle(c *gin.Context) {
	article := &model.Article{}

	err := c.ShouldBindWith(article, binding.JSON)

	article.Id = util.GetUlid()
	if err != nil {
		logger.Error(err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	insertErr := u.articleRepository.CreateArticle(article)

	if insertErr != nil {
		c.JSON(http.StatusInternalServerError, insertErr.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (u *ArticleUsecase) GetArticle(c *gin.Context) {
	id := c.Params.ByName("id")
	article, err := u.articleRepository.ShowArticle(id)

	if err != nil {
		logger.Error(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, article)
}

func (u *ArticleUsecase) UpdateArticle(c *gin.Context) {
	article := &model.Article{}
	err := c.ShouldBindWith(article, binding.JSON)

	if err != nil {
		logger.Error(err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	insertErr := u.articleRepository.UpdateArticle(article)

	if insertErr != nil {
		logger.Error(insertErr)
		c.JSON(http.StatusInternalServerError, insertErr.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (u *ArticleUsecase) DeleteArticle(c *gin.Context) {
	id := c.Params.ByName("id")
	err := u.articleRepository.DeleteArticle(id)

	if err != nil {
		logger.Error(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (u *ArticleUsecase) GetRandomArticle(c *gin.Context) {
	articles, err := u.articleRepository.ListArticles()

	if err != nil {
		logger.Error(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	// 記事の配列からランダムで生地を1つ選出する
	index := rand.Intn(len(articles))
	result := articles[index]

	c.JSON(http.StatusOK, result)
}
