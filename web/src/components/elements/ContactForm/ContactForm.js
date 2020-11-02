import React from "react";
import { ReCaptcha } from "react-recaptcha-v3";

import { RECAPTCHA_SITE_KEY } from "../../../config";

export function ContactForm({
  handleSubmit,
  handleChange,
  name,
  email,
  message,
  recaptchaRef,
  verifyCallback,
  errorMessage,
  btnRef,
}) {
  return (
    <form
      id="atbdp-contact-form"
      className="form-vertical"
      onSubmit={handleSubmit}
    >
      <div className="form-group">
        <input
          type="text"
          name="name"
          required={true}
          className="form-control"
          id="atbdp-contact-name"
          placeholder="Name"
          value={name}
          onChange={(event) =>
            handleChange(event.target.name, event.target.value)
          }
        />
      </div>

      <div className="form-group">
        <input
          type="email"
          name="email"
          required={true}
          className="form-control"
          id="atbdp-contact-email"
          placeholder="Email"
          value={email}
          onChange={(event) =>
            handleChange(event.target.name, event.target.value)
          }
        />
      </div>

      <div className="form-group">
        <textarea
          className="form-control"
          name="message"
          id="atbdp-contact-message"
          rows="3"
          placeholder="Message"
          required={true}
          value={message}
          onChange={(event) =>
            handleChange(event.target.name, event.target.value)
          }
        />
      </div>

      <div className="form-group">
        <ReCaptcha
          ref={recaptchaRef}
          sitekey={RECAPTCHA_SITE_KEY}
          action="report"
          verifyCallback={verifyCallback}
        />
        <small>
          This site is protected by reCAPTCHA and the Google{" "}
          <a href="https://policies.google.com/privacy" target="_blank">
            Privacy Policy
          </a>{" "}
          and{" "}
          <a href="https://policies.google.com/terms" target="_blank">
            Terms of Service
          </a>{" "}
          apply.
        </small>
      </div>

      <div className="form-group">
        {errorMessage && (
          <span className="error-message" role="alert">
            {errorMessage}
          </span>
        )}
      </div>

      <button
        type="submit"
        ref={btnRef}
        className="btn btn-gradient btn-gradient-one btn-block"
      >
        Send Message
      </button>
    </form>
  );
}

export function ContactResponse({ name }) {
  return (
    <div>
      <p>Thank you for getting in touch!</p>
      <p>
        We appreciate you contacting us, {name}. We will get back in touch with
        you soon! Have a great day!
      </p>
    </div>
  );
}
