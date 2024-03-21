CREATE TABLE mst_services(
	id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	service_name VARCHAR(100) NOT NULL,
	unit VARCHAR(100) NOT NULL,
	price INT
);

CREATE TABLE mst_users(
	id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	full_name VARCHAR(100) NOT NULL,
	phone_number VARCHAR(16) NOT NULL,
	username VARCHAR(50),
	password VARCHAR(100),
	role VARCHAR(40),
	date_created TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE tx_enigma_laundry(
	id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	id_users VARCHAR(100) NOT NULL,
	id_services VARCHAR(100) NOT NULL,
	transaction_in INT NOT NULL,
	transaction_out INT NOT NULL,
	amount INT NOT NULL,
	created_at INT,
	updated_at INT
);

INSERT INTO mst_services(service_name, unit, price) VALUES('Cuci + Setrika', 'KG', 7000);
INSERT INTO mst_services(service_name, unit, price) VALUES('Laundry Sprei', 'Buah', 50000);
INSERT INTO mst_services(service_name, unit, price) VALUES('Laundry Karpet', 'Buah', 25000);
INSERT INTO mst_users(full_name, phone_number, username, password, role, date_created) VALUES('Syifa', '081232671626', 'aby', '1234','USER', '2024-02-20');
INSERT INTO mst_users(full_name, phone_number, username, password, role, date_created) VALUES('Rizky', '08773577488', 'rizky', '1234', 'USER', '2023-10-11');
INSERT INTO mst_users(full_name, phone_number, username, password, role, date_created) VALUES('Pepeng', '0812778822', 'rafel', '1234','USER', '2023-08-05');
INSERT INTO tx_enigma_laundry(id_users, id_services, transaction_in, transaction_out, amount) VALUES('7d4a8411-5bb2-4571-ab16-d64e4323cd96', '28efdaf0-23e0-4dc0-b694-8a65e218105c', '21b16e8b-18d4-4580-9764-5d572207be25', '2024-01-18', '2024-01-20', 5);
INSERT INTO tx_enigma_laundry(id_users, id_services, transaction_in, transaction_out, amount) VALUES('8a289a7f-0b3f-4efe-be03-edec53380e2e', 'fe506ddf-faf2-442e-8f55-938451362f12', '144677c0-8d28-429c-aeae-33a7dc5439ca', '2024-02-18', '2024-02-20', 1);
INSERT INTO tx_enigma_laundry(id_users, id_services, transaction_in, transaction_out, amount) VALUES('7d4a8411-5bb2-4571-ab16-d64e4323cd96', 'ff9f60b7-842e-46eb-9b0a-a8988dca2615', 'c33cd3fa-e393-4cf5-a9df-5e9e5fefbcc7', '2023-08-18', '2023-08-20', 2);


SELECT * FROM mst_services;
SELECT * FROM mst_users;
SELECT * FROM tx_enigma_laundry;