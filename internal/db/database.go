package db

import (
  "fmt"
	"crypto/rand"
	"meme_service/internal/shared/types"
)

var customerDatabase = map[string]types.Customer{}
var customerTokenDatabase = map[string]int{}

func AddCustomer(input CustomerInput) (CustomerOuput, error) {
  uuid, err := generateUUID()
  if err != nil {
    return nil, err
  }

  customer := types.Customer {
    ID: uuid, 
    Email: input.GetEmail(),
    Name: input.GetName(),
  }
  customerDatabase[customer.ID] = customer
  return customer, nil
}

func AddCustomerToken(input CustomerTokenInput) error {
  customerTokenDatabase[input.GetCustomerID()] = input.GetTokens()
  return nil
}

func ConsumerCustomerToken(input ConsumerTokenInput) error {
  tokens := customerTokenDatabase[input.GetCustomerID()]
  tokens = tokens -1
  if (tokens < 0) {
    return fmt.Errorf("There is not token available")
  }
  customerTokenDatabase[input.GetCustomerID()] = tokens;
  return nil
}

func GetCustomer() {

}

func ListCustomers() ([]types.Customer, error) {
  customerList := []types.Customer{}
  for _, v := range customerDatabase {
    customerList = append(customerList, v)
  }
  return customerList, nil
}

func generateUUID() (string, error) {
	uuid := make([]byte, 16)

	_, err := rand.Read(uuid)
	if err != nil {
		return "", err
	}

	uuid[6] = (uuid[6] & 0x0F) | 0x40

	uuid[8] = (uuid[8] & 0x3F) | 0x80

	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}
