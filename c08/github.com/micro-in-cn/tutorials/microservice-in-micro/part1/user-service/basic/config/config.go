package config

import (
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/source"
	"github.com/micro/go-micro/v2/config/source/file"
	log "github.com/micro/go-micro/v2/logger"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// 1.定义要用到的变量
// 2.加载配置文件
// 3.分别解析配置文件
var (
	err error
)

var (
	defaultRootPath         = "app"
	defaultConfigFilePrefix = "application-"
	etcdConfig              defaultEtcdConfig
	mysqlConfig             defaultMysqlConfig
	profiles                defaultProfiles
	m                       sync.RWMutex
	inited                  bool
	sp                      = string(filepath.Separator)
)

func Init() {
	m.Lock()
	defer m.Unlock()

	if inited {
		log.Info("[Init] 配置已经初始化过")
		return
	}

	// 加载yml配置
	// 先加载基础配置
	// filepath.Abs是绝对路径时直接返回，不是绝对路径时将与当前工作路径拼成绝对路径
	// filepath.Dir获取除文件名外的部分路径（最后一个分隔符没有，除非是根目录）
	// os.Chdir改变当前工作目录到指定的path
	appPath, _ := filepath.Abs(filepath.Dir(filepath.Join(".", sp)))
	pt := filepath.Join(appPath, "conf")
	os.Chdir(appPath)

	// 加载application.yml文件
	if err = config.Load(file.NewSource(file.WithPath(pt + sp + "application.yml"))); err != nil {
		panic(err)
	}
	// 找到需要引入的新配置文件
	if err = config.Get(defaultRootPath, "profiles").Scan(&profiles); err != nil {
		panic(err)
	}
	log.Infof("[Init] 加载配置文件：path:%s,%+v\n", pt+sp+"application.yml", profiles)

	// 开始导入新文件
	// file.NewSource
	// file.WithPath
	if len(profiles.GetInclude()) > 0 {
		include := strings.Split(profiles.GetInclude(), ",")
		sources := make([]source.Source, len(include))
		for i := 0; i < len(include); i++ {
			filePath := pt + string(filepath.Separator) + defaultConfigFilePrefix + strings.TrimSpace(include[i]) + ".yml"
			log.Infof("[Init] 加载配置文件：path: %s\n", filePath)
			sources[i] = file.NewSource(file.WithPath(filePath))
		}

		// 加载include的文件
		if err = config.Load(sources...); err != nil {
			panic(err)
		}
	}

	// 赋值
	config.Get(defaultRootPath, "etcd").Scan(&etcdConfig)
	config.Get(defaultRootPath, "mysql").Scan(&mysqlConfig)
	// 标记已经初始化
	inited = true
}

// GetMysqlConfig 获取mysql配置
func GetMysqlConfig() (ret MysqlConfig) {
	return mysqlConfig
}

// GetEtcdConfig 获取Etcd配置
func GetEtcdConfig() (ret EtcdConfig) {
	return etcdConfig
}