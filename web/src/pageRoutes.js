const app = {
  home: () => "/",
  ipo: (alias) => `/ipo/${alias}`,
  about: () => "/about",
  contact: () => "/contact",
  privacyPolicy: () => "/privacy-policy",
  termsAndConditions: () => "/terms-and-conditions",
};

export default {
  ...app,
};
