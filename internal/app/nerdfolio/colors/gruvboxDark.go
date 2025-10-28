package colors

const GruvboxDark = `

:root {
  --bg0_h: #1d2021;
  --bg0: #282828;
  --bg0_s: #32302f;
  --bg1: #3c3836;
  --bg2: #504945;
  --bg3: #665c54;
  --bg4: #7c6f64;

  --fg0: #fbf1c7;
  --fg1: #ebdbb2;
  --fg2: #d5c4a1;
  --fg3: #bdae93;
  --fg4: #a89984;

  --gray: #928374;

  --red: #fb4934;
  --red-dim: #cc241d;
  --green: #b8bb26;
  --green-dim: #98971a;
  --yellow: #fabd2f;
  --yellow-dim: #d79921;
  --blue: #83a598;
  --blue-dim: #458588;
  --purple: #d3869b;
  --purple-dim: #b16286;
  --aqua: #8ec07c;
  --aqua-dim: #689d6a;
  --orange: #fe8019;
  --orange-dim: #d65d0e;
}

html, body {
  background-color: var(--bg0);
  color: var(--fg1);
  font-family: system-ui, sans-serif;
}

h1 { color: var(--yellow); }
h2 { color: var(--orange); }
h3 { color: var(--blue); }
h4 { color: var(--purple); }
h5 { color: var(--aqua); }
h6 { color: var(--green); }

p { color: var(--fg2); }

a { color: var(--blue); text-decoration: none; }
a:hover { color: var(--aqua); }

strong, b { color: var(--orange); }
em, i { color: var(--red); }

code, pre {
  color: var(--green);
  background-color: var(--bg1);
  border-radius: 4px;
  padding: 0.25rem 0.5rem;
}

blockquote {
  color: var(--purple);
  border-left: 4px solid var(--yellow);
  padding-left: 0.75rem;
}

hr {
  border: none;
  height: 1px;
  background-color: var(--bg3);
}

table {
  color: var(--fg2);
  background-color: var(--bg1);
  border-collapse: collapse;
}
th {
  color: var(--fg0);
  background-color: var(--bg2);
  padding: 0.5rem 1rem;
}
td {
  color: var(--fg3);
  padding: 0.5rem 1rem;
  border-top: 1px solid var(--bg3);
}

button, input[type="button"], input[type="submit"] {
  color: var(--bg0);
  background-color: var(--orange);
  border: none;
  border-radius: 4px;
  padding: 0.4rem 0.8rem;
}

input, textarea, select {
  color: var(--fg1);
  background-color: var(--bg1);
  border: 1px solid var(--bg3);
  border-radius: 4px;
  padding: 0.3rem 0.5rem;
}

ul, ol { color: var(--fg2); }
li::marker { color: var(--yellow); }

img { background-color: var(--bg1); }

::selection {
  background-color: var(--yellow);
  color: var(--bg0);
}`
