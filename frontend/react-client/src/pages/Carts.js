import { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import { fetchCarts } from "../api/cart";

const Carts = () => {
  const [carts, setCarts] = useState([]);
  const [error, setError] = useState(null);

  useEffect(() => {
    fetchCarts()
      .then(setCarts)
      .catch((err) => setError(err.message));
  }, []);

  if (error) return <div>Error: {error}</div>;

  return (
    <div>
      <h2>All Carts</h2>
      {carts.length === 0 ? (
        <p>No carts found.</p>
      ) : (
        <ul>
          {carts.map((cart) => (
            <li key={cart.id}>
              <Link to={`/cart/${cart.id}`}>
                Cart #{cart.id} ({cart.products.length} products)
              </Link>
            </li>
          ))}
        </ul>
      )}
    </div>
  );
};

export default Carts;
