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
    companyName: ipo.companyName,
    marketName: ipo.marketName,
  };
}
