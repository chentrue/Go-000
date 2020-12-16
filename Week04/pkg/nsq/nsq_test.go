package nsq

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"testing"
)

func TestProduct(t *testing.T) {
	config := nsq.NewConfig()
	producer, err := nsq.NewProducer("192.168.30.247:4150", config)
	if err != nil {
		fmt.Println(err)
	}
	var testData = "[{\"id\":1006,\"identifier\":\"[{\\\"to\\\":\\\"15800566405\\\",\\\"content\\\":\\\"这是一条测试短信\\\"}]\",\"title\":\"短信\",\"type\":\"sms\",\"subType\":\"group\",\"content\":\"this is a test msg\",\"company\":2,\"description\":\"\",\"build_sn\":\"\",\"ext_info\":{}},{\"id\":1001,\"identifier\":\"10141\",\"title\":\"app\",\"type\":\"app\",\"content\":\"apptitle\",\"company\":2,\"description\":\"\",\"build_sn\":\"\",\"ext_info\":{\"img\":\"\",\"rom\":\"HUAWEI\",\"brand\":\"HUAWEI\",\"h5_url\":\"\",\"intent\":\"eceibs:\\/\\/com.neweceibs.zo\\/year_course_detail?trainplan_id=1609&resources_id=3653&content_id=3159\",\"sys_id\":0,\"scenario\":\"365\",\"scenario_params\":{\"content_id\":\"3159\",\"resources_id\":\"3653\",\"trainplan_id\":\"1609\"}},\"ext_data\":{\"h5_url\":\"\",\"page_type\":\"message_detail\"},\"intent\":\"eceibs:\\/\\/com.neweceibs.zo\\/message_detail?h5_url=&page_type=message_detail\"},{\"id\":1005,\"identifier\":\"gchenguang@ceibsdigital.com\",\"title\":\"邮件\",\"type\":\"email\",\"content\":\"这是一封测试邮件\",\"company\":2,\"description\":\"\",\"build_sn\":\"\",\"ext_info\":{}}]"
	body := []byte(testData)
	topicName := "service_remind_noreal"
	err = producer.Publish(topicName, body)
	if err != nil {
		fmt.Println(err)
	}
}

func TestConsume(t *testing.T) {

}
