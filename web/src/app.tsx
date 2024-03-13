import './app.css';
// @refresh reload
import { Router } from '@solidjs/router';
import { FileRoutes } from '@solidjs/start';
import { Suspense } from 'solid-js';
import { MetaProvider } from '@solidjs/meta';
import Layout from './components/containers/layout';

export default function App() {
  return (
    <MetaProvider>
      <Router
        root={(props) => (
          <>
            <Suspense>
              <Layout {...props} />
            </Suspense>
          </>
        )}
      >
        <FileRoutes />
      </Router>
    </MetaProvider>
  );
}
