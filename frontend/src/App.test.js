import React from 'react';
import { render, screen, waitFor } from '@testing-library/react';
import { BrowserRouter as Router } from 'react-router-dom';
import App from './App';

describe('App', () => {
  it('renders the home page', async () => {
    render(
      <Router>
        <App />
      </Router>
    );

    await waitFor(() => expect(screen.getByText(/Home/i)).toBeInTheDocument());
  });

  it('renders the register page', async () => {
    render(
      <Router>
        <App />
      </Router>
    );

    await waitFor(() => expect(screen.getByText(/Register/i)).toBeInTheDocument());
  });
});
