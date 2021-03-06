import React from "react";

import IpoDetails from "../../content/Details/Details";
import IpoFeature from "../../content/Features/Features";
import { ContactIpo, SimilarIpos } from "../../elements/Widget";

const View = ({ ipo }) => {
  return (
    <section className="directory_listiing_detail_area single_area section-bg section-padding-strict">
      <div className="container">
        <div className="row">
          <div className="col-lg-8">
            <IpoDetails ipo={ipo} />

            <IpoFeature ipo={ipo} />
            <small>
              theIPOguide doesn't recommend any type of purchase or sale of
              shares or similar, all data provided in this website is only for
              informational purposes.
            </small>
          </div>
          <div className="col-lg-4">
            <ContactIpo ipo={ipo} />

            <SimilarIpos ipo={ipo} />
          </div>
        </div>
      </div>
    </section>
  );
};

export default View;
