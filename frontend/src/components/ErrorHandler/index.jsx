import React from "react";


function ErrorHandler({ error }) {

  const errorMessage = error?.response?.data?.error || "An unexpected error occurred.";

  return (
<div className="d-flex flex-column justify-content-center align-items-center vh-100">
  <div>
    <h1 className="text-center text-danger">
      ╭∩╮( ͡° ͜ʖ ͡°)╭∩╮ {errorMessage}
    </h1>
  </div>
  <img
    src="https://media1.tenor.com/m/pFz1Q12_hXEAAAAd/cat-holding-head-cat.gif"
    alt=""
  />
</div>
  );
}

export default ErrorHandler;
