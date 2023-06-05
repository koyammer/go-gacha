package handler

import "go-echo-starter/store"

type Handler struct {
	authStore store.AuthStore
	userStore store.UserStore
	postStore store.PostStore
}

func NewHandler(
	as store.AuthStore,
	us store.UserStore,
	ps store.PostStore,
) *Handler {
	return &Handler{
		authStore: as,
		userStore: us,
		postStore: ps,
	}
}
