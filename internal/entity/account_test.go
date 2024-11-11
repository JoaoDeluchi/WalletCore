package entity

func TestCreateNewAccount(t *testing.T){
	client := NewClient("test account", "test@account")
	got := NewAccount(client)

	assert.NotNil(t, got)
}

func TestCreateNewAccountWhenClientIsNil(t *testing.T){
	got := NewAccount(nil)
	assert.Nil(t, got)
}

func TestAccountCreditAndDebit(t *testing.T){
	client := NewClient("test account", "test@account")
	account := NewAccount(client)

	account.Credit(20)
	assert.Equal(t, account.Balance, 20)

	account.Debit(10)
	assert.Equal(t, account.Balance, 10)
}