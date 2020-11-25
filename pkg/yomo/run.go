package yomo

import (
	"fmt"
	"time"

	"github.com/yomorun/yomo/pkg/pprof"

	"github.com/yomorun/yomo/pkg/plugin"
	"github.com/yomorun/yomo/pkg/util"

	"github.com/yomorun/yomo/internal/framework"
)

var logger = util.GetLogger("yomo::run")

// Run a server for YomoObjectPlugin
func Run(plugin plugin.YomoObjectPlugin, endpoint string) {
	logger.Infof("plugin service [%s] start... [%s]", plugin.Name(), endpoint)

	// pprof
	go pprof.Run()

	// activation service
	framework.NewServer(endpoint, plugin)
}

// RunStream run a server for YomoStreamPlugin
func RunStream(plugin plugin.YomoStreamPlugin, endpoint string) {
	logger.Infof("plugin service [%s] start... [%s]", plugin.Name(), endpoint)

	// activation service
	panic("not impl")
}

// RunDev makes test plugin connect to a demo YoMo server
func RunDev(plugin plugin.YomoObjectPlugin, endpoint string) {

	go func() {
		time.Sleep(2 * time.Second)
		yomoPluginClient, err := util.QuicClient(endpoint)
		if err != nil {
			panic(err)
		}

		st, _ := yomoPluginClient.NewBidirectionalStream()

		go func() {
			for {
				if st.IsReadable() {
					buf := make([]byte, 3*1024)
					//index := 0
					for {
						n, fin, err := st.Read(buf)
						if err != nil {
							fmt.Println("client:", err)
							break
						}

						if n > 0 {
							fmt.Println(string(buf[:n]))
						}
						if fin {
							break
						}
					}
				}
			}
		}()

		i := 0
		for {
			i++
			st.Write([]byte(("{\"id\":" + fmt.Sprint(i) + ",\"name\":\"yomo!\",\"test\":{\"tag\":[\"5G\",\"ioT\"]}}\n")))
			st.Send()
			time.Sleep(1 * time.Second)
		}

	}()
	logger.Infof("plugin service [%s] start... [%s]", plugin.Name(), endpoint)

	framework.NewServer(endpoint, plugin)

}
