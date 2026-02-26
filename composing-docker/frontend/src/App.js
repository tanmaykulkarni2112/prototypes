import React, { useState, useEffect } from 'react';
import axios from 'axios';
import './App.css';

function App() {
  const [items, setItems] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    fetchItems();
  }, []);

  const fetchItems = async () => {
    try {
      setLoading(true);
      // When running in Docker, use the backend service name
      const backendUrl = process.env.REACT_APP_BACKEND_URL || 'http://localhost:3001';
      const response = await axios.get(`${backendUrl}/api/items`);
      setItems(response.data);
      setError(null);
    } catch (err) {
      setError('Failed to fetch items from backend');
      console.error(err);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="app">
      <div className="container">
        <h1>Items from Backend</h1>
        
        {loading && <div className="status">Loading...</div>}
        {error && <div className="error">{error}</div>}
        
        {!loading && !error && (
          <div className="items-container">
            <div className="items-list">
              {items.length === 0 ? (
                <div className="empty">No items found</div>
              ) : (
                <ul>
                  {items.map((item) => (
                    <li key={item.id} className="item">
                      <div className="item-header">
                        <h3>{item.name}</h3>
                        <span className="item-id">ID: {item.id}</span>
                      </div>
                      <p className="item-description">{item.description}</p>
                    </li>
                  ))}
                </ul>
              )}
            </div>
          </div>
        )}

        <button onClick={fetchItems} className="refresh-btn">
          Refresh Data
        </button>
      </div>
    </div>
  );
}

export default App;
