package serviceClient

import (
	"github.com/Diode222/Frigg/conf"
	"github.com/Diode222/Frigg/manager"
	pb "github.com/Diode222/Frigg/proto_gen"
	"sync"
)

var splittedWordClient pb.WordSplitServiceClient
var splittedWordClientOnce sync.Once

func GetSplittedWordClient(etcdAddr string) pb.WordSplitServiceClient {
	splittedWordClientOnce.Do(func() {
		splittedWordClient = manager.GetServiceManger(etcdAddr).GetClient(conf.ALGORITHM_SERVICE_NAME, pb.NewWordSplitServiceClientWrapper).(pb.WordSplitServiceClient)
	})
	return splittedWordClient
}
