import React, { Fragment } from "react";
import { NavLink } from "react-router-dom";

import pageRoutes from "../../../pageRoutes";
import { useSimilarListIpos } from "../../../pages/ListIpos/useListIpos";

export function ContactIpo({ ipo }) {
  const {
    companyName,
    companyAddress,
    companyCountry,
    companyPhone,
    companyEmail,
    companyWeb,
    companyLogo,
    companyFacebook,
    companyTwitter,
    companyLinkedin,
    companyPinterest,
    companyInstagram,
    companyCeo,
  } = ipo;
  return (
    <Fragment>
      <div className="widget atbd_widget widget-card">
        <div className="atbd_widget_title">
          <h4>
            <span className="la la-user"></span>Contact Info
          </h4>
        </div>
        <div className="widget-body atbd_author_info_widget">
          <div className="atbd_avatar_wrapper">
            <div className="atbd_review_avatar">
              <img src={companyLogo} alt={companyName} />
            </div>
            {companyCeo && (
              <div className="atbd_name_time">
                <h4>{companyCeo}</h4>
                <span className="review_time">CEO</span>
              </div>
            )}
          </div>
          <div className="atbd_widget_contact_info">
            <ul>
              {companyAddress && (
                <li>
                  <span className="la la-map-marker"></span>
                  <span className="atbd_info">
                    {companyAddress}, {companyCountry}
                  </span>
                </li>
              )}
              {companyPhone && (
                <li>
                  <span className="la la-phone"></span>
                  <span className="atbd_info">{companyPhone}</span>
                </li>
              )}
              {companyEmail && (
                <li>
                  <span className="la la-envelope"></span>
                  <span className="atbd_info">{companyEmail}</span>
                </li>
              )}
              {companyWeb && (
                <li>
                  <span className="la la-globe"></span>
                  <a href={companyWeb} className="atbd_info" target="_blank">
                    {companyWeb
                      .replace(/(https:\/\/)|(http:\/\/)/, "")
                      .replace("www.", "")}
                  </a>
                </li>
              )}
            </ul>
          </div>
          <div className="atbd_social_wrap">
            {companyFacebook && (
              <p>
                <a href={companyFacebook} target="_blank">
                  <span className="fab fa-facebook-f"></span>
                </a>
              </p>
            )}
            {companyTwitter && (
              <p>
                <a href={companyTwitter} target="_blank">
                  <span className="fab fa-twitter"></span>
                </a>
              </p>
            )}
            {companyLinkedin && (
              <p>
                <a href={companyLinkedin} target="_blank">
                  <span className="fab fa-linkedin-in"></span>
                </a>
              </p>
            )}
            {companyPinterest && (
              <p>
                <a href={companyPinterest} target="_blank">
                  <span className="fab fa-pinterest"></span>
                </a>
              </p>
            )}
            {companyInstagram && (
              <p>
                <a href={companyInstagram} target="_blank">
                  <span className="fab fa-instagram"></span>
                </a>
              </p>
            )}
          </div>
        </div>
      </div>
    </Fragment>
  );
}

export function SimilarIpos({ ipo }) {
  const { alias } = ipo;
  const { similarStatus, similarIpos } = useSimilarListIpos(alias);
  if (similarStatus !== "ready") {
    return "";
  }

  return (
    <Fragment>
      <div className="widget atbd_widget widget-card">
        <div className="atbd_widget_title">
          <h4>
            <span className="la la-building"></span> Similar IPOs
          </h4>
          <NavLink to={pageRoutes.home()}>View All</NavLink>
        </div>
        {/*<!-- ends: .atbd_widget_title -->*/}
        <div className="atbd_categorized_listings atbd_similar_listings">
          <ul className="listings">
            {Object.values(similarIpos)
              .slice(0, 4)
              .map((value, key) => {
                return (
                  <li key={key}>
                    <div className="atbd_left_img">
                      <NavLink to={pageRoutes.ipo(value.alias)}>
                        <img
                          src={value.companyLogo}
                          style={{ width: "45px" }}
                          alt={value.companyName}
                        />
                      </NavLink>
                    </div>
                    <div className="atbd_right_content">
                      <div className="cate_title">
                        <h4>
                          <NavLink to={pageRoutes.ipo(value.alias)}>
                            {value.companyName}
                          </NavLink>
                        </h4>
                      </div>
                      <p className="directory_tag">{value.marketName}</p>
                    </div>
                  </li>
                );
              })}
          </ul>
        </div>
      </div>
    </Fragment>
  );
}
