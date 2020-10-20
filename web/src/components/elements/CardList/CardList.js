import React, { Fragment } from "react";
import { NavLink } from "react-router-dom";

const CardList = ({ list }) => {
  return (
    <Fragment>
      {Object.values(list).map((ipo, key) => {
        const {
          companyName,
          alias,
          companyCountry,
          companySector,
          companyLogo,
          priceFrom,
          marketName,
          expectedDate,
        } = ipo;
        return (
          <div className="col-lg-12" key={key}>
            <div className="atbd_single_listing atbd_listing_list">
              <article className="atbd_single_listing_wrapper">
                <figure className="atbd_listing_thumbnail_area">
                  <div className="atbd_listing_image">
                    <NavLink to={"/ipo/" + alias}>
                      <img src={companyLogo} alt={companyName} />
                    </NavLink>
                  </div>
                  {/*<!-- ends: .atbd_listing_image -->*/}
                </figure>
                {/*<!-- ends: .atbd_listing_thumbnail_area -->*/}
                <div className="atbd_listing_info">
                  <div className="atbd_content_upper">
                    <h4 className="atbd_listing_title">
                      <NavLink to={"/ipo/" + alias}>{companyName}</NavLink>
                    </h4>
                    <div className="atbd_listing_data_list">
                      <ul>
                        <li>
                          <span className="la la-money"></span>
                          {marketName}
                        </li>
                        <li>
                          <span className="la la-map-marker"></span>
                          {companyCountry}
                        </li>
                        <li>
                          <span className="la la-calendar-check-o"></span>
                          {new Intl.DateTimeFormat("en-GB", {
                            year: "numeric",
                            month: "long",
                            day: "2-digit",
                          }).format(expectedDate)}
                        </li>
                      </ul>
                    </div>
                  </div>
                  <div className="atbd_listing_bottom_content">
                    <div className="atbd_content_left">
                      <div className="atbd_listing_category">
                        <NavLink to={"/ipo/" + alias}>{companySector}</NavLink>
                      </div>
                    </div>
                  </div>
                  <div className="atbd_listing_meta">
                    <span className="atbd_meta atbd_listing_price">
                      {priceFrom}
                    </span>
                  </div>
                </div>
                {/*<!-- ends: .atbd_listing_info -->*/}
              </article>
              {/*<!-- atbd_single_listing_wrapper -->*/}
            </div>
          </div>
        );
      })}
    </Fragment>
  );
};

export default CardList;
