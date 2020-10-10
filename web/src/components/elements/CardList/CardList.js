import React, { Fragment } from "react";
import { NavLink } from "react-router-dom";

const CardList = ({ list }) => {
  return (
    <Fragment>
      {Object.values(list).map((ipo, key) => {
        const { companyName, marketName } = ipo;
        return (
          <div className="col-lg-12" key={key}>
            <div className="atbd_single_listing atbd_listing_list">
              <article className="atbd_single_listing_wrapper">
                <figure className="atbd_listing_thumbnail_area">
                  <div className="atbd_listing_image">
                    <a href=" ">
                      <img src="" alt="listingimage" />
                    </a>
                  </div>
                  {/*<!-- ends: .atbd_listing_image -->*/}
                </figure>
                {/*<!-- ends: .atbd_listing_thumbnail_area -->*/}
                <div className="atbd_listing_info">
                  <div className="atbd_content_upper">
                    <h4 className="atbd_listing_title">
                      <NavLink to={"/company" + companyName}>
                        {companyName}
                      </NavLink>
                    </h4>
                    <div className="atbd_listing_meta">
                      <span className="atbd_meta atbd_listing_price">
                        {"$ 20"}
                      </span>
                    </div>
                    {/*<!-- End atbd listing meta -->*/}
                    <div className="atbd_listing_data_list">
                      <ul>
                        <li>
                          <p>
                            <span className="la la-map-marker"></span>
                            {marketName}
                          </p>
                        </li>
                      </ul>
                    </div>
                    {/*<!-- End atbd listing meta -->*/}
                  </div>
                  {/*<!-- end .atbd_content_upper -->*/}
                  <div className="atbd_listing_bottom_content">
                    <div className="atbd_content_left">
                      <div className="atbd_listing_category">
                        <a href=" ">Technology</a>
                      </div>
                    </div>
                  </div>
                  {/*<!-- end .atbd_listing_bottom_content -->*/}
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
