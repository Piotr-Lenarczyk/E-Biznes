#!/bin/bash

BASE_URL="http://localhost:9000"

echo "Testing API endpoints..."

# Test creating a product
echo "Creating a product..."
PRODUCT_ID=$(curl -s -X POST -H "Content-Type: application/json" -d '{"name": "Smartwatch", "price": 299.99}' $BASE_URL/products | jq -r '.id')
echo "Created Product ID: $PRODUCT_ID"

# Test getting all products
echo "Getting all products..."
curl -s -X GET $BASE_URL/products | jq

# Test getting a single product
echo "Getting product with ID: $PRODUCT_ID"
curl -s -X GET $BASE_URL/products/$PRODUCT_ID | jq

# Test updating a product
echo "Updating product with ID: $PRODUCT_ID"
curl -s -X PUT -H "Content-Type: application/json" -d "{\"name\": \"Updated Smartwatch\", \"price\": 279.99}" $BASE_URL/products/$PRODUCT_ID | jq

# Test deleting a product
echo "Deleting product with ID: $PRODUCT_ID"
curl -s -X DELETE $BASE_URL/products/$PRODUCT_ID
echo "Deleted."

# Test creating a category
echo "Creating a category..."
CATEGORY_ID=$(curl -s -X POST -H "Content-Type: application/json" -d '{"name": "Books"}' $BASE_URL/categories | jq -r '.id')
echo "Created Category ID: $CATEGORY_ID"

# Test getting all categories
echo "Getting all categories..."
curl -s -X GET $BASE_URL/categories | jq

# Test creating a basket
echo "Creating a basket..."
BASKET_ID=$(curl -s -X POST -H "Content-Type: application/json" -d "{\"productIds\": [$PRODUCT_ID]}" $BASE_URL/baskets | jq -r '.id')
echo "Created Basket ID: $BASKET_ID"

# Test getting all baskets
echo "Getting all baskets..."
curl -s -X GET $BASE_URL/baskets | jq

# Test getting a single basket
echo "Getting basket with ID: $BASKET_ID"
curl -s -X GET $BASE_URL/baskets/$BASKET_ID | jq

# Test updating a basket
echo "Updating basket with ID: $BASKET_ID"
curl -s -X PUT -H "Content-Type: application/json" -d "{\"productIds\": [$PRODUCT_ID]}" $BASE_URL/baskets/$BASKET_ID | jq

# Test deleting a basket
echo "Deleting basket with ID: $BASKET_ID"
curl -s -X DELETE $BASE_URL/baskets/$BASKET_ID
echo "Deleted."

echo "API tests completed."
