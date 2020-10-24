const app = {
  home: () => "/",
  ipo: (alias) => `/ipo/${alias}`,
  privacyPolicy: () => "/privacy-policy",
  termsAndConditions: () => "/terms-and-conditions",
};

export default {
  ...app,
};
