import React, { Fragment } from "react";

import Header from "../../components/layout/Header";
import Footer from "../../components/layout/Footer";
import { AdvSearch } from "../../components/elements/AdvanceSearch/AdvanceSearch";
import SectionTitle from "../../components/elements/SectionTitle";
import CardCategoryGrid4 from "../../components/elements/CardCategoryGrid4";
import ContentHome from "../../components/elements/ContentHome";
import PlaceList from "../../components/elements/PlaceList";

const Home = () => {
  const backgroundImage = {
    backgroundImage: "url('/assets/images/wallstreet-bull.jpg')",
    opacity: 1,
  };
  return (
    <Fragment>
      {/* Header section start */}
      <section className="intro-wrapper bgimage overlay overlay--dark">
        <div className="bg_image_holder" style={backgroundImage}>
          <img
            src="/assets/images/wallstreet-bull.jpg"
            alt="Find your IPO and invest"
          />
        </div>
        <div className="mainmenu-wrapper">
          <Header class="menu--light" />
        </div>
        {/* <!-- ends: .mainmenu-wrapper --> */}
        <AdvSearch />
      </section>
      {/* Header section end */}

      {/* Category section start */}
      <section className="categories-cards section-padding-two">
        <div className="container">
          <SectionTitle
            title="What Kind of Activity do you Want to try?"
            content="Discover best things to do restaurants, shopping, hotels, cafes and places around the world by categories."
          />
          <div className="row">
            <CardCategoryGrid4 />
          </div>
        </div>
      </section>
      {/* Category section end */}

      <ContentHome />

      {/* Place section start */}
      <section className="places section-padding">
        <div className="container">
          <div className="row">
            <PlaceList />
          </div>
        </div>
      </section>
      {/* Place section end */}

      <Footer />
    </Fragment>
  );
};

export default Home;
