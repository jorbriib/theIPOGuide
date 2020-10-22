import React, { Fragment, useEffect, useRef, useState } from "react";
import { loadReCaptcha, ReCaptcha } from "react-recaptcha-v3";

import { RECAPTCHA_SITE_KEY } from "../../../../config";
import { sendReport } from "./client";
import { formatMessage } from "../../../lang";

const Report = () => {
  const [message, setMessage] = useState("");
  const [token, setToken] = useState("");
  const [errorMessage, setErrorMessage] = useState("");
  const btnRef = useRef();
  const recaptchaRef = useRef();

  useEffect(() => {
    loadReCaptcha(RECAPTCHA_SITE_KEY);
  }, []);

  const verifyCallback = function (response) {
    setToken(response);
  };

  const handleSubmit = async (event) => {
    event.preventDefault();
    btnRef.current.setAttribute("disabled", "disabled");
    setErrorMessage("");

    if (!message) {
      showError("Forms.ErrorMessages.EmptyReportMessage");
      return;
    }

    if (!token) {
      showError("Forms.ErrorMessages.EmptyToken");
      return;
    }

    const error = await sendReport({
      url: window.location.href,
      message: message,
      token: token,
    });
    if (error !== undefined) {
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
    setErrorMessage(formatMessage({ id: aliasMessage }));
    btnRef.current.removeAttribute("disabled");
  };

  const cleanForm = function () {
    setMessage("");
    btnRef.current.removeAttribute("disabled");
    recaptchaRef.current.execute();

    // eslint-disable-next-line no-undef
    $("#atbdp-report-abuse-modal").modal("hide");
  };

  return (
    <Fragment>
      <div
        className="modal fade"
        id="atbdp-report-abuse-modal"
        tabIndex="-1"
        role="dialog"
        aria-hidden="true"
        aria-labelledby="atbdp-report-abuse-modal-label"
      >
        <div className="modal-dialog modal-dialog-centered" role="document">
          <div className="modal-content">
            <form
              id="atbdp-report-abuse-form"
              className="form-vertical"
              onSubmit={handleSubmit}
            >
              <div className="modal-header">
                <h3 className="modal-title" id="atbdp-report-abuse-modal-label">
                  Send feedback
                </h3>
                <button type="button" className="close" data-dismiss="modal">
                  <span aria-hidden="true">×</span>
                </button>
              </div>
              <div className="modal-body">
                <div className="form-group has-error">
                  <p>
                    We try to have the most accurate data for you, but sometimes
                    we fail :( Please share with us your concerns or
                    improvements:
                  </p>
                  <label
                    htmlFor="atbdp-report-abuse-message"
                    className="not_empty"
                  >
                    Your message<span className="atbdp-star">*</span>
                  </label>
                  <textarea
                    className="form-control"
                    id="atbdp-report-abuse-message"
                    rows="4"
                    placeholder="Write here..."
                    required={true}
                    name="message"
                    value={message}
                    onChange={(event) => setMessage(event.target.value)}
                  />
                </div>
                <div className="form-group">
                  <ReCaptcha
                    ref={recaptchaRef}
                    sitekey={RECAPTCHA_SITE_KEY}
                    action="report"
                    verifyCallback={verifyCallback}
                  />
                </div>
                <div className="form-group">
                  {errorMessage && (
                    <span className="error-message" role="alert">
                      {errorMessage}
                    </span>
                  )}
                </div>
              </div>
              <div className="modal-footer">
                <button
                  type="submit"
                  className="btn btn-sm"
                  data-dismiss="modal"
                >
                  Close
                </button>
                <button
                  type="submit"
                  ref={btnRef}
                  className="btn btn-primary btn-sm"
                >
                  Submit
                </button>
              </div>
            </form>
          </div>
        </div>
      </div>
    </Fragment>
  );
};

export default Report;
