import axios from "axios";

import { API_URL } from "../../../config";

export const sendContactMessage = async ({ name, email, message, token }) => {
  try {
    const data = new URLSearchParams();
    data.append("name", name);
    data.append("email", email);
    data.append("message", message);
    data.append("token", token);

    const options = {
      headers: { "content-type": "application/x-www-form-urlencoded" },
    };

    await axios.post(`${API_URL}/contact`, data, options);
  } catch (error) {
    const status = error.response && error.response.status;
    const message =
      error.response && error.response.data && error.response.data.message;
    return { error: status, message: message };
  }
};
