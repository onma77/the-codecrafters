// CodeCrafters — Operation Gopher Protocol
// Module: File Pipeline
// Author: [peace oigoga]
// Squad:  [Peace Oigoga, Faith Ejembi, Nzekwe Chinaza, Treasure Gabriel, Owulo Celebrate, Ortse Alphonsus, Chekwube Okafor, Emaikwu Godwin, Endurance Ochefije, Chubiyojo Akoh]

// ═══════════════════════════════════════════
// SQUAD PIPELINE CONTRACT
// Squad: [Squad Name]
// ───────────────────────────────────────────
// Input line types:
//   [line 1: ALL CAP]
//   [line 2: lower]
//   [line 3: Trimspace]
//   [line 4: TODO ]
//
// Transformation rules (in order):
//   1. [Trim all leading and trailing whitespace ]
//   2. [Replace TODO: with ✦ ACTION]
//   3. [Convert ALL CAPS lines to Title Case ]
//   4. [Convert all lowercase lines to uppercase ]
//   5. [ Remove lines that are only dashes or blanks 		]
//
// Output format:
//   [Header: yes/no — SENTINEL FIELD REPORT — PROCESSED]
//   [Line numbering format- 1]
//
// Terminal summary fields:
//    ✦ Lines read    : [number]
//   ✦ Lines written : [number]
//  ✦ Lines removed : [number]
//   ✦ Rules applied :
//   1. [Trim all leading and trailing whitespace ]
//   2. [Replace TODO: with ✦ ACTION]
//   3. [Convert ALL CAPS lines to Title Case ]
//   4. [Convert all lowercase lines to uppercase ]
//   5. [Remove lines that are only dashes or blanks 	]
// ═══════════════════════════════════════════

package main

import (
	"fmt"
	"strings"
)

func CAP(e string) string {
	return strings.ToUpper(e)
}

func lower(t string) string {
	return strings.ToLower(t)
}

func TODO(h string) string {
	words := strings.Fields(h)
	for p := 0; p < len(words); p++ {
		if words[p] == "TODO" {
			words[p] = "ACTION"
		}
	}
	return strings.Join(words, " ")
}

func dashBlank(m string) string {
	m = strings.ReplaceAll(m, " ", "")
	m = strings.ReplaceAll(, "_", "")

	return m
	// if d == " " {
	// 	return true
	// }
	//if
}

// func trimSpace(s string) string {

// }

func main() {
	fmt.Println(TODO("sabhvdwfde TODO agfyguetf,avyjgfyuy"))
	fmt.Println(CAP("ngfyugffljhejhfegfyegfyeogfrfeljhvryg"))
	fmt.Println(lower("LNDJKFDBFGDLGDNKJDJDBH"))
	fmt.Println(dashBlank("vdfsff          egdkyeftyrire          uyskgfyuytyf"))
}
