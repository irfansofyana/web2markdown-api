package handler

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/go-shiori/go-readability"

	md "github.com/JohannesKaufmann/html-to-markdown"
)

// ConversionRequest struct
type ConversionRequest struct {
	URL string `json:"url"`
}

// ConversionResult struct
type ConversionResult struct {
	URL    string `json:"url,omitempty"`
	Title  string `json:"title,omitempty"`
	Domain string `json:"domain,omitempty"`
	Data   string `json:"data,omitempty"`
	Error  error  `json:"error,omitempty"`
}

// Convert handler function to convert any web content from a certain URL to markdown
func Convert(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "/api/convert only support POST method", http.StatusMethodNotAllowed)
		return
	}

	var cReq ConversionRequest
	if err := json.NewDecoder(r.Body).Decode(&cReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	markdown, title, domain, err := getWebContent(cReq.URL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := ConversionResult{URL: cReq.URL, Domain: domain, Data: markdown, Title: title, Error: nil}
	writeResponse(w, http.StatusOK, resp)
}

func getWebContent(webURL string) (markdown string, title string, domain string, err error) {
	u, err := url.Parse(webURL)
	if err != nil {
		return
	}

	domain = u.Hostname()

	resp, err := http.Get(webURL)
	if err != nil {
		return
	}

	article, err := readability.FromReader(resp.Body, u)
	if err != nil {
		return
	}

	title = article.Title

	converter := md.NewConverter("", true, nil)
	markdownStr, err := converter.ConvertString(article.Content)
	if err != nil {
		return
	}

	markdown = markdownStr
	return
}

func writeResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
