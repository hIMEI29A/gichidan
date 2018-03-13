// Copyright 2017 hIMEI

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/*
 * File cliface.go contains functions for iteracting with user's terminal.
 * Errors messages is also created here.
 */

package main

import (
	"fmt"
	"time"
)

func makeErrString(errConst string) string {
	errString := BOLD + RED + errConst + RESET
	return errString
}

func makeMessage(messageConst string) string {
	message := BOLD + YEL + messageConst + RESET
	return message
}

func makeValMessage(value string) string {
	message := BOLD + CYN + value + RESET
	return message
}

func makeUrlMessage(url string) string {
	message := BOLD + GRN + url + RESET
	return message
}

func getTotalStats(bloodyRoots map[string]string, finalHosts []*Host, total int) string {
	stats := makeMessage(FOUND) + makeValMessage(iToa(total)) + "\n"

	for i, s := range bloodyRoots {
		stats += makeMessage(BYREQ) + makeValMessage(i) + ": " + makeValMessage(s) + "\n"
	}

	stats += makeValMessage(iToa(len(finalHosts))) + makeMessage(WILL)

	return stats
}

func pressAny() {
	var input string
	fmt.Println(makeMessage(PRESS))
	fmt.Scanln(&input)
}

func banner() {
	fmt.Println(GRN, "           ███         ", CYN, "█████     ", GRN, "███    ", CYN, "█████          ")
	time.Sleep(100 * time.Millisecond)
	fmt.Println(WHT, "  v1.1.1", CYN, "░░░           ░░███       ░░░      ░░███          ", WHT, "© hIMEI")
	time.Sleep(100 * time.Millisecond)
	fmt.Println(CYN, "  ███████ ████   ██████  ░███████   ████   ███████   ██████   ████████  ")
	time.Sleep(100 * time.Millisecond)
	fmt.Println(CYN, " ███░░███░░███  ███░░███ ░███░░███ ░░███  ███░░███  ░░░░░███ ░░███░░███ ")
	time.Sleep(100 * time.Millisecond)
	fmt.Println(CYN, "░███ ░███ ░███ ░███ ░░░  ░███ ░███  ░███ ░███ ░███   ███████  ░███ ░███ ")
	time.Sleep(100 * time.Millisecond)
	fmt.Println(CYN, "░███ ░███ ░███ ░███  ███ ░███ ░███  ░███ ░███ ░███  ███░░███  ░███ ░███ ")
	time.Sleep(100 * time.Millisecond)
	fmt.Println(CYN, "░░███████ █████░░██████  ████ █████ █████░░████████░░████████ ████ █████")
	time.Sleep(100 * time.Millisecond)
	fmt.Println(CYN, " ░░░░░███░░░░░  ░░░░░░  ░░░░ ░░░░░ ░░░░░  ░░░░░░░░  ░░░░░░░░ ░░░░ ░░░░░ ")
	time.Sleep(100 * time.Millisecond)
	fmt.Println(CYN, " ███ ░███        ", GRN, "___onion secrets for console cowboys___")
	time.Sleep(100 * time.Millisecond)
	fmt.Println(CYN, "░░██████")
	time.Sleep(100 * time.Millisecond)
	fmt.Println(CYN, "░░░░░░", RESET)
}
