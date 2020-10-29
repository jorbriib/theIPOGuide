const app = {
  home: () => "/",
  ipo: (alias) => `/ipo/${alias}`,
  contact: () => "/contact",
  privacyPolicy: () => "/privacy-policy",
  termsAndConditions: () => "/terms-and-conditions",
};

export default {
  ...app,
};
