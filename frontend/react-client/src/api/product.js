// src/api/product.js
const API_URL = "http://localhost:8080";

export async function fetchProducts() {
  const response = await fetch(`${API_URL}/products`);
  if (!response.ok) throw new Error("Failed to fetch products");
  return response.json();
}
