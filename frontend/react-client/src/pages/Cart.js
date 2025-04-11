import React, { useState, useEffect } from "react";
import { fetchCart, updateCart } from "../api/cart";
import { fetchProducts } from "../api/product";
import { useParams } from "react-router-dom";

export default function Cart() {
    const { id } = useParams();
    const [cart, setCart] = useState(null);
    const [products, setProducts] = useState([]);
    const [selectedProducts, setSelectedProducts] = useState([]);

    useEffect(() => {
        const loadData = async () => {
            const cartData = await fetchCart(id);
            setCart(cartData);
            setSelectedProducts(cartData.products.map((p) => p.id));

            const productsData = await fetchProducts();
            setProducts(productsData);
        };

        loadData();
    }, [id]);

    const handleCheckboxChange = (productId) => {
        setSelectedProducts((prevSelected) =>
            prevSelected.includes(productId)
                ? prevSelected.filter((id) => id !== productId)
                : [...prevSelected, productId]  // Just store product IDs, not objects
        );
    };

    const handleUpdateCart = async () => {
        console.log("Updated selectedProducts:", selectedProducts);

        try {
            await updateCart(id, selectedProducts); // Pass product IDs directly
            const updatedCart = await fetchCart(id);
            setCart(updatedCart);
        } catch (error) {
            console.error("Failed to update cart:", error);
        }
    };

    if (!cart) return <div>Loading...</div>;

    return (
        <div>
            <h2>Cart #{cart.id}</h2>

            <h3>Products in Cart:</h3>
            <ul>
                {cart.products.map((product) => (
                    <li key={product.id}>{product.name}</li>
                ))}
            </ul>

            <h3>All Products:</h3>
            {products.map((product) => (
                <div key={product.id}>
                    <input
                        type="checkbox"
                        checked={selectedProducts.includes(product.id)}
                        onChange={() => handleCheckboxChange(product.id)}
                    />
                    {product.name}
                </div>
            ))}

            <button onClick={handleUpdateCart}>Update Cart</button>
        </div>
    );
}
