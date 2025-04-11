/** @type {import('tailwindcss').Config} */
module.exports = {
  mode: 'jit',
  content: ["./ui/html/**/*.html"], // This is where your HTML templates / JSX files are located
  theme: {
    extend: {
      fontFamily: {
        // Fredoka!
      },
      colors: {
        'feathers': '#FDF6E3'
      }
    },
  },
  plugins: [],
};
