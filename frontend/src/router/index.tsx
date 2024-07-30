// src/router/index.tsx
import Confirm from '@/features/auth/pages/confirm';
import Signup from '@/features/auth/pages/signup';
import Dashboard from '@/features/dashboard/pages/server-selector';
import Welcome from '@/features/dashboard/pages/welcome-screen';
import Settings from '@/features/settings/pages/layout';
import ServerSettings from '@/features/settings/pages/server-settings';
import { createBrowserRouter } from 'react-router-dom';
import App from '../App';

const router: ReturnType<typeof createBrowserRouter> = createBrowserRouter([
  {
    path: "/",
    element: <App />,
    children: [
      {
        path: "/auth",
        children: [
          { path: "/auth/signup", element: <Signup /> },
          { path: "/auth/confirm", element: <Confirm /> },
        ]
      },
      {
        path: "/dashboard",
        element: <Dashboard />
      },
      {
        path: "/welcome",
        element: <Welcome />
      },
      {
        path: "/settings",
        element: <Settings />,
        children: [
          { path: "/settings/:serverId", element: <ServerSettings /> },
        ]
      },
    ],
  },
]);

export default router;