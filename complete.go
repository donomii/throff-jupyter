package main

import (
	"strings"

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
func handleCompleteRequest(ir *throfflib.Engine, receipt msgReceipt) error {
	// Extract the data from the request.
	reqcontent := receipt.Msg.Content.(map[string]interface{})
	code := reqcontent["code"].(string)
	cursorPos := int(reqcontent["cursor_pos"].(float64))
	startpos := cursorPos
	for startpos = cursorPos - 1; startpos > 0 && code[startpos] != byte(32); startpos -= 1 {
	}
	startpos += 1
	query := code[startpos:cursorPos]

	e := throfflib.MakeEngine()
	e = e.RunString(throfflib.BootStrapString(), "Bootstrap code")
	ne := e.RunString(" KEYS ENVIRONMENTOF [ here ] ", "Get Functions")

	t := ne.DataStackTop()
	val := t.GetString()

	words := strings.Split(val, " ")

	var matches []string
	for _, v := range words {
		if strings.HasPrefix(strings.ToLower(v), strings.ToLower(query)) {
			matches = append(matches, v)
		}
	}

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
