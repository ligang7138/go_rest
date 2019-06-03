package kafka

import (
	"testing"
)

func newProducer() (*Producer, error) {
	address := []string{
		"172.25.0.36:9092",
		"172.25.0.37:9092",
	}
	p, err := NewProducer(address)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func TestNewProducer(t *testing.T) {
	p, err := newProducer()
	if err != nil {
		t.Error(err)
	}

	t.Logf("pass and id is %v", p)

}

func TestProducer_SyncProducer(t *testing.T) {
	p, err := newProducer()
	if err != nil {
		t.Error(err)
	}
	err = p.SyncProducer("test", "TestProducer_SyncProducer")
	if err != nil {
		t.Error(err)
	} else {
		t.Log("pass!!!")
	}
}

func TestProducer_AsyncProducer(t *testing.T) {

	ch := make(chan int)

	<-ch
}
