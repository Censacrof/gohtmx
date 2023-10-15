/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["web/**/*.{html,js}"],
  theme: {
    extend: {
      gridTemplateRows: {
        "page-layout": "0.2fr 1fr 0.2fr",
      },
      gridTemplateColumns: {
        "page-layout": "0.2fr 1fr 0.2fr",
      },
    },
  },
  plugins: [],
};
