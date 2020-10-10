import React, { Fragment } from "react";

import CardList from "../../elements/CardList/CardList";

const List = ({ list }) => {
  return (
    <Fragment>
      <section className="all-listing-wrapper section-bg">
        <div className="container">
          <div className="row">
            <div className="col-lg-12">
              <div className="atbd_generic_header">
                <div className="atbd_generic_header_title">
                  <h4>All Items</h4>
                  <p>Total Listing Found: {Object.values(list).length}</p>
                </div>
                {/*<!-- ends: .atbd_generic_header_title -->*/}
              </div>
              {/*<!-- ends: .atbd_generic_header -->*/}
            </div>{" "}
            {/*<!-- ends: .col-lg-12 -->*/}
            <div className="col-lg-12 listing-items">
              <div className="row">
                <div className="col-lg-4 order-lg-0 order-1 mt-5 mt-lg-0">
                  <div className="listings-sidebar">
                    <div className="search-area default-ad-search">
                      <form action="#">
                        <div className="form-group">
                          <input
                            type="text"
                            placeholder="What are you looking for?"
                            className="form-control"
                          />
                        </div>
                        {/*<!-- ends: .form-group -->*/}
                      </form>
                      {/*<!-- ends: form -->*/}
                    </div>
                  </div>
                </div>{" "}
                {/* wiget */}
                <div className="col-lg-8 order-lg-1 order-0">
                  <div className="row">
                    {Object.values(list).length ? (
                      <CardList list={list} />
                    ) : (
                      <div className="col-lg-12">
                        <div className="alert alert-warning" role="alert">
                          Companies not found!
                        </div>
                      </div>
                    )}
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
