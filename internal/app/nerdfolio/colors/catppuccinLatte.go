package colors

const CatppuccinLatte = `

:root {
  --rosewater: #dc8a78;
  --flamingo: #dd7878;
  --pink: #ea76cb;
  --mauve: #8839ef;
  --red: #d20f39;
  --maroon: #e64553;
  --peach: #fe640b;
  --yellow: #df8e1d;
  --green: #40a02b;
  --teal: #179299;
  --sky: #04a5e5;
  --sapphire: #209fb5;
  --blue: #1e66f5;
  --lavender: #7287fd;
  --text: #4c4f69;
  --subtext1: #5c5f77;
  --subtext0: #6c6f85;
  --overlay2: #7c7f93;
  --overlay1: #8c8fa1;
  --overlay0: #9ca0b0;
  --surface2: #acb0be;
  --surface1: #bcc0cc;
  --surface0: #ccd0da;
  --base: #eff1f5;
  --mantle: #e6e9ef;
  --crust: #dce0e8;
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
  color: var(--base);
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
  color: var(--base);
}`
