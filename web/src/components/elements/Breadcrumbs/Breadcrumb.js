import React, { Fragment } from "react";
import { NavLink } from "react-router-dom";

const BreadcrumbWraper = ({ title, onlyTitle }) => {
  return (
    <Fragment>
      <div className="breadcrumb-wrapper content_above">
        <div className="container">
          <div className="row">
            <div className="col-lg-12 text-center">
              <h1 className="page-title">{title}</h1>
              {onlyTitle ? null : (
                <nav aria-label="breadcrumb">
                  <ol className="breadcrumb">
                    <li className="breadcrumb-item">
                      <NavLink to="/">Home</NavLink>
                    </li>
                    <li className="breadcrumb-item active" aria-current="page">
                      {title}
                    </li>
                  </ol>
                </nav>
              )}
            </div>
          </div>
        </div>
      </div>
    </Fragment>
  );
};

export default BreadcrumbWraper;
