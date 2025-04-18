package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/fatih/color"
)

// Message templates with dark humor for encouraging Friday deployments
var messageTemplates = []string{
	"Push To Production On Friday, What Could Go Wrong?",
	"Friday Deploys: Because Weekends Are Overrated",
	"Deploy On Friday, Debug On Monday",
	"Breaking Prod On Friday Is A Power Move",
	"Friday 4PM: Perfect Time For Major System Changes",
	"Deploy On Friday, Blame It On Monday",
	"Your Weekend Plans Are Less Important Than This Merge",
	"Friday Deploys: Creating Job Security Since Forever",
	"Hit Deploy And Run: The Friday Special",
	"Nothing Says 'I Trust My Code' Like A Friday Deployment",
	"Friday Deployments: The Russian Roulette Of Software Engineering",
	"Deploy On Friday, Turn Off Notifications",
	"If It Works On Friday, It Works Forever* (*Terms And Conditions Apply)",
	"Push Now, Pray Later: The Friday Mantra",
}

// ASCII art styles for the meme boxes
var boxStyles = []struct {
	topLeft     string
	topRight    string
	bottomLeft  string
	bottomRight string
	horizontal  string
	vertical    string
}{
	{"┌", "┐", "└", "┘", "─", "│"},     // Simple box
	{"╔", "╗", "╚", "╝", "═", "║"},     // Double line box
	{"+-", "-+", "+-", "-+", "-", "|"}, // ASCII simple
	{"*", "*", "*", "*", "*", "*"},     // Star box
	{"▛", "▜", "▙", "▟", "▀", "▌"},     // Block style
}

// Improved and expanded ASCII art figures for beneath the box
var figures = [][]string{
	{
		"          _\\/_        ",
		"           /\\         ",
		"           /\\         ",
		"          /  \\        ",
		"         |    |       ",
		"         |    |       ",
		"       __|    |__     ",
		"      /__|    |__\\    ",
	},
	{
		"        (•̀ᴗ•́)و       ",
		"         /|\\         ",
		"        / | \\        ",
		"         / \\         ",
		"        /   \\        ",
		"       👟    👟       ",
	},
	{
		"       \\(^o^)/       ",
		"         |__|         ",
		"          ||          ",
		"          ||          ",
		"         /  \\         ",
		"        /    \\        ",
	},
	{
		"      _____           ",
		"     /     \\          ",
		"    | ^   ^ |         ",
		"    |   ω   |         ",
		"     \\_____/          ",
		"        ||            ",
		"       /||\\           ",
		"      / || \\          ",
		"        ⅃⅃            ",
	},
	{
		"        (҂◡_◡)        ",
		"       ᕦ(╭ರ ⊙ ͜ʖ⊙ರ)ᕥ    ",
		"          |           ",
		"         /|\\          ",
		"        / | \\         ",
		"         J L          ",
	},
	{
		"       (╯°□°)╯        ",
		"       ┻━━┻           ",
		"        /|\\           ",
		"       / | \\          ",
		"        / \\           ",
		"       /   \\          ",
	},
	{
		"        _ノ乙            ",
		"       ( ͡° ͜ʖ ͡°)       ",
		"       /  ⌒\\         ",
		"      /   |  \\        ",
		"     /   /|   \\       ",
		"         / \\          ",
		"        /   \\         ",
	},
	{
		"        ______        ",
		"       /      \\       ",
		"      | ⊙  ⊙  |       ",
		"      |   ▽    |      ",
		"       \\______/       ",
		"         |  |         ",
		"        ┌|  |┐        ",
		"        └|  |┘        ",
		"         |__|         ",
		"         /  \\         ",
	},
	{
		"        (ง •̀_•́)ง      ",
		"          |           ",
		"          |           ",
		"         /|\\          ",
	},
	{
		"       \\( ⌐■_■)/      ",
		"         \\| |/        ",
		"          | |         ",
		"         /| |\\        ",
		"        / | | \\       ",
		"          ⅃ Ⅎ          ",
	},
	{
		"         /|\\          ",
		"        /*|*\\         ",
		"       /* | *\\        ",
		"      /*  |  *\\       ",
		"     /*   |   *\\      ",
		"    /*    |    *\\     ",
		"   /*************\\    ",
		"          |  |         ",
		"          |  |         ",
		"         /|  |\\       ",
	},
	{
		"      ┌─┐             ",
		"      ┴─┴             ",
		"   ಠ_ಠ                ",
		"   <|>                ",
		"   /ω\\                ",
	},
	{
		"        _______       ",
		"       /       \\      ",
		"      | ⇀ ‿ ↼  |      ",
		"      |  SOON   |     ",
		"       \\_______/      ",
		"       W  | |  W      ",
		"          | |         ",
		"          | |         ",
		"         // \\\\        ",
	},
	{
		"       .---------.    ",
		"      / .-------. \\   ",
		"     / /         \\ \\  ",
		"     | |         | |  ",
		"    _| |_________| |_ ",
		"  .' |_|         |_| '.",
		"  '._____ ____ _____.'",
		"  |     .'____'.     |",
		"  '.__.'.'    '.'.__.'",
		"  '.__  |      |  __.'",
		"  |   '.'.____.'.'   |",
		"  '.____'.____.'____.'",
		"  '.________________.'",
	},
	{
		"         🔥  🔥         ",
		"        🔥    🔥        ",
		"       🔥 PROD 🔥       ",
		"        🔥    🔥        ",
		"         🔥  🔥         ",
	},
}

// Different background styles for the meme
var backgroundPatterns = []string{
	"", // No background
	"░░░░░░░░░░░░░░░░░░", // Light dots
	"▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒", // Medium dots
	"▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓", // Heavy dots
	"················",   // Dots
	"~~~~~~~~~~~~~~~~",   // Waves
	"////////////////",   // Slashes
	"××××××××××××××××",   // Crosses
}

// Chaos descriptors for the random failure probabilities
var chaosDescriptors = []string{
	"Probable System Downtime: %d%%",
	"Weekend Ruined Probability: %d%%",
	"On-Call Alert Chance: %d%%",
	"Production Fire Risk: %d%%",
	"Angry Manager Likelihood: %d%%",
	"Stack Overflow Visits Needed: %d",
	"Career Impact Rating: %d/10",
	"Blame Deflection Difficulty: %d/10",
	"Monday Morning Regret Factor: %d/10",
	"Coffee Cups Required For Aftermath: %d",
}

func main() {
	// Define command line flags
	customMsg := flag.String("message", "", "Custom message for the meme")
	colorful := flag.Bool("color", true, "Use colorful output")
	boxStyle := flag.Int("style", -1, "Box style (0-4, -1 for random)")
	figureStyle := flag.Int("figure", -1, "Figure style (0-14, -1 for random)")
	bgPattern := flag.Int("bg", -1, "Background pattern (0-7, -1 for random)")
	chaosLevel := flag.Int("chaos", -1, "Chaos level data (0-5, -1 for random)")
	width := flag.Int("width", 50, "Width of the meme box")
	help := flag.Bool("help", false, "Display help information")
	listFigures := flag.Bool("list-figures", false, "Display all available figures")

	flag.Parse()

	if *help {
		displayHelp()
		return
	}

	if *listFigures {
		displayAllFigures()
		return
	}

	// Seed the random number generator
	rand.New(rand.NewSource(time.Now().UnixNano()))

	// Choose the message (either custom or random template)
	message := *customMsg
	if message == "" {
		message = messageTemplates[rand.Intn(len(messageTemplates))]
	}

	// Select styles based on flags or randomly
	selectedBoxStyle := *boxStyle
	if selectedBoxStyle < 0 || selectedBoxStyle >= len(boxStyles) {
		selectedBoxStyle = rand.Intn(len(boxStyles))
	}

	selectedFigure := *figureStyle
	if selectedFigure < 0 || selectedFigure >= len(figures) {
		selectedFigure = rand.Intn(len(figures))
	}

	selectedBg := *bgPattern
	if selectedBg < 0 || selectedBg >= len(backgroundPatterns) {
		selectedBg = rand.Intn(len(backgroundPatterns))
	}

	// Generate the chaos metrics
	chaosMetrics := generateChaosMetrics(*chaosLevel)

	// Generate and print the meme
	generateMeme(message, boxStyles[selectedBoxStyle], figures[selectedFigure],
		backgroundPatterns[selectedBg], *width, *colorful, chaosMetrics)
}

func displayHelp() {
	fmt.Println("Friday Deployment Meme Generator - Because who needs a peaceful weekend?")
	fmt.Println("")
	fmt.Println("This tool generates dark humor memes about deploying to production on Friday.")
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("  friday-deploy [options]")
	fmt.Println("")
	fmt.Println("Options:")
	fmt.Println("  -message string   Custom message for the meme (default: random from templates)")
	fmt.Println("  -color            Use colorful output (default: true)")
	fmt.Println("  -style int        Box style (0-4, -1 for random) (default: -1)")
	fmt.Println("  -figure int       Figure style (0-14, -1 for random) (default: -1)")
	fmt.Println("  -bg int           Background pattern (0-7, -1 for random) (default: -1)")
	fmt.Println("  -chaos int        Chaos level data (0-5, -1 for random) (default: -1)")
	fmt.Println("  -width int        Width of the meme box (default: 50)")
	fmt.Println("  -list-figures     Display all available figures")
	fmt.Println("  -help             Display this help information")
	fmt.Println("")
	fmt.Println("Examples:")
	fmt.Println("  friday-deploy")
	fmt.Println("  friday-deploy -message \"YOLO Deploy Friday\" -style 2 -figure 3")
	fmt.Println("  friday-deploy -color=false -width 60")
	fmt.Println("")
	fmt.Println("Remember: With great power comes great deniability.")
}

func displayAllFigures() {
	fmt.Println("Available Figures:")
	fmt.Println("")

	for i, figure := range figures {
		fmt.Printf("Figure %d:\n", i)
		for _, line := range figure {
			fmt.Println(line)
		}
		fmt.Println("")
	}
}

func generateChaosMetrics(level int) []string {
	metrics := make([]string, 0)

	// Choose how many metrics to display (1-4)
	numMetrics := 1
	if level < 0 || level > 5 {
		numMetrics = rand.Intn(3) + 2 // 2-4 metrics
	} else {
		numMetrics = level
	}

	// Create a copy of the descriptors to avoid repeats
	descriptors := make([]string, len(chaosDescriptors))
	copy(descriptors, chaosDescriptors)

	// Shuffle the descriptors
	rand.Shuffle(len(descriptors), func(i, j int) {
		descriptors[i], descriptors[j] = descriptors[j], descriptors[i]
	})

	// Select and format the metrics
	for i := 0; i < numMetrics && i < len(descriptors); i++ {
		var value int
		if strings.Contains(descriptors[i], "/10") {
			value = rand.Intn(5) + 6 // 6-10 for /10 scales (always bad)
		} else if strings.Contains(descriptors[i], "%%") {
			value = rand.Intn(50) + 50 // 50-99%
		} else {
			value = rand.Intn(8) + 3 // 3-10 for countable things
		}
		metrics = append(metrics, fmt.Sprintf(descriptors[i], value))
	}

	return metrics
}

func generateMeme(message string, boxStyle struct {
	topLeft, topRight, bottomLeft, bottomRight, horizontal, vertical string
}, figure []string, bgPattern string, width int, colorful bool, chaosMetrics []string) {
	// Validate and adjust width
	messageLen := len(message)
	minWidth := messageLen + 4 // minimum width needed for message + padding
	if width < minWidth {
		width = minWidth
	}

	// Setup colors if enabled
	titleColor := color.New(color.FgRed, color.Bold)
	boxColor := color.New(color.FgCyan)
	figureColor := color.New(color.FgYellow)
	bgColor := color.New(color.FgHiBlack)
	metricColor := color.New(color.FgMagenta, color.Bold)

	if !colorful {
		titleColor.DisableColor()
		boxColor.DisableColor()
		figureColor.DisableColor()
		bgColor.DisableColor()
		metricColor.DisableColor()
	}

	// Create the top border with validation
	topBorder := boxStyle.topLeft + strings.Repeat(boxStyle.horizontal, width-2) + boxStyle.topRight

	// Create the message line with safe padding calculation
	totalPadding := width - 2 - messageLen     // total available padding
	paddingLeft := totalPadding / 2            // half for left side
	paddingRight := totalPadding - paddingLeft // remainder for right side

	messageLine := boxStyle.vertical +
		strings.Repeat(" ", paddingLeft) +
		message +
		strings.Repeat(" ", paddingRight) +
		boxStyle.vertical

	// Create the bottom border
	bottomBorder := boxStyle.bottomLeft + strings.Repeat(boxStyle.horizontal, width-2) + boxStyle.bottomRight

	// Print the background if there is one
	if bgPattern != "" {
		bgPadding := (width - len(bgPattern)) / 2
		if bgPadding >= 0 {
			padding := strings.Repeat(" ", bgPadding)
			for i := 0; i < 3; i++ {
				_, err := bgColor.Println(padding + bgPattern)
				if err != nil {
					return
				}
			}
		}
	}

	// Print the box with the message
	_, err := boxColor.Println(topBorder)
	if err != nil {
		return
	}

	_, _ = titleColor.Println(messageLine)

	_, _ = boxColor.Println(bottomBorder)

	// Print the figure with safe position calculation
	for _, line := range figure {
		figurePos := (width - len(line)) / 2
		if figurePos >= 0 {
			_, err := figureColor.Println(strings.Repeat(" ", figurePos) + line)
			if err != nil {
				return
			}
		} else {
			_, err := figureColor.Println(line)
			if err != nil {
				return
			} // print without padding if the width is too small
		}
	}

	// Print chaos metrics
	fmt.Println()
	for _, metric := range chaosMetrics {
		_, err := metricColor.Printf("  • %s\n", metric)
		if err != nil {
			return
		}
	}

	// Add a snarky footer
	timeNow := time.Now()
	if timeNow.Weekday() == time.Friday && timeNow.Hour() >= 16 {
		fmt.Println()
		_, err := color.New(color.FgHiRed, color.Bold).Println("  ⚠️  IT'S ACTUALLY FRIDAY AFTERNOON RIGHT NOW. DO IT. DO IT.")
		if err != nil {
			return
		}
	}
}
