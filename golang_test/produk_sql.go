package golang_test

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"kelpin_golang/config"
	"kelpin_golang/models"
)

func GetData(ctx context.Context, sortingBy string) ([]models.Produk, error) {
	var produks []models.Produk

	db, err := config.ConnectDB()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}
	queryText := ""
	switch sortingBy {
	case "ASC":
		queryText = fmt.Sprintf("SELECT * FROM table_product Order By product_name ASC")
	case "DESC":
		queryText = fmt.Sprintf("SELECT * FROM table_product Order By product_name DESC")
	case "high-price":
		queryText = fmt.Sprintf("SELECT * FROM table_product Order By product_price DESC")
	case "low-price":
		queryText = fmt.Sprintf("SELECT * FROM table_product Order By product_price ASC")
	default:
		queryText = fmt.Sprintf("SELECT * FROM table_product Order By product_id ASC")
	}

	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var produk models.Produk

		if err = rowQuery.Scan(&produk.ProdukId,
			&produk.ProdukName,
			&produk.ProdukPrice,
			&produk.ProdukDescription,
			&produk.ProdukQuantity); err != nil {
			return nil, err
		}

		if err != nil {
			log.Fatal(err)
		}

		if err != nil {
			log.Fatal(err)
		}

		produks = append(produks, produk)
	}

	return produks, nil

}

func CreateData(ctx context.Context, produk models.Produk) error {
	db, err := config.ConnectDB()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("INSERT INTO table_product (product_name, product_price, product_description, product_quantity) values('%v',%v,'%v',%v)",
		produk.ProdukName,
		produk.ProdukPrice,
		produk.ProdukDescription,
		produk.ProdukQuantity)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}
	return nil
}

func UpdateData(ctx context.Context, produk models.Produk) error {
	db, err := config.ConnectDB()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("UPDATE table_product set product_name = '%s', product_price = '%d', product_description = '%s', product_quantity = '%d' where product_id = %d",
		produk.ProdukName,
		produk.ProdukPrice,
		produk.ProdukDescription,
		produk.ProdukQuantity,
		produk.ProdukId)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}
	return nil
}

func DeleteData(ctx context.Context, produk models.Produk) error {
	db, err := config.ConnectDB()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("DELETE FROM table_product where product_id = %d", produk.ProdukId)

	s, err := db.ExecContext(ctx, queryText)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	check, err := s.RowsAffected()
	fmt.Println(check)
	if check == 0 {
		return errors.New("Produk tidak tersedia")
	}

	return nil
}
