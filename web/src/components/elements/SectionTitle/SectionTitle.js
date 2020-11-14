import React, { Fragment } from "react";

const SectionTitle = ({title, content}) => {
  return (
    <Fragment>
      <div className="row">
        <div className="col-lg-12">
          <div className="section-title">
            <h2>{title}</h2>
            <p>{content}</p>
          </div>
        </div>
      </div>
    </Fragment>
  );
};

export default SectionTitle;
