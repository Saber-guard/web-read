package util

import (
	"fmt"
	c "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestStringUtil_SplitByLen(t *testing.T) {
	c.Convey("test1", t, func() {
		su := StringUtil{}
		input := "1234567890a"
		res := su.SplitByLen(input, 5)
		fmt.Println("res=",res)
		c.So(len(res), c.ShouldEqual, 3)
	})

	c.Convey("test2", t, func() {
		su := StringUtil{}
		input := "1234567890"
		res := su.SplitByLen(input, 5)
		fmt.Println("res=",res)
		c.So(len(res), c.ShouldEqual, 2)

	})

	c.Convey("test3", t, func() {
		su := StringUtil{}
		input :="新近落成的中国共产党历史展览馆巍然矗立，气势恢宏。下午3时20分许，习近平等党和国家领导同志来到这里，步入展厅参观展览。展览以“不忘初心、牢记使命”为主题，精心设计了“建立中国共产党 夺取新民主主义革命伟大胜利”"

		res := su.SplitByLen(input, 105)
		fmt.Printf("resLen=%v,res=%v\n",len(res),res)
		c.So(len(res), c.ShouldEqual, 2)
	})
}
