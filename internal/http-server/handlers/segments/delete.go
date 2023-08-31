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

type SegmentRemover interface {
	DeleteSegment(segment_name string) (int64, error)
}

func Delete(log *slog.Logger, segmentRemover SegmentRemover) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		const op = "handlers.segments.delete"

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

		name := req.Name

		if name == "" {
			log.Error("empty segment name")
			render.JSON(w, r, resp.Error("empty segment name"))
			return
		}
		id, err := segmentRemover.DeleteSegment(name)
		if err != nil {
			log.Error("failed to create segment", err)
			render.JSON(w, r, resp.Error("failed to  to create segment"))
			return
		}

		log.Info("segment added", id)

		render.JSON(w, r, Response{
			Response: resp.OK(),
		})
	}
}

