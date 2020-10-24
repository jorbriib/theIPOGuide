import React, { Fragment } from "react";
import { NavLink } from "react-router-dom";

import pageRoutes from "../../../pageRoutes";

export function BreadcrumbWrapper({ title, onlyTitle }) {
  return (
    <Fragment>
      <div className="breadcrumb-wrapper content_above">
        {title && (
          <div className="container">
            <div className="row">
              <div className="col-lg-12 text-center">
                <h1 className="page-title">{title}</h1>
                {onlyTitle ? null : (
                  <nav aria-label="breadcrumb">
                    <ol className="breadcrumb">
                      <li className="breadcrumb-item">
                        <NavLink to={pageRoutes.home()}>Home</NavLink>
                      </li>
                      <li
                        className="breadcrumb-item active"
                        aria-current="page"
                      >
                        {title}
                      </li>
                    </ol>
                  </nav>
                )}
              </div>
            </div>
          </div>
        )}
      </div>
    </Fragment>
  );
}

export function BreadcrumbSingle({ ipo }) {
  const {
    companyName,
    companyCountry,
    companySector,
    priceFrom,
    priceTo,
    marketName,
  } = ipo;

  return (
    <Fragment>
      <div className="col-lg-8 col-md-7">
        <ul className="list-unstyled listing-info--badges">
          <li>
            <span className="atbd_badge atbd_badge_popular">{marketName}</span>
          </li>
        </ul>
        <ul className="list-unstyled listing-info--meta">
          <li>
            <span className="atbd_meta atbd_listing_average_pricing">
              {priceFrom} {priceTo && <span>- {priceTo}</span>}
            </span>
          </li>
          <li>
            <div className="atbd_listing_category">
              <a href="#">
                <span className="la la-map-marker"></span>
                {companyCountry}
              </a>
            </div>
          </li>
        </ul>
        {/*<!-- ends: .listing-info-meta -->*/}
        <h1>{companyName}</h1>
        <p className="subtitle">{companySector}</p>
      </div>
      <div className="col-lg-4 col-md-5 d-flex align-items-end justify-content-start justify-content-md-end">
        <div className="atbd_listing_action_area">
          <div className="atbd_action atbd_share dropdown">
            <span
              className="dropdown-toggle"
              id="social-links"
              data-toggle="dropdown"
              aria-haspopup="true"
              aria-expanded="false"
              role="menu"
            >
              <span className="la la-share"></span>Share
            </span>
            <div
              className="atbd_director_social_wrap dropdown-menu"
              aria-labelledby="social-links"
            >
              <ul className="list-unstyled">
                <li>
                  <a
                    href={
                      "https://www.facebook.com/sharer/sharer.php?u=" +
                      encodeURIComponent(window.location.href)
                    }
                    target="_blank"
                  >
                    <span className="fab fa-facebook-f color-facebook"></span>
                    Facebook
                  </a>
                </li>
                <li>
                  <a
                    href={
                      "https://twitter.com/intent/tweet?text=" +
                      encodeURIComponent(
                        "Looking for more information about " +
                          companyName +
                          " IPO?" +
                          " Look here: " +
                          window.location.href
                      )
                    }
                    target="_blank"
                  >
                    <span className="fab fa-twitter color-twitter"></span>
                    Twitter
                  </a>
                </li>
                <li className="no-desktop">
                  <a
                    href={
                      "whatsapp://send?text=" +
                      encodeURIComponent(
                        "Look here for more information about " +
                          companyName +
                          " IPO: " +
                          window.location.href
                      )
                    }
                    target="_blank"
                    data-action="share/whatsapp/share"
                  >
                    <span className="fab fa-whatsapp"></span>
                    WhatsApp
                  </a>
                </li>
                <li>
                  <a
                    href={
                      "https://www.linkedin.com/sharing/share-offsite/?url=" +
                      encodeURIComponent(window.location.href)
                    }
                    target="_blank"
                  >
                    <span className="fab fa-linkedin-in color-linkedin"></span>
                    LinkedIn
                  </a>
                </li>
              </ul>
            </div>
            {/* <!--Ends social share--> */}
          </div>
          <div className="atbd_action atbd_report">
            <div className="action_button">
              <a
                href=" "
                data-toggle="modal"
                data-target="#atbdp-report-abuse-modal"
              >
                <span className="la la-flag-o"></span> Feedback
              </a>
            </div>
          </div>
        </div>
        {/* <!-- ends: .atbd_listing_action_area --> */}
      </div>
    </Fragment>
  );
}
