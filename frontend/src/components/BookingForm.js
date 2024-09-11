// src/components/BookingForm.js
import React from 'react';

const BookingForm = () => {
  return (
    <form className="booking-form">
      <label htmlFor="startDate">Start Date:</label>
      <input type="date" id="startDate" name="startDate" />

      <label htmlFor="endDate">End Date:</label>
      <input type="date" id="endDate" name="endDate" />

      <button type="submit">Book</button>
    </form>
  );
};

export default BookingForm;
