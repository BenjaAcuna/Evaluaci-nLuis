package main

type Producto struct {
	IdProducto     int64  `json:"idProducto"`
	NombreProducto string `json:"nombreProducto"`
	TipoProducto   string `json:"tipoProducto"`
	PrecioProducto int64  `json:"precioProducto"`
	FechaProducto  string `json:"fechaProducto"`
}
