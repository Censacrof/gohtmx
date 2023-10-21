/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./cmd/**/*.{html,js,templ}", "./web/**/*.{html,js,templ}"],
  theme: {
    extend: {
      gridTemplateRows: {
        "page-layout": "0.2fr 1fr 0.2fr",
      },
      gridTemplateColumns: {
        "page-layout": "1fr min-content 1fr",
      },
    },
  },
  plugins: [],
};
