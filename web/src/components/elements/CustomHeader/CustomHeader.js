import React from "react";
import { Helmet } from "react-helmet-async";

import { APP_HOST } from "../../../config";

const CustomHeader = ({ title, description, url }) => {
  if (url === undefined) {
    url = window.location.pathname;
  }
  return (
    <Helmet>
      <title>{title} · theIPOguide</title>
      <link rel="sitemap" type="application/xml" href="/sitemap.xml" />
      <meta name="description" content={description} />

      <meta name="name" content={title + " · theIPOguide"} />
      <meta itemprop="description" content={description} />

      <meta
        data-react-helmet="true"
        property="og:site_name"
        content="theIPOguide"
      />
      <meta property="og:url" content={APP_HOST + url} />
      <meta property="og:type" content="website" />
      <meta property="og:title" content={title} />
      <meta property="og:description" content={description} />

      <meta name="twitter:card" content="summary" />
      <meta name="twitter:title" content={title} />
      <meta name="twitter:description" content={description} />
    </Helmet>
  );
};
export default CustomHeader;
