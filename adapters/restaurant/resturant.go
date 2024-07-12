package restaurant

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/fabianogoes/fiap-kitchen/domain/entities"
)

type ClientAdapter struct {
	config *entities.Config
}

func NewClientAdapter(config *entities.Config) ClientAdapter {
	return ClientAdapter{config: config}
}

func (p *ClientAdapter) ReadyForDelivery(orderID uint) error {
	fmt.Printf("ReadyForDelivery orderID: %d \n", orderID)

	url := fmt.Sprintf("%s/orders/%d/ready-for-delivery", p.config.RestaurantApiUrl, orderID)
	req, err := http.NewRequest(http.MethodPut, url, nil)
	if err != nil {
		log.Fatalf("An Error Occured to prepar request %v", err)
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	fmt.Printf("calling restaurant ready for delivery url %s \n", url)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("An Error Occured to call restaurant ready for delivery %v", err)
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("An Error Occured to read response body %v", err)
		return err
	}

	sb := string(body)
	log.Println(sb)

	return nil
}
