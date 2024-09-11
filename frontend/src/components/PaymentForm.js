// src/components/PaymentForm.js
import React from 'react';

const PaymentForm = () => {
  return (
    <form className="payment-form">
      <label htmlFor="cardNumber">Card Number:</label>
      <input type="text" id="cardNumber" name="cardNumber" />

      <label htmlFor="expiryDate">Expiry Date:</label>
      <input type="text" id="expiryDate" name="expiryDate" />

      <label htmlFor="cvv">CVV:</label>
      <input type="text" id="cvv" name="cvv" />

      <button type="submit">Pay</button>
    </form>
  );
};

export default PaymentForm;
