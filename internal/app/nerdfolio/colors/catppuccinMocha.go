package colors

const CatppuccinMocha = `
:root {
  --rosewater: #f5e0dc;
  --flamingo: #f2cdcd;
  --pink: #f5c2e7;
  --mauve: #cba6f7;
  --red: #f38ba8;
  --maroon: #eba0ac;
  --peach: #fab387;
  --yellow: #f9e2af;
  --green: #a6e3a1;
  --teal: #94e2d5;
  --sky: #89dceb;
  --sapphire: #74c7ec;
  --blue: #89b4fa;
  --lavender: #b4befe;
  --text: #cdd6f4;
  --subtext1: #bac2de;
  --subtext0: #a6adc8;
  --overlay2: #9399b2;
  --overlay1: #7f849c;
  --overlay0: #6c7086;
  --surface2: #585b70;
  --surface1: #45475a;
  --surface0: #313244;
  --base: #1e1e2e;
  --mantle: #181825;
  --crust: #11111b;
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
