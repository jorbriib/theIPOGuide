import React, { Fragment } from "react";
import { NavLink } from "react-router-dom";

import pageRoutes from "../../../pageRoutes";

const NavItem = () => {
  return (
    <Fragment>
      <ul className="navbar-nav">
        <li>
          <NavLink to={pageRoutes.home()}>Home</NavLink>
        </li>
        <li>
          <NavLink to={pageRoutes.home()}>Find IPOs</NavLink>
        </li>
        <li>
          <NavLink to={pageRoutes.about()}>About</NavLink>
        </li>
        <li>
          <NavLink to={pageRoutes.contact()}>Contact</NavLink>
        </li>
      </ul>
    </Fragment>
  );
};

export default NavItem;
