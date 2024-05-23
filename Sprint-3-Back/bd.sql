-- Tabla de Usuarios
CREATE TABLE Usuarios (
    id_usuario SERIAL PRIMARY KEY,
    nombre VARCHAR(255),
    correo_electronico VARCHAR(255),
    contrasena VARCHAR(255),
    saldo_ficticio DECIMAL(10, 2)
);

-- Tabla de Monedas
CREATE TABLE Monedas (
    id_moneda SERIAL PRIMARY KEY,
    nombre VARCHAR(255),
    simbolo VARCHAR(10)
);

-- Tabla de Transacciones
CREATE TABLE Transacciones (
    id_transaccion SERIAL PRIMARY KEY,
    id_usuario INTEGER,
    id_moneda INTEGER,
    tipo VARCHAR(10),
    cantidad DECIMAL(10, 2),
    precio_por_unidad DECIMAL(10, 2),
    fecha_transaccion DATE,
    FOREIGN KEY (id_usuario) REFERENCES Usuarios(id_usuario),
    FOREIGN KEY (id_moneda) REFERENCES Monedas(id_moneda)
);

-- Tabla de Precios Actuales
CREATE TABLE Precios_Actuales (
    id_precio SERIAL PRIMARY KEY,
    id_moneda INTEGER,
    precio_actual DECIMAL(10, 2),
    fecha_actualizacion DATE,
    FOREIGN KEY (id_moneda) REFERENCES Monedas(id_moneda)
);