import { useEffect, useState } from "react";

import * as Client from "./client";

export function useListIpos(params, client = Client) {
  const [state, setState] = useState({ status: "idle", ipos: [], total: 0 });
  const [relations, setRelations] = useState(params);

  useEffect(() => {
    async function getIpos() {
      const { error, ipos = [], total = 0 } = await client.fetchIPOs(relations);
      if (error) {
        setState((prevState) => ({ ...prevState, status: "ready" }));
        return;
      }
      setState({ status: "ready", ipos, total });
    }

    getIpos();
  }, [relations, client]);

  return {
    status: state.status,
    ...relations,
    setRelations: setRelations,
    ipos: state.ipos.map(Ipo),
    total: state.total,
  };
}

export function useRelationIpos(client = Client) {
  const [state, setState] = useState({
    status: "idle",
    markets: [],
    countries: [],
    sectors: [],
  });

  useEffect(() => {
    async function getRelationIpos() {
      const {
        error,
        markets = [],
        countries = [],
        sectors = [],
      } = await client.fetchRelationIPOs();
      if (error) {
        setState((prevState) => ({ ...prevState, status: "ready" }));
        return;
      }

      setState({ status: "ready", markets, countries, sectors });
    }

    getRelationIpos();
  }, [client]);

  return {
    status: state.status,
    markets: state.markets.map(Market),
    countries: state.countries.map(Country),
    sectors: state.sectors.map(Sector),
  };
}

export function useSimilarListIpos(alias, client = Client) {
  const [state, setState] = useState({ status: "idle", ipos: [] });

  useEffect(() => {
    async function getSimilarIpos(alias) {
      const { error, ipos = [] } = await client.fetchSimilarIPOs(alias);
      if (error) {
        setState((prevState) => ({ ...prevState, status: "ready" }));
        return;
      }
      setState({ status: "ready", ipos });
    }

    if (alias) {
      getSimilarIpos(alias);
    }
  }, [alias, client]);
  return {
    similarStatus: state.status,
    similarIpos: state.ipos.map(Ipo),
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

export function Market(market) {
  return {
    code: market.code,
    name: market.name,
    currency: market.currency,
  };
}

export function Country(country) {
  return {
    code: country.code,
    name: country.name,
  };
}

export function Sector(sector) {
  return {
    alias: sector.alias,
    name: sector.name,
  };
}
