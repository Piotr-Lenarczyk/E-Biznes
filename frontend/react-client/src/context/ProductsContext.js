import React, { createContext, useState, useEffect, useContext, useMemo } from 'react';
import PropTypes from 'prop-types';
import { fetchProducts } from '../api/product';

const ProductsContext = createContext();

export const useProducts = () => useContext(ProductsContext);

export const ProductsProvider = ({ children }) => {
    const [products, setProducts] = useState([]);

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

    const value = useMemo(() => ({
        products,
    }), [products]);

    return (
        <ProductsContext.Provider value={value}>
            {children}
        </ProductsContext.Provider>
    );
};

ProductsProvider.propTypes = {
    children: PropTypes.node.isRequired,
};
