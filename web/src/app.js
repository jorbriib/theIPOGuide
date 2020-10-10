import React, { Fragment } from "react";
import { Route, Switch } from "react-router-dom";

import NotFound from "./pages/NotFound/NotFound";
import ListIpos from "./pages/ListIpos/ListIpos";
import pageRoutes from "./pageRoutes";

const App = () => {
  return (
    <Fragment>
      <Switch>
        <Route path={pageRoutes.home()} exact component={ListIpos} />
        <Route path="*" component={NotFound} />
      </Switch>
    </Fragment>
  );
};

export default App;
