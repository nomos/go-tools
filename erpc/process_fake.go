package erpc

import (
	"github.com/nomos/go-lokas"
	"github.com/nomos/go-lokas/log"
	"github.com/nomos/go-lokas/network/dockerclient"
	"github.com/nomos/go-lokas/network/etcdclient"
	"github.com/nomos/go-lokas/network/ossclient"
	"github.com/nomos/go-lokas/network/redisclient"
	"github.com/nomos/go-lokas/protocol"
	"github.com/nomos/go-lokas/util"
	"github.com/nomos/qmgo"
)

var _ lokas.IServiceRegisterMgr = util.Nil[*serviceRegisterMgrFake]()

type serviceRegisterMgrFake struct {
}

func (this *serviceRegisterMgrFake) Register(info *lokas.ServiceInfo) error {
	return nil
}
func (this *serviceRegisterMgrFake) Unregister(serviceType string, serviceId uint16) error {
	return nil
}
func (this *serviceRegisterMgrFake) UpdateServiceInfo(info *lokas.ServiceInfo) error {
	return nil
}

var _ lokas.IProcess = util.Nil[*processFake]()

func GetFakeProcess() lokas.IProcess {
	return &processFake{}
}

type processFake struct {
	*serviceRegisterMgrFake
}

func (this *processFake) GetServerRegisterMgr() lokas.IServiceRegisterMgr {
	return this
}

func (this *processFake) GetProcessIdByActor(actorId util.ID) (util.ProcessId, error) {
	return 0, nil
}

func (this *processFake) RegisterActors() error {
	return nil
}

func (this *processFake) RegisterActorRemote(actor lokas.IActor) error {
	return nil
}

func (this *processFake) UnregisterActorRemote(actor lokas.IActor) error {
	return nil
}

func (this *processFake) RegisterActorLocal(actor lokas.IActor) error {
	return nil
}

func (this *processFake) UnregisterActorLocal(actor lokas.IActor) error {
	return nil
}

func (this *processFake) GetActorIdsByTypeAndServerId(serverId int32, typ string) []util.ID {
	return nil
}

func (this *processFake) AddActor(actor lokas.IActor) {
}

func (this *processFake) RemoveActor(actor lokas.IActor) {
}

func (this *processFake) RemoveActorById(id util.ID) lokas.IActor {
	return nil
}

func (this *processFake) GetActorIds() []util.ID {
	return nil
}

func (this *processFake) GetActor(id util.ID) lokas.IActor {
	return nil
}

func (this *processFake) StartActor(actor lokas.IActor) error {
	return nil
}

func (this *processFake) RouteMsg(msg *protocol.RouteMessage) {
}

func (this *processFake) Send(id util.ProcessId, msg *protocol.RouteMessage) error {
	return nil
}

func (this *processFake) Add(modules lokas.IModule) lokas.IModule {
	return nil
}

func (this *processFake) RegisterModule(ctor lokas.IModuleCtor) {
}

func (this *processFake) LoadAllModule(config lokas.IProcessConfig) error {
	return nil
}

func (this *processFake) LoadMod(name string, conf lokas.IConfig) error {
	return nil
}

func (this *processFake) UnloadMod(name string) error {
	return nil
}

func (this *processFake) Get(name string) lokas.IModule {
	return nil
}

func (this *processFake) Load(config lokas.IProcessConfig) error {
	return nil
}

func (this *processFake) Start() error {
	return nil
}

func (this *processFake) Stop() error {
	return nil
}

func (this *processFake) PId() util.ProcessId {
	return 0
}

func (this *processFake) GetId() util.ID {
	return 0
}

func (this *processFake) Type() string {
	return ""
}

func (this *processFake) GenId() util.ID {
	return 0
}

func (this *processFake) GetLogger() *log.ComposeLogger {
	return nil
}

func (this *processFake) GetMongo() *qmgo.Database {
	return nil
}

func (this *processFake) GetRedis() *redisclient.Client {
	return nil
}

func (this *processFake) GetEtcd() *etcdclient.Client {
	return nil
}

func (this *processFake) GetOss() *ossclient.Client {
	return nil
}

func (this *processFake) GetDocker() (*dockerclient.TLSClient, error) {
	return nil, nil
}

func (this *processFake) GlobalMutex(key string, ttl int) (*etcdclient.Mutex, error) {
	return nil, nil
}

func (this *processFake) Config() lokas.IConfig {
	return nil
}

func (this *processFake) GameId() string {
	return ""
}

func (this *processFake) ServerId() int32 {
	return 0
}

func (this *processFake) GameServerId() string {
	return ""
}

func (this *processFake) Version() string {
	return ""
}
