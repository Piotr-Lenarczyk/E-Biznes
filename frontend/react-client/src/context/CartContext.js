import React, { createContext, useState, useEffect, useContext, useMemo } from 'react';
import PropTypes from 'prop-types';
import { fetchCart, createCart, updateCart } from '../api/cart';

const CartContext = createContext();

export const useCart = () => useContext(CartContext);

export const CartProvider = ({ children }) => {
    const [cart, setCart] = useState(null);

    useEffect(() => {
        const loadCart = async () => {
            try {
                const cartId = localStorage.getItem('cartId');
                if (cartId) {
                    const cartData = await fetchCart(cartId);
                    setCart(cartData);
                } else {
                    const newCart = await createCart([]);
                    localStorage.setItem('cartId', newCart.id);
                    setCart(newCart);
                }
            } catch (error) {
                console.error("Error loading cart:", error);
                const newCart = await createCart();
                localStorage.setItem('cartId', newCart.id);
                setCart(newCart);
            }
        };

        loadCart();
    }, []);

    const addProductToCart = (productId) => {
        if (cart) {
            const updatedCart = { ...cart, products: [...cart.products, { id: productId }] };
            setCart(updatedCart);
            updateCart(cart.id, updatedCart.products.map(p => p.id));
        }
    };

    const removeProductFromCart = (productId) => {
        if (cart) {
            const updatedCart = {
                ...cart,
                products: cart.products.filter((product) => product.id !== productId),
            };
            setCart(updatedCart);
            updateCart(cart.id, updatedCart.products.map(p => p.id));
        }
    };

    const value = useMemo(() => ({
        cart,
        addProductToCart,
        removeProductFromCart,
    }), [cart]);

    return (
        <CartContext.Provider value={value}>
            {children}
        </CartContext.Provider>
    );
};

CartProvider.propTypes = {
    children: PropTypes.node.isRequired,
};