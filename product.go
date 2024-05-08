package main

func createProduct(product *Product) error {
	_, err := db.Exec("INSERT INTO public.products(name, price) VALUES ($1, $2);", product.Name, product.Price)
	return err
}

func getProduct(id int) (Product, error) {
	var p Product
	row := db.QueryRow(`SELECT id, name, price FROM products WHERE id = $1;`, id)
	err := row.Scan(&p.Id, &p.Name, &p.Price)
	if err != nil {
		return Product{}, err
	}
	return p, nil
}

func getAllProducts() ([]Product, error) {
	var p []Product
	rows, err := db.Query(`SELECT id, name, price FROM products`)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var productItem Product
		err := rows.Scan(&productItem.Id, &productItem.Name, &productItem.Price)
		if err != nil {
			return nil, err
		}
		p = append(p, productItem)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return p, nil
}

func updateProduct(id int, product *Product) (Product, error) {
	// exec รันคำสั่งแต่ไม่คืนค่าอะไรออกมา
	// QueryRow รันคำสั่ง และคืนค่าบางอย่างออกมา
	var p Product
	row := db.QueryRow("UPDATE public.products SET name=$1, price=$2 WHERE id=$3 RETURNING id, name, price;", product.Name, product.Price, id)
	err := row.Scan(&p.Id, &p.Name, &p.Price)
	if err != nil {
		return Product{}, err
	}
	return p, err
}

func deleteProduct(id int) error {
	// exec รันคำสั่งแต่ไม่คืนค่าอะไรออกมา
	// QueryRow รันคำสั่ง และคืนค่าบางอย่างออกมา
	_, err := db.Exec("DELETE FROM public.products WHERE id=$1", id)
	return err
}
