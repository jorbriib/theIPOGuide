import axios from "axios";

import { API_URL } from "../../../../config";

export const sendReport = async ({ url, message, token }) => {
  try {
    const data = new URLSearchParams();
    data.append("url", url);
    data.append("message", message);
    data.append("token", token);

    const options = {
      headers: { "content-type": "application/x-www-form-urlencoded" },
    };

    await axios.post(`${API_URL}/report`, data, options);
  } catch (error) {
    const status = error.response && error.response.status;
    const message =
      error.response && error.response.data && error.response.data.message;

    return { error: status, message: message };
  }
};
