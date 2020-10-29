import React, { Fragment, useEffect, useRef, useState } from "react";
import { loadReCaptcha } from "react-recaptcha-v3";

import Header from "../../components/layout/Header";
import { BreadcrumbWrapper } from "../../components/elements/Breadcrumbs";
import { formatMessage } from "../../lang";
import { RECAPTCHA_SITE_KEY } from "../../../config";
import { sendContactMessage } from "./client";
import Footer from "../../components/layout/Footer";
import {
  ContactForm,
  ContactResponse,
} from "../../components/elements/ContactForm";
import { WidgetContactInfo } from "../../components/elements/Widget";

const Contact = () => {
  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [message, setMessage] = useState("");
  const [token, setToken] = useState("");
  const [errorMessage, setErrorMessage] = useState("");
  const [submitted, setSubmitted] = useState(false);
  const btnRef = useRef();
  const recaptchaRef = useRef();

  useEffect(() => {
    loadReCaptcha(RECAPTCHA_SITE_KEY);
  }, []);

  const backgroundImage = {
    backgroundImage: "url('/assets/images/wallstreet-bull.jpg')",
    opacity: 1,
  };

  const verifyCallback = function (response) {
    setToken(response);
  };

  const handleChange = function (name, value) {
    switch (name) {
      case "name":
        setName(value);
        break;
      case "email":
        setEmail(value);
        break;
      case "message":
        setMessage(value);
        break;
    }
  };

  const handleSubmit = async (event) => {
    event.preventDefault();
    btnRef.current.setAttribute("disabled", "disabled");
    setErrorMessage("");

    if (!name) {
      showError("Forms.ErrorMessages.EmptyName");
      return;
    }

    if (!email) {
      showError("Forms.ErrorMessages.EmptyEmail");
      return;
    }

    const regExp = /^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$/i;
    if (!regExp.test(email.toLowerCase())) {
      showError("Forms.ErrorMessages.InvalidEmail");
      return;
    }

    if (!message) {
      showError("Forms.ErrorMessages.EmptyMessage");
      return;
    }

    if (!token) {
      showError("Forms.ErrorMessages.EmptyToken");
      return;
    }

    const error = await sendContactMessage({
      name: name,
      email: email,
      message: message,
      token: token,
    });
    if (error !== undefined) {
      recaptchaRef.current.execute();
      showError(
        error.message
          ? error.message
          : "Forms.ErrorMessages.DefaultErrorMessage"
      );
      return;
    }
    cleanForm();
  };

  const showError = function (aliasMessage) {
    setErrorMessage(
      formatMessage({ id: aliasMessage, defaultMessage: aliasMessage })
    );
    btnRef.current.removeAttribute("disabled");
  };

  const cleanForm = function () {
    setMessage("");
    setName("");
    setEmail("");
    btnRef.current.removeAttribute("disabled");
    recaptchaRef.current.execute();
    setSubmitted(true);
  };

  return (
    <Fragment>
      {/* Header section start */}
      <section className="header-breadcrumb bgimage overlay overlay--dark">
        <div className="bg_image_holder" style={backgroundImage}>
          <img
            src="/assets/images/wallstreet-bull.jpg"
            alt="Find your IPO and invest"
          />
        </div>
        <div className="mainmenu-wrapper">
          <Header class="menu--light" />
        </div>
        {/* <!-- ends: .mainmenu-wrapper --> */}
        <BreadcrumbWrapper
          title="Contact with us"
          onlyTitle={true}
          subtitle="We're here to help and answer any question you might have"
        />
      </section>
      {/* Header section end */}
      <section className="contact-area section-bg p-top-60 p-bottom-40">
        <div className="container">
          <div className="row">
            <div className="col-lg-8">
              <div className="widget atbd_widget widget-card contact-block">
                <div className="atbd_widget_title">
                  <h4>
                    <span className="la la-envelope"></span> Contact Form
                  </h4>
                </div>
                <div className="atbdp-widget-listing-contact contact-form">
                  {!submitted ? (
                    <ContactForm
                      verifyCallback={verifyCallback}
                      btnRef={btnRef}
                      email={email}
                      name={name}
                      message={message}
                      errorMessage={errorMessage}
                      handleChange={handleChange}
                      handleSubmit={handleSubmit}
                      recaptchaRef={recaptchaRef}
                    />
                  ) : (
                    <ContactResponse name={name} />
                  )}
                </div>
              </div>
            </div>

            <div className="col-lg-4">
              <div className="widget atbd_widget widget-card">
                <div className="atbd_widget_title">
                  <h4>
                    <span className="la la-phone"></span>Contact Info
                  </h4>
                </div>
                <WidgetContactInfo />
              </div>
            </div>
          </div>
        </div>
      </section>

      <Footer />
    </Fragment>
  );
};

export default Contact;
