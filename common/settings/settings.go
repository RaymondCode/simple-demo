package settings

//  配置文件
import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Init 此初始化需要在main中调用
func Init() (err error) {
	viper.SetConfigFile("config.yaml")
	viper.AddConfigPath(".") // 还可以在工作目录中查找配置
	err = viper.ReadInConfig()
	if err != nil { // 处理读取配置文件的错误
		fmt.Printf("viper.ReadInConfig failed err:  %s", err)
		return err
	}
	//动态检测config.yaml文件变化
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Printf("config file changed...\n")
	})
	return err
}
