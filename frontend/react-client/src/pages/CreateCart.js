// src/pages/CreateCart.js
import { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { fetchProducts } from "../api/product";

const CreateCart = () => {
  const [products, setProducts] = useState([]);
  const [selectedProducts, setSelectedProducts] = useState([]);
  const navigate = useNavigate();

  useEffect(() => {
    async function loadProducts() {
      try {
        const productsData = await fetchProducts();
        setProducts(productsData);
      } catch (error) {
        console.error("Error fetching products:", error);
      }
    }
    loadProducts();
  }, []);

  const handleCheckboxChange = (event) => {
    const productId = parseInt(event.target.value);
    if (event.target.checked) {
      setSelectedProducts((prev) => [...prev, { id: productId }]);
    } else {
      setSelectedProducts((prev) =>
          prev.filter((product) => product.id !== productId)
      );
    }
    console.log("Selected Products:", selectedProducts); // Debugging
  };

  const handleSubmit = async (event) => {
    event.preventDefault();
    console.log("Form submitted");

    try {
      const response = await fetch("http://localhost:8080/carts", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ products: selectedProducts }),
      });

      if (!response.ok) {
        throw new Error("Failed to create cart");
      }

      const newCart = await response.json();
      navigate(`/cart/${newCart.id}`);
    } catch (error) {
      console.error("Error creating cart:", error);
    }
  };

  return (
      <div>
        <h2>Create Cart</h2>
        <form onSubmit={handleSubmit}>
          {products.length === 0 ? (
              <p>Loading products...</p>
          ) : (
              products.map((product) => (
                  <div key={product.id}>
                    <label>
                      <input
                          type="checkbox"
                          value={product.id}
                          onChange={handleCheckboxChange}
                      />
                      {product.name}
                    </label>
                  </div>
              ))
          )}
          <button type="submit">Create Cart</button>
        </form>
      </div>
  );
};

export default CreateCart;
