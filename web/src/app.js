import React, { Fragment } from "react";
import { Route, Switch } from "react-router-dom";
import CookieConsent from "react-cookie-consent";

import NotFound from "./pages/NotFound/NotFound";
import ListIpos from "./pages/ListIpos/ListIpos";
import pageRoutes from "./pageRoutes";
import ViewIpo from "./pages/ViewIpo";
import PrivacyPolicy from "./pages/PrivacyPolicy";
import TermsAndConditions from "./pages/TermsAndConditions";

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
        <Route
          path={pageRoutes.privacyPolicy()}
          exact
          component={PrivacyPolicy}
        />
        <Route
          path={pageRoutes.termsAndConditions()}
          exact
          component={TermsAndConditions}
        />
        <Route path="*" component={NotFound} />
      </Switch>
      <CookieConsent buttonText="Accept" onAccept={() => {}}>
        By clicking “Accept”, you agree to the storing of cookies on your device
        to enhance site navigation, analyze site usage, and assist in our
        marketing efforts.{" "}
        <a href={pageRoutes.privacyPolicy()}>Cookie Notice</a>
      </CookieConsent>
    </Fragment>
  );
};

export default App;
