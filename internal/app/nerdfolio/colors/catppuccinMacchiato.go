package colors

const CatppuccinMacchiato = `

:root {
  --rosewater: #f4dbd6;
  --flamingo: #f0c6c6;
  --pink: #f5bde6;
  --mauve: #c6a0f6;
  --red: #ed8796;
  --maroon: #ee99a0;
  --peach: #f5a97f;
  --yellow: #eed49f;
  --green: #a6da95;
  --teal: #8bd5ca;
  --sky: #91d7e3;
  --sapphire: #7dc4e4;
  --blue: #8aadf4;
  --lavender: #b7bdf8;
  --text: #cad3f5;
  --subtext1: #b8c0e0;
  --subtext0: #a5adcb;
  --overlay2: #939ab7;
  --overlay1: #8087a2;
  --overlay0: #6e738d;
  --surface2: #5b6078;
  --surface1: #494d64;
  --surface0: #363a4f;
  --base: #24273a;
  --mantle: #1e2030;
  --crust: #181926;
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
