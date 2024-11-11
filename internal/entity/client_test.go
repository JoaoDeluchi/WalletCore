package entity 

import ( 
	"testing"
	"github.com/stretchr/testify/assert"
)


func TestCreateNewClient(t *testing.T) {
	client, err := NewClient("John", "jj@test.com")
	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "John", client.Name)
	assert.Equal(t, "jj@test.com", client.Email)
}

func TestCreateNewClientWhenArgsAreInvalid(t *testing.T) {
	client, err := NewClient("", "")
	assert.NotNil(t, err)
	assert.nil(t, client)
}


func TestUpdateClient(t *testing.T) {
	client, err := NewClient("John", "jj@test.com")
	err := client.Update("Updated John", "jj_updated@test.com")

	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "Updated John", client.Name)
	assert.Equal(t, "jj_updated@test.com", client.Email)
	
}

func TestUpdatetWhenArgsAreInvalid(t *testing.T) {
	client, err := NewClient("Not Updated John", "jj@test.com")
	err := client.Update("", "jj_updated@test.com")

	assert.Error(t, err, "name is required")
	
	assert.Equal(t, "Not Updated John", client.Name)
}