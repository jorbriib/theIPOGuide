import get from "lodash/get";

import translations from "./en";

const formatMessage = (
  messageDescriptor,
  values,
  translationsMap = translations
) => {
  const { id, defaultMessage } = messageDescriptor;
  let message = get(translationsMap, id);
  if (!message) {
    message = defaultMessageFallback(defaultMessage);
  }

  message = replacePlaceholders(message, values);

  return message;
};

const defaultMessageFallback = (fallbackMsg) => {
  if (fallbackMsg !== undefined) {
    return fallbackMsg;
  }
};

const replacePlaceholders = (message, values) => {
  if (!values) {
    return message;
  }

  let replaced = message;
  Object.keys(values).forEach((key) => {
    replaced = replaced.replace(`{${key}}`, values[key]);
    replaced = replaced.replace(`%${key}%`, values[key]);
  });

  return replaced;
};

export default formatMessage;
