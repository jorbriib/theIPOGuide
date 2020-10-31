import React, { Fragment } from "react";

import Footer from "../../components/layout/Footer";
import Header from "../../components/layout/Header";
import { BreadCrumbAbout } from "../../components/elements/Breadcrumbs";
import {
  ContentBlock1About,
  ContentBlock2About,
  CounterAbout,
} from "../../components/elements/ContentAbout";
import pageRoutes from "../../pageRoutes";
import CustomHeader from "../../components/elements/CustomHeader";

const About = () => {
  return (
    <Fragment>
      <CustomHeader
        title="About"
        description="Explore IPOs around the world and invest your money"
        url={pageRoutes.about()}
      />
      {/* Header section start */}
      <section className="about-wrapper bg-gradient-ps">
        <div className="mainmenu-wrapper">
          <Header class="menu--light" />
        </div>
        {/* <!-- ends: .mainmenu-wrapper --> */}
        <BreadCrumbAbout />
      </section>
      {/* Header section end */}
      <ContentBlock1About />
      <CounterAbout />
      <ContentBlock2About />

      <Footer />
    </Fragment>
  );
};

export default About;
