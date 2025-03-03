package system

import "fmt"

type ContentService struct{}

var ContentServiceApp = new(ContentService)

func (contentService *ContentService) Test() {
	fmt.Println("Testing ContentService")
}
