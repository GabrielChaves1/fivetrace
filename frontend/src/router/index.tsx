// src/router/index.tsx
import { createBrowserRouter } from 'react-router-dom';
import App from '../App';
import Signup from '../pages/sign-up';

const router: ReturnType<typeof createBrowserRouter> = createBrowserRouter([
  {
    path: "/",
    element: <App />,
    children: [
      {
        path: "/signup",
        element: <Signup />
      },
    ]
  }
]);

export default router;