import React, { Fragment } from "react";

const IpoFeature = ({ ipo }) => {
  const {
    companySymbol,
    marketName,
    marketCurrency,
    priceFrom,
    priceTo,
    companyCountry,
    companySector,
    companyIndustry,
    expectedDate,
    shares,
    companyEmployees,
    companyFounded,
    companyFiscalYearEnd,
    exchangeCommissionUrl,
    ipoUrl,
  } = ipo;

  return (
    <Fragment>
      <div className="atbd_content_module atbd_listing_features">
        <div className="atbd_content_module__tittle_area">
          <div className="atbd_area_title">
            <h4>
              <span className="la la-list-alt"></span>IPO details
            </h4>
          </div>
        </div>
        <div className="atbdb_content_module_contents">
          <ul className="atbd_custom_fields features-table">
            <li>
              <div className="atbd_custom_field_title">
                <p>Symbol: </p>
              </div>
              <div className="atbd_custom_field_content">
                <p>{companySymbol}</p>
              </div>
            </li>
            <li>
              <div className="atbd_custom_field_title">
                <p>Market: </p>
              </div>
              <div className="atbd_custom_field_content">
                <p>{marketName}</p>
              </div>
            </li>
            <li>
              <div className="atbd_custom_field_title">
                <p>Price ranges: </p>
              </div>
              <div className="atbd_custom_field_content">
                <p>
                  {priceFrom} {priceTo && <span>- {priceTo}</span>}{" "}
                  {marketCurrency && <span>({marketCurrency})</span>}
                </p>
              </div>
            </li>
            <li>
              <div className="atbd_custom_field_title">
                <p>Expected date: </p>
              </div>
              <div className="atbd_custom_field_content">
                <p>
                  {new Intl.DateTimeFormat("en-GB", {
                    year: "numeric",
                    month: "long",
                    day: "2-digit",
                  }).format(expectedDate)}
                </p>
              </div>
            </li>
            <li>
              <div className="atbd_custom_field_title">
                <p>Shares offered: </p>
              </div>
              <div className="atbd_custom_field_content">
                <p>{shares ? shares.toLocaleString() : "-"}</p>
              </div>
            </li>
            <li>
              <div className="atbd_custom_field_title">
                <p>Company country: </p>
              </div>
              <div className="atbd_custom_field_content">
                <p>{companyCountry}</p>
              </div>
            </li>
            <li>
              <div className="atbd_custom_field_title">
                <p>Sector: </p>
              </div>
              <div className="atbd_custom_field_content">
                <p>{companySector ? companySector : "-"}</p>
              </div>
            </li>
            <li>
              <div className="atbd_custom_field_title">
                <p>Industry: </p>
              </div>
              <div className="atbd_custom_field_content">
                <p>{companyIndustry ? companyIndustry : "-"}</p>
              </div>
            </li>

            <li>
              <div className="atbd_custom_field_title">
                <p>Employees: </p>
              </div>
              <div className="atbd_custom_field_content">
                <p>
                  {companyEmployees ? companyEmployees.toLocaleString() : "-"}
                </p>
              </div>
            </li>
            <li>
              <div className="atbd_custom_field_title">
                <p>Company founded: </p>
              </div>
              <div className="atbd_custom_field_content">
                <p>{companyFounded ? companyFounded : "-"}</p>
              </div>
            </li>
            <li>
              <div className="atbd_custom_field_title">
                <p>Fiscal Year End: </p>
              </div>
              <div className="atbd_custom_field_content">
                <p>{companyFiscalYearEnd ? companyFiscalYearEnd : "-"}</p>
              </div>
            </li>
            <li>
              <div className="atbd_custom_field_title">
                <p>Exchange commission: </p>
              </div>
              <div className="atbd_custom_field_content">
                <p>
                  {exchangeCommissionUrl ? (
                    <a href={exchangeCommissionUrl} target="_blank">
                      {
                        new URL(exchangeCommissionUrl.replace("www.", ""))
                          .hostname
                      }
                    </a>
                  ) : (
                    "-"
                  )}
                </p>
              </div>
            </li>
            {ipoUrl && (
              <li>
                <div className="atbd_custom_field_title">
                  <p>Extended info: </p>
                </div>
                <div className="atbd_custom_field_content">
                  <p>
                    <a href={ipoUrl} target="_blank">
                      {new URL(ipoUrl.replace("www.", "")).hostname}
                    </a>
                  </p>
                </div>
              </li>
            )}
          </ul>
        </div>
      </div>
    </Fragment>
  );
};
export default IpoFeature;
