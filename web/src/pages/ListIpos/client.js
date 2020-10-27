import axios from "axios";

import { API_URL } from "../../../config";

export const fetchIPOs = async (relations) => {
  try {
    const response = await axios.get(`${API_URL}/ipos`, {
      params: {
        markets: Object.values(relations.markets).join(","),
        countries: Object.values(relations.countries).join(","),
        sectors: Object.values(relations.sectors).join(","),
        page: relations.page,
      },
    });
    return {
      total: response.data.total,
      ipos: response.data.list.map(Ipo),
    };
  } catch (error) {
    const status = error.response && error.response.status;
    return { error: status || error.message };
  }
};

export const fetchRelationIPOs = async () => {
  try {
    const response = await axios.get(`${API_URL}/ipos/relations`);
    return {
      markets: response.data.markets.map(Market),
      countries: response.data.countries.map(Country),
      sectors: response.data.sectors.map(Sector),
    };
  } catch (error) {
    const status = error.response && error.response.status;
    return { error: status || error.message };
  }
};

export const fetchSimilarIPOs = async (alias) => {
  try {
    const response = await axios.get(`${API_URL}/ipos/${alias}/similar`);
    return {
      ipos: response.data.map(Ipo),
    };
  } catch (error) {
    const status = error.response && error.response.status;
    return { error: status || error.message };
  }
};

export function Ipo(ipo) {
  return {
    alias: ipo.alias,
    companySymbol: ipo.company.symbol,
    companyName: ipo.company.name,
    companySector: ipo.company.sector,
    companyCountry: ipo.company.country,
    companyLogo: ipo.company.logo,
    priceFrom: ipo.priceFrom,
    priceTo: ipo.priceTo,
    expectedDate: new Date(ipo.expectedDate),
    marketName: ipo.market.name,
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
