package process

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

type Processor struct {
}

func NewProcessor() *Processor {
	return &Processor{}
}

func (p *Processor) LoadFromFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return p.LoadFromReader(file)
}

func (p *Processor) LoadFromReader(reader io.Reader) error {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		log.Println(line)
	}

	return nil
}

func (p *Processor) Update() {
	var possibilities = []string{"ubg", "grgw", "uwgb", "rubwug", "urbbw", "uuuuu", "uruwbug", "bgw", "wwwr", "ug", "gbwrug", "bug", "ubw", "gbb", "rur", "grrurw", "brbu", "ggrr", "bugrw", "ggrbg", "rwu", "gruru", "guur", "rbw", "brwg", "rru", "gwurbu", "bwwur", "urgb", "rrb", "rgug", "uwwr", "wuu", "gbbuu", "uurbuw", "guwww", "uwub", "rrbr", "gruubugg", "rrrw", "uugr", "rrgu", "gwuu", "rurrr", "wwuu", "rwbww", "ubbbb", "wbrug", "wrwuww", "urg", "uuurbrw", "wbgug", "gwrbg", "ugr", "rwgbub", "bwg", "wbrrgr", "uwuw", "wuwgbwrb", "wbw", "urur", "ubrgrr", "rgw", "ggugww", "grb", "bru", "gubr", "bgrr", "ugw", "rgbwg", "ugwgrr", "rguwwr", "rubu", "rurbrgu", "gbgwbr", "ugg", "wwrugbr", "wbruw", "ru", "rww", "rgrgwbr", "uwgr", "bbwb", "wgwwggu", "gwr", "wgw", "rrrubgu", "rrwbgbb", "burw", "rrr", "gg", "ubbwbwwb", "bgg", "brgb", "uwuru", "rrw", "uuu", "wur", "rrwg", "buug", "bubr", "rurrg", "uwwb", "wu", "rugb", "rguub", "buuw", "gwgub", "wwug", "rubrrw", "uguu", "bgbbbub", "ggwgbw", "rug", "uurwr", "wburrb", "wbu", "ugu", "gw", "ggrw", "gwrbwr", "rbggrwbu", "wugbg", "gbg", "rw", "wbrwurb", "rguuugrb", "bbb", "ugb", "bwguwuw", "ggr", "uru", "brr", "urw", "brrgg", "uub", "rbwu", "bugrg", "bwrgu", "rrbbgw", "wggrbgwu", "wrbbgb", "grrbgb", "brruwr", "rbgwrb", "guw", "rr", "wgguuwg", "rbg", "brwgrrrg", "gwwu", "gggwb", "wwbggwbg", "ggb", "wrg", "uwgrg", "gwb", "gww", "rurbbw", "wgub", "gu", "u", "rwgub", "gbr", "bbrbuwbw", "grw", "gwu", "rrgrgw", "uuwgrwr", "bwbruggb", "wggb", "bwb", "wrrugggb", "ubbuubg", "rubwu", "wbbr", "gbubr", "rbggww", "bbg", "uwbbu", "bbgb", "brw", "wrwggwr", "gr", "uubwwu", "gbrb", "wgwwrb", "urwgb", "brgbbg", "rwbgwg", "www", "rbuw", "bruwg", "grwbu", "gbww", "rbuugwrw", "wrbb", "bwgb", "gbubuu", "bu", "bgb", "ugwbru", "gwgb", "rbrwbgu", "wwubg", "uw", "grubw", "wrw", "wrb", "ubuw", "bgr", "grg", "bww", "uwggr", "bugg", "wugb", "urbwgb", "gggw", "ubrgruw", "uwrrb", "rbuu", "rwuubu", "guu", "gugu", "wbgggu", "bbggggr", "gug", "wbgub", "bb", "uww", "uug", "bgbgrbr", "gggr", "bgurg", "wrr", "rubgg", "bwrru", "ubguu", "wg", "rbgrr", "guugg", "ubu", "wruu", "wgb", "rgrburu", "wbwgr", "wb", "grbru", "bgwu", "gub", "brwr", "ggg", "wbggb", "wuwurwb", "uuw", "ubbbr", "gggbbg", "ruur", "bgug", "wwwg", "uggg", "rgg", "wwb", "ggwrgww", "bub", "wubg", "rggbbb", "brgbu", "ruuu", "wrub", "gur", "uggbr", "uur", "rgbbb", "gubruwwu", "bwu", "burugg", "rwg", "ugbw", "bwbgggw", "rburbgug", "wug", "ubgrb", "rurbb", "rwuug", "wr", "ggww", "rgwgu", "wrugu", "gwwbr", "uggbug", "buruuu", "guwgrg", "wru", "wwrbugg", "grgg", "rrbrr", "urgrgbg", "rbu", "uwu", "bugub", "gwru", "bbrg", "w", "wgu", "buw", "brbgwrb", "bgrwbwgu", "uwr", "wwurbr", "wuw", "wbgbgwb", "urbw", "guuggub", "rggggwu", "rrg", "rwbg", "rbgg", "ggwwbb", "gruruu", "gbwgr", "gubuu", "bgbr", "ururg", "wbr", "rwb", "rrrguug", "rgu", "wgru", "rubuwr", "ugbrg", "guuw", "urb", "wwwuurgb", "wwg", "urbrw", "grww", "brug", "bruu", "ggu", "grrurb", "rgwgwr", "guwrbrgu", "wrgugr", "gb", "wwr", "uwbguw", "brggub", "ugbb", "wgbg", "g", "wrrrgw", "ubbrb", "brb", "ruwrgub", "bgrwuwwg", "wgr", "bbgwrr", "ur", "buu", "ubr", "rrwwb", "wubbg", "uubug", "rrgw", "wwu", "wgwrw", "uggur", "urr", "rwr", "rgwrgw", "brggrr", "bbrubbw", "bgbgu", "guguw", "bbw", "gru", "bg", "ubwugu", "ubb", "brubug", "wugg", "brub", "ggwb", "bugrwgwr", "wgg", "ggw", "rrwr", "rub", "gbw", "rbbg", "gwrwgwr", "rbr", "grwr", "rgb", "b", "bbubrw", "bwwgu", "gwg", "uuguwgg", "brwgu", "rwuuguw", "ub", "ruw", "bwbb", "gbu", "uu", "bgrrggw", "uwwbgrg", "bw", "bbu", "wbb", "bwbu", "bbr", "wuuw", "grr", "rrbgr", "rgr", "gggbrru", "bgu", "rrguu", "uwg", "wbgw", "gbbbbggr", "rbrg", "wbbb", "rbrbrw", "brgbb", "uggb", "wwgu", "uugu", "ruu", "rrgb", "wub", "guwwww", "bbugwgu", "buwr", "rg", "uwubr", "gwwbbrr", "br", "bwr", "grrwrg", "wbbub", "ubugbr", "uugb", "rrgg"}
	var words = []string{
		"rgrruuurwuggbwgrggbrwruugrrrguwuubuwwrruwbbuubrwuwubgbwgur",
		"rubbrgrrbgwrgrrwgugwuwuwwwugbwruurwrugrrurbbgrur",
		"rbbrruggwwuurrgwgbububgrburrrwgbuubuwwgrrwwgbwbbr",
		"bgbrgrubgugrgugrwwwwuwruggbuubbgrwbgrbbbwu",
		"ugurgwruwgwbbbgrwugbwrugbwrrgrgubuubuubrurwgwugburguwwubgw",
		"rruggrwbuuwrbruwuwgugrwbgrwgwbbuwguwrgggbbuwurr",
		"ubbwuwgurwwwrrrrbgggwurrurwgwuwgbgwburrrrruguggwrurr",
		"bwrrbwgbwbrwubwubgwbbrruurrgbbruwgrwgbugguwg",
		"rwbuggwrrrburgwubbugwgrurgrrgruburwgwubrurgwbww",
		"ugrwbwwbgguugurwwrwrwrwbrbrwrububbbrguurbuww",
		"urgurbwwuuuuugurrwuuuubugruuwugbwwuugbwuggwrwwubbwrrgwwbgw",
		"ggurugbgbburwbwgbuubwbugwwbbubbbbgwuguurwgrbrrwuggubwwbrru",
		"ubuuggwgrbrgbbbrbrwuwrggubbbggguguwwrbbrbwgugggwbrgb",
		"rgwurrrwwbgrrrwbbrrubrrwrgbgurrgubuwrrrrgwguwwug",
		"ubuwgwguuguwrrguuggruuggrbrurwrrbgbuguurwwrbgruruwg",
		"gbbggubgwuguugbugurgrgggwwgrgrurwbguruwwbu",
		"wuwwbgbrwbrwrrruwuurgggrwgguburgugrgbuwrruuwrrwubbrbg",
		"wbbbuggbbruguuwurrgwgwurbuguuubrgubwggwbugguubub",
		"rbburuububugwgwugwrbubwrbrrubuubrwuwwbuguurruwr",
		"wwubwrgbbrwrwwubbgurruwwbgrgrburrbrgrrbgwgruurbgurwuburb",
		"bguwwrgrbrbrrgwgbbrrubuguurwwwrbggrwbuggbrwbw",
		"brgbbgrbgbbgwrguuwrbbrgwbwurwruwurugbuuubw",
		"wguwuwrugrubgwwrrrbwuwwuuuwugubwgwbbuugbwubgbwwuuurwwbuwrg",
		"uubgbrbubwwbgbwuuwgugubrguwgrbrwbwgwgubrbgbrgurw",
		"urbgbruugruuuwggbwbrrwwrbbbgrrugwgrgrugbrbrwwubwrbuuw",
		"uguggurwbrbbrugrgugwbbrbbgrwwbrwgburbbbruwubgwwr",
		"wwgbugguwuuurrgbbggrrubwguwguburgwwgrrrbrbrrwgbbubww",
		"wrrrubbrgurbgrwwrggrwgbwrgwwurbuwwwrgguubgrbgwwbbwbuwgu",
		"rbbrbuwbwwurbwrgwgbbwwrwrrgwgwgurgubwbbw",
		"uwbgrbwgrgbubrwrrgwbubgggrgubuwbrrguubbgbbrgggurubggwbwb",
		"ggbguurgrburbggwrguwbuwbruurbruurrgrbuurbbwguu",
		"rbrgrwbwbwwbgbrbrggugbgruuugbruwggbggggwbrbbw",
		"bbwubwburrubuggrbbgbgwrbgbggrbwgrururbwrugruuwrubugwbrrwb",
		"rrububugwrbruwubwuubuwgurbwruwuuwwbbrbrwruruwuuugb",
		"rbbubggrguwruuwuwrwbbuwguruwbggrbuuubwbuug",
		"rbbruuwuggwrwuwgrburwrwwwwgbbwbgrbggwuwgwgwb",
		"rbbrrrurgwggggubbugbgwuwbuguwuwubbbuwbwwgguggrbuuuugbrbwuwww",
		"rbburruurubuguwggurgrbggugrggubwgbuggbubg",
		"ugbrbrurruwubuubburwbwuwwwbwgrgguuuwwrbruubruuuwu",
		"uwrbrgbwugugwruruuwuugggbwbgwgwgwgrrgrgrgubwbrgwrb",
		"rbbbrwbrwrruburrgrgwbgugbbrbgubgggugubruwwuurrgbrbrww",
		"rbbrrbwrwbbbugrrwrurrgwrguuguuugwuuuwwuurwbguwugbgrwwrwrrw",
		"bguubbgbbbwwbuugbggbuuubbuwuuuuwwuwubrgbgbbgrgurugw",
		"rbbbbuwbbwggbbgwwbwurbgggguubwurwwrgrgwwwuubgbr",
		"grguubbgwwgugurwrwwuwgrurbubwgwrbugwggruwgbwrrrrrbrrbrrbu",
		"rbbbwgbggrwugbwbburbuuwguwwbrggwrbwburwrbrbuwburbugwug",
		"rbbbwgwbrwrbbuuwgruggugruwuuwugbugrrwbbbbrruruwbwgrrur",
		"rbbuwubrwrgwgrbrbbwgggruubrgrbuwruuuurruwwbbuwg",
		"rbgguwgbburuubbwwgwwuwwuwugbbwrbwwrubwgrwbrruubgwgu",
		"rbbbgwgurgubgbuwwwbgwugbruwrbrgbrwbgrrurrwbgwbbuwu",
		"ugubgwrgwgrrgwbbwwbgbuwgbrrrgbguuurgwgrgwggbu",
		"buuwrrwgurugbrubbwuwuuububbuwgubwuuuubbbubwrrwg",
		"uubggruguuwggubrgbruwrgwuubuguwrurgwgbuuwbuw",
		"gbrrrgurruururrbgugbwubruuwrurbugggbwrwgbbbwgbbuggu",
		"wwwbgwwwbggbugwwggrrbuwbugbbrguubwubwgugbrwuurbgg",
		"bwurbwbgrburwrbuwuuuruururrbbbgubggbrurrbwbuwg",
		"rrbwrbbgwbuubwgruwubrrgrgrurrwgrrgguugwgugbbwrg",
		"gugugrgguwbwbbwbgruwggrrbbgburbwbrwgbguubwubrgwgg",
		"wrgurbwrgrgwrwuuguggwuuwbgwurrrguuwurrbwwrgbuubburubbbbwrg",
		"uurbbbrbrubwrrrbrrgbuurubwbwrbuburwwwrwbub",
		"wrugwrugrrwguwwbguwuuurwuurbwwbgugbruugwbrbbwb",
		"guwgrwgwuwuwubwgwrrruugurgbuwrrrwgubgwwrgbubwwwuruu",
		"rbbbwuruwgggrrwwgubrwrbrrrwbuggrwgbbrrgbwb",
		"ggbgwbwwrbbuugbrwrbuwbbwwggwbrrbwbbwgurwwuuu",
		"wgrrwuubrburbbbuguwwurbrubgburgwgugbubrwwrugggbgggrwu",
		"rbbwugwgrwgrbwwgrrgrruwurubbgrwggbuggubrwwgrrrwwwwwugwrubrw",
		"bgruwwuwrrggwgwgbgbgrwgubuubrwgrrrwbgbrbgg",
		"bwbggwuggbubgurgburubgwrrwrbrguguwwururbwwuwrrgrub",
		"grrwguuwwruwbrbwgwbrwubrwuuggwgruwwrgbuubbrggubbbu",
		"rbbugbuwbburrwgwugwbrguruguwbgruubuggrgrgbgrw",
		"wgwuggugrrggrrrrgrwrgbuwruwugbbbrurbgruurbbrbbrb",
		"buwrrgwuugwwwgbbugrgwrwguggwwwuuuggbugbrubbuuuruwbwubuur",
		"rbbuwrrwubuggrwgugguwbuuuubgguuugbbugurggbgrbgbggbwrg",
		"urwbburwbrbrgbwrbbwgrgwwgugwuuubrrwbgbwbgrguugbwugrg",
		"buwrrgwwrbrbrrbrgbruwuurwggbbbrbguburubrwbwwbugbw",
		"rbbwggurgbuwgbrbwguurwggrbwgbuurubrrggguuruwuuguggbuw",
		"rgurrgwrwuwbubggrrugbrrrbbbgbbubwubrwgubbugrwwb",
		"rbbubwrubwgwbwgbwrrbugrwuugbguuguwggbwuugwbuwwbbbu",
		"bwbwrrubgugrgwbgbuggbrburubbguwggrwgbrguburbu",
		"rbrrbbwurwwbruwrugwbrggubbuwururrgbwburgrguwruwgugbubw",
		"gwrbwrbgbgbrwbgbbgrbbrwbbbbuwrbuwgurgbgrbwrgbgwgbubbbrrbuw",
		"wwgrrururrrbgruwrrguggrwrgrgurggwuuurguwwbggrugr",
		"rbbuwugrwwgwrbgubgbgrrrwggwubwgurgbrurwbrruwggggurgw",
		"wuggrguugbugbbwwrwbgrwuwwguwbbbbugwuggrugugbgb",
		"rbbuuuwurruwbrrrwruwuubbbugwuugbgwubwwwrw",
		"rbbwuwgwggubugrwbbggugrgrubugbgruwwuugrwruububgubgwg",
		"rbbbwbgggbuwuwwbrgwrgbrurbuurbwbwuububbuuwgbwgrgrww",
		"rbbbrwrbgbwuuwwwguggrbwugubgwrubwrbubrbgwwubugrugw",
		"ggwbugubrubugurwuwbubbbggrwggrgrgbgrrbubwuwrrwuwubgruuubw",
		"rbbwwuwrgrgwwwgwbwbggruwrbwuugrgbubugrubuwrwb",
		"wrwurbgbruggrgurwwuwbbrgugrwbwbgwuuwrbbuurbwubwwrwwugwbgr",
		"rbbwbggbubbuwwbggrgggrwbuguwgwrbgrwuwwggwbuuurbgrurb",
		"grrbgrwrgurugbwwbwuuuwuwuubugrguubwwgbguwgbrgbw",
		"wrubwruguwrgrgbgrwwururbggrbbugrwgbrwugbggwrbugg",
		"rbugwugbrbggugbwurrbuggburrubgwrggbgugrrrgu",
		"wgguuwwruwbggurruburgbuuwbuuwwgrbbrruburwrgrrwbwwuur",
		"rbbbgwwugugbbbwrbgbggrwwwwrgrbrgwrugbbwbbbbggburrgubwrugruu",
		"rgwbwubwgurbwwwuwrbuwubrrguwbruwgwrrwgububbrgwwguuwwbwgb",
		"rbbuggubgbwgbbrbubbgugrbrrubrwuuuwbbubwbwgwbwgu",
		"wgbggbbrubbbuggbggrbrrgruubuggurwbwuwggbbbug",
		"rbbwrruuuubwubwuurgrgbrrwbrggggrggwwwrbwbrurwubugwg",
		"rbbbrruuwggbwwbuwrbwugrrrwgggbgguuuwbwwrwrwwrubru",
		"bwbrbbwburbgugbwgbbbuuwgwurbwrubgwuuwubrbgwubwbwb",
		"gbwugbuububguuurbgwrwwwwgwburuuggbrwrrwwwgwbrguw",
		"ugrbwrwuuuwuwwgrggwbuwwrwbrbrrrrbgwwwurbbw",
		"wwguwwburuwrrugggbuburwugururuggrrgurbrrbrurrwwr",
		"rbbbruwrurwbubrurgwuubugggrgurggbbwbwwbburrwuwrrwuwwwrbb",
		"ggbrbrbrbbwurgbugwubbrbguugbwuwbgbwguwrrgwubbbwug",
		"rugrrgrbbrrbugbwgbgbrbrggwwbuuwbugrgubbuwrg",
		"rbbruuurgubuwgwubruggruwbrubugbbbrwwrbbugwugb",
		"wuwwgbgbubgwwrgbgrbwbubwurwbgrwgwugwwrgrguggur",
		"gwruubgguwbgguwugurwggrggbgrrwbrbgbwbwbrbwgwrwwgugbgrggu",
		"gbwgrrbrwbgbburbgbrwrugbugbbrwubrrbwgwugguurbrgwbgg",
		"rbbrgwwubrrbggbwwbwbbwbbgwwwgwbwruurbgbbbbggrwuu",
		"ugwwbbgrbwrwgrgrwrrbguwuuuuguugrrggubrwgrgrurwur",
		"grubgwwwgrgggurbuuggguuwrrgwrwugrugrrwwgugbwwbggbuuwbgrwu",
		"rbbrbrwuwguwguwuwuwurubrrgwrrguwbuugwwwuurgbrb",
		"rbbbbwrwbwbbbgbugrwgwrbubgggwgbwrwbbuwuwbbgbbr",
		"grrgbgwwggrwgwwburrubwurrbuurwwrwgwbwwuubbgu",
		"rbbrwwgbwwubgrurbbwwrugurugwbgrbruuwugrwbwgruwguwbgugrwggwu",
		"brgugrbwbwgbgrwguruwuurggwbubuurrwwuwurbrgbrwurwwgwrrbu",
		"uggrgrrbubgwburrrbbwbubggbugwrggrgbugwubwwwurrggrwguwb",
		"rbbrwrbbggrgbwrguwbbgwugbrgbrrrwrwguwuwguwwrwgru",
		"rbbrwrrrgubgwbrugurwwwwbgwwuggwbbubruwwrwgrgbbur",
		"rbbuwggbubrwbugbbbuurbubbrburbgugwbuggbwgwgbrbgww",
		"rbbugrugggurbwbrgrrwugrrugbwwwggbgubggbbbbuuwwwbgrrbu",
		"wbwrrgwwbguwurgrbgwrwuggrgubbbwubwgbwugggburr",
		"ugbrgrggwrrwwwrugbrrrrgwwuggbbrurwbrbgrrbgbbw",
		"rbbwgwggbuubwwrbwugrwbwrgwguwwrwrrugggbbgrugbrrubbgb",
		"gugbwgwurugrwrwwubbrrruwwubuurwgugggbugwwwbwr",
		"ubwruuugurbuwbuggurwrwbgbwruwuwbgbwwwuurgb",
		"rbbbwgrruguwwgrbwbuuuggbgrbgrrwwwgwrwbgr",
		"ugguwgwrugubgguuwwrwwggrwwrbrwubbwrugburgbwubbgwruurwbww",
		"rgwgbbgrrruwgbguggwrbrrwbbbbugggbrggrgwgruuu",
		"rbbubrrubwurbggwggrurwbrgbrgwgrgbugbbwgwgwgrbbgugbwwwgg",
		"wbgbwrrburugwwgbgggububggwgruwwrbwgwuwuruuwgurrbbrr",
		"wgrwbuugbrwgrrrgwwwwubgwugwwbubbrwwbrbgbrwrbgwuwwugrwbwu",
		"rbbrwggwggrwbguuubwgbgrbgbbrgbubgurrrrgrgrbwgwgb",
		"rbbugguurwguburgguggubrggrrbrbuurggbwbruggbbgbwrburwurgwrb",
		"brwgrbbbrrbrggurggruwwgubwbwburuuwgggugurrrwggbgubbb",
		"urrwbrwgbrbgwugbgrrguwrruwuwgugurgurrgguwwww",
		"ugrgbbgwbugbwuugrububuuuggbgrbrugwruwgrrrwgbrgbbbbrurwu",
		"brwwgwbwrurrwubwbuwrwrbrrgrbgwwbuwwrrugugrrgrw",
		"rbruggubwwgggwbgrwbwgurrrgbrggruuwburrwwggbbgruu",
		"rbbbrbrwbubgguubruggwgwburugbugbuubbuugrubwrrbrgwrbgb",
		"guwgrurwgrrrbwurrwwwwrugubrrbburburbgugwrwgwwwug",
		"rbbwwgbwuuubwrururuuwuwuuuuwuuubwgbgubgwggwub",
		"guugbgrgrugwugwrwbgrruuugubwwubgwuwbbbbgrgbw",
		"rbbwbugrwubrrrguwwbbbrruurwrwwbgwrrgguubwr",
		"rbrbrrurwggggrwrrguuugrbbbwgrgbgbwguwbwuruuwbbubuubggrwww",
		"rwrbwurrwrbubbubrbrgwuwbugrgwwubwwwbwgbbrurggrbbrwwwrwwuwr",
		"urruwwgwubrbwuubgugbgrrrrwruwgwrrrrbwwgwwrgwubwrbwru",
		"grrrggbwbbgubwbwgbrurbbrwwgugggbgubggwurbbugbuubbgwbgwgg",
		"rbbubrwgwruugguwbbugrgggrububuuurbrggggwurwbwrwurgwgruwwrr",
		"rgbwbrwrwwrubgrwuwwbrubugubwugrrruwubuurwwrubwu",
		"rbbruuwbrbgguuggbbbuwugbwrgrrrugwwbgrgbwbwruwbbwgwgbuw",
		"rbgrbgwbuuwwrbgguugruwuburrbugwwburbubugbggubgb",
		"grgurgubguwwrruwubgurbwwurgwgbuggbrrggbwugwrwrrguwuu",
		"wbgwugrgbwwrwuwbrbrrrbrggrrgrruwrwwgrrurwgrwwubugubwuu",
		"uubbbbubgbggrwgurggrrgbuwgbgugugrwrubgwwruugbwrwwuw",
		"ubbgbruruugbubuuwruugwubuwgrbrgrwggugbbrrgwgrwrwwrw",
		"wuwuuwggwgwgwbgwruuwuwrwuwbwrwrggwwrrwubwwg",
		"rbbbrggrrbrgwwuubrbuwgwrbgubuuuwburubrrrbbburwwuuubuw",
		"rbbbbubugrgurbrgbubrwbbgwgrwbgurubggrgwgrgrg",
		"guubwuwbuwubbuurgrwugrrrggrrguuruburbbuurburgwwwb",
		"rbbuwrwrbgugurgrgbgbrrrugbgruwrubbgbrrgurbrbgg",
		"buwgurgrruuuwggrwuburbgggbgubrbubuwgurwbbbwgb",
		"uguwruwbwurrguwwwurrwrwrgwgbbugrgwwguwbgggbwwrr",
		"uwguurguwuwgbwrbburrruwugburrbwugrwguwwwrubrg",
		"guuburgguwrbubbrbwrwwggwwwwrrbwwrgguuuuurru",
		"wuguguggugrwruugugbbugrrrrbrugugbbwgwubrrwrbubrrrwugguw",
		"rburguuubrrgbugrbgwuguwbbrgwgwrwrrwgbrbwwrrubw",
		"rbbuubrguwrubgrgggbubgwuubrubugbrwrwrubbrbrwgbwwgwurgug",
		"rbbburwgggwgbwrwrwrbggrwburubbbuuwwwuuwgbgwurwuwgbuwguggwgb",
		"wgbrgwrrwgubuuruugwwrurwugrbwrbugwgwgbwbguur",
		"bbwuurgrwurwguguwbrrrbrugbruuwwrrbuwuwrwrrr",
		"rbubuggubwwrrgbbwbgrururgwwuugbbbbbgbggrgbwgwubrrbwg",
		"rbbbrrgwwguruwgbbgrbrurrwbuwrgruwgrrwrwgwugwgrgu",
		"wuwbwwbugbgbwgurrrwurggwwugrrrburubugggwuuuwwrwgbbrr",
		"wrgubgwwgrwugubububbgwwuburguuwbgwguubrbrwuuu",
		"bubgbwbguwrurbbwgwbbwuggwwbuwgrgrrurrbrrrwrrrguwgruwurg",
		"rbbrwuwubgbrwugbbbggubbuguurgruuuwgbbgwwugrbbwwrrgbrbwgwuggw",
		"rbbbugwrrubrubwruubgrwggrbruwbrbuwbgwugrgbuburrwru",
		"rbbubbgrwbwguubrwgwgwrrbbgbwgrwugwuuuruwugbwrurwgrugggu",
		"rbbbrurbrrbbbrgrrrbugwbbgbuuwwubrgruwburbubwwwuww",
		"bgbuburrrrguwugbuuugwugggbbubgbbguwrbrguuuubuuwr",
		"rbbuuwguurgwwguggwggbwguguwrgrgbwgruugrwuuubrwburrwrbrgr",
		"wwrbggwrwrgbbbrbuuurgbgwurrrgwurgbbgrrgrrgugbbrrbbwrwwwgww",
		"ubwubrburrrbrgbbgugrrbgrurbbbbrbubbugurgwbgubwugbbg",
		"rbbbgbgubbgurguuurwgrwgurguwguwwrwwgburbuwbu",
		"uuuwwgwrwbubrwwuwrugbguubuuugbgrgurgugubbgrguwbu",
		"brguggrwguggguggrrgbwrugrwrbwgbwurwggwrbbbgugbbrbu",
		"rbbugrbugbbrrruurgwbwgggbguwbburuugurrwwwwgrbwbuurbwwrbrbbwg",
		"rbbuwgwbbugrbgrurbgwwwugrbwurubwuubbbwrrbbgruwgrwr",
		"grggbwgwbwgggrwrbrrgbbgwuuurbgruuubbbgubwwu",
		"rbrrwbrrwrgwbwwrwgwrwbbwuwguwguugwbbbgbugbugururbwwggg",
		"rbbwgrurwwwggubgbwrwwwrrbrbuugrbrurrubwrggrgugbwrurgru",
		"rbbbbwwugrurrrgwbgrwrggguwggggrrurgwgrbwbuwrgrwbgrugwbr",
		"rbbwrwwwwbwbrbbbrggbubwrubwgwubuubuuurwu",
		"rbbrwuubwwwugwwgrrbrggrbugwuurrwrgbwggrggwrgwuubrwbbuwgrrr",
		"ubwggrwguwwugwuwguuggwggbugbwrbbburgwuubuuwbuuurbrubwgbwgb",
		"grwuwrrrwurgbbrugwbgbwgrgggrgrggubugbbgrwurbgwgrb",
		"ubbwgwuwwurwgbubwrruggrwgbwgurrrwbwggrrbbubgb",
		"rbbrgbrubwbubbubururgrwubbbggwuugugubuggbuguww",
		"rburbbwwwurguuugrbuuruwgbgrggrgubbbugrwgbgwrgburrurgbbg",
		"gbbrwgugrggbwgugrubwbgbuwrbwwubbwuggurrrbwgwbbgw",
		"guwrgrgubbwbwwbrrubbbwwgbbbbrgubgwbbbwgwgwgwwwguruwwgbuggg",
		"rrwurbugbugwrwruwuwggubuwgwuuugwrugwbguuguwgrgbgggbb",
		"rbbbugubgrrubrrguggwrwbuuwwgrgbugwwugrubguuuwwuwugwr",
		"gwrwwwuuuurruwgwbuubrrurubwwrwwggrbwuwgggr",
		"wgbugwuwruwwuguuurgurrbubwrwuwwrbrgrubrugrubrbbu",
		"guwgwggwwurguwbrguwgwwgbbuuuuuggurbrwruwbugg",
		"rrwgrrgrwwwwggggwgwwbubbbbrgwguuwbbggrugwrgwurbwb",
		"rbbbrbbwurrugbwgrbbrbgugugbwwggbgugrwggwuwwgguurwb",
		"wrbbuuwbbgrbubruwgwrurggggggburggrwwbgurbugbggbwwbrubbrbw",
		"bugrrrwruuwrrburgrrbbbwuuurwrgubrwuwguubwurbrwrwbgub",
		"rbbwrwuwgggwbwuurgwuwruuuwuwgwgbrrguwuuwbwbr",
		"wuwgwuwwguburuugbbbrwruwurgggwrbrwwburruwubgggggwrubruww",
		"rbbuuwgugwwwgrrrgbubgrrubbuuwruuuwbggwugbrwubwwbrubrbwggb",
		"wwurwwwurwuwwbgubbwggwuguwgbwurugguugrrwwggb",
		"rbburgugurrrbwbubuururrbgwrrbruwuubbwrgwubwuwrgurbrggg",
		"uwgrgwggugrurwbwgrggwguugurubgruwgwrbbubbwuubrbrgggwb",
		"bbwwbbbwrguwwwwguwgrrubrgrrubugwuuguuwgguwuwbw",
		"rrubrbbuburggwwwggbwwrbwgrgrwurrrrgbuurrbru",
		"rgggbgwgwrwuwgbgggbugrrbbrurwgbgwbbwuuuwuwb",
		"bwuurggubrbwwwbwbwubggrgbwwwbrurwrrubuwwuwuu",
		"rbbwwrbbbgubbuwuggubgwrguwgrwbgruwrrwugwuubwwwrurgwgurugu",
		"rbbuuwggwwuwguwgruwwwggwbbwggbwbubrbrgubuguwuuubwrruuugwgru",
		"gurruwwbwgbuwbgbgbgbrgwbggbgwurrgrwwburgurbbrr",
		"wwuubwwurwrwgbgbrwwbwuugwbbbwuwwbgbgrwuwrur",
		"gbruuuurwwugwguuwrwuurggwburwbgwbburwwrbwbw",
		"wbbwuwrbbrrgubbuwbuwrruwrgbbrgwbgwrubbbgbgwr",
		"rguguuuuguwrrurrbrgbbrrbrwrugbgrubuugrbrwr",
		"wgbwuwwurbuwgrrbgubwgwbrgrwruugrrwrbgwrgrgrwwurbgrwbburbrw",
		"ggurruguwrbuwwrwggwubgrburbgbrwgbbwbuugrwwwgw",
		"grgbwgbgbuwburuwrrgwbgrbgubwbubugurrruwwguww",
		"rbbrwrurgrurwubrrrggrrbwugbgggbbrrbgbwggbrbbgrbguwrrrbw",
		"rbbugwrwrwbrwggwuwuwrrrwwuwugrwwwbuwwrrbbwubbubbbgr",
		"rubrguubwrbbwggurrugwuwgrbbgrruwburuuuwggwgwrbwrgwruub",
		"wgrubuugbrubbwrugwwubgrrburguuuwwbwbrurbrwwr",
		"uwrurwgrurubrbgbgbgwugrbwurgrugwrbbbwuwruwugbbg",
		"uuwbburbrbrbbuuubrwwrgbrwuwbgggrwwbrgwgggwgbrrbuugb",
		"uwuurgrgrbbrrrguugbbrwwwgburwuuwruugwggbwbwwbwwrrb",
		"ubwuwbbuwbrrrgrguuubugbwugrurrrwggrbgwubrwwrggugbwggwwu",
		"rbbuwrwwrgbuuwrbbbuwbruggrrgrwwwrugurubgwbwurbbwb",
		"rbbwuubuwrburrgwuururwrrgwbgrrugbgubrbuggbgbgwurgg",
		"wbuwruwwubgwgbugrwwgwwugruuurggubrwwuwwbrbrrwgwrrrwbgb",
		"rbuwwbbuuwubgrugubgugrbwuwgwubrwwbwugbuuwgg",
		"rbbrwgwwuwgggrbwggwbwrurwurrrwbwrwwrwubuwu",
		"rbbuwrwwwrrgbuuubruugguwwbuwgrrbgbbuwggrugrwrgbuu",
		"wuuubbggbrbgwgbuwwwubuuwbuguggbwuwrgubbwwggbrrurgwrw",
		"rbbbwrbrwrubrwwbggubwwbwwbggbrwgwwurrrurwbr",
		"wbwgwwugwbbwbuwrbwgrbrgrbwbwuuwurbuwrbwbbgwbgw",
		"ggwuwugurwurggwugrbgwrwuwrwugwuguwbbubbgwbguguburuubgubrr",
		"rbbbuwrwruugbubwubwbgbbrgrwgwwrwgguugwuugb",
		"gwwbbrbuwbwugguburubbruurbwwrugbwggwgbubgrrugrbru",
		"wwwbuuwwbwbgbwrubwwbgubbbgbbwbrbgwrbruguguur",
		"rbbrubruwuurgrrwbwurwugwbbguubrgguwrbrguggbubg",
		"bbgubbugwgurbgbruuguwgwwbrwgrwbbwrwbrgguwbuurwrbbub",
		"rbbwbwwbbwbruuurugguuuuwwuwbrurrrwuugwrbgwrrwgrwuwuwwru",
		"gbrbbrgbugrrbgwwbbwrwguuggbuwgrrrrbgrggwubbggrwugwugwbwr",
		"wugugrrrrgrwgrgggwggbgurwugubwurgwgugurwwbwuwubgbgbub",
		"rbbwgbrugwwgruubrgwububbwbugwbgbuwbbgwrgbuubrwgwbrwgbrb",
		"bggggrrrbrbgggbwrrwwbburbrwwruurbrrgwrrubrwrbbgwwuwuubrr",
		"rbbrgguuwgbrugugbwrrrbwuwuwubrwubugruurbuwgwubrbbgugu",
		"ubbgrgrbggwruubruwbrbrbwbubwuuwwwbwbgugrrrwrrubuugrwgrwg",
		"gbgwwubgrrwwururbuugwwurwubbwggubbwuuwbrububgbbrggwbwrwb",
		"rbbbrgwruwwbbubruurguwbgggrrguwggubwurbruuwrgrubuw",
		"gwwuubwrwwwwgwrwrbugruwugbwbrgrgrwrgwrbrgbugbgbb",
		"ubgwbrbbwgrgubuuwurgwrwwubwuwrugbrbwwwgbrrugbrruuwg",
		"wwwwwuwugrwwuwwbbggubburuurwwrgrwrggwbrwbrug",
		"rwgwwbwuuwguguguwrugruwwrwbgbuwgugubuugbrrwgggwuwbrbgrrub",
		"rbbuwrwggwrrbgwrbrrwwbubwwubruuwrgurbgbwrwbrg",
		"ggrrbuuwrbbgbbrguuubbgggbbbrggbgrburugrbwbg",
		"rgrwurruuwwrrrwgubuwubwrbwbuwururrbrrwruru",
		"wubuguruwuwwbwguwbgrurwrwwugbrgbwgwrbwgrrbgrrgrruwugwbug",
		"wguwubgwbugrbubgbbgbwbbwubbbwwruwguurrugrrurrubgg",
		"rbbuubgbbuggrwwrgbuwwwrbbwbwbgrgwruwrgwwruubugbbguwrbuwwb",
		"bubwbugbrwbgwgwbubugwbwgwrbrbbbbrguugbugwrrrbuurru",
		"bgruggrgwbuuwubbubggwgguggubbbgwuwurwuurgwbb",
		"rbbbwrwrugrwwrgwrbgguuuwwbugruuwrgbrgbbrbwurbugrguuwwwrubbrw",
		"wbwrburrgwbbwrbuwgwgrrwbwgwbbrbburrbgbrrbuwgwwrr",
		"uubguggwwwuurbbuggurwwrgrubgubuwwwgubwrgrruwgr",
		"rruwbbbuwbgwruuguggwgwrbrwrugruwwwgugubrgwuwwru",
		"bggbwrgggbbuugrwgugrgwuuwbbuuwubguurbruurbgwbgbgubwb",
		"brgbbgggbururguugrugbwbgwrgrrgwburburrwburbwugrrwwuww",
		"ubwbrbbuwgwrbbgwrguruurwurgggbbgrgwbgwwwurrbwbuwrgw",
		"rbbwwbggwbgbwugrwrugwbrbugugwrgwrbrgwbgwbgrbwrbwbwrr",
		"wgwgbrrrurgugurwgugbwugwwwrrwbgburgrwbbwbbgugbru",
		"gbbrubwuwwwbubggwwgrwgwuuwwrrugwbwwguubwuurwggrruub",
		"wrrurbwguwgrrbwurrwbwurgrgbgwbrubwuuwwbuubguguububuggrub",
		"uubbwwgggwrrbgbwbgubrruruuwbgrbbgrbubbbrwwu",
		"rbbbrgubgrbgrgrbgwurwgggrrggrggurwwruwbwwrwwuubbrbrrgrbuwuw",
		"rbbwwrwwgwuubbrbubrbrbrrbbwuugubbuggwwgwgwgruwgrubgbuguuwrb",
		"rbbbwrbruburbgwrburwwurbwruwrrbuwbwrbrruuuwruwbubbwugg",
		"wbgbgwugggwguuggbbrgwbrrwbruwwwgggwruwbggrwbgbwuwurb",
		"bgwwbgwgrwbwuwuwugrbbgrgbbrrgwwbuwbgwggrwwgg",
		"wrrubguwwgubwwuwguggbugrurwwgbbwrbgrrgrbwgbbbb",
		"rbbuwugubwwrbbwgbubrurbrugrruugwruuuubgrbb",
		"brbrruuurwwbrrrgrgwururuwrurbwrugwbugwubbgbggbwgw",
		"rbbbugwbwbbbrbwrrwwgbbwgwwgrbrgrrwgwrwwwgbwrg",
		"wugwuurbggwwrgwwbrggbbrbwrrgwgurbuubugugwruwwwbbwuwwbwurb",
		"rbbrgbrgubggwrbrwgbbbubgbrrurrgwbrggbbrwbuur",
		"rbbugbwwuwgbwrbgrwruuwgrrrurubgggwgwuurrbbwuu",
		"wgubrgwgugbrubgbgggrwwubwguwwbuurgubruwwubuuwgugwbbruugu",
		"gwggwbwbwwbuuwruwgwwuruugubrrurbggrgwbbuggwwwgwurgrwu",
		"bwbubuwggwbgrwbbrgrwrwguruwwwbwwwggugrgbrggrrbrgbw",
		"urwgrurwrgbgrubuuurgwwrwwwbrwuubgwwbrwwwuwbruur",
		"rbbuuuwbgurbwwbuwuugwrubbubuwbuburrrugbbg",
		"bbwbubwubbrrugrwurwgubrgbbgurwbbbguguwuugbrbbwbwwgrg",
		"gguwgrrrbgwggugwgwgrgbbrbwururbrgbbbwrgwuurrwuu",
		"rruuwbbwrwbbwrwrrrbbuguuwurruwbguwgggrgwbbrugugu",
		"wbrrrruuwrrrurgrburgrugbwrruwwwwgrwuuuggrbgruguuwb",
		"rbbwuuwwbbgggrwgugbgrwugrrgbugbwgwrrbrgwruubrg",
		"rbbwwrwuugbuwurgrgbgbbwugbrrbruwguururggbgruwb",
		"bwugbubruuwubrwuurubwbrgguugbbwrgrwuggggrbb",
		"wwrgubwbwgrbbbgbgbbubwwbggwbggwurwwrggrwbbuugwgrruuuw",
		"bubruguuwwrwbggbgrbbbrwwgrbrwubgurbuwgwbguurbrbgggbgrwwbwb",
		"bgrgwbrwwwbbrwbwurrbgrwgugwwgwwburbbwwrwbrrguruuuur",
		"rbbuubuwugbbuguguggbrwgbgrrubuurwguwugrbggbubbbrbrwrbugg",
		"rbbrbubgrrbrbgrurugggbwgruuwugggwwubwurrbgruugww",
		"rbbbrugguggubbrwgrrwwgurrurwrbbbbwgrrgbuwgurrrr",
		"brwwwbbgrbguwbbwurwwgrgrwwwgrruuwbruruguub",
		"rbbrbuwwwbwwbgbbwuwurubrbwuuwurguuwrrggwbu",
		"rgwwbgrgbuguwuuggurbugwwugrrrurrbuwwwbubrwgbgg",
		"urgwgbgbrruugrbwwubrruurbbrurwrugwbbrubgwuuwubguur",
		"wrbguwbuwbuwwurwgbbuuwubwgwuwbrwrgrbwbwbggwrb",
		"guugwuwbrrgwugbgrbrgbwrugwgbgrrubbwbbbgububwwuruggruwbrr",
		"rbwrbbruuuwgwgrwgwggrwrwwuguwgwwrbguwrrrwguruu",
		"bgwguubwubbbrwrrrugwwbrbggwgwgbubrgurggwbgbr",
		"rbbbbrbugwbrwgbrgugrrwurwuwwugbwrrwbwguuwbgburg",
		"rbbrggurrrgurugbgwgbrrwbuubbrwggbwgbugugbwwuubbwbggwg",
		"ggbbubwrwgrwrwbgbugwruruurruguubbbguwrgburgwgwurwbwruuwr",
		"rbbuuggruuwrrggrwwgrbgugruwurbbwrbguuguuuburbwbbrwrgwb",
		"bbwwguwuwburwgubwwwgwwububbrruwrbguugbwuubbrgwgguubggggrb",
		"rugwuggrgubgbrrwwbrrgwgrwbwwuguwgwbrrrubggggurwuwg",
		"rbbruwggrbgwuuwgbwguwurbubwuwurwuwwwrwrggbrgrugg",
		"wgurbrbbugbbrbgrrubggbgrgrrruugwbwbbgwgubwwbuubuwurgug",
		"ggwbrurgwwwgwwbbwgrbgwgrgbrrgbrbgrgurgugbrur",
		"rbbbwbggwgrbwggbrwuwuruuuwrbrrugrbuwgwwwu",
		"ggrbrrugugguuwugbbwgwgwbugrwbuubwbuggwbgurbw",
		"rbbubwrbrbrwwburggurrruwwrbgrgwruwgrgruuwbug",
		"rbbrugbwggbrwuwrrwuuwgubbruuwggbbwbrrrwrurwuuw",
		"wugwubrgbrububbgbwrrggwgguugwrguuburubwwrgruugbrubwwuggur",
		"rbbbwrbgbbgrggrwuwgguubruuuurburbrbggwurugwubugrur",
		"wrbgwwuwbgbbbbuwruwburwrgrugbruugubgwbbuwub",
		"rbbuwugbguwrrbbwwggwubrgrwuubugwubbwbrwuuggrgugbguwbb",
		"wwggwrwgbbgggrgrggbbguubgrrburgwrbuubwgwgrwubgbgugwrbr",
		"wgbbrurrrwbgbbwgbbwbwggbgugwrggrrruwrgbubuugrbrwwub",
		"rurbgurgbrrgggggurgwuwbbguurbrgbgwwwrgbgugbrwrguwwbw",
		"bwbgruuguwrrruurbrwuuwbugwuwuggwrrbbruruugggub",
		"rbbbbuwrbuguwwgurwruwrrrbgurugrurrwgbwrbuuggbrbbg",
		"rbbwgrrrbrwrbgubggwwgwggbwwwgrruwugugruub",
		"bgbgrbbwbuuwgwrbgwgbwgrwbgwbrgwrrguuugwwgurbrgubrrbb",
		"rbbubwbwrwgwbrwbgbbgubgurwubruwrgggwrurw",
		"gbgrggurwgwgbbwwgbwrwwgwbrrgwgbugbwubruwbugugrrwrrwuugwwrb",
		"ruwbwrbrburgwwwuuuuguwwwbwwwugwubugrwgwrrrgww",
		"uruuwwrbbrurgguggwrubbuuggrwbwurruuuwbwubuwbbuw",
		"uguuurgurrgrwwwbrgggwrurrubgggwugugwrgrwrwrugwurbubwgg",
		"wbbrrgrbwuwrugurubbuuuwuwuruwrgrgwbrruwwurgrwr",
		"bbbwurbgbwbrrgwgwwrbgurwbgrrwbuwwbwrrrbgug",
		"rbbubbwbwwbrrrwwuuggwrgruruuurrgwbugubbw",
		"uuwbwwugrgugwgrgwwbgwburbrugrrrbggwuburuwb",
		"brbbbwwgwbwurgubbwbggbrbbbuuwuwwbuwguwrrwrgbg",
		"rbbbwwugwrbguwgbuwbwgggugwurugwrgwrwbggbgbururbrrbgw",
		"rwrgurwrrgwwrrgwgrgwwwubrrbrurwrwbuugruwgrbgww",
		"wurrwwrgbrwugrrrggbbbuwwrbrugbbrggbruuugugwbbburgrrw",
		"ruurrubgbuggwugrwrwburrrbuggurubrbuwwguuwbr",
		"rgrggurgurgbrwurwwrggrubrgbwwurrwgrrubrwgggbb",
		"gwwrurwwuwurbrwwruwrwgubrurbwugrbrgurguwrbrggbgugrrwrrgbu",
		"gbwrgbrwuwugwbrwrbggrbbuwurrguwwwrgugrbwbruggbrwuurrb",
		"uwrwurwbbugbwbwubggbbrguurbuwgubwwguwgbbbbggrruu",
		"bbruburrbrbuurrrrgbgbrurbugwrbrwrgrbwbwbrguurwuwr",
		"rbbuwwrwrurbuwrgwwwwguuwubggrwbwwurbuurrwwwuggbbubrgurrguub",
		"rrbwruguubgbgbuubrrugurubwrggbrrgrgwwwguburbbgrgrurwrgb",
		"rbbuurwrwbrwwggrbwubrgwwrggugrurubwbbrrwuurrbruruw",
		"rbbuwguuruwwbgurrbgbbuwbwubuurbbrrwuugbwbrgrg",
		"wbbgrwgugbrgwubgrgwwrrgrrgwwgwuwuburwrbbwubguwug",
		"rbbrggbrurrgrbbgbbgbgburbuugwrwruwggurubgburgbbrur",
		"wugbrgbguwwwuuubwwwgwurbrubwwubgwgguwrbuwgbuurubrgwu",
		"wrbwruwwrwrgrrwwgrbrrubrwbwurwbuubbbguuurbbubrrr",
		"rbbrrggbbggwubwbwbbwwbbbrgurbuwwwuurbburgrgugur",
		"rbbuugruuwwbwgwguggbubwbbrbguguruggburrwbguwgbgggwwbbrw",
		"bburbuubuuurwruwbggbbbggrwwrbuwggguwuuwbuubuuwgugburwbrw",
		"uuwgwrwgbruurwurrgwwgrrguwwbrggrwbuwguggruuuurr",
		"rbbwugbgbbgrurbrwgwbbrwgrrrggggbuwgurwgurrbwr",
		"wwbrurbwruwgurwwwurwugbuwbbggbwbwurbrwuugwgbuwbbuwgw",
		"wurbggwbwuwgbgrbgwgbuuubwggwwwrburbuugwrwwugwrbgbgruwrburr",
		"wrwbbwururruguwrugbuuggrgbuburggrrrgrgugugwgrrrwggubgwrugr",
		"rbbbbrruwwburgrgbbgggububbbburrwwwbwbgrububgrwbuwbrrgwrbbrr",
		"bwgrbrbubugugrbbgwbwrwwbugubrrwguurgbrbgggburbbgurbb",
		"wbgbbbbwbburwgrbwuugubwrruuuwbrrwbrwbwwgbbruurgwwbuu",
		"brgbrbrrwubgrwwbuugwbgrbguwwgrguwbrgrbugubwwbrwgbug",
		"rbbuwgggwwbbbgurbrwrbugrruwbbburbbwbwbgwr",
		"rbbbwwrbbuggruugbwgubrruggurbbwgurrwgwuubbwrwggwgbrrw",
		"rbbwubgbbgrgbrggbbgwrrwgubrbbuwgruuwgwrubguwbwwuruuggurbuuw",
		"uuuwwggrrbgrbgwwrrbrurrgbrwwruwwrwuuurbrwbr",
		"gwbubbrgbburugrbruuuuwgbwrrrbrrrwrrgbgrwruguurrrugrgrggwb",
		"rbuugrgbbrgubuuuuwwwrwwubburrrrgbubbbwrwwuubwrwb",
		"rbbwububruuwgwrurbburrwgurbuggrrwbruwbwgwbuwbb",
	}
	// var possibilities = []string{"r", "wr", "b", "g", "bwu", "rb", "gb", "br"}
	// var words = []string{"rrbgbr"}
	// var words = []string{"rgrruuurwuggbwgrggbrwruugrrrguwuubuwwrruwbbuubrwuwubgbwgur"}

	var findWord func(word string)
	var foundMatches int = 0

	findWord = func(word string) {
		if len(word) == 0 {
			foundMatches++
			return
		}

		for i := 0; i < len(possibilities); i++ {
			possibility := possibilities[i]
			if len(word) < len(possibility) {
				continue
			}

			start := word[0:len(possibility)]
			if start == possibility {
				newWord := word[len(possibility):]
				// fmt.Printf("found match, new word: %s\n", newWord)
				findWord(newWord)
			}
		}
		// for {
		// 	startLength := len(word)
		// 	for _, possibility := range possibilities {
		// 		if len(word) < len(possibility) {
		// 			continue
		// 		}

		// 		start := word[0:len(possibility)]
		// 		if start == possibility {
		// 			fmt.Println("found match", possibility)
		// 			word = word[len(possibility):]
		// 			fmt.Println("new word", word)
		// 			break
		// 		}
		// 	}
		// 	if len(word) == 0 {
		// 		fmt.Println("found match")
		// 		foundMatches++
		// 		break
		// 	}
		// 	if startLength == len(word) {
		// 		fmt.Println("no match")
		// 		break
		// 	}
		// }
	}

	var i = 0
	for _, word := range words {
		fmt.Printf("Processing word %d of %d\n", i, len(words))
		n := len(word)
		dp := make([]int, n+1)
		dp[0] = 1
		for i := 1; i <= n; i++ {
			for _, possibility := range possibilities {
				possibility_len := len(possibility)
				if i >= possibility_len && word[i-possibility_len:i] == possibility {
					dp[i] += dp[i-possibility_len]
				}
			}
		}
		foundMatches += dp[n]
	}

	log.Println(foundMatches)
}
