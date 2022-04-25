CREATE TABLE Productos
(
    idProducto  int unsigned not null primary key auto_increment,
    nombreProducto  VARCHAR(45)    NOT NULL,
    tipoProducto VARCHAR(45)    NOT NULL,
    precioProducto INTEGER    NOT NULL,
    fechaProducto  VARCHAR(45)    NOT NULL
);