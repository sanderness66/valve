// TUBE -- decode tube designations from stdin or cmdline
//
// SvM 03-JAN-2021
//
// todo c: complete ph_bases[]

package main

import (
	"fmt"
	"github.com/chzyer/readline"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const maxdev = 5 // would be nice if it weren't hard-coded but I've never seen more than two
type tube struct {
	heater  string
	dev     [maxdev]string
	base    string
	variant string // we fill this in but don't acutally print is as it has no deeper meaning
	aliases string
	notes   string // might contain extra info, or error msg if parse_* funcs return false
}

// return kind of tube designation
func findcontinent(s string) string {
	if matched, _ := regexp.MatchString(`^[A-Z][A-Z]+[0-9]`, s); matched {
		return "philips"
	} else if matched, _ := regexp.MatchString(`^Z[0-9]+[A-Z]`, s); matched {
		return "philipsz"
	} else if matched, _ := regexp.MatchString(`^[0-9]+[A-Z]`, s); matched {
		return "retma"
	} else {
		return "other"
	}
}

func parse_philips(s string) (tube, bool) {
	var devs, num, rest string
	var tube tube

	re := regexp.MustCompile(`^[A-Z]|[A-Z]+|[0-9]+`)
	parts := re.FindAllString(s, -1)

	switch len(parts) {
	case 4:
		rest = parts[3]
		fallthrough
	case 3:
		tube.heater = parts[0]
		devs = parts[1]
		num = parts[2]
	default:
		tube.notes = "unknown (couldn't parse designation as Philips-Mullard)"
		return tube, false
	}

	tube.heater = ph_heater[tube.heater]

	for devnum, d := range devs {
		dd := string(d)
		if devnum == maxdev {
			tube.notes = "warning: found more than " + strconv.Itoa(maxdev) + " devices"
			break
		}
		tube.dev[devnum] = ph_device[dd]
	}

	xx, _ := strconv.Atoi(num)
	tube.base = ph_base[xx/10]
	tube.variant = strconv.Itoa(xx%10) + rest

	for key, val := range other_to_ph {
		if val == s {
			tube.aliases += key + " "
		}
	}

	return tube, true
}

func parse_philipsz(s string) (tube, bool) {
	re := regexp.MustCompile(`^[A-Z]|[0-9]+|[A-Z]+`)
	parts := re.FindAllString(s, -1)

	var tube tube
	if len(parts) == 3 {
		tube.heater = "N/A"
		tube.dev[0] = ph_z[parts[2]]
		return tube, true
	}
	tube.notes = "unknown (couldn't parse designation as Philips-Mullard Z series)"
	return tube, false
}

func parse_retma(s string) (tube, bool) {
	re := regexp.MustCompile(`[A-Z]+|[0-9]+`)
	parts := re.FindAllString(s, -1)

	// retma names are not too informative, so try to find a philips equivalent and use that instead first
	// use parts[0:3] for this since any suffix in parts[4] won't be in the lookup map
	if len(parts) >= 3 {
		key := strings.Join(parts[0:3], "")
		eq := other_to_ph[key]
		if eq != "" {
			tube, res := parse_philips(eq)
			if res {
				// add eq as it won't be in eq's own aliases list
				tube.aliases += eq
				// but remove (suffix-less) self as it will be
				tube.aliases = strings.ReplaceAll(tube.aliases, key+" ", "")
				return tube, true
			}
		}
	}

	// no equiv found -> just find out what little we can
	var rest string
	switch len(parts) {
	case 4:
		rest = parts[3]
		fallthrough
	case 3:
		var tube tube
		tube.heater = parts[0] + "V"
		tube.variant = parts[1] + rest
		tube.notes = parts[2] + " elements"
		return tube, true
	default:
		var tube tube
		tube.notes = "unknown (couldn't parse designation as R.E.T.M.A.)"
		return tube, false
	}
}

func parse_other(s string) (tube, bool) {
	// just hope there's a philips eq
	eq := other_to_ph[s]
	if eq != "" {
		tube, res := parse_philips(eq)
		if res {
			// add eq as it won't be in eq's own aliases list
			tube.aliases += eq
			// but remove (suffix-less) self as it will be
			tube.aliases = strings.ReplaceAll(tube.aliases, s+" ", "")
			return tube, true
		}
	}

	var tube tube
	tube.notes = "tube not found"
	return tube, false
}

func feed(s string) {
	if s == "" || s[0] == '#' {
		fmt.Println(s)
		return
	}
	s = strings.ToUpper(s)

	var des tube
	var res bool
	typ := findcontinent(s)
	switch typ {
	case "philips":
		des, res = parse_philips(s)
	case "philipsz":
		des, res = parse_philipsz(s)
	case "retma":
		des, res = parse_retma(s)
	case "other":
		des, res = parse_other(s)
	}

	if res {
		fmt.Println(s)
		fmt.Printf("heater:\t\t%s\n", des.heater)
		if des.base != "" {
			fmt.Printf("base:\t\t%s\n", des.base)
		}

		for devnum, d := range des.dev {
			if d != "" {
				fmt.Printf("function:\t%d. %s\n", devnum+1, d)
			}
		}
		if des.notes != "" {
			fmt.Printf("notes:\t\t%s\n", des.notes)
		}
		if des.aliases != "" {
			fmt.Printf("aliases:\t%s\n", strings.TrimSpace(des.aliases))
		}
	} else {
		fmt.Printf("%s: %s\n", s, des.notes)
	}
	fmt.Println()
}

func main() {
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			feed(arg)
		}
	} else {
		rl, err := readline.New(". ")
		if err != nil {
			panic(err)
		}
		defer rl.Close()

		for {
			line, err := rl.Readline()
			if err != nil {
				break
			}
			feed(line)
		}
	}
}
