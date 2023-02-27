package handler

import (
	"awesomeProject-L0/internal/model"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"
)

// getOrderById обработка запроса на получение order
func (h *Handler) getOrderById(w http.ResponseWriter, r *http.Request) {

	orderId, _ := strconv.Atoi(mux.Vars(r)["id"])

	order, err := h.service.GetOrderById(orderId)
	if err != nil {
		openTemplate("order not found", w)
		return
	}
	openTemplate(order, w)

}

func (h *Handler) startPage(w http.ResponseWriter, r *http.Request) {

	openTemplate("", w)
}

func openTemplate(data any, w http.ResponseWriter) {

	var page, name string

	switch data.(type) {
	case string:
		page = "index.html"
		name = "msg"
	case model.Order:
		page = "order.html"
		name = "order"
	}

	path := filepath.Join("ui", page)
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	err = tmpl.ExecuteTemplate(w, name, data)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

}
