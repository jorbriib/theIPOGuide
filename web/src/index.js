import React from "react";
import ReactDOM from "react-dom";
import { Route, BrowserRouter } from "react-router-dom";
import { HelmetProvider } from "react-helmet-async";

import App from "./app";

ReactDOM.render(
  <React.StrictMode>
    <HelmetProvider>
      <BrowserRouter>
        <Route path="/" component={App} />
      </BrowserRouter>
    </HelmetProvider>
  </React.StrictMode>,
  document.getElementById("theIPOGuide")
);
