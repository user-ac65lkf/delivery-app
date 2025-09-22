package handlers

import (
	"context"
	"google.golang.org/grpc/metadata"
	"log"
	"net/http"
	"strings"
)

func getContext() context.Context {
	return context.Background()
}

func getToken(r *http.Request) string {
	prefix := "Bearer "
	authHeader := r.Header.Get("authorization")
	reqToken := strings.TrimPrefix(authHeader, prefix)

	return reqToken
}

func AddTokenToContext(ctx context.Context, token string) context.Context {
	md := metadata.Pairs("authorization", "Bearer "+token)
	ctx = metadata.NewOutgoingContext(ctx, md)

	return ctx
}

func handleToken(h *handler, w http.ResponseWriter, r *http.Request) (context.Context, int) {
	ctx := getContext()

	token := getToken(r)

	ctx = AddTokenToContext(ctx, token)

	courierId, err := h.serv.VerifyToken(ctx)
	if err != nil {
		log.Println(err)

		http.Error(w, err.Error(), http.StatusBadRequest)

		return nil, 0
	}

	return ctx, courierId
}
