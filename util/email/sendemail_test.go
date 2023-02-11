package email

import (
	"fmt"
	"testing"
)

func TestSendValidCode(t *testing.T) {
	ms := NewEmailService()
	if err, _ := ms.SendValidCode("1652091948@qq.com"); err != nil {
		fmt.Println(err.Error())
		t.Fail()
	}

}
