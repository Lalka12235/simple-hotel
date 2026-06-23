CREATE TABLE IF NOT EXISTS room_categories (
    id_room_categories SERIAL PRIMARY KEY,
    class_name VARCHAR(100) NOT NULL
);

CREATE TABLE IF NOT EXISTS rooms (
    id_room SERIAL PRIMARY KEY,
    id_room_categories INT REFERENCES room_categories(id_room_categories) ON DELETE CASCADE,
    capacity INT NOT NULL,
    price INT NOT NULL
);

CREATE TABLE IF NOT EXISTS clients (
    id_client SERIAL PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    surname VARCHAR(100),
    address TEXT,
    passport VARCHAR(50) NOT NULL,
    coment TEXT
);

CREATE TABLE IF NOT EXISTS bookings (
    id_booking SERIAL PRIMARY KEY,
    id_client INT REFERENCES clients(id_client) ON DELETE CASCADE,
    id_room INT REFERENCES rooms(id_room) ON DELETE CASCADE,
    id_room_categories INT REFERENCES room_categories(id_room_categories) ON DELETE CASCADE,
    check_in_time TIMESTAMP NOT NULL,
    check_out_time TIMESTAMP NOT NULL
);