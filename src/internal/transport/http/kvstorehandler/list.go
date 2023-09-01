package kvstorehandler

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/SenRecep/redisclone/src/internal/kverror"
)

func (h *kvstoreHandler) List(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.JSON(
			w,
			http.StatusMethodNotAllowed,
			map[string]string{"error": "method " + r.Method + " not allowed"},
		)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), h.CancelTimeout)
	defer cancel()

	serviceResponse, err := h.service.List(ctx)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			h.JSON(
				w,
				http.StatusGatewayTimeout,
				map[string]string{"error": err.Error()},
			)
			return
		}

		var kvErr *kverror.Error
		if errors.As(err, &kvErr) {
			clientMessage := kvErr.Message
			if kvErr.Data != nil {
				data, ok := kvErr.Data.(string)
				if ok {
					clientMessage = clientMessage + ", " + data
				}
			}

			if kvErr.Loggable {
				h.Logger.Error("kvstorehandler List service.List", "err", clientMessage)
			}
		}

		h.JSON(
			w,
			http.StatusInternalServerError,
			map[string]string{"error": err.Error()},
		)
		return
	}

	var handlerResponse ListResponse
	for _, item := range *serviceResponse {
		handlerResponse = append(handlerResponse, ItemResponse{
			Key:   item.Key,
			Value: item.Value,
		})
	}

	if len(handlerResponse) == 0 {
		h.JSON(
			w,
			http.StatusNotFound,
			map[string]string{"error": "nothing found"},
		)
		return
	}
	q := r.URL.Query()
	if len(q) != 0 && q.Has("order_by") {
		qv := q.Get("order_by")
		if len(qv) == 0 {
			goto OrderBy
		}
		isAsc := qv[0] == '-'
		orderFieldIsKey := strings.HasSuffix(qv, "key")
		handlerResponse.Order(isAsc, orderFieldIsKey)
	}
OrderBy:

	h.JSON(
		w,
		http.StatusOK,
		handlerResponse,
	)
}
