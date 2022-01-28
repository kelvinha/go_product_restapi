package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"kelpin_golang/golang_test"
	"kelpin_golang/models"
	"kelpin_golang/utilities"
)

func main() {
	http.HandleFunc("/produk", GetProduk)
	http.HandleFunc("/produk/create", PostProduk)
	http.HandleFunc("/produk/update", UpdateProduk)
	http.HandleFunc("/produk/delete", DeleteProduk)
	fmt.Println("starting web server at http://localhost:8080/")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}
}

func GetProduk(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		sortingBy := r.URL.Query().Get("sortingBy")

		produk, err := golang_test.GetData(ctx, sortingBy)

		if err != nil {
			http.Error(w, "Error", http.StatusMethodNotAllowed)
			return
		}

		utilities.ResponseJSON(w, produk, http.StatusOK)
		return

	}
	http.Error(w, "Method must be GET", http.StatusMethodNotAllowed)
	return
}

func PostProduk(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var produk models.Produk

		if err := json.NewDecoder(r.Body).Decode(&produk); err != nil {
			utilities.ResponseJSON(w, err, http.StatusBadRequest)
			return
		}

		if err := golang_test.CreateData(ctx, produk); err != nil {
			utilities.ResponseJSON(w, err, http.StatusInternalServerError)
			return
		}

		result := map[string]string{
			"status":  strconv.Itoa(http.StatusCreated),
			"message": "Success!",
		}

		utilities.ResponseJSON(w, result, http.StatusCreated)
		return
	}

	http.Error(w, "Method must be POST", http.StatusMethodNotAllowed)
	return
}

func UpdateProduk(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var produk models.Produk

		product_id := r.URL.Query().Get("product_id")

		if product_id == "" {
			utilities.ResponseJSON(w, "product_id tidak boleh kosong", http.StatusBadRequest)
			return
		}

		produk.ProdukId, _ = strconv.Atoi(product_id)

		if err := json.NewDecoder(r.Body).Decode(&produk); err != nil {
			utilities.ResponseJSON(w, err, http.StatusBadRequest)
			return
		}

		if err := golang_test.UpdateData(ctx, produk); err != nil {
			utilities.ResponseJSON(w, err, http.StatusInternalServerError)
			return
		}

		result := map[string]string{
			"status":  strconv.Itoa(http.StatusCreated),
			"message": "Success!",
		}

		utilities.ResponseJSON(w, result, http.StatusCreated)
		return
	}

	http.Error(w, "Method must be PUT", http.StatusMethodNotAllowed)
	return
}

func DeleteProduk(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var produk models.Produk

		product_id := r.URL.Query().Get("product_id")

		if product_id == "" {
			utilities.ResponseJSON(w, "product_id tidak boleh kosong", http.StatusBadRequest)
			return
		}
		produk.ProdukId, _ = strconv.Atoi(product_id)

		if err := golang_test.DeleteData(ctx, produk); err != nil {
			utilities.ResponseJSON(w, err, http.StatusInternalServerError)
			return
		}

		result := map[string]string{
			"status":  strconv.Itoa(http.StatusOK),
			"message": "Success!",
		}

		utilities.ResponseJSON(w, result, http.StatusOK)
		return
	}

	http.Error(w, "Method must be DELETE", http.StatusMethodNotAllowed)
	return
}
