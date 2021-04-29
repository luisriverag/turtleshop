package main

import "runtime"

var (
	// no color
	nc = "\033[0m"

	// normal color
	// black   = "\033[0;30m"
	// red     = "\033[0;31m"
	// green   = "\033[0;32m"
	// yellow  = "\033[0;33m"
	purple = "\033[0;34m"
	// magenta = "\033[0;35m"
	// cyan    = "\033[0;36m"
	white = "\033[0;37m"

	// bold color
	boldblack   = "\033[1;30m"
	boldred     = "\033[1;31m"
	boldgreen   = "\033[1;32m"
	boldyellow  = "\033[1;33m"
	boldpurple  = "\033[1;34m"
	boldmagenta = "\033[1;35m"
	boldcyan    = "\033[1;36m"
	boldwhite   = "\033[1;37m"

	// underline color
	underlineblack   = "\033[4;30m"
	underlinered     = "\033[4;31m"
	underlinegreen   = "\033[4;32m"
	underlineyellow  = "\033[4;33m"
	underlinepurple  = "\033[4;34m"
	underlinemagenta = "\033[4;35m"
	underlinecyan    = "\033[4;36m"
	underlinewhite   = "\033[4;37m"

	// background color
	bgblack   = "\033[40m"
	bgred     = "\033[41m"
	bggreen   = "\033[42m"
	bgyellow  = "\033[43m"
	bgpurple  = "\033[44m"
	bgmagenta = "\033[45m"
	bgcyan    = "\033[46m"
	bgwhite   = "\033[47m"

	// bright background color
	bbgblack   = "\033[100m"
	bbgred     = "\033[101m"
	bbggreen   = "\033[102m"
	bbgyellow  = "\033[103m"
	bbgpurple  = "\033[104m"
	bbgmagenta = "\033[105m"
	bbgcyan    = "\033[106m"
	bbgwhite   = "\033[107m"

	// bright color
	brightblack   = "\033[0;90m"
	brightred     = "\033[0;91m"
	brightgreen   = "\033[0;92m"
	brightyellow  = "\033[0;93m"
	brightpurple  = "\033[0;94m"
	brightmagenta = "\033[0;95m"
	brightcyan    = "\033[0;96m"
	brightwhite   = "\033[0;97m"

	// boldbright color
	boldbrightblack   = "\033[1;90m"
	boldbrightred     = "\033[1;91m"
	boldbrightgreen   = "\033[1;92m"
	boldbrightyellow  = "\033[1;93m"
	boldbrightpurple  = "\033[1;94m"
	boldbrightmagenta = "\033[1;95m"
	boldbrightcyan    = "\033[1;96m"
	boldbrightwhite   = "\033[1;97m"
)

func osCheck() {
	if runtime.GOOS == "windows" {
		nc = ""

		brightblack = ""
		brightred = ""
		brightgreen = ""
		brightyellow = ""
		brightpurple = ""
		brightmagenta = ""
		brightcyan = ""
		brightwhite = ""

		boldblack = ""
		boldred = ""
		boldgreen = ""
		boldyellow = ""
		boldpurple = ""
		boldmagenta = ""
		boldcyan = ""
		boldwhite = ""

		underlineblack = ""
		underlinered = ""
		underlinegreen = ""
		underlineyellow = ""
		underlinepurple = ""
		underlinemagenta = ""
		underlinecyan = ""
		underlinewhite = ""

		bgblack = ""
		bgred = ""
		bggreen = ""
		bgyellow = ""
		bgpurple = ""
		bgmagenta = ""
		bgcyan = ""
		bgwhite = ""

		bbgblack = ""
		bbgred = ""
		bbggreen = ""
		bbgyellow = ""
		bbgpurple = ""
		bbgmagenta = ""
		bbgcyan = ""
		bbgwhite = ""

		brightblack = ""
		brightred = ""
		brightgreen = ""
		brightyellow = ""
		brightpurple = ""
		brightmagenta = ""
		brightcyan = ""
		brightwhite = ""

		boldbrightblack = ""
		boldbrightred = ""
		boldbrightgreen = ""
		boldbrightyellow = ""
		boldbrightpurple = ""
		boldbrightmagenta = ""
		boldbrightcyan = ""
		boldbrightwhite = ""
	}
}
