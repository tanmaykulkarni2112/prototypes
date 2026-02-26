const express = require('express');
const cors = require('cors');

const app = express();
const PORT = process.env.PORT || 3001;

// Middleware
app.use(cors());
app.use(express.json());

// Sample data
const items = [
  { id: 1, name: 'JavaScript', description: 'Popular programming language' },
  { id: 2, name: 'Docker', description: 'Containerization platform' },
  { id: 3, name: 'Express', description: 'Web application framework' },
  { id: 4, name: 'React', description: 'JavaScript library for building UIs' },
  { id: 5, name: 'MongoDB', description: 'NoSQL database' }
];

// Routes
app.get('/api/items', (req, res) => {
  res.json(items);
});

app.get('/api/items/:id', (req, res) => {
  const item = items.find(i => i.id === parseInt(req.params.id));
  if (!item) {
    return res.status(404).json({ error: 'Item not found' });
  }
  res.json(item);
});

app.get('/health', (req, res) => {
  res.json({ status: 'Backend is running' });
});

app.listen(PORT, () => {
  console.log(`Backend server listening on port ${PORT}`);
});
