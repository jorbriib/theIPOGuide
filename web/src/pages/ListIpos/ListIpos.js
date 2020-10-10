import React, { Fragment } from "react";

import Header from "../../components/layout/Header";
import Footer from "../../components/layout/Footer";
import BreadcrumbWraper from "../../components/elements/Breadcrumbs";
import List from "../../components/container/List";
import useListIpos from "./useListIpos";

const ListIpos = () => {
  const { status, ipos } = useListIpos();

  const backgroundImage = {
    backgroundImage: "url('/assets/images/wallstreet-bull.jpg')",
    opacity: 1,
  };

  if (status !== "ready") {
    return "";
  }

  return (
    <Fragment>
      {/* Header section start */}
      <section className="header-breadcrumb bgimage overlay overlay--dark">
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
        <BreadcrumbWraper title="Find your IPO and invest" onlyTitle={true} />
      </section>
      {/* Header section end */}

      <List list={ipos} categories={[]} />
      <Footer />
    </Fragment>
  );
};

export default ListIpos;
