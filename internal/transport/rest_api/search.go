package rest_api

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

func (s *Server) Search(resp http.ResponseWriter, req *http.Request) {
	var rq string
	switch req.Method {
	case http.MethodPost:
		rqB, err := io.ReadAll(req.Body)
		if err != nil {
			_, _ = resp.Write([]byte("invalid body"))
			resp.WriteHeader(http.StatusBadRequest)
		}
		rq = strings.ToLower(string(rqB))
	default:
		_, _ = resp.Write([]byte("invalid method"))
		resp.WriteHeader(http.StatusBadRequest)
		return
	}

	var out []PostsResponse
	for idx, item := range postsLower {
		if strings.Contains(item.Name, rq) ||
			strings.Contains(item.Content, rq) {
			out = append(out, posts[idx])
		}
	}
	postsB, err := json.Marshal(out)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = resp.Write(postsB)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}
}
