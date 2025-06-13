// File: internal/mailer/mailer.go
package mailer

import (
    "bytes"
    "crypto/tls"
    "fmt"
    "net/smtp"
    "text/template"

    "github.com/bocaletto-luca/server-gen/internal/config"
    "github.com/bocaletto-luca/server-gen/internal/sysinfo"
    "gopkg.in/mail.v2"
)

const tpl = `
Time: {{.Timestamp}}
IPs: {{.IPs}}
Host: {{.Hostname}}
OS: {{.OS}}
CPU: {{.CPUPercent}}
Memory: {{.MemUsed}}/{{.MemTotal}}
Users: {{.Users}}
`

// Send emails with TLS, retry on failure.
func Send(smtpCfg config.SMTPConfig, data *sysinfo.Data) error {
    m := mail.NewMessage()
    m.SetHeader("From", smtpCfg.From)
    m.SetHeader("To", smtpCfg.To...)
    m.SetHeader("Subject", "[server-gen] System Report")
    body := new(bytes.Buffer)
    if err := template.Must(template.New("r").Parse(tpl)).Execute(body, data); err != nil {
        return fmt.Errorf("template exec: %w", err)
    }
    m.SetBody("text/plain", body.String())

    d := mail.NewDialer(smtpCfg.Host, smtpCfg.Port,
        smtpCfg.Username, smtpCfg.Password)
    d.TLSConfig = &tls.Config{InsecureSkipVerify: false}
    d.StartTLSPolicy = mail.MandatoryStartTLS

    // retry once
    if err := d.DialAndSend(m); err != nil {
        fmt.Println("first send failed, retrying:", err)
        time.Sleep(time.Second * 5)
        return d.DialAndSend(m)
    }
    return nil
}
