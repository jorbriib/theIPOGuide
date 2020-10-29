import React, { useState } from "react";
import { NavLink } from "react-router-dom";

import pageRoutes from "../../../pageRoutes";
import { useSearchByText } from "./useSearchByText";

const SearchByText = () => {
  const [text, setText] = useState("");
  const [focused, setFocused] = useState(false);

  const { ipos } = useSearchByText(text);

  const handleSubmit = (event) => {
    event.preventDefault();
    setText(event.target.value);
  };

  const onBlur = (event) => {
    event.preventDefault();
    setFocused(false);
  };

  const onFocus = (event) => {
    event.preventDefault();
    setFocused(true);
  };

  let searchModuleClasses = "nav_right_module search_module";
  if (focused) {
    searchModuleClasses = searchModuleClasses + " active";
  }

  return (
    <div className="search-wrapper">
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
              <NavLink to="#">We don't have results for that input :( </NavLink>
            </li>
          )}
        </ul>
      </div>
    </div>
  );
};

export default SearchByText;
