import React, { useState } from 'react';
import axios from 'axios';

const Payments = () => {
  const [paymentDetails, setPaymentDetails] = useState({
    amount: 0,
    cardNumber: '',
    expirationDate: '',
  });
  const [message, setMessage] = useState('');

  const handleChange = (e) => {
    const { name, value } = e.target;

    if (name === 'amount') {
      setPaymentDetails((prevDetails) => ({
        ...prevDetails,
        [name]: parseFloat(value),
      }));
    } else {
      setPaymentDetails((prevDetails) => ({
        ...prevDetails,
        [name]: value,
      }));
    }
  };

  const handlePaymentSubmit = (e) => {
    e.preventDefault();
    console.log('Payment form submitted');

    axios.post('http://localhost:8080/payments', paymentDetails, {
      headers: {
        'Content-Type': 'application/json',
        'Access-Control-Allow-Origin': '*',
      },
    })
        .then((response) => {
          console.log('Payment response:', response);
          if (response.status === 200) {
            setMessage('Payment Successful!');
          } else {
            setMessage('Payment Failed!');
          }
        })
        .catch((error) => {
          console.error('Error during payment:', error);
          setMessage('Payment Failed!');
        });
  };

  return (
      <div>
        <h2>Payments</h2>
        <form onSubmit={handlePaymentSubmit}>
          <div>
            <label htmlFor="amount">Amount: </label>
            <input
                type="number"
                id="amount"
                name="amount"
                value={paymentDetails.amount}
                onChange={handleChange}
                required
            />
          </div>
          <div>
            <label htmlFor="cardNumber">Card Number: </label>
            <input
                type="text"
                id="cardNumber"
                name="cardNumber"
                value={paymentDetails.cardNumber}
                onChange={handleChange}
                required
            />
          </div>
          <div>
            <label htmlFor="expirationDate">Expiration Date: </label>
            <input
                type="text"
                id="expirationDate"
                name="expirationDate"
                value={paymentDetails.expirationDate}
                onChange={handleChange}
                required
            />
          </div>
          <button type="submit">Submit Payment</button>
        </form>
        {message && <p>{message}</p>}
      </div>
  );
};

export default Payments;