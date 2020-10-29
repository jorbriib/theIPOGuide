import React, { Fragment } from "react";

import pageRoutes from "../../../pageRoutes";

export function ContentBlock1About() {
  return (
    <Fragment>
      <section className="about-contents section-padding">
        <div className="container">
          <div className="row">
            <div className="col-lg-12 contents-wrapper">
              <div className="contents">
                <div className="row align-items-center">
                  <div className="col-lg-5 col-sm-6">
                    <img
                      src="./assets/images/about-img1.jpg"
                      alt="About theIPOguide"
                    />
                  </div>
                  <div className="col-lg-6 offset-lg-1 col-sm-6 mt-5 mt-md-0">
                    <h1>
                      About the<span>IPO</span>guide
                    </h1>
                    <p>
                      The IPO guide is a new project that aims to show the IPOs
                      of the most popular markets in a simple and unified way.
                    </p>
                    <p>
                      Now we are small, but we intend to continue growing until
                      we are the worldwide reference guide for companies that
                      are going public.
                    </p>
                  </div>
                </div>
              </div>
              {/*<!-- ends: .contents -->*/}
            </div>
            {/*<!-- ends: .content-block -->*/}
          </div>
        </div>
      </section>
    </Fragment>
  );
}

export function ContentBlock2About() {
  return (
    <Fragment>
      <section className="about-contents why-wrapper section-padding">
        <div className="container">
          <div className="row">
            <div className="col-lg-12 contents-wrapper">
              <div className="contents">
                <div className="row align-items-center">
                  <div className="col-lg-5 col-sm-6">
                    <h1>
                      Why the<span>IPO</span>guide
                    </h1>
                    <ul className="list-unstyled list-features p-top-15">
                      <li>
                        <div className="list-count">
                          <span>1</span>
                        </div>
                        <div className="list-content">
                          <h4>Easy to use</h4>
                          <p>
                            Find companies by market, country, sector, name...
                            in a simple click.
                          </p>
                        </div>
                      </li>
                      <li>
                        <div className="list-count">
                          <span>2</span>
                        </div>
                        <div className="list-content">
                          <h4>USA, Europe, Canada, Japan, China...</h4>
                          <p>
                            You don't need to investigate every market, we have
                            the most important markets here.
                          </p>
                        </div>
                      </li>
                      <li>
                        <div className="list-count">
                          <span>3</span>
                        </div>
                        <div className="list-content">
                          <h4>Expanding community</h4>
                          <p>
                            the<span>IPO</span>guide is a personal project where
                            you can decide the next features. Contact with us{" "}
                            <a href={pageRoutes.contact()}>here</a>.
                          </p>
                        </div>
                      </li>
                    </ul>
                  </div>
                  <div className="col-lg-6 offset-lg-1 text-right col-sm-6 mt-5 mt-md-0">
                    <img
                      src="./assets/images/about-img2.jpg"
                      alt="Why theIPOguide"
                    />
                  </div>
                </div>
              </div>
              {/*<!-- ends: .contents -->*/}
            </div>
            {/*<!-- ends: .content-block -->*/}
          </div>
        </div>
      </section>
    </Fragment>
  );
}

export function CounterAbout() {
  return (
    <Fragment>
      <section className="counters-wrapper bg-gradient-pw section-padding">
        <div className="container">
          <div className="row">
            <div className="col-lg-12 text-center">
              <h1>Tens of IPOs</h1>
              <p>turn to theIPOguide every day to make spending decisions</p>
              <ul className="list-unstyled counter-items">
                <li>
                  <p>
                    <span className="count_up">50</span>+
                  </p>
                  <span>Companies</span>
                </li>
                <li>
                  <p>
                    <span className="count_up">6</span>+
                  </p>
                  <span>Markets</span>
                </li>
                <li>
                  <p>
                    <span className="count_up">8</span>+
                  </p>
                  <span>Countries</span>
                </li>
                <li>
                  <p>
                    <span className="count_up">12</span>+
                  </p>
                  <span>Sectors</span>
                </li>
              </ul>
            </div>
          </div>
        </div>
      </section>
    </Fragment>
  );
}
