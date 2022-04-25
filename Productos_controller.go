package main

func createProductos(productos Producto) error {
	bd, err := getDB()
	if err != nil {
		return err
	}
	_, err = bd.Exec("INSERT INTO productos (nombreProducto, tipoProducto, precioProducto, fechaProducto) VALUES (?, ?, ?, ?)", productos.NombreProducto, productos.TipoProducto, productos.PrecioProducto, productos.FechaProducto)
	return err
}

func deleteProductos(id int64) error {

	bd, err := getDB()
	if err != nil {
		return err
	}
	_, err = bd.Exec("DELETE FROM productos WHERE idProducto = ?", id)
	return err
}

// It takes the ID to make the update
func updateProductos(productos Producto) error {
	bd, err := getDB()
	if err != nil {
		return err
	}
	_, err = bd.Exec("UPDATE productos SET precioProducto = ?, fechaProducto = ? WHERE idProducto = ?", productos.PrecioProducto, productos.FechaProducto, productos.IdProducto)
	return err
}
func getProductos() ([]Producto, error) {
	//Declare an array because if there's error, we return it empty
	productos := []Producto{}
	bd, err := getDB()
	if err != nil {
		return productos, err
	}
	// Get rows so we can iterate them
	rows, err := bd.Query("SELECT idProducto, nombreProducto, tipoProducto, precioProducto, fechaProducto FROM productos")
	if err != nil {
		return productos, err
	}
	// Iterate rows...
	for rows.Next() {
		// In each step, scan one row
		var producto Producto
		err = rows.Scan(&producto.IdProducto, &producto.NombreProducto, &producto.TipoProducto, &producto.PrecioProducto, &producto.FechaProducto)
		if err != nil {
			return productos, err
		}
		// and append it to the array
		productos = append(productos, producto)
	}
	return productos, nil
}

func getProductosById(id int64) (Producto, error) {
	var productos Producto
	bd, err := getDB()
	if err != nil {
		return productos, err
	}
	row := bd.QueryRow("SELECT idProducto, nombreProducto, tipoProducto, precioProducto, fechaProducto FROM Productos WHERE idProducto = ?", id)
	err = row.Scan(&productos.IdProducto, &productos.NombreProducto, &productos.TipoProducto, &productos.PrecioProducto, &productos.FechaProducto)
	if err != nil {
		return productos, err
	}
	// Success!
	return productos, nil
}
