package mail

import (
	"fmt"
	"io"
	"net/mail"
	"os"
	"strings"

	// "strconv"

	"github.com/emersion/go-smtp"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

// MailServer implements the SMTP server interface for receiving
// emails and storing them in PocketBase. It implements the
// `smtp.Backend` interface.
type MailServer struct {
	App *pocketbase.PocketBase
}

// Session is used by servers to respond to an SMTP client.
func (bkd *MailServer) NewSession(c *smtp.Conn) (smtp.Session, error) {
	return &MailSession{App: bkd.App}, nil
}

// Serve starts the SMTP server and listens for incoming email connections.
func (bkd *MailServer) Serve() error {
	s := smtp.NewServer(bkd)

	s.Addr = os.Getenv("SMTP_PORT")
	if s.Addr == "" {
		s.Addr = ":1025"
	}

	// s.Addr = bkd.App.Settings().SMTP.Host + ":" + strconv.Itoa(bkd.App.Settings().SMTP.Port)
	// s.Domain = bkd.App.Settings().SMTP.Host
	s.AllowInsecureAuth = true

	fmt.Println("Starting SMTP server on " + s.Addr)
	return s.ListenAndServe()
}

// MailSession represents an SMTP session for handling incoming email
// transactions. It implements the `smtp.Session` interface.
type MailSession struct {
	App       *pocketbase.PocketBase
	sender    string
	recipient string
	data      []byte
}

// Discard currently processed message.
func (bkd *MailSession) Reset() {
	fmt.Println("Mail `RESET`")
}

// Free all resources associated with session.
func (bkd *MailSession) Logout() error {
	fmt.Println("Mail `LOGOUT`")
	return nil
}

// Set return path for currently processed message.
func (bkd *MailSession) Mail(from string, opts *smtp.MailOptions) error {
	// fmt.Printf("from: `%s`\n", from)
	// fmt.Printf("auth: %v\n", opts.Auth)
	// fmt.Printf("body type: %v\n", opts.Body)
	// fmt.Printf("envelope id: %s\n", opts.EnvelopeID)
	// fmt.Printf("require tls: %v\n", opts.RequireTLS)

	fmt.Printf("Mail `MAIL` From: %s\n", from)

	return nil
}

// Add recipient for currently processed message.
func (bkd *MailSession) Rcpt(to string, opts *smtp.RcptOptions) error {
	// fmt.Printf("to: `%s`\n", to)
	// fmt.Printf("DeliverBy: %v\n", opts.DeliverBy)
	// fmt.Printf("MTPriority: %v\n", opts.MTPriority)
	// fmt.Printf("Notify: %v\n", opts.Notify)
	// fmt.Printf("OriginalRecipient: %v\n", opts.OriginalRecipient)
	// fmt.Printf("OriginalRecipientType: %v\n", opts.OriginalRecipientType)
	// fmt.Printf("RequireRecipientValidSince: %v\n", opts.RequireRecipientValidSince)

	fmt.Printf("Mail `RCPT` To: %s\n", to)

	return nil
}

// Set currently processed message contents and send it.
//
// r must be consumed before Data returns.
func (bkd *MailSession) Data(r io.Reader) error {
	msg, err := mail.ReadMessage(r)
	if err != nil {
		return err
	}

	// create headers map
	headers := make(map[string]string)
	for k, v := range msg.Header {
		headers[k] = v[0]
	}

	// read body (without headers)
	body, err := io.ReadAll(msg.Body)
	if err != nil {
		return err
	}

	fmt.Printf("Received message: Message-Id: %s, From: %s, To: %s\n", headers["Message-Id"], headers["From"], headers["To"])

	// write to collection
	collection, err := bkd.App.FindCollectionByNameOrId("mails")
	if err != nil {
		return err
	}

	record := core.NewRecord(collection)

	// set general record fields
	record.Set("messageId", headers["Message-Id"])
	record.Set("sender", headers["From"])
	record.Set("recipient", headers["To"])
	record.Set("subject", headers["Subject"])
	record.Set("headers", headers)
	record.Set("body", string(body))
	record.Set("incoming", true)

	// find the recipient name (before the @ symbol)
	recipientName := strings.Split(headers["To"], "@")[0]

	// scan users collection for matching recipient email
	// if no error, set the recipient field to the user record id
	if userRecord, err := bkd.App.FindFirstRecordByFilter("users", "mailName = {:recipient}", dbx.Params{"recipient": recipientName}); err == nil {
		record.Set("owner", userRecord.Id)
	}

	fmt.Printf("Redirecting message: Message-Id: %s, Assigned Owner: %s\n", headers["Message-Id"], record.GetString("owner"))

	if err := bkd.App.Save(record); err != nil {
		return err
	}

	return nil
}
