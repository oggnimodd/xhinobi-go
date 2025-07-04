package helpers

import (
	"path/filepath"
	"strings"
)

// IsTextFile checks if the given filename suggests a text file by checking
// its name directly (for files like Dockerfile) and its extension.
func IsTextFile(filename string) bool {
	textExtensions := []string{
		"Dockerfile", "Makefile", "Rakefile", "ada", "adb", "ads", "applescript", "as", "ascx", "asm", "asmx", "asp", "aspx", "atom", "bas", "bash", "bashrc", "bat", "bbcolors", "bdsgroup", "bdsproj", "bib", "bowerrc", "c", "cbl", "cc", "cfc", "cfg", "cfm", "cfml", "cgi", "clj", "cls", "cmake", "cmd", "cnf", "cob", "coffee", "coffeekup", "conf", "cpp", "cpt", "cpy", "crt", "cs", "csh", "cson", "csr", "css", "csslintrc", "csv", "ctl", "curlrc", "cxx", "dart", "dfm", "diff", "dockerignore", "dof", "dpk", "dproj", "dtd", "eco", "editorconfig", "ejs", "el", "emacs", "eml", "ent", "erb", "erl", "eslintignore", "eslintrc", "ex", "exs", "f", "f03", "f77", "f90", "f95", "fish", "for", "fpp", "frm", "ftn", "gemrc", "gitattributes", "gitconfig", "gitignore", "gitkeep", "gitmodules", "go", "gpp", "gradle", "groovy", "groupproj", "grunit", "gtmpl", "gvimrc", "h", "haml", "hbs", "hgignore", "hh", "hpp", "hrl", "hs", "hta", "htaccess", "htc", "htm", "html", "htpasswd", "hxx", "iced", "inc", "ini", "ino", "int", "irbrc", "itcl", "itermcolors", "itk", "jade", "java", "jhtm", "jhtml", "js", "jscsrc", "jshintignore", "jshintrc", "json", "json5", "jsonld", "jsp", "jspx", "jsx", "ksh", "kt", "less", "lhs", "lisp", "log", "ls", "lsp", "lua", "m", "mak", "map", "markdown", "master", "md", "mdown", "mdwn", "mdx", "metadata", "mht", "mhtml", "mjs", "mk", "mkd", "mkdn", "mkdown", "ml", "mli", "mm", "mxml", "nfm", "nfo", "njk", "noon", "npmignore", "npmrc", "nvmrc", "ops", "pas", "pasm", "patch", "pbxproj", "pch", "pem", "pg", "php", "php3", "php4", "php5", "phpt", "phtml", "pir", "pl", "pm", "pmc", "pod", "pot", "prisma", "properties", "props", "pt", "pug", "py", "r", "rake", "rb", "rdoc", "rdoc_options", "resx", "rhtml", "rjs", "rlib", "rmd", "ron", "rs", "rss", "rst", "rtf", "rvmrc", "rxml", "s", "sass", "scala", "scm", "scss", "seestyle", "sh", "shtml", "sls", "spec", "sql", "sqlite", "ss", "sss", "st", "strings", "sty", "styl", "stylus", "sub", "sublime-build", "sublime-commands", "sublime-completions", "sublime-keymap", "sublime-macro", "sublime-menu", "sublime-project", "sublime-settings", "sublime-workspace", "sv", "svelte", "svc", "svg", "t", "tcl", "tcsh", "terminal", "tex", "text", "textile", "tg", "tmLanguage", "tmTheme", "tmpl", "toml", "tpl", "ts", "tsv", "tsx", "tt", "tt2", "ttml", "txt", "v", "vb", "vbs", "vh", "vhd", "vhdl", "vim", "viminfo", "vimrc", "vue", "webapp", "wxml", "wxss", "x-php", "xaml", "xht", "xhtml", "xml", "xs", "xsd", "xsl", "xslt", "yaml", "yml", "zsh", "zshrc",
	}

	lowerFilename := strings.ToLower(filename)
	extension := strings.ToLower(strings.TrimPrefix(filepath.Ext(filename), "."))

	// First, check for exact filename matches (e.g., "Dockerfile", "Makefile")
	for _, name := range textExtensions {
		if lowerFilename == strings.ToLower(name) {
			return true
		}
	}

	// If no exact match, check by extension
	if extension != "" {
		for _, ext := range textExtensions {
			if ext == extension {
				return true
			}
		}
	}

	return false
}
