
import React, { useState, useEffect } from 'react';
import './App.css';

function App() {
  const [keys, setKeys] = useState([]);
  const [newKey, setNewKey] = useState('');
  const [newValue, setNewValue] = useState('');

  useEffect(() => {
    fetchKeys();
  }, []);

  const fetchKeys = async () => {
    try {
      const response = await fetch('http://localhost:8080/get');
      if (response.ok) {
        const data = await response.json();
        setKeys(data);
      } else {
        console.error('Failed to fetch keys:', response.statusText);
      }
    } catch (error) {
      console.error('Error fetching keys:', error);
    }
  };

  const handleSetKey = async () => {
    try {
      const response = await fetch('http://localhost:8080/set', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ key: newKey, value: newValue }),
      });
      if (response.ok) {
        console.log('Key set successfully');
        fetchKeys(); 
      } else {
        console.error('Failed to set key:', response.statusText);
      }
    } catch (error) {
      console.error('Error setting key:', error);
    }
  };

  return (
    <div className="App">
      <h1>LRU Cache Management</h1>
      <div>
        <h2>Keys in Cache:</h2>
        <ul>
          {keys.map((key, index) => (
            <li key={index}>{key}</li>
          ))}
        </ul>
      </div>
      <div>
        <h2>Set Key/Value Pair:</h2>
        <input
          type="text"
          placeholder="Key"
          value={newKey}
          onChange={(e) => setNewKey(e.target.value)}
        />
        <input
          type="text"
          placeholder="Value"
          value={newValue}
          onChange={(e) => setNewValue(e.target.value)}
        />
        <button onClick={handleSetKey}>Set Key</button>
      </div>
    </div>
  );
}

export default App;
