package main

import (
	"database/sql"
	"errors"
)

// The product struct

type product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// methods that deal with single product

func (p *product) getProduct(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (p *product) updateProduct(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (p *product) deleteProduct(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (p *product) createProduct(db *sql.DB) error {
	return errors.New("Not implemented")
}

// standalone function to fetch a list of products

func getProducts(db *sql.DB, start, count int) ([]product, error) {
	return nil, errors.New("Not implemented")
}
