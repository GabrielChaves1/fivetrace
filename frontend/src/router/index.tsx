// src/router/index.tsx
import Confirm from '@/modules/auth/pages/confirm';
import Signup from '@/modules/auth/pages/signup';
import Dashboard from '@/modules/dashboard/pages/server-selector';
import Welcome from '@/modules/dashboard/pages/welcome-screen';
import Panel from '@/modules/panel/pages';
import PanelLogs from '@/modules/panel/pages/logs';
import Settings from '@/modules/settings/pages';
import ServerSettings from '@/modules/settings/pages/server';
import ServerSettingsGeneral from '@/modules/settings/pages/server/general';
import ServerSettingsIntegration from '@/modules/settings/pages/server/integration';
import ServerNotFound from '@/modules/settings/pages/server/not-found';
import ServerSettingsTeam from '@/modules/settings/pages/server/team';
import api from '@/services/axios';
import { createBrowserRouter } from 'react-router-dom';
import App from '../App';
import OrganizationSettings from '@/modules/settings/pages/organization';
import OrganizationSettingsGeneral from '@/modules/settings/pages/organization/general';

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
        loader: async () => {
          return await api.get("/servers")
        },
        element: <Dashboard />
      },
      {
        path: "/welcome",
        element: <Welcome />
      },
      {
        path: "/settings",
        loader: async () => {
          return await api.get("/servers")
        },
        element: <Settings />,
        children: [
          {
            path: "/settings/organization",
            element: <OrganizationSettings />,
            children: [
              { path: "/settings/organization/general", element: <OrganizationSettingsGeneral /> }
            ]
          },
          { 
            path: "/settings/servers/:serverId", 
            element: <ServerSettings />,
            loader: async ({ params }) => {
              const { serverId } = params;
              console.log(serverId);
              return await api.get(`/servers/${serverId}`)
            },
            children: [
              { path: "/settings/servers/:serverId/general", element: <ServerSettingsGeneral /> },
              { path: "/settings/servers/:serverId/team", element: <ServerSettingsTeam /> },
              { path: "/settings/servers/:serverId/integration", element: <ServerSettingsIntegration /> },
              { path: "/settings/servers/:serverId/*", element: <ServerNotFound /> },
            ]
          },
        ]
      },
      {
        path: "/panel/:serverId",
        element: <Panel />,
        children: [
          { path: "/panel/:serverId/logs", element: <PanelLogs /> }
        ]
      }
    ],
  },
]);

export default router;