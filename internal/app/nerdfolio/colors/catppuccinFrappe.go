package colors

const CatppuccinFrappe = `

:root {
  --rosewater: #f2d5cf;
  --flamingo: #eebebe;
  --pink: #f4b8e4;
  --mauve: #ca9ee6;
  --red: #e78284;
  --maroon: #ea999c;
  --peach: #ef9f76;
  --yellow: #e5c890;
  --green: #a6d189;
  --teal: #81c8be;
  --sky: #99d1db;
  --sapphire: #85c1dc;
  --blue: #8caaee;
  --lavender: #babbf1;
  --text: #c6d0f5;
  --subtext1: #b5bfe2;
  --subtext0: #a5adce;
  --overlay2: #949cbb;
  --overlay1: #838ba7;
  --overlay0: #737994;
  --surface2: #626880;
  --surface1: #51576d;
  --surface0: #414559;
  --base: #303446;
  --mantle: #292c3c;
  --crust: #232634;
}

html, body {
  background-color: var(--base);
  color: var(--text);
}

h1 { color: var(--mauve); }
h2 { color: var(--pink); }
h3 { color: var(--blue); }
h4 { color: var(--lavender); }
h5 { color: var(--sky); }
h6 { color: var(--teal); }

p { color: var(--subtext1); }

a { color: var(--sapphire); }
a:hover { color: var(--blue); }

strong, b { color: var(--yellow); }
em, i { color: var(--peach); }

code, pre { color: var(--green); background-color: var(--surface0); }

blockquote { color: var(--maroon); border-left-color: var(--mauve); }

hr { color: var(--surface1); background-color: var(--surface1); }

table { color: var(--subtext0); background-color: var(--surface0); }
th { color: var(--text); background-color: var(--surface1); }
td { color: var(--subtext1); }

button, input[type="button"], input[type="submit"] {
  color: var(--crust);
  background-color: var(--blue);
}

input, textarea, select {
  color: var(--text);
  background-color: var(--surface0);
}

ul, ol { color: var(--subtext1); }
li::marker { color: var(--peach); }

img { background-color: var(--mantle); }

::selection {
  background-color: var(--mauve);
  color: var(--crust);
}`
