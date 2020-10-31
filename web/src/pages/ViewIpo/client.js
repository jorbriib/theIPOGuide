import axios from "axios";

import { API_URL } from "../../../config";

export const fetchIPO = async (alias) => {
  try {
    const response = await axios.get(`${API_URL}/ipos/${alias}`);
    return {
      ipo: Ipo(response.data),
    };
  } catch (error) {
    const status = error.response && error.response.status;
    return { error: status || error.message };
  }
};

export function Ipo(ipo) {
  return {
    alias: ipo.alias,
    intro: ipo.intro,
    companySymbol: ipo.company.symbol,
    companyName: ipo.company.name,
    companyDescription: ipo.company.description,
    companySector: ipo.company.sector,
    companyIndustry: ipo.company.industry,
    companyAddress: ipo.company.address,
    companyCountry: ipo.company.country,
    companyLogo: ipo.company.logo,
    priceFrom: ipo.priceFrom,
    priceTo: ipo.priceTo,
    shares: ipo.shares,
    expectedDate: new Date(ipo.expectedDate),
    marketName: ipo.market.name,
    marketCurrency: ipo.market.currency,
    companyPhone: ipo.company.phone,
    companyEmail: ipo.company.email,
    companyWeb: ipo.company.website,
    companyFacebook: ipo.company.facebook,
    companyTwitter: ipo.company.twitter,
    companyLinkedin: ipo.company.linkedin,
    companyPinterest: ipo.company.pinterest,
    companyInstagram: ipo.company.instagram,
    companyEmployees: ipo.company.employees,
    companyFounded: ipo.company.founded,
    companyCeo: ipo.company.ceo,
    companyFiscalYearEnd: ipo.company.fiscalYearEnd,
    ipoUrl: ipo.company.ipoUrl,
    exchangeCommissionUrl: ipo.company.exchangeCommissionUrl,
  };
}
