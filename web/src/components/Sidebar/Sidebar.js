import React from "react";
import queryString from "query-string";
import { withRouter } from "react-router-dom";

import { useRelationIpos } from "../../pages/ListIpos/useListIpos";

const Sidebar = ({ onSubmit, history }) => {
  const { status, markets, countries, sectors } = useRelationIpos();
  const params = queryString.parse(history.location.search);
  const selectedMarkets = params.markets ? params.markets.split(",") : [];
  const selectedCountries = params.countries ? params.countries.split(",") : [];
  const selectedSectors = params.sectors ? params.sectors.split(",") : [];

  if (status !== "ready") {
    return "";
  }

  return (
    <div className="listings-sidebar">
      <div className="search-area default-ad-search">
        <form action="#">
          {Object.values(markets).length > 0 && (
            <div className="filter-checklist">
              <h5>Filter by Market</h5>
              <div className="checklist-items feature-checklist">
                {Object.values(markets).map((market, key) => {
                  return (
                    <div
                      className="custom-control custom-checkbox checkbox-outline checkbox-outline-primary"
                      key={key}
                    >
                      <input
                        type="checkbox"
                        className="custom-control-input"
                        id={market.code}
                        defaultChecked={selectedMarkets.includes(market.code)}
                        onChange={(event) => {
                          onSubmit(
                            "markets",
                            market.code,
                            event.target.checked
                          );
                        }}
                      />
                      <label
                        className="custom-control-label"
                        htmlFor={market.code}
                      >
                        {market.name}
                      </label>
                    </div>
                  );
                })}
              </div>
            </div>
          )}

          {Object.values(countries).length > 0 && (
            <div className="filter-checklist">
              <h5>Filter by Company Country</h5>
              <div className="checklist-items feature-checklist">
                {Object.values(countries).map((country, key) => {
                  return (
                    <div
                      className="custom-control custom-checkbox checkbox-outline checkbox-outline-primary"
                      key={key}
                    >
                      <input
                        type="checkbox"
                        className="custom-control-input"
                        id={country.code}
                        defaultChecked={selectedCountries.includes(
                          country.code
                        )}
                        onChange={(event) => {
                          onSubmit(
                            "countries",
                            country.code,
                            event.target.checked
                          );
                        }}
                      />
                      <label
                        className="custom-control-label"
                        htmlFor={country.code}
                      >
                        {country.name}
                      </label>
                    </div>
                  );
                })}
              </div>
            </div>
          )}

          {Object.values(sectors).length > 0 && (
            <div className="filter-checklist">
              <h5>Filter by Sector</h5>
              <div className="checklist-items feature-checklist">
                {Object.values(sectors).map((sector, key) => {
                  return (
                    <div
                      className="custom-control custom-checkbox checkbox-outline checkbox-outline-primary"
                      key={key}
                    >
                      <input
                        type="checkbox"
                        className="custom-control-input"
                        id={sector.alias}
                        defaultChecked={selectedSectors.includes(sector.alias)}
                        onChange={(event) => {
                          onSubmit(
                            "sectors",
                            sector.alias,
                            event.target.checked
                          );
                        }}
                      />
                      <label
                        className="custom-control-label"
                        htmlFor={sector.alias}
                      >
                        {sector.name}
                      </label>
                    </div>
                  );
                })}
              </div>
            </div>
          )}
        </form>
      </div>
    </div>
  );
};

export default withRouter(Sidebar);
