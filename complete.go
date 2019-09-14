package main

import (
	"log"
	"strings"

	interp "github.com/cosmos72/gomacro/fast"
	"github.com/donomii/throfflib"
)

type Completion struct {
	class,
	name,
	typ string
}

type CompletionResponse struct {
	partial     int
	completions []Completion
}

/************************************************************
* entry function
************************************************************/
func handleCompleteRequest(ir *interp.Interp, receipt msgReceipt) error {
	log.Println(ir, receipt)
	// Extract the data from the request.
	reqcontent := receipt.Msg.Content.(map[string]interface{})
	code := reqcontent["code"].(string)
	cursorPos := int(reqcontent["cursor_pos"].(float64))
	startpos := cursorPos
	for startpos = cursorPos - 1; startpos > 0 && code[startpos] != byte(32); startpos -= 1 {
	}
	startpos += 1
	query := code[startpos:cursorPos]
	log.Println("Req", query)

	//wordsStr := `ACTIVATE/CC ADDWORD COUNTDOWN IDEBUGON FILTER WHEN .L NUM2CHAR WITH ->STRING ACTORREC MULTIBUBBLESTEP ITROFF MERGESORT PWD STACKMAP ALIAS GETARRAY POPARRAY REPEAT NEWQUEUE GETHASH DO SHELLHELP MACRO ONTO HASHITERATE SWAP RANGE STRING GETLINE ADD1 SWATCH KEYS/VALS ZERO LENGTH COMPAREBYTES FOLD OVER SWAPARR NEWHASH GETBYTE READSOCKETLINE ITERATESLICE2 MULT SW TCPSERVER FOR IOTA NEVAL HASHSET OPEN-SQUARE-BRACE-CHAR OF DIRECTORY-LIST TRON NULLSTEP PRINTLN CASE H[RECURSE SUBSHELL REAL-GETHASH ACTOR NAMETESTBLOCK PRINT2 CALL/SELF , REVERSE ROLL ITERATE2 OPENSQLITE NEXTROW LESSTHAN CAR CMDSTDOUTSTDERR IDEBUGOFF TOK ( THIN DEBUGON -( SETHASH MERGE EMPTY? CALL/CC INVERTHASH BYTE2STR CLOSEFILE WITA IT DUMP SPLIT FI CDR STRING-JOIN TESTS test ; FLOOR ROT DUP ->CODE LOCATIONOF PRINTSOCKET NOT TESTBLOCK CALLWITHSELF ARRAYPUSH ADD MULTIBUBBLESORT FORCEGC 2DUP FILEOF SLICE CWD UNSHIFTARRAY SLURPFILE TRUE MODULUS OR ... MAPFILE TRUTHY ERRORHANDLER SETTYPE PICK ELSE DEFINE STARTPROCESS SCRUBLEX ]A .E ITERATESLICE PRINT3 ARRAY CALLA READFILELINE TROFF MAPFILERECURSE RUNTESTS BUBBLESTEP CALL CATCH OPEN-FUNCTION-CHAR ITERATE KILLPROC MAP2 LS SLICE-STRING SPACE OPENSOCKET STRING-CONCATENATE* SHELL MIME UNFUNC tests: CRLF ->FUNC BUBBLESORT OPENFILE RECURSE COUNTDOWN2ZERO2 d OS WATCH FOREVER PL DIVIDE SETLEX PRINTX ->LAMBDA NEWLINE COMMENT pn CLEARSTACK SHIFTARRAY A[ GET IF DUMPBYTES SUB1 ITRON GETTYPE MAP MODULUSI PROMPT_HOOK ->ARRAY DEFAULT CLOSE-SQUARE-BRACE-CHAR DNS.CNAME THREAD CLEAROUTPUT returns STATEMENT DEBUGOFF ->BYTES DNS.TXT EXEC SUB PUSH WRITEQ [ = READQ _ SETARRAY FUNCTION? INSTALLDYNA TEST REBIND ARG p LAMBDA FORRECURSE VALS STRING! ] .S THROW BIND WAITPROC LEXLOOKUP DESCRIPTION SAFETYON )- 2DROP EMIT CLOSE-FUNCTION-CHAR EXIT AND CURRENT-DIRECTORY H[ IFFY RELEASEPROC OPEN-CURLY CLOSE-CURLY ENVIRONMENTOF TIMESEC QUICKSTART HELP PUSHARRAY QUERY TUCK GETLEX ]H EQUAL SIN .C SETENV DROP HASHGET LN -f => GETFUNCTION EXIT/CODE A[RECURSE ) CODE DNS.REVERSE : -s PROMISE CR SEARCHBYTES NEWARRAY CHAT KEYS DNS.HOST STRING-CONCATENATE NIP GETSTRING expects: MMAPFILE KEYVALS RUNSTRING SLEEP APPEND FALSE`
	e := throfflib.MakeEngine()
	e = e.RunString(throfflib.BootStrapString(), "Bootstrap code")
	ne := e.RunString(" KEYS ENVIRONMENTOF [ here ] ", "Get Functions")

	t := ne.DataStackTop()
	val := t.GetString()
	log.Println("val ", val)

	words := strings.Split(val, " ")

	var matches []string
	for _, v := range words {
		log.Println("Comparing", strings.ToLower(v), strings.ToLower(query))
		if strings.HasPrefix(strings.ToLower(v), strings.ToLower(query)) {
			log.Println("Adding", v)
			matches = append(matches, v)
		}
	}

	log.Println("Matches", matches)

	// prepare the reply
	content := make(map[string]interface{})

	if len(matches) == 0 {
		content["ename"] = "ERROR"
		content["evalue"] = "no completions found"
		content["traceback"] = nil
		content["status"] = "error"
	} else {

		content["cursor_start"] = float64(startpos)
		content["cursor_end"] = float64(cursorPos)
		content["matches"] = matches
		content["status"] = "ok"
	}

	return receipt.Reply("complete_reply", content)
}
