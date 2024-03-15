package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	// replace with your actual package path
)

type PaymentMethod struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	IsActive bool   `json:"isActive"`
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/payment-methods", func(r chi.Router) {
		r.Get("/", HandleGetPaymentMethods)
		r.Post("/", HandlePostPaymentMethod)
		r.Get("/{id}/edit", HandleEditPaymentMethod)
		r.Post("/{id}", HandlePutPaymentMethod)
	})

	http.ListenAndServe(":3001", r)
}

var paymentMethods = []PaymentMethod{
	{Id: "1", Name: "Credit Card", IsActive: true},
	{Id: "2", Name: "PayPal", IsActive: true},
	// Add more payment methods here
}

func HandlePostPaymentMethod(w http.ResponseWriter, r *http.Request) {
	var newPaymentMethod PaymentMethod
	err := json.NewDecoder(r.Body).Decode(&newPaymentMethod)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	paymentMethods = append(paymentMethods, newPaymentMethod)
	w.WriteHeader(http.StatusCreated)
}
func HandleEditPaymentMethod(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	for _, pm :=  range paymentMethods {
		if pm.Id == id {
			page := form(id)
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			page.Render(r.Context(), w)
		}
	}

	var newPaymentMethod PaymentMethod
	err := json.NewDecoder(r.Body).Decode(&newPaymentMethod)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	paymentMethods = append(paymentMethods, newPaymentMethod)
	w.WriteHeader(http.StatusCreated)
}
func HandlePutPaymentMethod(w http.ResponseWriter, r *http.Request) {
	fmt.Println("starting...")
	id := chi.URLParam(r, "id")

	paymentMethodName := r.PostFormValue("paymentMethodName")

	// type requestBody struct {
	// 	PaymentMethodName string `json:"paymentMethodName"`
	// }

	// var body requestBody

	fmt.Println("i am at the decoder")
	// err := json.NewDecoder(r.Body).Decode(&body)
	fmt.Println("decoding...")
	// if err != nil {
	// 	fmt.Println("getting error:", err)
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	fmt.Println("i am after the decoder")

	fmt.Println("pm name", paymentMethodName)

	var updatedPaymentMethod PaymentMethod
	for _, pm :=  range paymentMethods {
		if pm.Id == id {
			pm.Name = paymentMethodName
			updatedPaymentMethod = pm
		}
	}




	mi := methodItem(updatedPaymentMethod)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	mi.Render(r.Context(), w)
}

func HandleGetPaymentMethods(w http.ResponseWriter, r *http.Request) {
	pms := Page(paymentMethods)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	pms.Render(r.Context(), w)
}
