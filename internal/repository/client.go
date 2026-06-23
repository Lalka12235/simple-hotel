package repository

import (
	"database/sql"
	"github.com/Lalka12235/simple-hotel.git/internal/model"
)

type ClientRepository struct {
	db *sql.DB
}

func NewClientRepository(db *sql.DB) *ClientRepository {
	return &ClientRepository{db: db}
}

func (r *ClientRepository) Create(c *model.Client) (int, error) {
	query := `INSERT INTO clients (first_name, last_name, surname, address, passport, coment) 
	          VALUES ($1, $2, $3, $4, $5, $6) RETURNING id_client`
	var id int
	err := r.db.QueryRow(query, c.FirstName, c.LastName, c.Surname, c.Address, c.Passport, c.Coment).Scan(&id)
	return id, err
}

func (r *ClientRepository) GetByID(id int) (*model.Client, error) {
	query := `SELECT id_client, first_name, last_name, surname, address, passport, coment FROM clients WHERE id_client = $1`
	var c model.Client
	err := r.db.QueryRow(query, id).Scan(&c.IdClient, &c.FirstName, &c.LastName, &c.Surname, &c.Address, &c.Passport, &c.Coment)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *ClientRepository) Update(c *model.Client) error {
	query := `UPDATE clients SET first_name=$1, last_name=$2, surname=$3, address=$4, passport=$5, coment=$6 WHERE id_client=$7`
	_, err := r.db.Exec(query, c.FirstName, c.LastName, c.Surname, c.Address, c.Passport, c.Coment, c.IdClient)
	return err
}

func (r *ClientRepository) Delete(id int) error {
	query := `DELETE FROM clients WHERE id_client = $1`
	_, err := r.db.Exec(query, id)
	return err
}