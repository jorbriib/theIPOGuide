import React, { Fragment } from "react";
import { NavLink } from "react-router-dom";

import NavItem from "../NavItem";
import pageRoutes from "../../../pageRoutes";

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
                    <div className="search-wrapper">
                      <div className="nav_right_module search_module">
                        <span className="icon-left" id="basic-addon9">
                          <i className="la la-search"></i>
                        </span>
                        <div className="search_area">
                          <form action="/">
                            <div className="input-group input-group-light">
                              <input
                                type="text"
                                className="form-control search_field top-search-field"
                                placeholder="What are you looking for?"
                                autoComplete="off"
                              />
                            </div>
                          </form>
                        </div>
                      </div>
                      <div className="search-categories">
                        <ul className="list-unstyled">
                          <li>
                            <a href="/">
                              <span className="la la-glass bg-danger"></span>{" "}
                              Find IPOs
                            </a>
                          </li>
                          <li>
                            <a href="/login">
                              <span className="la la-glass bg-danger"></span>{" "}
                              Login
                            </a>
                          </li>
                        </ul>
                      </div>
                    </div>
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
