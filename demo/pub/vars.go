package pub

import (
	//"github.com/go-redis/redis"
	"github.com/xjieinfo/xjgo/xjgorm"
	"github.com/xjieinfo/xjgo/xjhttp"
	//"github.com/xjieinfo/xjgo/xjrpc"
	//"gopkg.in/olivere/elastic.v6"
	"gorm.io/gorm"
)

var (
	AppConfig  xjhttp.AppConfig
	Gorm       *gorm.DB
	BaseMapper xjgorm.BaseMapper
	TxMapper   xjgorm.TxMapper
	//Redisdb    *redis.Client
	//EsClient   *elastic.Client
	//RpcClient  *xjrpc.Client
)
