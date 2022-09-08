package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/filbertkm/wikiclient/mwclient"
	"github.com/gorilla/mux"
)

type Description struct {
	l *log.Logger
}

// Description returned in the response
// swagger:response pageResponse
type pageResponse struct {
	Title string `json:"title"`
	Lang string `json:"lang"`
	Description string `json:"description"`
	Source string `json:"source"`
}

func NewDescription(l *log.Logger) *Description {
	return &Description{l}
}

// swagger:parameters title getDescription
type TitleParameterWrapper struct {
	// Wikipedia page title
	// in: path
	// required: true
	// example: [[["New_York_City"]]]
	Title string `json:"title"`
}

// swagger:parameters lang getDescription
type LangParameterWrapper struct {
	// Wikipedia site lang code
	// in: query
	// required: false
	// example: [[["en"]]]
	Lang string `json:"lang"`
}

// swagger:parameters fallback getDescription
type FallbackParameterWrapper struct {
	// Wikidata description fallback option, If false, then use local description only
	// in: query
	// required: false
	// example: [[[true]]]
	Lang string `json:"fallback"`
}

// swagger:route GET /page/{title} description getDescription
// Returns a short description for a Wikipedia page
// responses:
//	200: pageResponse
// 	400: Unexpected error
//	404: Page not found

// GetDescription returns wikipedia description for language (project)
func (d *Description) GetDescription(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	title := vars["title"]
	lang := r.URL.Query().Get("lang")
	fallback := r.URL.Query().Get("fallback")
	useFallback, _ := strconv.ParseBool(fallback)

	if lang == "" {
		lang = "en"
	}

	client := mwclient.NewClient(lang)
	pages, _ := client.GetPages(title)

	var pageList = []*mwclient.Page {}
	for _, page := range pages {
		if page.Source != "local" && !useFallback {
			http.Error(w, "Description not found", http.StatusNotFound)
			return
		}

		if !page.Missing {
			pageList = append(pageList, &page)
		}
	}

	if len(pageList) == 0 {
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}

	response := &pageResponse{
		Title: title,
		Lang: lang,
		Description: pageList[0].Description,
		Source: pageList[0].Source,
	}

	responseData, err := json.Marshal(response)

	if err != nil {
		http.Error(w, "Unexpected error", http.StatusBadRequest)
		return
	}

	w.Write(responseData)
}