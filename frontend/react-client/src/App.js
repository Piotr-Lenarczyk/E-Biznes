import { BrowserRouter as Router, Routes, Route, Link } from "react-router-dom";
import Products from "./pages/Products";
import Payments from "./pages/Payments";
import Cart from "./pages/Cart";
import Carts from "./pages/Carts";
import CreateCart from './pages/CreateCart';
import { ProductsProvider } from './context/ProductsContext';
import { CartProvider } from './context/CartContext';

function App() {
    return (
        <Router>
            <ProductsProvider>
                <CartProvider>
                    <nav>
                        <Link to="/">Products</Link> | <Link to="/payments">Payments</Link> | <Link to="/carts">All Carts</Link> | <Link to="/create-cart">Create Cart</Link>
                    </nav>

                    <Routes>
                        <Route path="/" element={<Products />} />
                        <Route path="/payments" element={<Payments />} />
                        <Route path="/cart/:id" element={<Cart />} />
                        <Route path="/carts" element={<Carts />} />
                        <Route path="/create-cart" element={<CreateCart />} />
                    </Routes>
                </CartProvider>
            </ProductsProvider>
        </Router>
    );
}

export default App;
