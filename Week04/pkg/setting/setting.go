package setting

import (
	"flag"
	"github.com/go-ini/ini"
	"github.com/nsqio/go-nsq"
	"path/filepath"
	"service-notification/pkg/logging"
	"time"
)
var Env string //环境变量
type App struct {
}
var AppSetting = &App{}
type Server struct {
	RunMode string
	MockSend bool
	HttpPort int
	ReadTimeout time.Duration
	WriteTimeout time.Duration
}
var ServerSetting = &Server{}

type Database struct {
	Type string
	User string
	Password string
	Host string
	Port int
	Name string
	TablePrefix string
}
var DbSetting = &Database{}

type Redis struct {
	Host string
	Password string
	MaxIdle int
	MaxActive int
	IdleTimeout time.Duration
}
var RedisSetting = &Redis{}

type Nsq struct {
	NsqD string
	NsqDPort string
	NsqLookupD string
	NsqLookupDPort string
	Concurrency int
	SmsTopic string
	EmailTopic string
	AppTopic string
	CallbackTopic string
	LogLevel nsq.LogLevel
}
var NsqSetting = &Nsq{}

type AliyunPush struct {
	RegionId	string
	AccessKeyId string
	AccessSecret string
	AndroidAppKey string
	AndroidAppSecret string
	IosAppKey string
	IosAppSecret string
}
var AliyunPushSetting = &AliyunPush{}

type Sms struct {
	Account string 	// 账户
	Password string // 密码
	Ip string   	//https 网关 ip 地址
	Port   int		//https网关端口号
	Version string	//版本号
	Domain string	//应用功能业务域
}
var SmsSetting = &Sms{}

func Setup() {
	flag.StringVar(&Env, "env", "dev", "环境变量,(默认dev)")
	flag.Parse()
	configPath := filepath.Join("conf", Env, "app.ini")
	Cfg, err := ini.Load(configPath)
	if err != nil {
		logging.Errorf("Fail to parse '$v':%v ",configPath,err)
	}
	err = Cfg.Section("app").MapTo(AppSetting)
	if err != nil {
		logging.Errorf("Map to AppSetting err:%v", err)
	}
	err = Cfg.Section("server").MapTo(ServerSetting)
	if err != nil {
		logging.Errorf("Map to ServerSetting err:%v", err)
	}
	ServerSetting.ReadTimeout  *= time.Second
	ServerSetting.WriteTimeout *= time.Second
	err = Cfg.Section("database").MapTo(DbSetting)
	if err != nil {
		logging.Errorf("Map to ServerSetting err:%v", err)
	}
	err = Cfg.Section("redis").MapTo(RedisSetting)
	if err != nil {
		logging.Errorf("Map to RedisSetting err :%v", err)
	}
	RedisSetting.IdleTimeout *= time.Second

	err = Cfg.Section("nsq").MapTo(NsqSetting)
	if err != nil {
		logging.Errorf("Map to NsqSetting err:%v", err)
	}

	err = Cfg.Section("aliyunpush").MapTo(AliyunPushSetting)
	if err != nil {
		logging.Errorf("Map to AliyunPush err:%v", err)
	}

	err = Cfg.Section("sms").MapTo(SmsSetting)
	if err != nil {
		logging.Errorf("Map to sms err:%v", err)
	}
}
