package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"time"
)
const (
	MongoDBHosts = "127.0.0.1"
	AuthUserName = "xdmp"
	AuthPassword = "20E6QK8V"
	MaxCon = 300
)
type Person struct {
	Name      string
	Phone     string
	City      string
	Age       int8
	IsMan     bool
	Interests []string
}
func main()  {
	mongoDBDialInfo :=&mgo.DialInfo{
		Addrs: []string{MongoDBHosts},
		Timeout: 60*time.Second,
		Username: AuthUserName,
		Password: AuthPassword,
		Database: "admin",
	}
	session,err := mgo.DialWithInfo(mongoDBDialInfo)
	if err !=nil{
		log.Fatalf("CreateSession failed:%\n",err)

	}
	session.SetPoolLimit(MaxCon)
	defer session.Close()
	err = createData(session,"test","people")
	//persons,err := findPerson(session,"test","people")
	if err !=nil{
		panic(err)
	}
	//fmt.Printf("%v",persons)
}
func createData(session *mgo.Session,dbname string,tablename string)error  {
	persons := []Person{
		Person{Name: "Tony", Phone: "123432", City: "Shanghai", Age: 33, IsMan: true, Interests: []string{"music", "tea", "collection"}},
		Person{Name: "Mary", Phone: "232562", City: "Beijing", Age: 43, IsMan: false, Interests: []string{"sport", "film"}},
		Person{Name: "Tom", Phone: "123432", City: "Suzhou", Age: 22, IsMan: true, Interests: []string{"music", "reading"}},
		Person{Name: "Bob", Phone: "123432", City: "Hangzhou", Age: 32, IsMan: true, Interests: []string{"shopping", "coffee"}},
		Person{Name: "Alex", Phone: "15772", City: "Shanghai", Age: 21, IsMan: true, Interests: []string{"music", "chocolate"}},
		Person{Name: "Alice", Phone: "43456", City: "Shanghai", Age: 42, IsMan: false, Interests: []string{"outing", "tea"}},
		Person{Name: "Ingrid", Phone: "123432", City: "Shanghai", Age: 22, IsMan: false, Interests: []string{"travel", "tea"}},
		Person{Name: "Adle", Phone: "123432", City: "Shanghai", Age: 20, IsMan: false, Interests: []string{"game", "coffee", "sport"}},
		Person{Name: "Smith", Phone: "54223", City: "Fuzhou", Age: 54, IsMan: true, Interests: []string{"music", "reading"}},
		Person{Name: "Bruce", Phone: "123432", City: "Shanghai", Age: 31, IsMan: true, Interests: []string{"film", "tea", "game", "shoping", "reading"}},
	}
	cloneSession := session.Clone()
	c := cloneSession.DB(dbname).C(tablename)
	for _,item := range persons{
		err := c.Insert(&item)
		if err !=nil{
			panic(err)
		}
	}
	return nil
}
func findPerson(session *mgo.Session,dbname string,tableName string) (persons []Person,err error) {
	query:=bson.M{"$group":bson.M{"_id":"city","num":bson.M{"$sum":"1"}}}
	err =session.DB(dbname).C(tableName).Find(query).All(&persons)

	if err !=nil{
		return persons,err
	}
	return persons,nil
}
type UserServiceInfo struct {
	ID               string                 `bson:"_id,omitempty" json:"id,omitempty"`
	CreateTime       string                 `bson:"create_time,omitempty" json:"create_time,omitempty"`
	UpdateTime       string                 `bson:"update_time,omitempty" json:"update_time,omitempty"`
	UserID           string                 `bson:"user_id,omitempty" json:"user_id,omitempty"`
	Nick             string                 `bson:"nick,omitempty" json:"nick,omitempty"`
	ServiceName      string                 `bson:"service_name,omitempty" json:"service_name,omitempty"`
	ServiceLevel     string                 `bson:"service_level,omitempty" json:"service_level,omitempty"`
	ExpireTime       string                 `bson:"expire_time,omitempty" json:"expire_time,omitempty"`
	InviteBy         string                 `bson:"invite_by,omitempty" json:"invite_by,omitempty"`
	Settings         map[string]interface{} `bson:"settings,omitempty" json:"settings,omitempty"`
	SendMsgSettings  map[string]interface{} `bson:"send_msg_settings,omitempty" json:"send_msg_settings,omitempty"`
	GuideStatus      map[string]bool        `bson:"guide_status,omitempty" json:"guide_status,omitempty"`
	LastChannel      string                 `bson:"last_channel,omitempty" json:"last_channel,omitempty"`
	GoodsSizeConfig  map[string]interface{} `bson:"goods_size_config,omitempty" json:"goods_size_config,omitempty"`
	IsCVD            bool                   `bson:"is_cvd,omitempty" json:"is_cvd,omitempty"`
	BusyRespSettings map[string]interface{} `bson:"busy_resp_settings,omitempty" json:"busy_resp_settings,omitempty"`

	PluginExpireTime   string `bson:"plugin_expire_time,omitempty" json:"plugin_expire_time,omitempty"`
	PluginServiceLevel string `bson:"plugin_service_level,omitempty" json:"plugin_service_level,omitempty"`
	KlkExpireTime      string `bson:"klk_expire_time,omitempty" json:"klk_expire_time,omitempty"`
	KlkServiceLevel    string `bson:"klk_service_level,omitempty" json:"klk_service_level,omitempty"`
	OriServiceLevel    string `bson:"ori_service_level,omitempty" json:"ori_service_level,omitempty"`
	Channel            string `bson:"channel,omitempty" json:"channel,omitempty"`

	IsExpire bool `bson:"-,omitempty" json:"is_expire,omitempty"`
}
//func findUserService(session *mgo.Session,dbname string,tableName string)(user *UserServiceInfo,err error)  {
//
//}