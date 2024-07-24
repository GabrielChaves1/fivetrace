import axios from 'axios';

const api = axios.create({
  baseURL: 'https://upbxrc60ub.execute-api.us-east-1.amazonaws.com/dev',
  headers: {
    'Content-Type': 'application/json',
  },
});

export default api;