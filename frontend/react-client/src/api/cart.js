const API_URL = "http://localhost:8080";

export async function fetchCart(cartId) {
  const response = await fetch(`${API_URL}/carts/${cartId}`);
  if (!response.ok) throw new Error("Failed to fetch cart");
  return response.json();
}

export async function fetchCarts() {
  const response = await fetch(`${API_URL}/carts`);
  if (!response.ok) throw new Error("Failed to fetch carts");
  return response.json();
}


export const createCart = async (productIds = []) => {
  try {
    const response = await fetch(`${API_URL}/carts`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ productIds }),
    });

    if (!response.ok) {
      const errorDetails = await response.text();
      console.error('Failed to create cart:', response.status, errorDetails);
      throw new Error('Failed to create cart');
    }

    return await response.json();
  } catch (error) {
    console.error('Error creating cart:', error);
    throw error;
  }
};

export async function updateCart(cartId, productIds) {
  console.log(productIds)
  const response = await fetch(`${API_URL}/carts/${cartId}`, {
    method: "PUT",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      products: productIds.map((id) => ({ id })), // Map product IDs into the expected format
    }),
  });

  if (!response.ok) throw new Error("Failed to update cart");
  return response.json();
}
