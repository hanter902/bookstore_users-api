package services

var (
	ItemsService itemsServiceInterface = &itemsService{}
)

type itemsServiceInterface interface {
	GetItems()
	SaveItem()
}

type itemsService struct {
}

func (s *itemsService) GetItems() {

}

func (s *itemsService) SaveItem() {

}
