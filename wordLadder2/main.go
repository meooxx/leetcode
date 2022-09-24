package main

import (
	"fmt"
)

func main() {
	fmt.Println(findLadders("hit", "cog", []string{
		"hot", "dot", "dog", "lot", "log", "cog",
	}))
	fmt.Println(findLadders("a", "c", []string{
		"a", "b", "c",
	}))
	fmt.Println(findLadders("red", "tax", []string{
		"ted", "tex", "red", "tax", "tad", "den", "rex", "pee",
	}))
	fmt.Println(findLadders("lost", "cost", []string{
		"most", "fist", "lost", "cost", "fish",
	}))
	fmt.Println(findLadders("nape", "mild", []string{
		"dose", "ends", "dine", "jars", "prow", "soap", "guns", "hops", "cray", "hove", "ella", "hour", "lens", "jive", "wiry", "earl", "mara", "part", "flue", "putt", "rory", "bull", "york", "ruts", "lily", "vamp", "bask", "peer", "boat", "dens", "lyre", "jets", "wide", "rile", "boos", "down", "path", "onyx", "mows", "toke", "soto", "dork", "nape", "mans", "loin", "jots", "male", "sits", "minn", "sale", "pets", "hugo", "woke", "suds", "rugs", "vole", "warp", "mite", "pews", "lips", "pals", "nigh", "sulk", "vice", "clod", "iowa", "gibe", "shad", "carl", "huns", "coot", "sera", "mils", "rose", "orly", "ford", "void", "time", "eloy", "risk", "veep", "reps", "dolt", "hens", "tray", "melt", "rung", "rich", "saga", "lust", "yews", "rode", "many", "cods", "rape", "last", "tile", "nosy", "take", "nope", "toni", "bank", "jock", "jody", "diss", "nips", "bake", "lima", "wore", "kins", "cult", "hart", "wuss", "tale", "sing", "lake", "bogy", "wigs", "kari", "magi", "bass", "pent", "tost", "fops", "bags", "duns", "will", "tart", "drug", "gale", "mold", "disk", "spay", "hows", "naps", "puss", "gina", "kara", "zorn", "boll", "cams", "boas", "rave", "sets", "lego", "hays", "judy", "chap", "live", "bahs", "ohio", "nibs", "cuts", "pups", "data", "kate", "rump", "hews", "mary", "stow", "fang", "bolt", "rues", "mesh", "mice", "rise", "rant", "dune", "jell", "laws", "jove", "bode", "sung", "nils", "vila", "mode", "hued", "cell", "fies", "swat", "wags", "nate", "wist", "honk", "goth", "told", "oise", "wail", "tels", "sore", "hunk", "mate", "luke", "tore", "bond", "bast", "vows", "ripe", "fond", "benz", "firs", "zeds", "wary", "baas", "wins", "pair", "tags", "cost", "woes", "buns", "lend", "bops", "code", "eddy", "siva", "oops", "toed", "bale", "hutu", "jolt", "rife", "darn", "tape", "bold", "cope", "cake", "wisp", "vats", "wave", "hems", "bill", "cord", "pert", "type", "kroc", "ucla", "albs", "yoko", "silt", "pock", "drub", "puny", "fads", "mull", "pray", "mole", "talc", "east", "slay", "jamb", "mill", "dung", "jack", "lynx", "nome", "leos", "lade", "sana", "tike", "cali", "toge", "pled", "mile", "mass", "leon", "sloe", "lube", "kans", "cory", "burs", "race", "toss", "mild", "tops", "maze", "city", "sadr", "bays", "poet", "volt", "laze", "gold", "zuni", "shea", "gags", "fist", "ping", "pope", "cora", "yaks", "cosy", "foci", "plan", "colo", "hume", "yowl", "craw", "pied", "toga", "lobs", "love", "lode", "duds", "bled", "juts", "gabs", "fink", "rock", "pant", "wipe", "pele", "suez", "nina", "ring", "okra", "warm", "lyle", "gape", "bead", "lead", "jane", "oink", "ware", "zibo", "inns", "mope", "hang", "made", "fobs", "gamy", "fort", "peak", "gill", "dino", "dina", "tier",
	}))
	fmt.Println(findLadders("aaaaa", "ggggg", []string{

		"aaaaa", "caaaa", "cbaaa", "daaaa", "dbaaa", "eaaaa", "ebaaa", "faaaa", "fbaaa", "gaaaa", "gbaaa", "haaaa", "hbaaa", "iaaaa", "ibaaa", "jaaaa", "jbaaa", "kaaaa", "kbaaa", "laaaa", "lbaaa", "maaaa", "mbaaa", "naaaa", "nbaaa", "oaaaa", "obaaa", "paaaa", "pbaaa", "bbaaa", "bbcaa", "bbcba", "bbdaa", "bbdba", "bbeaa", "bbeba", "bbfaa", "bbfba", "bbgaa", "bbgba", "bbhaa", "bbhba", "bbiaa", "bbiba", "bbjaa", "bbjba", "bbkaa", "bbkba", "bblaa", "bblba", "bbmaa", "bbmba", "bbnaa", "bbnba", "bboaa", "bboba", "bbpaa", "bbpba", "bbbba", "abbba", "acbba", "dbbba", "dcbba", "ebbba", "ecbba", "fbbba", "fcbba", "gbbba", "gcbba", "hbbba", "hcbba", "ibbba", "icbba", "jbbba", "jcbba", "kbbba", "kcbba", "lbbba", "lcbba", "mbbba", "mcbba", "nbbba", "ncbba", "obbba", "ocbba", "pbbba", "pcbba", "ccbba", "ccaba", "ccaca", "ccdba", "ccdca", "cceba", "cceca", "ccfba", "ccfca", "ccgba", "ccgca", "cchba", "cchca", "cciba", "ccica", "ccjba", "ccjca", "cckba", "cckca", "cclba", "cclca", "ccmba", "ccmca", "ccnba", "ccnca", "ccoba", "ccoca", "ccpba", "ccpca", "cccca", "accca", "adcca", "bccca", "bdcca", "eccca", "edcca", "fccca", "fdcca", "gccca", "gdcca", "hccca", "hdcca", "iccca", "idcca", "jccca", "jdcca", "kccca", "kdcca", "lccca", "ldcca", "mccca", "mdcca", "nccca", "ndcca", "occca", "odcca", "pccca", "pdcca", "ddcca", "ddaca", "ddada", "ddbca", "ddbda", "ddeca", "ddeda", "ddfca", "ddfda", "ddgca", "ddgda", "ddhca", "ddhda", "ddica", "ddida", "ddjca", "ddjda", "ddkca", "ddkda", "ddlca", "ddlda", "ddmca", "ddmda", "ddnca", "ddnda", "ddoca", "ddoda", "ddpca", "ddpda", "dddda", "addda", "aedda", "bddda", "bedda", "cddda", "cedda", "fddda", "fedda", "gddda", "gedda", "hddda", "hedda", "iddda", "iedda", "jddda", "jedda", "kddda", "kedda", "lddda", "ledda", "mddda", "medda", "nddda", "nedda", "oddda", "oedda", "pddda", "pedda", "eedda", "eeada", "eeaea", "eebda", "eebea", "eecda", "eecea", "eefda", "eefea", "eegda", "eegea", "eehda", "eehea", "eeida", "eeiea", "eejda", "eejea", "eekda", "eekea", "eelda", "eelea", "eemda", "eemea", "eenda", "eenea", "eeoda", "eeoea", "eepda", "eepea", "eeeea", "ggggg", "agggg", "ahggg", "bgggg", "bhggg", "cgggg", "chggg", "dgggg", "dhggg", "egggg", "ehggg", "fgggg", "fhggg", "igggg", "ihggg", "jgggg", "jhggg", "kgggg", "khggg", "lgggg", "lhggg", "mgggg", "mhggg", "ngggg", "nhggg", "ogggg", "ohggg", "pgggg", "phggg", "hhggg", "hhagg", "hhahg", "hhbgg", "hhbhg", "hhcgg", "hhchg", "hhdgg", "hhdhg", "hhegg", "hhehg", "hhfgg", "hhfhg", "hhigg", "hhihg", "hhjgg", "hhjhg", "hhkgg", "hhkhg", "hhlgg", "hhlhg", "hhmgg", "hhmhg", "hhngg", "hhnhg", "hhogg", "hhohg", "hhpgg", "hhphg", "hhhhg", "ahhhg", "aihhg", "bhhhg", "bihhg", "chhhg", "cihhg", "dhhhg", "dihhg", "ehhhg", "eihhg", "fhhhg", "fihhg", "ghhhg", "gihhg", "jhhhg", "jihhg", "khhhg", "kihhg", "lhhhg", "lihhg", "mhhhg", "mihhg", "nhhhg", "nihhg", "ohhhg", "oihhg", "phhhg", "pihhg", "iihhg", "iiahg", "iiaig", "iibhg", "iibig", "iichg", "iicig", "iidhg", "iidig", "iiehg", "iieig", "iifhg", "iifig", "iighg", "iigig", "iijhg", "iijig", "iikhg", "iikig", "iilhg", "iilig", "iimhg", "iimig", "iinhg", "iinig", "iiohg", "iioig", "iiphg", "iipig", "iiiig", "aiiig", "ajiig", "biiig", "bjiig", "ciiig", "cjiig", "diiig", "djiig", "eiiig", "ejiig", "fiiig", "fjiig", "giiig", "gjiig", "hiiig", "hjiig", "kiiig", "kjiig", "liiig", "ljiig", "miiig", "mjiig", "niiig", "njiig", "oiiig", "ojiig", "piiig", "pjiig", "jjiig", "jjaig", "jjajg", "jjbig", "jjbjg", "jjcig", "jjcjg", "jjdig", "jjdjg", "jjeig", "jjejg", "jjfig", "jjfjg", "jjgig", "jjgjg", "jjhig", "jjhjg", "jjkig", "jjkjg", "jjlig", "jjljg", "jjmig", "jjmjg", "jjnig", "jjnjg", "jjoig", "jjojg", "jjpig", "jjpjg", "jjjjg", "ajjjg", "akjjg", "bjjjg", "bkjjg", "cjjjg", "ckjjg", "djjjg", "dkjjg", "ejjjg", "ekjjg", "fjjjg", "fkjjg", "gjjjg", "gkjjg", "hjjjg", "hkjjg", "ijjjg", "ikjjg", "ljjjg", "lkjjg", "mjjjg", "mkjjg", "njjjg", "nkjjg", "ojjjg", "okjjg", "pjjjg", "pkjjg", "kkjjg", "kkajg", "kkakg", "kkbjg", "kkbkg", "kkcjg", "kkckg", "kkdjg", "kkdkg", "kkejg", "kkekg", "kkfjg", "kkfkg", "kkgjg", "kkgkg", "kkhjg", "kkhkg", "kkijg", "kkikg", "kkljg", "kklkg", "kkmjg", "kkmkg", "kknjg", "kknkg", "kkojg", "kkokg", "kkpjg", "kkpkg", "kkkkg", "ggggx", "gggxx", "ggxxx", "gxxxx", "xxxxx", "xxxxy", "xxxyy", "xxyyy", "xyyyy", "yyyyy", "yyyyw", "yyyww", "yywww", "ywwww", "wwwww", "wwvww", "wvvww", "vvvww", "vvvwz", "avvwz", "aavwz", "aaawz", "aaaaz"},
	))
}

// 1 构建这样的关系
//  0        1      2        3    levelNodes             
//          ted    tad  
// 											
//   red                    tax
//          rex    tex      
//          
// 2 从 tax 往 red 找, 保存所有结果
func findLadders(beginWord string, endWord string, wordList []string) [][]string {

	anwser := [][]string{}
	wordSet := map[string]bool{}
	for _, w := range wordList {
		wordSet[w] = true
	}
	if !wordSet[endWord] {
		return anwser
	}
	delete(wordSet, beginWord)
	levelNodes := [][]string{}
	queue := []string{beginWord}
	isNeighbor := func(word1, word2 string) bool {
		diff := 0
		for i := range word1 {
			if word1[i] != word2[i] {
				diff++
			}
		}
		return diff == 1
	}
	fond := false
	for len(queue) > 0 && !fond {
		qSize := len(queue)
		nodes := make([]string, len(queue))
		copy(nodes, queue)
		levelNodes = append(levelNodes, nodes)

		for i := 0; i < qSize; i++ {
			curWord := queue[0]
			queue = queue[1:]

			for word := range wordSet {
				if !isNeighbor(curWord, word) {
					continue
				}
				if word == endWord {
					fond = true
					break
				}
				delete(wordSet, word)
				queue = append(queue, word)
			}
		}

	}
	anwser = [][]string{{endWord}}
	for i := len(levelNodes) - 1; i >= 0; i-- {
		size := len(anwser)
		for _, word := range levelNodes[i] {
			for j := 0; j < size; j++ {
				lastLevelWords := anwser[j]
				lastWord := lastLevelWords[0]
				if isNeighbor(word, lastWord) {
					item := []string{word}
					item = append(item, lastLevelWords...)
					anwser = append(anwser, item)
				}

			}
		}
		anwser = anwser[size:]

	}
	return anwser
}

func findLadders1(beginWord string, endWord string, wordList []string) [][]string {

	wordSet := map[string]bool{}
	for _, w := range wordList {
		wordSet[w] = true
	}
	visited := map[string]bool{}
	visited[beginWord] = true
	queue := [][]string{{beginWord}}
	anwser := [][]string{}

	fond := false
	for len(queue) > 0 {
		qSize := len(queue)
		for ; qSize > 0; qSize-- {
			cur := queue[0]
			queue = queue[1:]
			lastWord := cur[len(cur)-1]
			neighbors := getNeighbors(lastWord, wordSet)
			for _, word := range neighbors {
				paths := make([]string, len(cur))
				copy(paths, cur)
				paths = append(paths, word)
				visited[word] = true
				if word == endWord {
					anwser = append(anwser, paths)
					fond = true
				} else {
					queue = append(queue, paths)
				}
			}
		}
		for key := range visited {
			delete(wordSet, key)
		}
		if fond {
			return anwser
		}
	}
	return anwser
}

func getNeighbors(word string, wordSet map[string]bool) []string {
	result := []string{}
	for i := range word {
		byteWord := []byte(word)
		for c := byte('a'); c <= 'z'; c++ {
			byteWord[i] = c
			newWord := string(byteWord)
			if wordSet[newWord] {
				result = append(result, newWord)
			}
		}
	}
	return result
}
