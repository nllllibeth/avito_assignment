package segments

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
	Name string `json:"name"`
}

type Response struct {
	resp.Response
	Name string `json:"name,omitempty"`
}

type SegmentCreator interface {
	CreateSegment(segment_name string) (int64, error)
}

func Create(log *slog.Logger, segmentCreator SegmentCreator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		const op = "handlers.segments.create.New"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req Request
		err := render.DecodeJSON(r.Body, &req)
		if errors.Is(err, io.EOF) {
			// Такую ошибку встретим, если получили запрос с пустым телом.
			// Обработаем её отдельно
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

		name := req.Name

		if name == "" {
			log.Error("empty segment name")
			render.JSON(w, r, resp.Error("empty segment name"))
			return
		}
		id, err := SegmentCreator.CreateSegment(segmentCreator, name)
		if err != nil {
			log.Error("failed to create segment", err)
			render.JSON(w, r, resp.Error("failed to  to create segment"))
			return
		}

		log.Info("segment added", id)

		responseOK(w, r, name)

	}
}

func responseOK(w http.ResponseWriter, r *http.Request, name string) {
	render.JSON(w, r, Response{
		Response: resp.OK(),
		Name:     name,
	})
}

