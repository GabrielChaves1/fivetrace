// src/router/index.tsx
import { createBrowserRouter } from 'react-router-dom';
import App from '../App';
import Signup from '../pages/sign-up';
import Confirm from '@/pages/confirm';
import Dashboard from '@/pages/dashboard';

const router: ReturnType<typeof createBrowserRouter> = createBrowserRouter([
  {
    path: "/",
    element: <App />,
    children: [
      {
        path: "/signup",
        element: <Signup />
      },
      {
        path: "/confirm",
        element: <Confirm />
      },
      {
        path: "/dashboard",
        element: <Dashboard />
      },
    ]
  }
]);

export default router;