# wiki - Get summary information on a topic from Wikipedia

## Todos
[ ] - Accept one argument that will map to a Wikipedia article
[ ] - Use `curl` to retrieve the article and output to stdout
[ ] - Limit output to first sentence
[ ] - Grab section headings from `curl` result and output to stdout
[ ] - Add support for multiple args to output the first sentence of a specified section and a list of subsection headings

### Stretch goals
[ ] - Add support for tab completion
[ ] - Add a flag to support querying multipl pages in parallel; could show each page after the other with a line of charaters to separate the pages; use `xargs` for parallelism (?) and store output as vars in an array, then once done querying output array values one at a time (?)

## TL;DR

`wiki` is a Bash program that will output summary information on the topic you provide it. The output includes:

- The first sentence of the article/section
- A list of section/subsection headings

## Usage

`wiki walrus` - shows first sentence of the "Walrus" article and a list of article sections

`wiki walrus anatomy` - shows first section of the "anatomy" section of the "Walrus" article and a list of the subsections in the "anatomy" section
