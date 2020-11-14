import React, { Fragment } from "react";
import { NavLink } from "react-router-dom";

const CardCategoryGrid4 = (category) => {
  return (
    <Fragment>
      {Object.values(category)
        .slice(0, 6)
        .map((value, key) => {
          const { img, totalIpos, category } = value;
          return (
            <div className="col-lg-4 col-sm-6" key={key}>
              <div className="category-single category--img">
                <figure className="category--img4">
                  <img src={img} alt={category} />
                  <figcaption className="overlay-bg">
                    <NavLink to="/at_demo" className="cat-box">
                      <div>
                        <h4 className="cat-name">{category}</h4>
                        <span className="badge badge-pill badge-success">
                          {totalIpos} IPOs
                        </span>
                      </div>
                    </NavLink>
                  </figcaption>
                </figure>
              </div>
              {/*<!-- ends: .category-single -->*/}
            </div>
          );
        })}
    </Fragment>
  );
};

export default CardCategoryGrid4;
