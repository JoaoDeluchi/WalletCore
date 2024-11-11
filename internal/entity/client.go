package entity 

import (
	"time"

	"github.com/google/uuid"
)

type Client struct {
	ID string 
	Name string 
	Email string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (c *Client) Validate() error{
	if c.Name == "" {
		return errors.New("name is required")
	}
	// should also have validation for email syntax
	if c.Email == "" {
		return errors.New("email is required")
	}
	return nil 
}

func (c *Client) Update(name, email string) error{
	c.Name = name
	// should also have validation for email syntax
	c.Email = email
	c.UpdatedAt = time.Now()

	err := client.Validate()
	if err != nil {
		return nil, err 
	}
	
	return nil 
}

func NewClient(name, email string) (*Client, error){
	client := &Client{
		ID: uuid.New().String(),
		Name: name,
		Email: email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := client.Validate()
	if err != nil {
		return nil, err 
	}

	return client, nil
}
