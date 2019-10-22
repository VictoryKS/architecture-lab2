package balancers

import (
	"database/sql"
	"fmt"
)

type Balancer struct {
	Id   int  `json:"id"`
	Name string `json:"name"`
}

type Machine struct {
	Id   int  `json:"id"`
	Name string `json:"name"`
	IsWorking int `json:"isWorking"`
}

type Relations struct {
	IdBalancer int  `json:"id_balancer"`
	IdMachine int  `json:"id_machine"`
	IsWorking int `json:"isWorking"`
}

type Store struct {
	Db *sql.DB
}

type Result struct {
    Id int `json:"id"`
		UsedMachines []int `json:"usedMachines"`
		TotalMachinesCount int `json:"totalMachinesCount"`
}

func NewStore(db *sql.DB) *Store {
	return &Store{Db: db}
}

func (s *Store) GetBalancer(id int) (*Result, error) {

	rows, err := s.Db.Query("SELECT relations.id_balancer, relations.id_machine, " +
		"machines.is_working FROM relations JOIN machines ON machines.id = relations.id_machine " +
		"WHERE relations.id_balancer = ($1)", id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
  var res Result
	res.Id = id
	res.UsedMachines = make([]int, 0)
	res.TotalMachinesCount = 0
	for rows.Next() {
		var c Relations
		if err := rows.Scan(&c.IdBalancer, &c.IdMachine, &c.IsWorking); err != nil {
			return nil, err
		}

		if (c.IsWorking == 1) {
			res.UsedMachines = append(res.UsedMachines, c.IdMachine)
		}
	  res.TotalMachinesCount++
	}
	return &res, nil
}

func (s *Store) MachineStatus(id int, isWorking int) error {
	if id <= 0 {
		return fmt.Errorf("incorrect balancer id")
	}
	_, err := s.Db.Exec("UPDATE machines SET is_working = ($2) WHERE id = ($1)", id, isWorking)
	return err
}
