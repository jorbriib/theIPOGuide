import React, { Fragment } from "react";
import { Route, Switch } from "react-router-dom";

import NotFound from "./pages/NotFound/NotFound";
import ListIpos from "./pages/ListIpos/ListIpos";
import pageRoutes from "./pageRoutes";
import ViewIpo from "./pages/ViewIpo";

const App = () => {
  return (
    <Fragment>
      <Switch>
        <Route path={pageRoutes.home()} exact component={ListIpos} />
        <Route
          path={pageRoutes.ipo(":alias")}
          exact
          render={({ match }) => <ViewIpo alias={match.params.alias} />}
        />
        <Route path="*" component={NotFound} />
      </Switch>
    </Fragment>
  );
};

export default App;
