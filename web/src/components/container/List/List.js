import React, { Fragment } from "react";

import CardList from "../../elements/CardList/CardList";
import Sidebar from "../../Sidebar";
import Pagination from "../../elements/Pagination";

const List = ({ list, onSubmit, total, page, onPage }) => {
  return (
    <Fragment>
      <section className="all-listing-wrapper section-bg">
        <div className="container">
          <div className="row">
            <div className="col-lg-12">
              <div className="atbd_generic_header">
                <div className="atbd_generic_header_title">
                  <h4>All Items</h4>
                  <p>Total IPOs Found: {total}</p>
                </div>
                <div
                  className="atbd_listing_action_btn btn-toolbar"
                  role="toolbar"
                >
                  <div className="dropdown">
                    <a
                      className="action-btn dropdown-toggle"
                      href=" "
                      role="button"
                      id="dropdownMenuLink2"
                      data-toggle="dropdown"
                      aria-haspopup="true"
                      aria-expanded="false"
                    >
                      Sort by <span className="caret"></span>
                    </a>
                    <div
                      className="dropdown-menu"
                      aria-labelledby="dropdownMenuLink2"
                    >
                      <a
                        className="dropdown-item"
                        href="#"
                        onClick={(e) => {
                          e.preventDefault();
                          return false;
                        }}
                      >
                        Latest IPOs
                      </a>
                    </div>
                  </div>
                </div>
              </div>
            </div>{" "}
            {/*<!-- ends: .col-lg-12 -->*/}
            <div className="col-lg-12 listing-items">
              <div className="row">
                <div className="col-lg-4 order-lg-0 order-1 mt-5 mt-lg-0">
                  <Sidebar onSubmit={onSubmit} />
                </div>{" "}
                {/* wiget */}
                <div className="col-lg-8 order-lg-1 order-0">
                  <div className="row">
                    {Object.values(list).length ? (
                      <CardList list={list} />
                    ) : (
                      <div className="col-lg-12">
                        <div className="alert alert-warning" role="alert">
                          IPOs not found!
                        </div>
                      </div>
                    )}
                  </div>
                  <div className="row">
                    <div className="col-lg-12">
                      <Pagination total={total} page={page} onPage={onPage} />
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </section>
    </Fragment>
  );
};

export default List;
