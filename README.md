<p align="center">
 <img src="https://raw.githubusercontent.com/oggnimodd/xhinobi/main/images/logo-rounded.png" width="175"/>
</p>

# xhinobi (Go Version)
## Go version of https://github.com/oggnimodd/xhinobi

xhinobi simplifies the process of aggregating text content from multiple files without the need to manually open and copy each file individually. Instead of laboriously opening each file and copying its content, Xhinobi streamlines this task by allowing users to gather text from multiple files automatically using command-line instructions. Xhinobi is a useful tool when you need to copy a large amount of text content from multiple files for use in a language model like ChatGPT or Phind. This is especially helpful when you're coding and need to provide the model with a large amount of context or data.

<img src="https://raw.githubusercontent.com/oggnimodd/xhinobi/main/images/demo.gif" />

## Usage
It is designed to work with other command-line tools like `fdfind` and `fzf`. Here is an example of how to use it:

```
fdfind --type f --exclude node_modules,dist | fzf -m | xhinobi
```

In this example, `fdfind` is used to find all files in the current directory excluding node_modules and dist. The output is piped to `fzf` for multi-selection and then piped to `xhinobi`. `xhinobi` will combine all the content from those files, minify them, and copy them to the clipboard.

## Options
xhinobi supports several options that can be used to customize its behavior:

- `-n` or `--prependFileName`: Prepend the file name before the content. 
- `-m` or `--minify`: Minify the output.
- `-i` or `--ignore`: Glob patterns to ignore (can be used multiple times).
- `-t` or `--tree`: Prepend the output with a directory tree (requires 'tree' command).
- `-o` or `--osc52`: Use the OSC52 escape sequence to copy to the clipboard. This is essential for copying text from a remote machine over SSH to your local clipboard, provided you use a compatible terminal (e.g., iTerm2, Kitty, WezTerm).