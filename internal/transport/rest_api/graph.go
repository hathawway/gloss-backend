package rest_api

import (
	"encoding/json"
	"net/http"
)

func (s *Server) Graph(resp http.ResponseWriter, req *http.Request) {
	var out GraphResponse

	switch req.Method {
	case http.MethodGet:
		out = graph
	default:
		_, _ = resp.Write([]byte("invalid method"))
		resp.WriteHeader(http.StatusBadRequest)
		return
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

type GraphResponse struct {
	Nodes   map[string]GraphNode  `json:"nodes"`
	Edges   map[string]GraphEdges `json:"edges"`
	Layouts GraphLayouts          `json:"layouts"`
}

type GraphNode struct {
	Name string `json:"name"`
}

type GraphEdges struct {
	Source string `json:"source"`
	Target string `json:"target"`
}
type GraphLayouts struct {
	Nodes map[string]GraphLayout `json:"nodes"`
}
type GraphLayout struct {
	X int `json:"x"`
	Y int `json:"y"`
}
