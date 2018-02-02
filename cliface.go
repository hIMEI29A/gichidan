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

func getTotalStats(bloodyRoots map[string]string, total int) string {
	stats := makeMessage(FOUND) + makeValMessage(string(total)) + "\n"

	for s, i := range bloodyRoots {
		stats += makeMessage(BYREQ) + makeValMessage(s) + ": " + makeValMessage(i)

	}

	return stats
}

func pressAny() {
	var input string
	fmt.Println(PRESS)
	fmt.Scanln(&input)
}

func banner() {
	fmt.Println(GRN, "           ███         ", RED, "█████     ", GRN, "███    ", RED, "█████          ")
	fmt.Println(WHT, "  v1.0.0", RED, "░░░           ░░███       ░░░      ░░███          ", WHT, "© hIMEI")
	fmt.Println(RED, "  ███████ ████   ██████  ░███████   ████   ███████   ██████   ████████  ")
	fmt.Println(RED, " ███░░███░░███  ███░░███ ░███░░███ ░░███  ███░░███  ░░░░░███ ░░███░░███ ")
	fmt.Println(RED, "░███ ░███ ░███ ░███ ░░░  ░███ ░███  ░███ ░███ ░███   ███████  ░███ ░███ ")
	fmt.Println(RED, "░███ ░███ ░███ ░███  ███ ░███ ░███  ░███ ░███ ░███  ███░░███  ░███ ░███ ")
	fmt.Println(RED, "░░███████ █████░░██████  ████ █████ █████░░████████░░████████ ████ █████")
	fmt.Println(RED, " ░░░░░███░░░░░  ░░░░░░  ░░░░ ░░░░░ ░░░░░  ░░░░░░░░  ░░░░░░░░ ░░░░ ░░░░░ ")
	fmt.Println(RED, " ███ ░███        ", GRN, "___onion secrets for console cowboys___")
	fmt.Println(RED, "░░██████")
	fmt.Println(RED, "░░░░░░", RESET)

}
