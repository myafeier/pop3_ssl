package pop3

import (
	"io/ioutil"
	"testing"
)

func TestNewClient(t *testing.T) {
	c, err := NewClient("pop3.mxhichina.com:995", "", "")
	if err != nil {
		t.Error(err)
	}

	result, err := c.List()
	if err != nil {
		t.Error(err)
	}
	t.Logf("num:%d\n", len(result))

	for _, v := range result {

		message, err := c.GetMail(v.Id)
		if err != nil {
			t.Error(err)
		}

		t.Log("Uid:", string(message.UID))
		t.Log(string(message.Message.Header.Get("From")))
		t.Log(string(message.Message.Header.Get("Subject")))
		t.Log(string(message.Message.Header.Get("Date")))
		t.Log(string(message.Message.Header.Get("Content-Type")))
		t.Log(string(message.Message.Header.Get("Received")))
		t.Log(string(message.Message.Header.Get("Message-ID")))

		m, _ := ioutil.ReadAll(message.Message.Body)

		t.Log(string(m))
		//converter,err:=iconv.Open("gb2312","utf-8")
	}

}
