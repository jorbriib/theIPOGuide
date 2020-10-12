import { useEffect, useState } from "react";

import * as Client from "./client";

export default function useListIpos(client = Client) {
  const [state, setState] = useState({ status: "idle", ipos: [] });

  useEffect(() => {
    async function getIpos() {
      const { error, ipos = [] } = await client.fetchIPOs();
      if (error) {
        setState((prevState) => ({ ...prevState, status: "ready" }));
        return;
      }

      setState({ status: "ready", ipos });
    }
    getIpos();
  }, [client]);

  return {
    status: state.status,
    ipos: state.ipos.map(Ipo),
  };
}

function Ipo(ipo) {
  return {
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
