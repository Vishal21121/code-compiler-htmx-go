/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      boxShadow: {
        "green": "0 0 5px #FF6C00, 0 0 25px #22C55E, 0 0 50px #22C55E, 0 0 100px #22C55E, 0 0 200px #22C55E",
      }
    },
  },
  plugins: [],
}

