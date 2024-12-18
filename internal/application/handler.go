package application

import (
	"encoding/json"
	"net/http"
)

// CalcHandler обрабатывает HTTP-запросы.
type CalcHandler struct {
	Service *CalcService
}

// NewCalcHandler создает новый экземпляр CalcHandler.
func NewCalcHandler(service *CalcService) *CalcHandler {
	return &CalcHandler{Service: service}
}

// Req представляет структуру запроса.
type Req struct {
	Expression string `json:"expression"`
}

// HandlerCalc обрабатывает запросы на вычисление выражений.
func (c *CalcHandler) HandlerCalc(w http.ResponseWriter, r *http.Request) {
	var req Req
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	result, err := c.Service.Calculate(req.Expression)
	if err != nil {
		switch err.Error() {
		case "invalid expression":
			http.Error(w, "invalid expression", http.StatusUnprocessableEntity)
		default:
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}
		return
	}

	response := map[string]float64{"result": result}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
