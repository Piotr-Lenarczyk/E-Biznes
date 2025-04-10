// src/App.js
import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom'; // Update import
import Products from './Products';
import Payments from './Payments';

function App() {
  return (
    <Router>
      <div className="App">
        <h1>My React App</h1>
        <nav>
          <ul>
            <li><a href="/products">Products</a></li>
            <li><a href="/payments">Payments</a></li>
          </ul>
        </nav>
        <Routes> {/* Use Routes instead of Switch */}
          <Route path="/products" element={<Products />} /> {/* Update route definition */}
          <Route path="/payments" element={<Payments />} /> {/* Update route definition */}
        </Routes>
      </div>
    </Router>
  );
}

export default App;
