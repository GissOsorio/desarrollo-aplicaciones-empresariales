import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App.tsx'
import './index.css'
import {Auth0Provider} from "@auth0/auth0-react";
import {BrowserRouter} from "react-router-dom";

ReactDOM.createRoot(document.getElementById('root')!).render(
  <BrowserRouter>
      <Auth0Provider
          domain="aplicaciones-empresariales.us.auth0.com"
          clientId="LDxQSJy5CLnrKsMFVQhLMaQyGrgAjQSj"
          authorizationParams={{
              redirect_uri: window.location.origin
          }}
      >
          <App />
      </Auth0Provider>
  </BrowserRouter>,
)
