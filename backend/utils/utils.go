package utils

// import (
// 	"context"
// 	"encoding/json"
// 	"hex/common/constants"
// 	"hex/utils/logger"
// 	"net/http"
// )

// type response struct {
// 	Status  bool        `json:"success"`
// 	Message string      `json:"message"`
// 	Data    interface{} `json:"data"`
// }

// func newResponse(status bool, message string, data interface{}) *response {
// 	return &response{
// 		Status:  status,
// 		Message: message,
// 		Data:    data,
// 	}
// }

// func (resp *response) bytes() []byte {
// 	data, _ := json.Marshal(resp)
// 	return data
// }

// func (resp *response) string() string {
// 	return string(resp.bytes())
// }

// func (resp *response) sendResponse(w http.ResponseWriter, r *http.Request, statusCode int, log logger.Logger) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.WriteHeader(statusCode)
// 	_, err := w.Write(resp.bytes())
// 	i
// }

// // 200
// func StatusOK(w http.ResponseWriter, r *http.Request, message string, data interface{}, log logger.Logger) {
// 	newResponse(true, message, data).sendResponse(w, r, http.StatusOK, log)
// }

// // 204
// func StatusNoContent(w http.ResponseWriter, r *http.Request, message string, log logger.Logger) {
// 	newResponse(false, message, nil).sendResponse(w, r, http.StatusNoContent, log)
// }

// // 400
// func StatusBadRequest(w http.ResponseWriter, r *http.Request, message string, err error, log logger.Logger) {
// 	data := map[string]interface{}{"error": err.Error()}
// 	newResponse(false, message, data).sendResponse(w, r, http.StatusBadRequest, log)
// }

// // 401
// func StatusUnauthorized(w http.ResponseWriter, r *http.Request, message string, err error, log logger.Logger) {
// 	data := map[string]interface{}{"error": err.Error()}
// 	newResponse(false, message, data).sendResponse(w, r, http.StatusUnauthorized, log)
// }

// // 403
// func StatusForbidden(w http.ResponseWriter, r *http.Request, message string, err error, log logger.Logger) {
// 	data := map[string]interface{}{"error": err.Error()}
// 	newResponse(false, message, data).sendResponse(w, r, http.StatusForbidden, log)
// }

// // 404
// func StatusNotFound(w http.ResponseWriter, r *http.Request, message string, err error, log logger.Logger) {
// 	data := map[string]interface{}{"error": err.Error()}
// 	newResponse(false, message, data).sendResponse(w, r, http.StatusNotFound, log)
// }

// // 405
// func StatusMethodNotAllowed(w http.ResponseWriter, r *http.Request, message string, log logger.Logger) {
// 	newResponse(false, message, http.StatusMethodNotAllowed).sendResponse(w, r, http.StatusMethodNotAllowed, log)
// }

// // 409
// func StatusConflict(w http.ResponseWriter, r *http.Request, message string, err error, log logger.Logger) {
// 	data := map[string]interface{}{"error": err.Error()}
// 	newResponse(false, message, data).sendResponse(w, r, http.StatusConflict, log)
// }

// // 429
// func StatusTooManyRequests(w http.ResponseWriter, r *http.Request, message string, err error, log logger.Logger) {
// 	data := map[string]interface{}{"error": err.Error()}
// 	newResponse(false, message, data).sendResponse(w, r, http.StatusTooManyRequests, log)
// }

// // 500
// func StatusInternalServerError(w http.ResponseWriter, r *http.Request, message string, err error, log logger.Logger) {
// 	data := map[string]interface{}{"error": err.Error()}
// 	newResponse(false, message, data).sendResponse(w, r, http.StatusInternalServerError, log)
// }
