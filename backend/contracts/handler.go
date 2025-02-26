package contracts

import (
	"backend/api"
	"backend/logger"
	"encoding/json"
	"net/http"
)

func HandlePurchaseToken(service TokenService) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		var request PurchaseTokenRequest
		ctx := req.Context()

		err := json.NewDecoder(req.Body).Decode(&request)
		if err != nil {
			logger.Error(ctx, "failed to decode request body", "error", err.Error())
			api.RespondWithError(rw, http.StatusBadRequest, api.Response{
				Error: "Invalid request format",
			})
			return
		}

		err = service.PurchaseToken(ctx, request)
		if err != nil {
			logger.Error(ctx, "failed to purchase token", "error", err.Error())
			api.RespondWithError(rw, http.StatusBadRequest, api.Response{
				Error: "Token purchase failed",
			})
			return
		}

		api.RespondWithJSON(rw, http.StatusOK, api.Response{
			Message: "success",
		})

	}
}
