const app = {
  home: () => "/",
  ipo: (alias) => `/ipo/${alias}`,
  about: () => "/about",
  contact: () => "/contact",
  privacyPolicy: () => "/privacy-policy",
  termsAndConditions: () => "/terms-and-conditions",
  disclaimer: () => "/terms-and-conditions#disclaimer",
};

export default {
  ...app,
};
