package colors

const TokyoNightStorm = `

:root {
  --red: #f7768e;
  --orange: #ff9e64;
  --yellow: #e0af68;
  --cream: #cfc9c2;
  --green: #9ece6a;
  --teal: #73dacb;
  --aqua: #b4f9f8;
  --cyan: #2ac3de;
  --sky: #7dcfff;
  --blue: #7aa2f7;
  --mauve: #bb9af7;
  --text: #c0caf5;
  --subtext: #a9b1d6;
  --overlay: #9aa5ce;
  --surface: #565f89;
  --base: #24283b;
}

html, body {
  background-color: var(--base);
  color: var(--text);
  font-family: system-ui, sans-serif;
}

h1 { color: var(--mauve); }
h2 { color: var(--blue); }
h3 { color: var(--sky); }
h4 { color: var(--teal); }
h5 { color: var(--green); }
h6 { color: var(--yellow); }

p { color: var(--subtext); }

a { color: var(--sky); text-decoration: none; }
a:hover { color: var(--blue); }

strong, b { color: var(--orange); }
em, i { color: var(--red); }

code, pre {
  color: var(--green);
  background-color: var(--surface);
  padding: 0.25rem 0.5rem;
  border-radius: 4px;
}

blockquote {
  color: var(--mauve);
  border-left: 4px solid var(--blue);
  padding-left: 0.75rem;
}

hr {
  border: none;
  height: 1px;
  background-color: var(--surface);
}

table {
  color: var(--subtext);
  background-color: var(--base);
  border-collapse: collapse;
}
th {
  color: var(--text);
  background-color: var(--surface);
  padding: 0.5rem 1rem;
}
td {
  color: var(--subtext);
  padding: 0.5rem 1rem;
  border-top: 1px solid var(--surface);
}

button, input[type="button"], input[type="submit"] {
  color: var(--base);
  background-color: var(--blue);
  border: none;
  border-radius: 4px;
  padding: 0.4rem 0.8rem;
}

input, textarea, select {
  color: var(--text);
  background-color: var(--surface);
  border: 1px solid var(--overlay);
  border-radius: 4px;
  padding: 0.3rem 0.5rem;
}

ul, ol { color: var(--subtext); }
li::marker { color: var(--cyan); }

img { background-color: var(--base); }

::selection {
  background-color: var(--mauve);
  color: var(--base);
}`
