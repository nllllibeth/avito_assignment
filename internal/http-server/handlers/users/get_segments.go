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

type Req struct {
	User_id int `json:"user_id"`
}

type Res struct {
	resp.Response
	Segments []string `json:"segments"`
}

type UserSegmentsGrabber interface {
	GetActiveSegments(user_id int) ([]string, error)
}

func GetSegments(log *slog.Logger, userSegmentGrabber UserSegmentsGrabber) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		const op = "handlers.users.get"

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

		user_id := req.User_id

		activeSegments, err := userSegmentGrabber.GetActiveSegments(user_id)
		if err != nil {
			log.Error("failed to get segments", err)
			render.JSON(w, r, resp.Error("failed to get segments"))
			return
		}
		log.Info("segments were succesfully getted")
		render.JSON(w, r, Res{
			Response: resp.OK(),
			Segments: activeSegments})
	}
}