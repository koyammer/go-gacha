package handler

import (
	"go-echo-starter/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

type postsResponse struct {
	Posts []model.Post `json:"posts"`
}

func (h *Handler) getPosts(c echo.Context) error {
	posts, err := h.postStore.AllPostsA()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, postsResponse{Posts: posts})
}
