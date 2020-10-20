import React, { Fragment } from "react";

const IpoDetails = ({ ipo }) => {
  return (
    <Fragment>
      <div className="atbd_content_module atbd_listing_details">
        <div className="atbd_content_module__tittle_area">
          <div className="atbd_area_title">
            <h4>
              <span className="la la-file-text-o"></span>About {ipo.companyName}
            </h4>
          </div>
        </div>
        <div className="atbdb_content_module_contents">
          {ipo.companyDescription}
        </div>
      </div>
    </Fragment>
  );
};

export default IpoDetails;
