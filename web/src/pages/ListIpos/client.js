import axios from "axios";

import { API_URL } from "../../../config";

export const fetchIPOs = async () => {
  try {
    const response = await axios.get(`${API_URL}/ipos`);
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
