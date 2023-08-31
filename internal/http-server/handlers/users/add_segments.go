package users

import (
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"

	resp "segmentation_service/internal/lib/api/response"
)

type Request struct {
	SegmentsToAdd []string `json:"segmentsToAdd"`
	SegmentsToDelete []string `json:"segmentsToDelete"`
	User_id int `json:"user_id"`
}

type Response struct {
	resp.Response
}

type UserSegmentsChanger interface {
	AddSegmentToUser(name string, user_id int) (error)
	RemoveSegmentFromUser (name string, user_id int) error
}

func AddSegment(log *slog.Logger, userSegmentChanger UserSegmentsChanger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		const op = "handlers.segments.create.New"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req Request
		err := render.DecodeJSON(r.Body, &req)
		if errors.Is(err, io.EOF) {
			log.Error("request body is empty")
			render.JSON(w, r, resp.Error("empty request"))
			return
		}
		if err != nil {
			log.Error("failed to decode request body", err)
			render.JSON(w, r, resp.Error("failed to decode request"))
			return
		}

		log.Info("request body decoded", slog.Any("request", req))

		segmentsToAdd := req.SegmentsToAdd
		if len(segmentsToAdd) != 0 {
			for _, name := range segmentsToAdd {
				err := userSegmentChanger.AddSegmentToUser(name, req.User_id)
				if err != nil {
					log.Error("failed to add segment", err)
					render.JSON(w, r, resp.Error("failed to add segment"))
					return
				}	
			}
		}

		segmentsToDelete := req.SegmentsToDelete
		if len(segmentsToAdd) != 0 {
			for _, name := range segmentsToDelete{
				err := userSegmentChanger.RemoveSegmentFromUser(name, req.User_id)
				if err != nil {
					log.Error("failed to add segment", err)
					render.JSON(w, r, resp.Error("failed to add segment"))
					return
				}	
			}
		}
	
	
		log.Info("segments modified")
		render.JSON(w, r, Response{
			Response: resp.OK()})
	}
}