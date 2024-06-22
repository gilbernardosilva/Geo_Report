import React from "react";


function ErrorHandler({ error }) {

  const errorMessage = error?.response?.data?.error || "An unexpected error occurred.";

  return (
    <div className="d-flex justify-content-center align-items-center vh-100">
      <h1 className="text-center text-danger">
        {errorMessage}
      </h1>
    </div>
  );
}

export default ErrorHandler;
