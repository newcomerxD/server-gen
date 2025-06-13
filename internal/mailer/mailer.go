// File: internal/mailer/mailer.go
package mailer

import (
    "bytes"
    "fmt"
    "net/smtp"
    "text/template"

    "github.com/bocaletto-luca/server-gen/internal/config"
    "github.com/bocaletto-luca/server-gen/internal/sysinfo"
)

const reportTemplate = `
System Report:
IPs: {{.IPs}}
Hostname: {{.Hostname}}
OS: {{.OS}}
CPU Usage: {{.CPUPercent}}
Memory Used: {{.MemUsed}} / {{.MemTotal}}
Users: {{.Users}}
`

// Send formats the system data and sends it via SMTP.
func Send(smtpCfg config.SMTPConfig, data *sysinfo.Data) error {
    auth := smtp.PlainAuth("", smtpCfg.Username, smtpCfg.Password, smtpCfg.Host)
    var body bytes.Buffer
    tpl := template.Must(template.New("report").Parse(reportTemplate))
    if err := tpl.Execute(&body, data); err != nil {
        return fmt.Errorf("execute template: %w", err)
    }

    addr := fmt.Sprintf("%s:%d", smtpCfg.Host, smtpCfg.Port)
    msg := []byte(
        fmt.Sprintf("From: %s\r\n", smtpCfg.From) +
            fmt.Sprintf("To: %s\r\n", smtpCfg.To) +
            "Subject: [server-gen] System Report\r\n\r\n" +
            body.String(),
    )

    return smtp.SendMail(addr, auth, smtpCfg.From, smtpCfg.To, msg)
}
