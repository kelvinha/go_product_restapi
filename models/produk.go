package models

type (
	Produk struct {
		ProdukId          int    `form:"product_id" json:"product_id"`
		ProdukName        string `form:"product_name" json:"product_name"`
		ProdukPrice       int    `form:"product_price" json:"product_price"`
		ProdukDescription string `form:"product_description" json:"product_description"`
		ProdukQuantity    int    `form:"product_quantity" json:"product_quantity"`
	}
)