import i18n from "i18next";
import { initReactI18next } from "react-i18next";
import Backend from "i18next-http-backend";
import LanguageDetector from "i18next-browser-languagedetector";

i18n
  .use(Backend) // Load translations from the server (your public/locales folder)
  .use(LanguageDetector) // Detect user's preferred language
  .use(initReactI18next) // Connect with React
  .init({
    debug: true, // Uncomment this in development for debugging
    fallbackLng: "en",
    supportedLngs: ["en", "pt"], // Explicitly list supported languages
    ns: ["translation"],
    defaultNS: "translation", 
    backend: {
      loadPath: "/locales/{{lng}}/{{ns}}.json",
    },
    interpolation: {
      escapeValue: false, // React already escapes variables
    },
  });

export default i18n;