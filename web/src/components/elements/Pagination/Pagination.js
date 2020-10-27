import React from "react";

import constants from "./../../../constants";

const Pagination = ({ total, page, onPage }) => {
  const currentUrlParams = new URLSearchParams(window.location.search);

  const perPage = constants.perPage;
  const numberOfPages = Math.ceil(total / perPage);

  return (
    <div className="project-pagination m-top-30">
      <div className="pagination-area">
        <nav
          aria-label="navigation pagination d-flex justify-content-end"
          role="navigation"
        >
          <ul className="pagination justify-content-center">
            {[...Array(numberOfPages)].map((value, key) => {
              currentUrlParams.set("page", key);
              return (
                key <= numberOfPages - 1 && (
                  <li
                    key={key}
                    className={`page-item click ${key === page && "active"}`}
                  >
                    <a
                      className="page-link"
                      href={
                        window.location.pathname +
                        "?" +
                        currentUrlParams.toString()
                      }
                    >
                      {key + 1}
                    </a>
                  </li>
                )
              );
            })}
          </ul>
        </nav>
      </div>
      {/*<!-- ends: .pagination-wrapper -->*/}
    </div>
  );
};
export default Pagination;
