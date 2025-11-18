# Nerdfolio

**Static site generator framework tailored for building nerdish personal websites.**

![Nerdfolio Demo](https://github.com/atasoya/nerdfolio/blob/main/documentation/nerdfolio-banner.png?raw=true)

## Features:

- Static HTML file generation.
- Native .md blog support.
- HTML templating that can handle templates, data, loops.
- Single command color scheme linking.


## What is Nerdfolio?
Nerdfolio is my hobby project born by the pain I encounter managing my own personal website [ata.soy](https://www.ata.soy/). Here I tried to keep everything minimal and developer friendly so I have fun managing my personal website.

First idea was I was done enough with js ecosystem. It is crazy that I need to have a runtime for building simple portfolios like this. But main problem was blog support. I don't want to be dependent to some random providers even databases for just rendering simple .md files. 

I just wanted it to work.

## Should I use Nerdfolio?
Nerdfolio is not for everyone. If you are trying to impress others with some animations and white/dark mode switches you should use something else. If you want visitors to focus on the source of your personal website which is the actual work you did, Nerdfolio can be a great option. Especially native .md blog support will save a lot of time.

## Quick Start

This command will create initial project structure for nerdfolio projects. Nerdfolio embraces conventions over configuration mindset.
```bash
nerdfolio create
```

Now you need to build it. By default nerdfolio will use CatppuccinMocha color scheme but you can set it to your favorite color scheme using `--colorScheme <colorScheme>`
```bash
nerdfolio build 

# nerdfolio build  --colorScheme gruvboxDark
```

Build is created in `nerdfolio/` directory. You can open index.html in `nerdfolio/` directory to see the end result or use `nerdfolio open` command to open it faster than ever.

# How to Install?

As I keep things minimal, currently building from the source is the only option. You can just build it and link it to your system or use 'make build-test' which I use when testing that builds and links on my Mac.

```bash
git clone https://github.com/atasoya/nerdfolio.git
cd nerdfolio
make build
```

# Documentation

## Project Structure

```
./
├── public/ 
│  
├── blogs/                  # Markdown blogs
│   ├── blog-1.md              
│   └── blog-2.md
│  
├── templates/              # Re-usable templates
│   ├── blog-1.md              
│   └── blog-2.md
│          
├── index.html              # HTML files
├── projects.html     
├── blogs.html      
│       
└── data.json               # JSON data 
```
You can check the [ata.soy](https://www.ata.soy/) repo to see how I use it.

## How Nerdfolio Build Works
1. Parse HTML and markdown files.
2. Parse templates.
3. Replace templates in HTML files.
4. Parse JSON data.
5. Handle loop templating.
6. Replace placeholders with actual data from data.json
7. Write everything to `nerdfolio/`
8. Write `nerdfolio.css` file to `nerdfolio/`
9. Copy `public/` to `nerdfolio/public/`


## Templates
You can create a new template by creating a new HTML file in `templates/` like `templates/new-template.html`. A template file is just a normal HTML file. You can still reach data.json from templates.

To use your `templates/new-template.html` you need to use special syntax of `{% new-template.html %}` inside of your HTML files.

**You can't call template from inside of another template.**

## Data

You can put any valid json as your `data.json`. Main point of `data.json` is to streamline the process of making updates and handling list of data. When parsing, `data.json` gets flattened so you can reach elements with `{{ name }}` or `{{ me.email }}`.

### Loops

You can also handle loops in your HTML with `{{#each contact }} {{this}} {{/each contact }}`. If it is an array you can reach its elements by {{this}}. If you are working with key value pairs you need to use `{{#each contact }} Key: {{key}} Value: {{value}}  {{/each contact }}`

## Blogs
All .md blogs files must be inside `blogs/` directory. Markdown blog `blogs/cool-blog.md` will generate `blogs/cool-blog.html`.

## Color Scheme
You can change the color scheme by using `--colorScheme <colorScheme>` flag in `nerdfolio build` command. You can also use custom color scheme using `--colorScheme custom`. You need to link `nerdfolio.css` or `custom.css` to your HTML files to be able to use them.

Color scheme contributions are very welcomed.

### Available Color Schemes
- catppuccinFrappe (default)
- catppuccinLatte
- catppuccinMacchiato
- catppuccinMocha
- gruvboxDark
- tokyoNightStorm
- custom 

## Hosting 
They are just HTML files so you can host them in any place but I highly recommend using github pages.

### Hosting on Vercel
Hosting on Vercel is straightforward:
- Connect your Github project to Vercel.
- Select `other` as the framework.
- Set `nerddolio` as  the root directory.
- Change build command to `echo "no build"`.
- Change output directory to `.`.

# Websites Using Nerdfolio
- [ata.soy](https://www.ata.soy/)
