package rest_api

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"net/http"
	"sort"
)

func (s *Server) Posts(resp http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		s.getPosts(resp, req)
		return
	default:
		_, _ = resp.Write([]byte("invalid method"))
		resp.WriteHeader(http.StatusBadRequest)
		return
	}
}

//go:embed posts/posts.json
var postsBytes []byte

//go:embed posts/graph.json
var graphBytes []byte

var postsLower []PostsResponse
var posts []PostsResponse
var graph GraphResponse

func init() {
	err := json.Unmarshal(postsBytes, &posts)
	if err != nil {
		panic(err)
	}

	sort.Slice(posts, func(i, j int) bool {
		return posts[i].Name < posts[j].Name
	})

	postsBytes, err = json.Marshal(posts)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(bytes.ToLower(postsBytes), &postsLower)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(graphBytes, &graph)
	if err != nil {
		panic(err)
	}

}

func (s *Server) getPosts(resp http.ResponseWriter, _ *http.Request) {
	_, err := resp.Write(postsBytes)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
	}
}

type PostsResponse struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}
