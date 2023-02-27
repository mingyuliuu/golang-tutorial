package main

import (
	"bufio"
	"email-verifier/utils"
	"fmt"
	"log"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Domain, HasMX, HasSPF, (SPR Record), HasDMARC, (DMARC Record)")

	for scanner.Scan() {
		utils.CheckDomain(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error: Could not read from input %v\n", err)
	}
}
