package main


import "fmt"


type item struct{
	observersList []Observer
	name  string
	inStock bool
}

func newItem(name string) *item{
	// obList := make([]Observer, 0)
	item := &item{
		// observersList: obList,
		name: name,
	}
	return item
}

func (i *item) updateAvailability(){
	fmt.Printf("Item:%s now in stock", i.name)
	i.inStock = true
	i.notifyAll()
}

func (i *item) register(obs Observer){
	i.observersList = append(i.observersList, obs)
}

func (i *item) deregister(obs Observer){
	i.observersList = i.removeFromSlice(i.observersList, obs)
}

func (i *item) notifyAll(){
	for _, observer := range i.observersList{
		observer.update(i.name)
	}
}


func (i *item) removeFromSlice(obsList []Observer, obs Observer) []Observer{
	currentLen := len(obsList) - 1
	for i, ob := range obsList{
		if ob.getID() == obs.getID(){
			obsList[currentLen], obsList[i] = obsList[i], obsList[currentLen]
			obsList = obsList[:currentLen]
			return obsList
		}
	}
	return obsList
}