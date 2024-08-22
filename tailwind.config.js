/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    './view/**/*.templ',
  ],
  theme: {
    extend: {
        colors: {
            "ahblue": '#449bde',
        }
      },
    },
  plugins: [
    require('@tailwindcss/forms'),
  ]
}