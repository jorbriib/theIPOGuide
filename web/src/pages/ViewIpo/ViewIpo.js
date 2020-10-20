import React, { Fragment } from "react";

import Header from "../../components/layout/Header";
import Footer from "../../components/layout/Footer";
import useViewIpo from "./useViewIpo";
import { BreadcrumbSingle } from "../../components/elements/Breadcrumbs";
import Report from "../../components/elements/Report/Report";
import View from "../../components/container/View";

const ViewIpo = (alias) => {
  const { status, ipo } = useViewIpo(alias.alias);

  const backgroundImage = {
    backgroundImage: "url('/assets/images/wallstreet-bull.jpg')",
    opacity: 1,
  };

  if (status !== "ready") {
    return "";
  }

  return (
    <Fragment>
      {/* Header section start */}
      <section className="listing-details-wrapper bgimage">
        <div className="bg_image_holder" style={backgroundImage}>
          <img
            src="/assets/images/wallstreet-bull.jpg"
            alt="Find your IPO and invest"
          />
        </div>
        <div className="mainmenu-wrapper">
          <Header class="menu--light" />
        </div>
        {/* <!-- ends: .mainmenu-wrapper --> */}
        <div className="listing-info content_above">
          <div className="container">
            <div className="row">
              <BreadcrumbSingle ipo={ipo} />
            </div>
          </div>
        </div>
      </section>
      {/* Header section end */}

      <View ipo={ipo} similars={[]} />
      <Report />
      <Footer />
    </Fragment>
  );
};

export default ViewIpo;
