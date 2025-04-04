curl -X GET http://localhost:8080/products

curl -X POST http://localhost:8080/products \
-H "Content-Type: application/json" \
-d '{"name": "Product 1", "price": 100}'

curl -X GET http://localhost:8080/products

curl -X GET http://localhost:8080/products/1

curl -X PUT http://localhost:8080/products/1 \
-H "Content-Type: application/json" \
-d '{"name": "Updated Product", "price": 200}'

curl -X DELETE http://localhost:8080/products/1
