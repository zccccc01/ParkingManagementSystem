// User.js
import React from 'react';
import { useParams } from 'react-router-dom';

const User = () => {
  const { userId } = useParams();
  return <h2>User ID: {userId}</h2>;
};

export default User;