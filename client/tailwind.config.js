/** @type {import('tailwindcss').Config} */ 

export const content = [
  "./index.html",
  "./src/**/*.{js,ts,jsx,tsx}",

  // Path to the Tremor module
  "./node_modules/@tremor/**/*.{js,ts,jsx,tsx}",
];
export const theme = {
  transparent: "transparent",
  current: "currentColor",
  extend: {
    colors: {
      // light mode
      tremor: {
        brand: {
          faint: "#eff6ff",
          muted: "#bfdbfe",
          subtle: "#60a5fa",
          DEFAULT: "#3b82f6",
          emphasis: "#1d4ed8",
          inverted: "#ffffff", // white
        },
        background: {
          muted: "#f9fafb",
          subtle: "#f3f4f6",
          DEFAULT: "#ffffff",
          emphasis: "#374151", // gray-700
        },
        border: {
          DEFAULT: "#e5e7eb", // gray-200
        },
        ring: {
          DEFAULT: "#e5e7eb", // gray-200
        },
        content: {
          subtle: "#9ca3af",
          DEFAULT: "#6b7280",
          emphasis: "#374151",
          strong: "#111827",
          inverted: "#ffffff", // white
        },
      },
      // dark mode
      "dark-tremor": {
        brand: {
          faint: "#0B1229",
          muted: "#172554",
          subtle: "#1e40af",
          DEFAULT: "#3b82f6",
          emphasis: "#60a5fa",
          inverted: "#030712", // gray-950
        },
        background: {
          muted: "#131A2B",
          subtle: "#1f2937",
          DEFAULT: "#111827",
          emphasis: "#d1d5db", // gray-300
        },
        border: {
          DEFAULT: "#1f2937", // gray-800
        },
        ring: {
          DEFAULT: "#1f2937", // gray-800
        },
        content: {
          subtle: "#4b5563",
          DEFAULT: "#6b7280",
          emphasis: "#e5e7eb",
          strong: "#f9fafb",
          inverted: "#000000", // black
        },
      },
    },
    boxShadow: {
      // light
      "tremor-input": "0 1px 2px 0 rgb(0 0 0 / 0.05)",
      "tremor-card": "0 1px 3px 0 rgb(0 0 0 / 0.1), 0 1px 2px -1px rgb(0 0 0 / 0.1)",
      "tremor-dropdown": "0 4px 6px -1px rgb(0 0 0 / 0.1), 0 2px 4px -2px rgb(0 0 0 / 0.1)",
      // dark
      "dark-tremor-input": "0 1px 2px 0 rgb(0 0 0 / 0.05)",
      "dark-tremor-card": "0 1px 3px 0 rgb(0 0 0 / 0.1), 0 1px 2px -1px rgb(0 0 0 / 0.1)",
      "dark-tremor-dropdown": "0 4px 6px -1px rgb(0 0 0 / 0.1), 0 2px 4px -2px rgb(0 0 0 / 0.1)",
    },
    borderRadius: {
      "tremor-small": "0.375rem",
      "tremor-default": "0.5rem",
      "tremor-full": "9999px",
    },
    fontSize: {
      "tremor-label": ["0.75rem"],
      "tremor-default": ["0.875rem", { lineHeight: "1.25rem" }],
      "tremor-title": ["1.125rem", { lineHeight: "1.75rem" }],
      "tremor-metric": ["1.875rem", { lineHeight: "2.25rem" }],
    },
  },
};
export const safelist = [
  {
    pattern: /^(bg-(?:slate|gray|zinc|neutral|stone|red|orange|amber|yellow|lime|green|emerald|teal|cyan|sky|blue|indigo|violet|purple|fuchsia|pink|rose)-(?:50|100|200|300|400|500|600|700|800|900|950))$/,
    variants: ["hover", "ui-selected"],
  },
  {
    pattern: /^(text-(?:slate|gray|zinc|neutral|stone|red|orange|amber|yellow|lime|green|emerald|teal|cyan|sky|blue|indigo|violet|purple|fuchsia|pink|rose)-(?:50|100|200|300|400|500|600|700|800|900|950))$/,
    variants: ["hover", "ui-selected"],
  },
  {
    pattern: /^(border-(?:slate|gray|zinc|neutral|stone|red|orange|amber|yellow|lime|green|emerald|teal|cyan|sky|blue|indigo|violet|purple|fuchsia|pink|rose)-(?:50|100|200|300|400|500|600|700|800|900|950))$/,
    variants: ["hover", "ui-selected"],
  },
  {
    pattern: /^(ring-(?:slate|gray|zinc|neutral|stone|red|orange|amber|yellow|lime|green|emerald|teal|cyan|sky|blue|indigo|violet|purple|fuchsia|pink|rose)-(?:50|100|200|300|400|500|600|700|800|900|950))$/,
  },
  {
    pattern: /^(stroke-(?:slate|gray|zinc|neutral|stone|red|orange|amber|yellow|lime|green|emerald|teal|cyan|sky|blue|indigo|violet|purple|fuchsia|pink|rose)-(?:50|100|200|300|400|500|600|700|800|900|950))$/,
  },
  {
    pattern: /^(fill-(?:slate|gray|zinc|neutral|stone|red|orange|amber|yellow|lime|green|emerald|teal|cyan|sky|blue|indigo|violet|purple|fuchsia|pink|rose)-(?:50|100|200|300|400|500|600|700|800|900|950))$/,
  },
];
export const plugins = [require("@headlessui/tailwindcss")];