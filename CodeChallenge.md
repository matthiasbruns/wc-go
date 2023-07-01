# Write Your Own wc Tool

_Markdown version of https://codingchallenges.fyi/challenges/challenge-wc_

This challenge is to build your own version of the Unix command line tool wc!

The Unix command line tools are a great metaphor for good software engineering and they follow the Unix Philosophies of:

- Writing simple parts connected by clean interfaces - each tool does just one thing and provides a simple CLI that handles text input from either files or file streams.
- Design programs to be connected to other programs - each tool can be easily connected to other tools to create incredibly powerful compositions.

Following these philosophies has made the simple unix command line tools some of the most widely used software engineering tools - allowing us to create very complex text data processing pipelines from simple command line tools. There’s even a Coursera course on [Linux and Bash for Data Engineering](https://gb.coursera.org/learn/linux-and-bash-for-data-engineering-duke).

You can read more about the Unix Philosophy in the excellent book [The Art of Unix Programming.](http://www.catb.org/~esr/writings/taoup/html/)

## The Challenge - Building wc

The functional requirements for wc are concisely described by it’s man page - give it a go in your local terminal now:

```bash
man wc
```

The TL/DR version is: **wc** – word, line, character, and byte count.

## Step Zero

Like all good software engineering we’re zero indexed! In this step you’re going to set your environment up ready to begin developing and testing your solution.

I’ll leave you to setup your IDE / editor of choice and programming language of choice. After that here’s what I’d like you to do to be ready to test your solution.

Download the this [text from Project Gutenberg](https://www.gutenberg.org/cache/epub/132/pg132.txt) and save it as `test.txt`.

## Step One

In this step your goal is to write a simple version of wc, let’s call it ccwc (cc for Coding Challenges) that takes the command line option -c and outputs the number of bytes in a file.

If you’ve done it right your output should match this:

```bash
> ccwc -c test.txt
 341836 test.txt
```

## Step Two

In this step your goal is to support the command line option -l that outputs the number of lines in a file.

If you’ve done it right your output should match this:

```bash
> ccwc -l test.txt
    7137 test.txt
```

## Step Three

In this step your goal is to support the command line option -w that outputs the number of words in a file. If you’ve done it right your output should match this:

```bash
> ccwc -w test.txt
   58159 test.txt
```

## Step Four

In this step your goal is to support the command line option -m that outputs the number of characters in a file. If the current locale does not support multibyte characters this will match the -c option.

You can learn more about programming for locales [here](https://learn.microsoft.com/en-us/globalization/locale/locale-and-culture)

For this one your answer will depend on your locale, so if can, use wc itself and compare the output to your solution:

```bash
> wc -m test.txt
 339120 test.txt
> ccwc -m test.txt
 339120 test.txt
```

## Step Five

In this step your goal is to support the default option - i.e. no options are provided, which is the equivalent to the -c, -l and -w options. If you’ve done it right your output should match this:

```bash
> ccwc test.txt
    7137   58159  341836 test.txt
```

## The Final Step

In this step your goal is to support being able to read from standard input if no filename is specified. If you’ve done it right your output should match this:

```bash
> cat test.txt | ccwc -l
    7137
```

## Help Others by Sharing Your Solutions!

If you think your solution is an example other developers can learn from please share it, put it on GitHub, GitLab or elsewhere. Then let me know - ping me a message via [Twitter](https://twitter.com/johncrickett) or [LinkedIn](https://www.linkedin.com/in/johncrickett/) or just post about it there and tag me.

## Get The Challenges By Email

If you would like to recieve the coding challenges by email, you can subscribe to the weekly newsletter on SubStack here:

<iframe src="https://codingchallenges.substack.com/embed" width="480" height="320" frameborder="0" scrolling="no" style="border: 1px solid rgb(238, 238, 238); background: white;"></iframe>