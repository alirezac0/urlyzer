# URLyzer
URLyzer is a tool written in Golang that extracts words from URLs and generates a wordlist for further analysis. It can also perform basic mutations on the words to create more variations. URLyzer does not send any HTTP requests to the target, so it is good for passive reconnaissance.

## Installation
You can install URLyzer using the following command:

```bash
GO111MODULE=off go get -v github.com/alirezac0/urlyzer/cmd/urlyzer
```


## Usage
URLyzer can accept URLs as input from a file or from stdin. You can use the following syntax:

```bash
cat urls.txt | urlyzer

or

urlyzer -l urls.txt
```

URLyzer will output the wordlist to stdout by default. You can pipe the output of tools like [gau](https://github.com/lc/gau) to urlyzer.

```bash
gau google.com | urlyzer -e ".jpg,.js,.woff2"
```

To display the help for the tool use the `-h` flag:

```bash
urlyzer -h
```

| Flag | Description | Example |
|------|-------------|---------|
|`-h`| Show the help message and exit | urlyzer -h|
|`-e`| Exclude some words or extensions from the wordlist. You can provide a comma-separated list of words or extensions (with a dot) | urlyzer -e .jpg,.woff |
|`-l`| Gets URLs from a file | urlyzer -l urls.txt |


## Example
Suppose you have a file called urls.txt with the following content:
```
https://example.com/login?user=admin&pass=1234
https://example.com/blog/post/2021/06/27/how-to-use-urlyzer
https://example.com/search?q=urlyzer&lang=en
https://example.com/assets/logo.jpg
https://example.com/scripts/main.js
```

You can run URLyzer on this file using the following command:

urlyzer -l -l example.txt

This will generate the following wordlist

```
to-urlyzer
scripts
to
use-how
q
27
to-how
lang
assets.jpg
assets
06
how-urlyzer
use-urlyzer
urlyzer-use
scripts.js
login
blog
post
2021
how-to
1234
urlyzer-to
search
main
main.js
to-use
how-use
urlyzer-how
en
logo
user
admin
how
use
urlyzer
use-to
logo.jpg
pass
```

You can then use this wordlist for further analysis, such as directory bruteforcing, parameter fuzzing, etc.

## To Do List
Some features that are planned to be added in the future are:

•  [ ] Support more mutation techniques.

•  [ ] Add an option to filter out words by length or frequency.

•  [ ] Add an option to save the wordlist in different formats, such as JSON, CSV, etc.
