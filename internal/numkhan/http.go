package numkhan

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/gorm"
)

type HttpService struct {
	Service
	Router *chi.Mux
}

func NewHttpService(db *gorm.DB) *HttpService {
	s := &HttpService{
		Service{Db: db},
		chi.NewRouter(),
	}
	s.MountHandlers()
	return s
}

func (s *HttpService) MountHandlers() {
	s.Router.Use(middleware.Logger)

	s.Router.Get("/", s.GetCandidates)
}

func (h *HttpService) GetCandidates(w http.ResponseWriter, r *http.Request) {
	candidates := h.Candidates()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(candidates)
}
