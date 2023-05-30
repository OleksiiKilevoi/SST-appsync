package dao


type User struct {
	ID          string `dynamo:"id,hash"`
	Email       string `dynamo:"email" index:"email,hash"`
	PhoneNumber string `dynamo:"phoneNumber"`
}

// func (a *Account) ToAccountModel() *model.Account {
// 	if a == nil {
// 		return nil
// 	}
// 	return &model.Account{
// 		ID:          a.ID,
// 		CreatedAt:   a.CreatedAt,
// 		Email:       a.Email,
// 		PhoneNumber: a.PhoneNumber,
// 		FirstName:   a.FirstName,
// 		LastName:    a.LastName,
// 		Country:     a.Country,
// 		City:        a.City,
// 		Address:     a.Address,
// 	}
// }

// func (a *Account) ToPendingAccountModel() *model.PendingAccount {
// 	if a == nil {
// 		return nil
// 	}
// 	r := new(model.PendingAccount)
// 	r.ID = a.ID
// 	r.CreatedAt = a.CreatedAt
// 	r.Email = a.Email
// 	if a.PhoneNumber != "" {
// 		r.PhoneNumber = &a.PhoneNumber
// 	}
// 	if a.FirstName != "" {
// 		r.FirstName = &a.FirstName
// 	}
// 	if a.LastName != "" {
// 		r.LastName = &a.LastName
// 	}
// 	if a.Country != "" {
// 		r.Country = &a.Country
// 	}
// 	if a.City != "" {
// 		r.City = &a.City
// 	}
// 	if a.Address != "" {
// 		r.Address = &a.Address
// 	}
// 	r.Status = model.AccountStatus(a.Status)

// 	return r
// }

// type Card struct {
// 	ID           string `dynamo:"id,hash"`
// 	EnfuceID     string `dynamo:"enfuceId"`
// 	AccountID    string `dynamo:"accountId" index:"accountId,hash"`
// 	MaskedNumber string `dynamo:"maskedNumber"`
// }

// type Payee struct {
// 	ID            string  `dynamo:"id,hash"`
// 	BelongsTo     string  `dynamo:"belongsTo" index:"belongsTo,hash"`
// 	AccountNumber string  `dynamo:"accountNumber"`
// 	Name          string  `dynamo:"name"`
// 	Note          *string `dynamo:"note"`
// 	SortCode      string  `dynamo:"sortCode"`
// }

// func (p *Payee) ToPayeeModel() *model.Payee {
// 	return &model.Payee{
// 		ID:            p.ID,
// 		Name:          p.Name,
// 		AccountNumber: p.AccountNumber,
// 		Note:          p.Note,
// 		SortCode:      p.SortCode,
// 	}
// }

// type Vault struct {
// 	ID          string          `dynamo:"id,hash"`
// 	AccountID   string          `dynamo:"accountId" index:"accountId,hash"`
// 	MainVaultID *string         `dynamo:"mainVaultId" index:"mainVaultId,hash"`
// 	Name        string          `dynamo:"name"`
// 	Balance     float64         `dynamo:"balance"`
// 	Color       string          `dynamo:"color"`
// 	Type        model.VaultType `dynamo:"type"`
// }

// func (v *Vault) ToVaultModel() *model.Vault {
// 	if v == nil {
// 		return nil
// 	}
// 	return &model.Vault{
// 		ID:        v.ID,
// 		Name:      v.Name,
// 		Balance:   v.Balance,
// 		Color:     v.Color,
// 		Type:      v.Type,
// 		SubVaults: []*model.Vault{},
// 	}
// }

// type Transaction struct {
// 	ID              string   `json:"id"`
// 	ExternalID      string   `json:"external_id"`
// 	Ref             *string  `json:"ref"`
// 	Amount          float64  `json:"amount"`
// 	ForeignAmount   *float64 `json:"foreign_amount"`
// 	ForeignCurrency *string  `json:"foreign_currency"`
// 	Timestamp       string   `json:"timestamp"`
// 	// Always populated for incoming transactions
// 	SourceAccountID     *string `json:"source_account_id"`
// 	SourceVaultID       *string `json:"source_vault_id"`
// 	SourceAccountNumber *string `json:"source_account_number"`
// 	SourceSortCode      *string `json:"source_sort_code"`
// 	SourceName          *string `json:"source_name"`
// 	// Always populated for outgoing transactions
// 	TargetAccountID     *string                   `json:"target_account_id"`
// 	TargetVaultID       *string                   `json:"target_vault_id"`
// 	TargetAccountNumber *string                   `json:"target_account_number"`
// 	TargetSortCode      *string                   `json:"target_sort_code"`
// 	TargetName          *string                   `json:"target_name"`
// 	Status              model.TransactionStatus   `json:"status"`
// 	Type                model.TransactionType     `json:"type"`
// 	Category            *string                   `json:"category"`
// 	Security            model.TransactionSecurity `json:"security"`
// }

// func (t *Transaction) ToTransactionModel(
// 	sourceAccount *model.Account,
// 	sourceVault *model.Vault,
// 	targetAccount *model.Account,
// 	targetVault *model.Vault,
// ) *model.Transaction {
// 	if t == nil {
// 		return nil
// 	}
// 	return &model.Transaction{
// 		ID:                  t.ID,
// 		Ref:                 t.Ref,
// 		Amount:              t.Amount,
// 		ForeignAmount:       t.ForeignAmount,
// 		ForeignCurrency:     t.ForeignCurrency,
// 		Timestamp:           t.Timestamp,
// 		SourceAccount:       sourceAccount,
// 		SourceVault:         sourceVault,
// 		TargetAccount:       targetAccount,
// 		TargetVault:         targetVault,
// 		TargetAccountNumber: t.TargetAccountNumber,
// 		TargetName:          t.TargetName,
// 		Status:              t.Status,
// 		Type:                t.Type,
// 		Category:            t.Category,
// 		Security:            t.Security,
// 	}
// }
