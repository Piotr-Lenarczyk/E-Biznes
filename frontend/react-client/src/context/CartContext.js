import React, { createContext, useState, useEffect, useContext } from 'react';
import { fetchCart, createCart, updateCart } from '../api/cart'; // Assuming `createCart` is in `api/cart.js`

const CartContext = createContext();

export const useCart = () => useContext(CartContext);

export const CartProvider = ({ children }) => {
    const [cart, setCart] = useState(null);

    useEffect(() => {
        const loadCart = async () => {
            try {
                // Try to fetch the cart based on dynamic ID (you can modify this logic to get cart ID from URL, localStorage, etc.)
                const cartId = localStorage.getItem('cartId'); // Example of using localStorage to get cart ID
                if (cartId) {
                    const cartData = await fetchCart(cartId); // Try fetching the cart if cartId exists
                    setCart(cartData);
                } else {
                    const newCart = await createCart(); // If no cartId, create a new cart
                    localStorage.setItem('cartId', newCart.id); // Store the new cart ID for future reference
                    setCart(newCart);
                }
            } catch (error) {
                // If the cart doesn't exist (or any error occurs), create a new cart
                const newCart = await createCart();
                localStorage.setItem('cartId', newCart.id); // Store the new cart ID
                setCart(newCart);
            }
        };

        loadCart();
    }, []);

    const addProductToCart = (productId) => {
        if (cart) {
            const updatedCart = { ...cart, products: [...cart.products, { id: productId }] };
            setCart(updatedCart);
            updateCart(cart.id, updatedCart.products.map(p => p.id)); // Update the backend with the new cart
        }
    };

    const removeProductFromCart = (productId) => {
        if (cart) {
            const updatedCart = {
                ...cart,
                products: cart.products.filter((product) => product.id !== productId),
            };
            setCart(updatedCart);
            updateCart(cart.id, updatedCart.products.map(p => p.id)); // Update the backend with the updated cart
        }
    };

    return (
        <CartContext.Provider value={{ cart, addProductToCart, removeProductFromCart }}>
            {children}
        </CartContext.Provider>
    );
};
