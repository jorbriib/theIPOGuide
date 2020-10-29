import React, { Fragment } from "react";

import pageRoutes from "../../../pageRoutes";

const Footer = () => {
  return (
    <Fragment>
      <footer className="footer-three footer-grey p-top-65">
        <div className="footer-top p-bottom-25">
          <div className="container">
            <div className="row">
              <div className="col-lg-3 col-sm-3">
                <div className="widget widget_pages">
                  <ul className="list-unstyled">
                    <li className="page-item">
                      <a href={pageRoutes.about()}>About Us</a>
                    </li>
                  </ul>
                </div>
              </div>
              <div className="col-lg-3 col-sm-3">
                <div className="widget widget_pages">
                  <ul className="list-unstyled">
                    <li className="page-item">
                      <a href={pageRoutes.contact()}>Contact Us</a>
                    </li>
                  </ul>
                </div>
              </div>
              <div className="col-lg-3 col-sm-3">
                <div className="widget widget_pages">
                  <ul className="list-unstyled">
                    <li className="page-item">
                      <a href={pageRoutes.privacyPolicy()}>Privacy Policy</a>
                    </li>
                  </ul>
                </div>
              </div>
              <div className="col-lg-3 col-sm-3">
                <div className="widget widget_pages">
                  <ul className="list-unstyled">
                    <li className="page-item">
                      <a href={pageRoutes.termsAndConditions()}>
                        Terms and Conditions
                      </a>
                    </li>
                  </ul>
                </div>
              </div>
            </div>
          </div>
        </div>
        {/* ends: .Footer-top */}
        <div className="footer-bottom">
          <div className="container">
            <div className="row">
              <div className="col-lg-12">
                <div className="footer-bottom--content">
                  <a to={pageRoutes.home()} className="footer-logo">
                    <img
                      src="/assets/images/logo-black.png"
                      alt="theIPOguide"
                    />
                  </a>
                  <p className="m-0 copy-text">
                    ©{new Date().getFullYear()} theIPOGuide
                  </p>
                  <ul className="list-unstyled lng-list">
                    <li>
                      <a href={pageRoutes.termsAndConditions()}>Disclaimer</a>
                    </li>
                  </ul>
                </div>
              </div>
            </div>
          </div>
        </div>
        {/* ends: .Footer-bottom */}
      </footer>
      {/* ends: .Footer */}
    </Fragment>
  );
};
export default Footer;
