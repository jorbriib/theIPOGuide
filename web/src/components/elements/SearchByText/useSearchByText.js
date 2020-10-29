import { useEffect, useState } from "react";

import * as Client from "./client";

export function useSearchByText(text, client = Client) {
  const [state, setState] = useState({ status: "idle", ipos: [] });

  useEffect(() => {
    async function searchByText() {
      if (text.length < 3) {
        setState({ status: "ready", ipos: [] });
        return;
      }

      const { error, ipos = [] } = await client.searchByText(text);
      if (error) {
        setState((prevState) => ({ ...prevState, status: "ready" }));
        return;
      }
      setState({ status: "ready", ipos });
    }

    searchByText();
  }, [text, client]);

  return {
    status: state.status,
    ipos: state.ipos.map(Ipo),
  };
}

function Ipo(ipo) {
  return {
    alias: ipo.alias,
    companySymbol: ipo.companySymbol,
    companyName: ipo.companyName,
    companySector: ipo.companySector,
    companyCountry: ipo.companyCountry,
    companyLogo: ipo.companyLogo,
    marketName: ipo.marketName,
    priceFrom: ipo.priceFrom,
    priceTo: ipo.priceTo,
    expectedDate: ipo.expectedDate,
  };
}
