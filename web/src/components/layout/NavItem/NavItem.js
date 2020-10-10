import React, { Fragment } from 'react';
import { NavLink } from 'react-router-dom';

const NavItem = () => {
    return (
        <Fragment>
            <ul className="navbar-nav">
                <li>
                    <NavLink to="/">Find IPOs</NavLink>
                </li>
            </ul>
        </Fragment>
    )
}

export default NavItem;