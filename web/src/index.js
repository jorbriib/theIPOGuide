import React from 'react';
import ReactDOM from 'react-dom';
import { Route, BrowserRouter, Switch } from "react-router-dom";

import App from "./app";

ReactDOM.render(
    <React.StrictMode>
        <BrowserRouter>
            <Route path="/" component={App} />
        </BrowserRouter>
    </React.StrictMode>,
    document.getElementById('theIPOGuide')
);
