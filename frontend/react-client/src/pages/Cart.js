// src/pages/Cart.js
import React, { useState, useEffect } from "react";
import { fetchCart, updateCart } from "../api/cart";
import { fetchProducts } from "../api/product";
import { useParams } from "react-router-dom";

const Cart = () => {
    const { id } = useParams();
    const [cart, setCart] = useState(null);
    const [products, setProducts] = useState([]);
    const [selectedProducts, setSelectedProducts] = useState([]);

    // Load the cart and products data when the component is mounted
    useEffect(() => {
        const loadData = async () => {
            try {
                const cartData = await fetchCart(id);
                setCart(cartData);
                setSelectedProducts(cartData.products.map((p) => p.id));

                const productsData = await fetchProducts();
                setProducts(productsData);
            } catch (error) {
                console.error("Error fetching data:", error);
            }
        };

        loadData();
    }, [id]);

    // Handle checkbox change to add or remove products
    const handleCheckboxChange = (productId) => {
        setSelectedProducts((prevSelected) =>
            prevSelected.includes(productId)
                ? prevSelected.filter((id) => id !== productId) // Remove product if already selected
                : [...prevSelected, productId] // Add product if not selected
        );
    };

    // Handle updating the cart with selected products
    const handleUpdateCart = async () => {
        try {
            await updateCart(id, selectedProducts); // Update cart in backend
            const updatedCart = await fetchCart(id); // Fetch updated cart data
            setCart(updatedCart); // Update the state with new cart data
        } catch (error) {
            console.error("Error updating cart:", error);
        }
    };

    if (!cart || !products) return <div>Loading...</div>;

    return (
        <div>
            <h2>Cart #{cart.id}</h2>

            <h3>Products in Cart:</h3>
            <ul>
                {cart.products.map((product) => (
                    <li key={product.id}>{product.name}</li>
                ))}
            </ul>

            <h3>All Available Products:</h3>
            {products.map((product) => (
                <div key={product.id}>
                    <input
                        type="checkbox"
                        checked={selectedProducts.includes(product.id)} // Check if product is selected
                        onChange={() => handleCheckboxChange(product.id)} // Handle checkbox change
                    />
                    {product.name}
                </div>
            ))}

            <button onClick={handleUpdateCart}>Update Cart</button>
        </div>
    );
};

export default Cart;
