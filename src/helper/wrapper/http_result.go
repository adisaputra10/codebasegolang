package wrapper

import (
	"encoding/json"
	"net/http"

	"gitlab.playcourt.id/new-mypertamina/myptm-external-insurance-service/src/helper/errors"
)

// HTTPResponseJSON is a function
func HTTPResponseJSON(w http.ResponseWriter, result *Result, code int, message string) {
	httpResult := new(HTTPResult)
	if result.Err != nil {
		httpResult.Message = result.Message
		httpResult.Data = nil
		httpResult.Code = errors.GetHTTPStatusCodeByError(result.Err)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(httpResult.Code)
		json.NewEncoder(w).Encode(httpResult)
		return
	}

	httpResult.Success = true
	httpResult.Data = result.Data
	httpResult.Message = message
	httpResult.Code = code
	httpResult.Meta = result.PartialMeta

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpResult.Code)
	json.NewEncoder(w).Encode(httpResult)
	return
}
