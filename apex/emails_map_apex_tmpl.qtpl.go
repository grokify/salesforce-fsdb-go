// This file is automatically generated by qtc from "emails_map_apex_tmpl.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line emails_map_apex_tmpl.qtpl:1
package apex

//line emails_map_apex_tmpl.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line emails_map_apex_tmpl.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line emails_map_apex_tmpl.qtpl:1
func StreamApexEmailsMapTemplate(qw422016 *qt422016.Writer, data map[string]map[string]string, emailSubjectTmpl, emailBodyTmpl, replyToEmail, senderDisplayName string) {
	//line emails_map_apex_tmpl.qtpl:1
	qw422016.N().S(`
// BEGIN auto-generated Apex code (https://github.com/grokify/go-salesforce/tree/master/apex)
String emailSubjectTmpl = '`)
	//line emails_map_apex_tmpl.qtpl:3
	qw422016.N().S(EscapeSingleQuote(emailSubjectTmpl))
	//line emails_map_apex_tmpl.qtpl:3
	qw422016.N().S(`';
String emailBodyTmpl = '`)
	//line emails_map_apex_tmpl.qtpl:4
	qw422016.N().S(EscapeSingleQuote(emailBodyTmpl))
	//line emails_map_apex_tmpl.qtpl:4
	qw422016.N().S(`';
String replyToEmail = '`)
	//line emails_map_apex_tmpl.qtpl:5
	qw422016.N().S(EscapeSingleQuote(replyToEmail))
	//line emails_map_apex_tmpl.qtpl:5
	qw422016.N().S(`';
String senderDisplayName = '`)
	//line emails_map_apex_tmpl.qtpl:6
	qw422016.N().S(EscapeSingleQuote(senderDisplayName))
	//line emails_map_apex_tmpl.qtpl:6
	qw422016.N().S(`';
Map<String,Map<String,String>> emailsData = `)
	//line emails_map_apex_tmpl.qtpl:7
	qw422016.N().S(MapStringMapStringStringToApex(data, true))
	//line emails_map_apex_tmpl.qtpl:7
	qw422016.N().S(`;

List<Messaging.SingleEmailMessage> emails = new List<Messaging.SingleEmailMessage>();

for (String key : emailsData.keySet()) {
  Map<String,String> emailData = emailsData.get(key);

  Messaging.SingleEmailMessage email = new Messaging.SingleEmailMessage();
  Boolean hasRecipients = false;

  String sendTo = emailData.get('to_');
  if (string.isNotBlank(sendTo)) {
    email.setToAddresses(sendTo.split(';'));
    hasRecipients = true;
  }
  String sendCc = emailData.get('cc_');
  if (string.isNotBlank(sendCc)) {
    email.setCcAddresses(sendCc.split(';'));
    hasRecipients = true;
  }
  String sendBcc = emailData.get('bcc_');
  if (string.isNotBlank(sendBcc)) {
    email.setBccAddresses(sendBcc.split(';'));
    hasRecipients = true;
  }

  if (hasRecipients) {
    if (string.isNotBlank(replyToEmail)) {
      email.setReplyTo(replyToEmail);
    }
    if (string.isNotBlank(senderDisplayName)) {
      email.setSenderDisplayName(senderDisplayName);    
    }

    String emailSubject = emailSubjectTmpl;
    String emailBody = emailBodyTmpl;

    for (String emailTmplKey : emailData.keySet()) {
      Integer lastChar = emailTmplKey.charAt(emailTmplKey.length()-1);

      if (lastChar != 95) {
        String emailTmplVal = emailData.get(emailTmplKey);
        emailSubject = emailSubject.replace('{{'+emailTmplKey+'}}', emailTmplVal);
        emailBody = emailBody.replace('{{'+emailTmplKey+'}}', emailTmplVal);
      }
    }
    email.setSubject(emailSubject);
    email.setHtmlBody(emailBody);
    emails.add(email);
  }
}

if (emails.size()>0) {
  Messaging.sendEmail(emails);
}
// END auto-generated Apex code (https://github.com/grokify/go-salesforce/tree/master/apex)
`)
//line emails_map_apex_tmpl.qtpl:63
}

//line emails_map_apex_tmpl.qtpl:63
func WriteApexEmailsMapTemplate(qq422016 qtio422016.Writer, data map[string]map[string]string, emailSubjectTmpl, emailBodyTmpl, replyToEmail, senderDisplayName string) {
	//line emails_map_apex_tmpl.qtpl:63
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line emails_map_apex_tmpl.qtpl:63
	StreamApexEmailsMapTemplate(qw422016, data, emailSubjectTmpl, emailBodyTmpl, replyToEmail, senderDisplayName)
	//line emails_map_apex_tmpl.qtpl:63
	qt422016.ReleaseWriter(qw422016)
//line emails_map_apex_tmpl.qtpl:63
}

//line emails_map_apex_tmpl.qtpl:63
func ApexEmailsMapTemplate(data map[string]map[string]string, emailSubjectTmpl, emailBodyTmpl, replyToEmail, senderDisplayName string) string {
	//line emails_map_apex_tmpl.qtpl:63
	qb422016 := qt422016.AcquireByteBuffer()
	//line emails_map_apex_tmpl.qtpl:63
	WriteApexEmailsMapTemplate(qb422016, data, emailSubjectTmpl, emailBodyTmpl, replyToEmail, senderDisplayName)
	//line emails_map_apex_tmpl.qtpl:63
	qs422016 := string(qb422016.B)
	//line emails_map_apex_tmpl.qtpl:63
	qt422016.ReleaseByteBuffer(qb422016)
	//line emails_map_apex_tmpl.qtpl:63
	return qs422016
//line emails_map_apex_tmpl.qtpl:63
}
