import React, { Fragment } from "react";

import Header from "../../components/layout/Header";
import Footer from "../../components/layout/Footer";
import pageRoutes from "../../pageRoutes";
import CustomHeader from "../../components/elements/CustomHeader";

const NotFound = () => {
  const backgroundImage = {
    backgroundImage: "url('/assets/images/wallstreet-bull.jpg')",
    opacity: 1,
  };

  return (
    <Fragment>
      <CustomHeader
        title="404 Page not found"
        description="Check if you typed the address correctly, go back to your previous page or try using our site search or menu to find something specific."
      />
      <section className="intro-wrapper bgimage overlay overlay--dark">
        <div className="bg_image_holder" style={backgroundImage}>
          <img
            src="/assets/images/wallstreet-bull.jpg"
            alt="404 Page not found"
          />
        </div>
        <div className="mainmenu-wrapper">
          <Header class="menu--light" />
        </div>
        {/* <!-- ends: .mainmenu-wrapper --> */}

        <div className="not_found_area">
          <div className="container">
            <div className="row">
              <div className="col-lg-8 offset-lg-2">
                <div className="not_found_title_area">
                  <h2 className="title">404 Page not found</h2>
                  <p className="sub_title">
                    Check if you typed the address correctly, go back to your
                    previous page or try using our site search or menu to find
                    something specific.
                  </p>
                </div>
              </div>
            </div>
          </div>
        </div>
      </section>
      <Footer />
    </Fragment>
  );
};

export default NotFound;
