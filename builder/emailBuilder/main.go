package main

import "emailbuilder/email"

func main() {
	eb := email.NewEmailBuilder("abc.ef@somemail.com", "Weekly targets").Bcc("ghi.jk@somemail.com").Body("Hi XYZ, This is to inform you that the target for this week has beenm achieved. Thanks and regards, ABC").Priority(email.LOW)
	eb.Build().ToString()
}
