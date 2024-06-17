import React from "react";
import { useTranslation } from "react-i18next";
import Flag from 'react-world-flags'


function LanguageSwitcher() {
  const { i18n } = useTranslation();

  const changeLanguage = (lng) => {
    i18n.changeLanguage(lng);
  };

  return (
    <div className="language-switcher"> 
      <img
        src={<Flag code={ "PT" } />}
        alt="English"
        onClick={() => changeLanguage("en")}
        style={{ cursor: "pointer" }}
      />
      <img
        src={<Flag code={ "PT" } />}
        alt="Português"
        onClick={() => changeLanguage("pt")}
        style={{ cursor: "pointer" }}
      />
    </div>
  );
}

export default LanguageSwitcher;