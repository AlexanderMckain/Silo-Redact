/** @type {import('tailwindcss').Config} */
export default {
  content: ['./index.html', './src/**/*.{html,js,svelte,ts}'],
  theme: {
    extend: {
      colors: {
        'silo-accent': '#58a6ff',
        'silo-border': '#30363d',
        'silo-muted': '#8b949e',
      },
      typography: {
        silo: {
          css: {
            '--tw-prose-body': '#d4d4d8',
            '--tw-prose-headings': '#fafafa',
            '--tw-prose-links': '#22d3ee',
          },
        },
      },
    },
  },
  plugins: [require('@tailwindcss/typography')],
}
