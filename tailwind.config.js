/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./view/**/*.{templ,go}", "./view/components/**/*.{templ,go}"],
  theme: {
    extend: {
      colors: {
        cement: "#161616",
        eerienight: "#242424",
      },
      fontFamily: {
        signika: ["Signika", "sans-serif"],
      },
      gridTemplateColumns: {
        edit: "150px 1.5fr",
      },
    },
  },
  plugins: [],
};
