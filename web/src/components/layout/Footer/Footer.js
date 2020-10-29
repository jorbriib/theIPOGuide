import React, { Fragment } from "react";
import { NavLink } from "react-router-dom";

import pageRoutes from "../../../pageRoutes";

const Footer = () => {
  return (
    <Fragment>
      <footer className="footer-three footer-grey p-top-95">
        <div className="footer-top p-bottom-25">
          <div className="container">
            <div className="row">
              <div className="col-lg-12 col-sm-12">
                <div className="widget widget_pages">
                  <h2 className="widget-title">Company Info</h2>
                  <ul className="list-unstyled">
                    <li className="page-item">
                      <NavLink to="/about">About Us</NavLink>
                    </li>
                    <li className="page-item">
                      <a href={pageRoutes.contact()}>Contact Us</a>
                    </li>
                    <li className="page-item">
                      <a href={pageRoutes.privacyPolicy()}>Privacy Policy</a>
                    </li>
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
                  <NavLink to="/" className="footer-logo">
                    <img
                      src="/assets/images/logo-black.png"
                      alt="theIPOguide"
                    />
                  </NavLink>
                  <p className="m-0 copy-text">
                    ©{new Date().getFullYear()} theIPOGuide
                  </p>
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
