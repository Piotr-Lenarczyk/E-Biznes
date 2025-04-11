import React, { useState, useEffect } from 'react';
import axios from 'axios';

const Products = () => {
  const [products, setProducts] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    // Replace with your backend API URL
    axios.get('http://localhost:8080/products')
      .then(response => {
        setProducts(response.data);
        setLoading(false);
      })
      .catch(error => {
        console.error('Error fetching products:', error);
        setLoading(false);
      });
  }, []);

  if (loading) {
    return <p>Loading...</p>;
  }

  return (
    <div>
      <h2>Products</h2>
      <ul>
        {products.length > 0 ? (
          products.map(product => (
            <li key={product.id}>
              {product.name} - ${product.price}
            </li>
          ))
        ) : (
          <p>No products available.</p>
        )}
      </ul>
    </div>
  );
}

export default Products;

