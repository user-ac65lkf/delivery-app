BEGIN;

CREATE TABLE IF NOT EXISTS "products"
(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    count INT NOT NULL
);

INSERT INTO products (name, description, price, count) VALUES
        ('Смартфон Samsung Galaxy S121', 'Современный смартфон с высокой производительностью.', 799.99, 50),
        ('Ноутбук HP Envy 13', 'Легкий и мощный ноутбук для работы и развлечений.', 1099.99, 30),
        ('Фотокамера Canon EOS 90D', 'Профессиональная зеркальная фотокамера с 32 Мп сенсором.', 1299.99, 20),
        ('Телевизор LG OLED65CXPUA', '65-дюймовый OLED-телевизор с 4K разрешением и HDR.', 1799.99, 10),
        ('Наушники Sony WH-1000XM4', 'Беспроводные наушники с шумоподавлением.', 349.99, 40),
        ('Электрический чайник Bosch TWK7203', 'Мощный чайник с регулируемой температурой.', 59.99, 100),
        ('Кофеварка Philips 3200 Series', 'Автоматическая кофеварка с функцией капучино.', 349.99, 30),
        ('Пылесос Dyson V11 Absolute', 'Беспроводной пылесос с высокой мощностью всасывания.', 499.99, 15),
        ('Умные часы Apple Watch Series 6', 'Смарт-часы с измерением уровня кислорода в крови.', 399.99, 25),
        ('Игровая консоль PlayStation 5', 'Мощная игровая консоль нового поколения.', 499.99, 10),
        ('Холодильник Samsung RF28R7201SR', 'Большой холодильник с функцией автоматической очистки.', 1499.99, 5),
        ('Стиральная машина LG WM4000HWA', 'Стиральная машина с большой загрузкой и инверторным двигателем.', 799.99, 15),
        ('Микроволновая печь Panasonic NN-SU696S', 'Микроволновая печь с грилем и функцией размораживания.', 119.99, 20),
        ('Пылесос-робот iRobot Roomba i7+', 'Робот-пылесос с навигацией и самоочисткой контейнера.', 699.99, 8),
        ('Автомобиль Toyota Camry', 'Седан с бензиновым двигателем и автоматической коробкой передач.', 24999.99, 3);

COMMIT;

BEGIN;

CREATE TABLE IF NOT EXISTS "user" (
                        id SERIAL PRIMARY KEY,
                        name VARCHAR(255) NOT NULL,
                        surname VARCHAR(255) NOT NULL,
                        phone VARCHAR(15) NOT NULL,
                        login VARCHAR(255) NOT NULL,
                        password_hash VARCHAR(255) NOT NULL,
                        role VARCHAR(50) NOT NULL,
                        address TEXT,
                        coordinate_address_x DECIMAL(10, 6),
                        coordinate_address_y DECIMAL(10, 6)
);

CREATE TABLE IF NOT EXISTS "order" (
                         ID SERIAL PRIMARY KEY,
                         user_id INT REFERENCES "user"(id),
                         products_id INT[],
                         address TEXT,
                         coordinate_address_x DECIMAL(10, 6),
                         coordinate_address_y DECIMAL(10, 6),
                         coordinates_point_x DECIMAL(10, 6),
                         coordinates_point_y DECIMAL(10, 6),
                         create_at TIMESTAMP,
                         start_at TIMESTAMP,
                         delivery_at TIMESTAMP,
                         courier_id INT,
                         delivery_status VARCHAR(50)
);


CREATE TABLE IF NOT EXISTS "delivery" (
                          ID SERIAL PRIMARY KEY,
                          user_id INT REFERENCES "user" (id),
                          order_id INT REFERENCES "order" (ID)
);


CREATE TABLE IF NOT EXISTS "notes"(
                      id  SERIAL PRIMARY KEY,
                      user_id INT NOT NULL,
                      title  VARCHAR(50) NOT NULL,
                      description TEXT
);

COMMIT ;

-- migrate -path migrations -database "postgresql://delivery:delivery@localhost:6543/delivery?sslmode=disable" up
