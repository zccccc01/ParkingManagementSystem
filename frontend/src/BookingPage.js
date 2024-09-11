import React, { useState } from 'react';
import { Button, Form } from 'react-bootstrap';
import styles from './BookingPage.module.scss';

const BookingPage = () => {
  const [bookingDetails, setBookingDetails] = useState({
    name: '',
    email: '',
    phone: '',
  });

  const handleInputChange = (event) => {
    const { name, value } = event.target;
    setBookingDetails((prevState) => ({
      ...prevState,
      [name]: value,
    }));
  };

  const handleSubmit = (event) => {
    event.preventDefault();
    console.log('Booking Details:', bookingDetails);
    // 提交逻辑
  };

  return (
    <div className={styles.bookingContainer}>
      <h2>Book Your Parking Spot</h2>
      <Form onSubmit={handleSubmit}>
        <Form.Group controlId="formName">
          <Form.Label>Name</Form.Label>
          <Form.Control
            type="text"
            name="name"
            placeholder="Enter your name"
            value={bookingDetails.name}
            onChange={handleInputChange}
            required
          />
        </Form.Group>

        <Form.Group controlId="formEmail">
          <Form.Label>Email</Form.Label>
          <Form.Control
            type="email"
            name="email"
            placeholder="Enter your email"
            value={bookingDetails.email}
            onChange={handleInputChange}
            required
          />
        </Form.Group>

        <Form.Group controlId="formPhone">
          <Form.Label>Phone Number</Form.Label>
          <Form.Control
            type="tel"
            name="phone"
            placeholder="Enter your phone number"
            value={bookingDetails.phone}
            onChange={handleInputChange}
            required
          />
        </Form.Group>

        <Button variant="primary" type="submit">
          Submit
        </Button>
      </Form>
    </div>
  );
};

export default Booking;
