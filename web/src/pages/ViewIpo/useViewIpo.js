import { useEffect, useState } from "react";

import * as Client from "./client";

export default function useViewIpo(alias, client = Client) {
  const [state, setState] = useState({ status: "idle", ipo: null });

  useEffect(() => {
    async function getIpo(alias) {
      const { error, ipo = {} } = await client.fetchIPO(alias);
      if (error) {
        setState((prevState) => ({ ...prevState, status: "ready" }));
        return;
      }

      setState({ status: "ready", ipo });
    }

    if (alias) {
      getIpo(alias);
    }
  }, [alias, client]);

  return {
    status: state.status,
    ipo: state.ipo ? Ipo(state.ipo) : {},
  };
}

function Ipo(ipo) {
  return {
    alias: ipo.alias,
    companySymbol: ipo.companySymbol,
    companyName: ipo.companyName,
    companyDescription: ipo.companyDescription,
    companySector: ipo.companySector,
    companyIndustry: ipo.companyIndustry,
    companyAddress: ipo.companyAddress,
    companyCountry: ipo.companyCountry,
    companyLogo: ipo.companyLogo,
    marketName: ipo.marketName,
    priceFrom: ipo.priceFrom,
    priceTo: ipo.priceTo,
    shares: ipo.shares,
    expectedDate: ipo.expectedDate,
    companyPhone: ipo.companyPhone,
    companyEmail: ipo.companyEmail,
    companyWeb: ipo.companyWeb,
    companyFacebook: ipo.companyFacebook,
    companyTwitter: ipo.companyTwitter,
    companyLinkedin: ipo.companyLinkedin,
    companyPinterest: ipo.companyPinterest,
    companyInstagram: ipo.companyInstagram,
    companyEmployees: ipo.companyEmployees,
    companyFounded: ipo.companyFounded,
    companyCeo: ipo.companyCeo,
    companyFiscalYearEnd: ipo.companyFiscalYearEnd,
    ipoUrl: ipo.ipoUrl,
    exchangeCommissionUrl: ipo.exchangeCommissionUrl,
  };
}
