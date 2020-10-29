import React, { useState, Fragment, useRef, useEffect } from "react";
import { NavLink } from "react-router-dom";

import pageRoutes from "../../../pageRoutes";
import { useSearchByText } from "./useSearchByText";

const SearchByText = () => {
  const [text, setText] = useState("");
  const [focused, setFocused] = useState(false);
  const [isSearching, setIsSearching] = useState(false);
  const searchInput = useRef();
  const { ipos } = useSearchByText(text);

  const handleSubmit = (event) => {
    event.preventDefault();
    setText(event.target.value);
  };

  const onBlur = (event) => {
    event.preventDefault();
    setFocused(false);
    setIsSearching(false);
  };

  const onFocus = (event) => {
    event.preventDefault();
    setFocused(true);
  };

  const searchMobileOnClick = (event) => {
    event.preventDefault();
    setIsSearching(true);
    setFocused(true);
  };

  useEffect(() => {
    if (isSearching) {
      searchInput.current.focus();
    }
  }, [isSearching]);

  let searchModuleClasses = "nav_right_module search_module";
  if (focused) {
    searchModuleClasses = searchModuleClasses + " active";
  }

  let searchWrapper = "search-wrapper";
  if (isSearching) {
    searchWrapper = searchWrapper + " float";
  }

  return (
    <Fragment>
      <div className={searchWrapper}>
        <div className={searchModuleClasses}>
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
                  ref={searchInput}
                  onBlur={onBlur}
                  onFocus={onFocus}
                  onChange={handleSubmit}
                />
              </div>
            </form>
          </div>
        </div>
        <div className="search-categories">
          <ul className="list-unstyled">
            {Object.values(ipos)
              .slice(0, 4)
              .map((ipo, key) => {
                return (
                  <li key={key}>
                    <NavLink to={pageRoutes.ipo(ipo.alias)}>
                      <span>
                        <img
                          src={ipo.companyLogo}
                          style={{ width: "30px" }}
                          alt={ipo.companyName}
                        />
                      </span>
                      {ipo.companyName}
                    </NavLink>
                  </li>
                );
              })}
            {text.length > 2 && ipos.length === 0 && (
              <li>
                <NavLink to="#">
                  We don't have results for that input :({" "}
                </NavLink>
              </li>
            )}
          </ul>
        </div>
      </div>
      <div className={"offcanvas-menu d-none"}>
        <a
          href="#"
          onClick={searchMobileOnClick}
          className="offcanvas-menu__search"
        >
          <i className="la la-search"></i>
        </a>
      </div>
    </Fragment>
  );
};

export default SearchByText;
