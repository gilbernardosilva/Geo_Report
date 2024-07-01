import React from "react";


function ErrorHandler({ error }) {

  const errorMessage = error?.response?.data?.error || "An unexpected error occurred.";

  return (
<div className="d-flex flex-column justify-content-center align-items-center vh-100">
  <div>
    <h1 className="text-center text-danger">
      {errorMessage}
    </h1>
  </div>
</div>
  );
}

export default ErrorHandler;
