/** @type {import('tailwindcss').Config} */
module.exports = {
  mode: 'jit',
  content: ["./ui/html/**/*.html"], // This is where your HTML templates / JSX files are located
  theme: {
    extend: {
      fontFamily: {
        fredoka: ["Fredoka", "sans-serif"]
      },
      colors: {
        'feathers': '#FDF6E3',
        'feathers-dark': '#E6DFC9',
        'beak': '#C87824'
      },
      screens: {
        portrait: { raw: '(orientation: portrait)' },
        landscape: { raw: '(orientation: landscape)' },
      },
    },
  },
  plugins: [],
};
