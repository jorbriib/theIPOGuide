import axios from "axios";

import { API_URL } from "../../../config";

export const searchByText = async (text) => {
  try {
    const response = await axios.get(`${API_URL}/ipos/search`, {
      params: {
        text,
      },
    });
    return {
      ipos: response.data.list.map(Ipo),
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
