import React from 'react';
import { useProducts } from '../context/ProductsContext';

export default function Products() {
    const { products } = useProducts();

    return (
        <div>
            <h2>Products</h2>
            <ul>
                {products.map((product) => (
                    <li key={product.id}>{product.name}</li>
                ))}
            </ul>
        </div>
    );
}
