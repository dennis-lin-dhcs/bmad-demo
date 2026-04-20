import React from 'react';
import ReactDOM from 'react-dom/client';
import {
  Box,
  CssBaseline,
  ThemeProvider,
  Toolbar,
  createTheme,
} from '@mui/material';
import { HashRouter, Navigate, Route, Routes, useLocation, useNavigate } from 'react-router-dom';
import Header from './components/Header';
import AppDrawer from './components/AppDrawer';
import Home from './pages/Home';
import Projects from './pages/Projects';
import Project1 from './pages/Project1';
import Project2 from './pages/Project2';
import Author from './pages/Author';
import Demo from './pages/Demo';
import LMStudioSettings from './pages/LMStudioSettings';
import OllamaSettings from './pages/OllamaSettings';
import ExternalEndpointSettings from './pages/ExternalEndpointSettings';

const drawerWidth = 300;

const theme = createTheme({
  palette: {
    mode: 'light',
    primary: {
      main: '#1976d2',
    },
  },
});

function AppShell() {
  const navigate = useNavigate();
  const location = useLocation();
  const [mobileOpen, setMobileOpen] = React.useState(false);

  const handleDrawerToggle = () => {
    setMobileOpen((prev: boolean) => !prev);
  };

  const handleNavigate = (path: string) => {
    navigate(path);
    setMobileOpen(false);
  };

  return (
    <Box sx={{ display: 'flex', minHeight: '100vh' }}>
      <CssBaseline />

      <Header drawerWidth={drawerWidth} onMenuClick={handleDrawerToggle} />

      <AppDrawer
        drawerWidth={drawerWidth}
        mobileOpen={mobileOpen}
        onClose={handleDrawerToggle}
        pathname={location.pathname}
        onNavigate={handleNavigate}
      />

      <Box
        component="main"
        sx={{
          flexGrow: 1,
          p: 3,
          width: { sm: `calc(100% - ${drawerWidth}px)` },
          backgroundColor: 'background.default',
        }}
      >
        <Toolbar />
        <Routes>
          <Route path="/" element={<Navigate to="/home" replace />} />
          <Route path="/home" element={<Home />} />
          <Route path="/projects" element={<Projects />} />
          <Route path="/projects/project-1" element={<Project1 />} />
          <Route path="/projects/project-2" element={<Project2 />} />
          <Route path="/about/author" element={<Author />} />
          <Route path="/about/demo" element={<Demo />} />
          <Route path="/settings/lm-studio" element={<LMStudioSettings />} />
          <Route path="/settings/ollama" element={<OllamaSettings />} />
          <Route path="/settings/external-endpoint" element={<ExternalEndpointSettings />} />
          <Route path="*" element={<Navigate to="/home" replace />} />
        </Routes>
      </Box>
    </Box>
  );
}

function RootApp() {
  return (
    <ThemeProvider theme={theme}>
      <HashRouter>
        <AppShell />
      </HashRouter>
    </ThemeProvider>
  );
}

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <RootApp />
  </React.StrictMode>
);
