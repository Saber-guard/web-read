package service

import (
	"fmt"
	"github.com/faiface/beep"
	c "github.com/smartystreets/goconvey/convey"
	"sync"
	"testing"
)

func TestVoiceService_TextToVoice(t *testing.T) {
	vs := VoiceService{}
	// text := "引用未定义环境变量会被替换为空字符串"
	text := "新近落成的中国共产党历史展览馆巍然矗立，气势恢宏。下午3时20分许，习近平等党和国家领导同志来到这里，步入展厅参观展览。展览以“不忘初心、牢记使命”为主题，精心设计了“建立中国共产党 夺取新民主主义革命伟大胜利”"
	wg := &sync.WaitGroup{}
	wg.Add(1)
	voiceArr := make([]beep.Streamer, 2)
	formatArr := make([]beep.Format, 2)
	vs.TextToVoice(1, text, "test_file_name01", voiceArr, formatArr, 4, wg)

	fmt.Printf("voiceArr[1]--bytearr=%v\n", voiceArr[1])
}

func TestVoiceService_urlToVoice(t *testing.T) {
	c.Convey("test1", t, func() {
		vs := VoiceService{}
		// url:="http://mp.weixin.qq.com"
		url := "https://mp.weixin.qq.com/s/MRoSibzspGuoaY-FDCTBcw"
		fileName, err := vs.urlToVoice(url)

		c.So(err, c.ShouldEqual, nil)
		c.So(fileName, c.ShouldNotEqual, "")
	})
}
