package main

import (
	"fmt"
	"github.com/mdobydullah/go-spinner"
	"net"
	"net/mail"
	"net/smtp"
	"strings"
	"time"
)

func validateEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func getDomain(email string) string {
	parts := strings.Split(email, "@")
	domain := parts[1]

	return domain
}

func verifyDomain(spin *spinner.Spinner, email string) {
	spin = spinner.NewSpinner("Verify domain: Processing...")
	spin.Start()
	spin.SetSpeed(100 * time.Millisecond)
	spin.SetCharset([]string{"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷"})

	domain := getDomain(email)

	// Check if there are any A records
	_, err := net.LookupIP(domain)
	if err != nil {
		spin.Stop()
		fmt.Printf("x Verify domain: %s", err)
		return
	}

	spin.Stop()
	fmt.Println("✓ Verify domain: Completed")
}

func verifyEmail(spin *spinner.Spinner, email string) {
	spin = spinner.NewSpinner("Verify: Processing...")
	spin.Start()
	spin.SetSpeed(100 * time.Millisecond)
	spin.SetCharset([]string{"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷"})

	// Extract the domain from the email address
	parts := strings.Split(email, "@")
	domain := parts[1]

	// Check if there are any MX records
	_, err := net.LookupIP(domain)
	if err != nil {
		spin.Stop()
		fmt.Printf("x Verify: %s", err)
		return
	}

	//var firstMXRecord string
	//if len(mxrecords) > 0 {
	//	// Print the first MX record
	//	firstMXRecord = mxrecords[0].Host
	//} else {
	//	spin.Stop()
	//	fmt.Printf("No MX records found for: %s", domain)
	//	return
	//}
	//
	//sendHeloCommand(firstMXRecord)

	spin.Stop()
	fmt.Println("✓ Verify: Completed")

	//for _, mx := range mxrecords {
	//	fmt.Println(mx.Host, mx.Pref)
	//}
}

func sendHeloCommand(mxServer string) {
	// Connect to the MX server on port 25 (SMTP)
	conn, err := net.Dial("tcp", mxServer+":25")
	if err != nil {
		fmt.Println("Error connecting to the server:", err)
		return
	}
	defer conn.Close()

	// Create an SMTP client
	client, err := smtp.NewClient(conn, mxServer)
	if err != nil {
		fmt.Println("Error creating SMTP client:", err)
		return
	}
	defer client.Quit()

	// Send the HELO command
	err = client.Hello("localhost")
	if err != nil {
		fmt.Println("Error sending HELO command:", err)
		return
	}

	//fmt.Println("HELO command sent successfully.")
}
