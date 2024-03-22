# Aplikasi Enigma Laundry API

## Deskripsi

Fitur-fitur yang diminta adalah:

1.  Struktur/Design Database yang memenuhi kaidah normalisasi berdasarkan nota dibawah ini dengan kriteria sbb :

        - Hasil design dalam bentuk file Script DDL Postgre SQL
        - Design database minimal memiliki 2 tabel master dan 1 tabel transaksi

2.  Aplikasi berbasis API menggunakan bahasa pemrograman Golang dengan kriteria sbb :

        - Aplikasi memiliki fitur untuk melakukan GET, POST, PUT, dan DELETE pada tabel master
          1. Manajemen Users
          2. Manajemen Services
        - Aplikasi memiliki fitur untuk melakukan GET dan POST pada table Transaksi
          1. Manajemen Transaksi

3.  Dokumentasi cara menjalankan aplikasi dan penggunaan aplikasi dalam bentuk readme.md atau dokumen ektensi word atau pdf

---

## Instalation

Clone or download this repo to your local directory.

## Usage

1. Import DDL.sql on folder psql to your database.
2. Configure your db connection on config/config.go

```go
// Database Config
const (
	host     = ""   // Your connection hostname
	port     =      // Your connection port
	user     = ""   // Your connection username
	password = ""   // Your connection password
	dbname   = "enigma_laundry_api"
)
```

3. Open terminal on your program local directory.

```go
go run .
```

## API Spec

### Customer API

#### Create Users

Request :

- Method : `POST`
- Endpoint : `/users`
- Header :
  - Content-Type : application/json
  - Accept : application/json
- Body :

```json
{
  "full_name": "string",
  "phone_number": "string",
  "username": "string",
  "password": "string",
  "role": "string"
}
```

Response :

- Status : 201 Created
- Body :

```json
{
  "message": {
    "code": 201,
    "description": "OK"
  },
  "data": {
    "full_name": "string",
    "phone_number": "string",
    "username": "string",
    "password": "string",
    "role": "string"
  }
}
```

### Generate Token

Request :

- Method : `POST`
- Endpoint : `/users/login`
- Header :
  - Content-Type : application/json
  - Accept : application/json
- Body :

```json
{
    "username": "string",
    "password": "string"
}
```

Response :

- Status : 201 Created
- Body :

```json
{
    "status": {
        "code": 200,
        "description": "OK"
    },
    "data": {
        "accesToken": "string",
        "userId": "string"
    }
}
```

#### List Customer

Request :

- Method : GET
- Endpoint : `/customers/:id`
  - Header :
  - Accept : application/json
- Query Param :
  - id : string,

Response :

- Status Code : 200 OK
- Body:

```json
{
    "status": {
        "code": 201,
        "description": "OK"
    },
    "data": {
        "id": "string",
        "full_name": "string",
        "phone_number": "string",
        "username": "string",
        "password": "string",
        "role": "string",
        "date_created": "time"
    }
}
```

#### Update Customer

Request :

- Method : PUT
- Endpoint : `/customers/:id`
- Header :
  - Content-Type : application/json
  - Accept : application/json
- Body :

```json
{
    "id": "string",
    "full_name": "string",
    "phone_number": "string",
    "username": "string",
    "password": "string",
    "role": "string"
}
```

Response :

- Status : 200 OK
- Body :

```json
{
    "status": {
        "code": 201,
        "description": "OK"
    },
    "data": {
        "id": "string",
        "full_name": "string",
        "phone_number": "string",
        "username": "string",
        "password": "string",
        "role": "string",
        "date_created": "time"
    }
}
```

#### Delete Customer

Request :

- Method : DELETE
- Endpoint : `/users/:id`
- Header :
  - Accept : application/json
- Body :

Response :

- Status : 200 OK
- Body :

```json
{
    "status": {
        "code": 201,
        "description": "OK"
    },
    "data": {
        "id": "string",
        "full_name": "string",
        "phone_number": "string",
        "username": "string",
        "password": "string",
        "role": "string",
        "date_created": "time"
    }
}
```

### Service API

#### Create Service

Request :

- Method : POST
- Endpoint : `/services`
- Header :
  - Content-Type : application/json
  - Accept : application/json
- Body :

```json
{
	"service_name": "string",
  "unit": "string" (satuan product,cth: Buah atau Kg),
  "price": int,
}
```

Response :

- Status Code: 201 Created
- Body:

```json
{
    "status": {
        "code": 201,
        "description": "OK"
    },
    "data": {
        "id": "string",
        "service_name": "string",
        "unit": "string",
        "price": int
    }
}
```

#### List Service

Request :

- Method : GET
- Endpoint : `/services/:id`
  - Header :
  - Accept : application/json
- Query Param :
  - id : string,

Response :

- Status Code : 200 OK
- Body:

```json
{
    "status": {
        "code": 201,
        "description": "OK"
    },
    "data": {
        "id": "string",
        "service_name": "string",
        "unit": "string",
        "price": int
    }
}
```

#### Update Service

Request :

- Method : PUT
- Endpoint : `/services/:id`
- Header :
  - Content-Type : application/json
  - Accept : application/json
- Body :

```json
{
    "id": "string",
    "service_name": "string",
    "unit": "string",
    "price": int
}
```

Response :

- Status Code: 200 OK
- Body :

```json
{
    "status": {
        "code": 201,
        "description": "OK"
    },
    "data": {
        "id": "string",
        "service_name": "string",
        "unit": "string",
        "price": int
    }
}
```

#### Delete Service

Request :

- Method : DELETE
- Endpoint : `/services/:id`
- Header :
  - Accept : application/json
- Body :

Response :

- Status : 200 OK
- Body :

```json
{
    "status": {
        "code": 201,
        "description": "OK"
    },
    "data": {
        "id": "string",
        "service_name": "string",
        "unit": "string",
        "price": int
    }
}
```

### Transaction API

#### Create Transaction

Request :

- Method : POST
- Endpoint : `/transactions`
- Header :
  - Content-Type : application/json
  - Accept : application/json
- Body :

```json
{
  "id_users": "string",
  "id_services": "string",
  "transaction_in": "string",
  "transaction_out": "string",
  "amount": int
}
```

Request :

- Status Code: 201 Created
- Body :

```json
{
    "status": {
        "code": 201,
        "description": "OK"
    },
    "data": {
        "id": "string",
        "id_users": "string",
        "id_services": "string",
        "transaction_in": int,
        "transaction_out": int,
        "amount": int,
        "created_at": int,
        "updated_at": int
    }
}
```

#### List Transaction

Request :

- Method : DELETE
- Endpoint : `/transactions/:id`
- Header :
  - Accept : application/json
- Query Param :
  - id : string,
- Body :

Response :

- Status Code: 200 OK
- Body :

```json
{
    "status": {
        "code": 201,
        "description": "OK"
    },
    "data": {
        "id": "string",
        "id_users": "string",
        "id_services": "string",
        "transaction_in": int,
        "transaction_out": int,
        "amount": int,
        "created_at": int,
        "updated_at": int
    }
}
```