import React, { Fragment } from "react";
import { NavLink } from "react-router-dom";

import NavItem from "../NavItem";
import pageRoutes from "../../../pageRoutes";
import SearchByText from "../../elements/SearchByText/SearchByText";

const Header = () => {
  return (
    <Fragment>
      <div className={"menu-area menu1 menu--light"}>
        <div className="top-menu-area">
          <div className="container-fluid">
            <div className="row">
              <div className="col-lg-12">
                <div className="menu-fullwidth">
                  <div className="logo-wrapper order-lg-0 order-sm-1">
                    <div className="logo logo-top">
                      <NavLink to={pageRoutes.home()}>
                        <img
                          src="/assets/images/logo.png"
                          alt="logoImage"
                          className="img-fluid"
                        />
                      </NavLink>
                    </div>
                  </div>
                  {/*<!-- ends: .logo-wrapper -->*/}
                  <div className="menu-container order-lg-1 order-sm-0">
                    <div className="d_menu">
                      <nav className="navbar navbar-expand-lg mainmenu__menu">
                        <button
                          className="navbar-toggler"
                          type="button"
                          data-toggle="collapse"
                          data-target="#direo-navbar-collapse"
                          aria-controls="direo-navbar-collapse"
                          aria-expanded="false"
                          aria-label="Toggle navigation"
                        >
                          <span className="navbar-toggler-icon icon-menu">
                            <i className="la la-reorder"></i>
                          </span>
                        </button>
                        {/*<!-- Collect the nav links, forms, and other content for toggling -->*/}
                        <div
                          className="collapse navbar-collapse"
                          id="direo-navbar-collapse"
                        >
                          <NavItem />
                        </div>
                        {/*<!-- /.navbar-collapse -->*/}
                      </nav>
                    </div>
                  </div>
                  <div className="menu-right order-lg-2 order-sm-2">
                    <SearchByText />
                    {/*<!-- ends: .search-wrapper -->*/}
                  </div>
                  {/*<!-- ends: .menu-right -->*/}
                </div>
              </div>
            </div>
            {/* <!-- end /.row --> */}
          </div>
          {/* <!-- end /.container --> */}
        </div>
        {/* <!-- end  --> */}
      </div>
    </Fragment>
  );
};

export default Header;
