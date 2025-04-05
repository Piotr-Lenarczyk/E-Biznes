#!/bin/bash

# Category tests
echo "GET /categories"
curl -X GET http://localhost:8080/categories

echo "POST /categories"
curl -X POST http://localhost:8080/categories \
-H "Content-Type: application/json" \
-d '{"name": "Electronics"}'

echo "POST /categories"
curl -X POST http://localhost:8080/categories \
-H "Content-Type: application/json" \
-d '{"name": "Furniture"}'

echo "POST /categories"
curl -X POST http://localhost:8080/categories \
-H "Content-Type: application/json" \
-d '{"name": "Gardening"}'

echo "GET /categories"
curl -X GET http://localhost:8080/categories

echo "GET /categories/1"
curl -X GET http://localhost:8080/categories/1

echo "PUT /categories/1"
curl -X PUT http://localhost:8080/categories/1 \
-H "Content-Type: application/json" \
-d '{"name": "Updated Electronics"}'

echo "DELETE /categories/1"
curl -X DELETE http://localhost:8080/categories/1


# Product tests
echo "GET /products"
curl -X GET http://localhost:8080/products

echo "POST /products"
curl -X POST http://localhost:8080/products \
-H "Content-Type: application/json" \
-d '{"name": "Product 1", "price": 100, "categories": [{"id": 2}, {"id": 3}]}'

echo "GET /products"
curl -X GET http://localhost:8080/products

echo "GET /products/1"
curl -X GET http://localhost:8080/products/1

echo "PUT /products/1"
curl -X PUT http://localhost:8080/products/1 \
-H "Content-Type: application/json" \
-d '{"name": "Updated Product", "price": 200, "categories": [{"id": 3}, {"id": 2}]}'

echo "DELETE /products/1"
curl -X DELETE http://localhost:8080/products/1


# Cart tests
echo "POST /products"
curl -X POST http://localhost:8080/products \
-H "Content-Type: application/json" \
-d '{"name": "Product 2", "price": 100, "categories": [{"id": 2}]}'

echo "POST /products"
curl -X POST http://localhost:8080/products \
-H "Content-Type: application/json" \
-d '{"name": "Product 3", "price": 200, "categories": [{"id": 3}]}'

echo "GET /carts"
curl -X GET http://localhost:8080/carts

echo "POST /carts"
curl -X POST http://localhost:8080/carts \
-H "Content-Type: application/json" \
-d '{"products": [{"id": 2}]}'

echo "GET /carts/1"
curl -X GET http://localhost:8080/carts/1

echo "PUT /carts/1"
curl -X PUT http://localhost:8080/carts/1 \
-H "Content-Type: application/json" \
-d '{"products": [{"id": 2}, {"id": 3}]}'

echo "DELETE /carts/1"
curl -X DELETE http://localhost:8080/carts/1
