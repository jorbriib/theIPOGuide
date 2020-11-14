import React, { Fragment } from "react";
import { NavLink } from "react-router-dom";

export function AdvSearch() {
  return (
    <Fragment>
      <div className="directory_content_area">
        <div className="container">
          <div className="row">
            <div className="col-lg-10 offset-lg-1">
              <div className="search_title_area">
                <h2 className="title">Find the Best Places to Be</h2>
                <p className="sub_title">
                  All the top IPOs – from different countries, markets and
                  sectors..
                </p>
              </div>
              {/* ends: .search_title_area */}
              <form action="/" className="search_form">
                <div className="atbd_seach_fields_wrapper">
                  <div className="single_search_field search_category">
                    <select className="search_fields" id="at_biz_dir-category">
                      <option value>Select a category</option>
                      <option value="automobile">Automobile</option>
                      <option value="education">Education</option>
                      <option value="event">Event</option>
                    </select>
                  </div>
                  <div className="single_search_field search_location">
                    <select className="search_fields" id="at_biz_dir-location">
                      <option value>Select a location</option>
                      <option value="ab">AB Simple</option>
                      <option value="australia">Australia</option>
                      <option value="australia-australia">Australia</option>
                    </select>
                  </div>
                  <div className="atbd_submit_btn">
                    <button
                      type="submit"
                      className="btn btn-block btn-gradient btn-gradient-one btn-md btn_search"
                    >
                      Search
                    </button>
                  </div>
                </div>
              </form>
            </div>
            {/* ends: .col-lg-10 */}
          </div>
        </div>
      </div>
      {/* ends: .directory_search_area */}
    </Fragment>
  );
}
