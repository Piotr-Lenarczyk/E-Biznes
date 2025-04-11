import { Link } from "react-router-dom";

export default function Navbar() {
  return (
    <nav>
      <ul>
        <li><Link to="/">Products</Link></li>
        <li><Link to="/payments">Payments</Link></li>
        <li><Link to="/cart">Cart</Link></li>
	<li><Link to="/create-cart">Create New Cart</Link></li>
      </ul>
    </nav>
  );
}
