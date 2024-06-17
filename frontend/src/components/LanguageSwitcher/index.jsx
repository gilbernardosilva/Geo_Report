import React from "react";
import { useTranslation } from "react-i18next";
import Flag from 'react-world-flags'
import './index.css';


function LanguageSwitcher({ onLanguageChange }) {
  const { i18n } = useTranslation();

  const changeLanguage = (lng) => {
    i18n.changeLanguage(lng);
    onLanguageChange(lng === "en" ? "Sign Up" : "Inscreva-se");
  };

  return (
    <div className="language-switcher"> 
      <div className="flag-container">
        <Flag code="GB" height="30" className="round-flag" onClick={() => changeLanguage('en')} />
      </div>
      <div className="flag-container">
        <Flag code="pt" height="30" className="round-flag" onClick={() => changeLanguage('pt')} />
      </div>
    </div>       
  );
}

export default LanguageSwitcher;