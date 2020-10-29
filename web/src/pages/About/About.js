import React, { Fragment } from "react";

import Footer from "../../components/layout/Footer";
import Header from "../../components/layout/Header";
import { BreadCrumbAbout } from "../../components/elements/Breadcrumbs";
import {
  ContentBlock1About,
  ContentBlock2About,
  CounterAbout,
} from "../../components/elements/ContentAbout";

const About = () => {
  return (
    <Fragment>
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
