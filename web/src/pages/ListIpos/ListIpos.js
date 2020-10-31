import React, { Fragment } from "react";
import { withRouter } from "react-router-dom";
import queryString from "query-string";

import Header from "../../components/layout/Header";
import Footer from "../../components/layout/Footer";
import List from "../../components/container/List";
import { useListIpos } from "./useListIpos";
import { BreadcrumbWrapper } from "../../components/elements/Breadcrumbs";
import pageRoutes from "../../pageRoutes";
import CustomHeader from "../../components/elements/CustomHeader";

const ListIpos = ({ history }) => {
  const params = queryString.parse(history.location.search);

  let selectedMarkets = {};
  if (params.markets !== undefined) {
    params.markets.split(",").forEach(function (value) {
      selectedMarkets[value] = value;
    });
  }

  let selectedCountries = {};
  if (params.countries !== undefined) {
    params.countries.split(",").forEach(function (value) {
      selectedCountries[value] = value;
    });
  }

  let selectedSectors = {};
  if (params.sectors !== undefined) {
    params.sectors.split(",").forEach(function (value) {
      selectedSectors[value] = value;
    });
  }

  let page = 0;
  if (params.page !== undefined) {
    page = parseInt(params.page);
  }

  const {
    status,
    total,
    markets,
    countries,
    sectors,
    setRelations,
    ipos,
  } = useListIpos({
    markets: params.markets ? selectedMarkets : {},
    countries: params.countries ? selectedCountries : {},
    sectors: params.sectors ? selectedSectors : {},
    page: page,
  });

  const backgroundImage = {
    backgroundImage: "url('/assets/images/wallstreet-bull.jpg')",
    opacity: 1,
  };

  const onSubmit = (type, value, isActived) => {
    const currentUrlParams = new URLSearchParams(window.location.search);
    switch (type) {
      case "markets":
        if (isActived) {
          markets[value] = value;
        } else {
          delete markets[value];
        }
        break;
      case "countries":
        if (isActived) {
          countries[value] = value;
        } else {
          delete countries[value];
        }
        break;
      case "sectors":
        if (isActived) {
          sectors[value] = value;
        } else {
          delete sectors[value];
        }
        break;
    }
    page = 0;
    setRelations({ markets, countries, sectors, page });

    Object.values(markets).length
      ? currentUrlParams.set("markets", Object.values(markets).join(","))
      : currentUrlParams.delete("markets");
    Object.values(countries).length
      ? currentUrlParams.set("countries", Object.values(countries).join(","))
      : currentUrlParams.delete("countries");
    Object.values(sectors).length
      ? currentUrlParams.set("sectors", Object.values(sectors).join(","))
      : currentUrlParams.delete("sectors");
    currentUrlParams.set("page", page);
    history.push(window.location.pathname + "?" + currentUrlParams.toString());
  };

  const onPage = (page) => {
    setRelations({ markets, countries, sectors, page });

    const currentUrlParams = new URLSearchParams(window.location.search);
    currentUrlParams.set("page", page);

    history.push(window.location.pathname + "?" + currentUrlParams.toString());
  };

  if (status !== "ready") {
    return "";
  }

  return (
    <Fragment>
      <CustomHeader
        title="Find your IPO"
        description="Find your IPO and invest. The most important IPOs from the most important countries and markets."
        url={pageRoutes.ipos()}
      />
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
        <BreadcrumbWrapper title="Find your IPO and invest" onlyTitle={true} />
      </section>
      {/* Header section end */}

      <List
        list={ipos}
        onSubmit={onSubmit}
        total={total}
        page={page}
        onPage={onPage}
      />
      <Footer />
    </Fragment>
  );
};

export default withRouter(ListIpos);
