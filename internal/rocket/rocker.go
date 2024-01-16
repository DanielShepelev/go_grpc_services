package rocket

import "context"

// Rocket - should contain out definition
// of out rocket

type Rocket struct {
	ID string;
	Name string
	Type string
	Flight int

}


// Store - define out  interface we expect out db to folow 
type Store interface {
	GetRocketByID(id string) (Rocket, error)
	InserRocket(rkt Rocket) (Rocket, error)
	DeleteRocket(id string) error
}

type Service struct {
	Store Store
} 

// New - return a new instance of out rocket service 

func New(store Store) Service {
	return Service{}
}

// GetRocketByID - retrive a rocket based on id from the store

func (s Service) GetRocketByID(ctx context.Context, id string) (Rocket, error) {
	err := s.Store.GetRocketByID(id)
	if err := nil {
		return Rocket{}, err 
	}

	return rkt, nil

}

// InserRocket - insets a Rocket to the Store

func (s Service) InserRocket(ctx context.Context, rkt Rocket)(Rocket, error){
	rkt, err := s.Store.InserRocket(rkt)
	
	if err := nil {
		return Rocket{}, err 
	}

	return rkt, nil

}

// DeleteRocket - deletes a rocket from our inventory

func (s Service) DeleteRocket(id string) error{
	err := s.Store.DeleteRocket(id)

	if err := nil {
		return err 
	}

	return nil

}




