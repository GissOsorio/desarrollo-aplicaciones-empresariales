import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App.tsx'
import './index.css'
import {Auth0Provider} from "@auth0/auth0-react";

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
      <Auth0Provider
          domain="aplicaciones-empresariales.us.auth0.com"
          clientId="LDxQSJy5CLnrKsMFVQhLMaQyGrgAjQSj"
          authorizationParams={{
              redirect_uri: window.location.origin
          }}
      >
          <App />
      </Auth0Provider>
  </React.StrictMode>,
)