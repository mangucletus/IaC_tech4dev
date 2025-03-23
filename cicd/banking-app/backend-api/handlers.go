package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func getAccounts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	db.mu.RLock()
	accounts := make([]Account, 0, len(db.accounts))
	for _, account := range db.accounts {
		accounts = append(accounts, account)
	}
	db.mu.RUnlock()
	json.NewEncoder(w).Encode(accounts)
}

func createAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newAccount Account
	err := json.NewDecoder(r.Body).Decode(&newAccount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	db.mu.Lock()
	db.accountID++
	newAccount.ID = strconv.Itoa(db.accountID)
	db.accounts[newAccount.ID] = newAccount
	db.mu.Unlock()
	
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newAccount)
}

func getAccountById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	// Extract account ID from URL
	id := r.PathValue("id")
	
	db.mu.RLock()
	account, exists := db.accounts[id]
	db.mu.RUnlock()
	
	if !exists {
		http.Error(w, "Account not found", http.StatusNotFound)
		return
	}
	
	json.NewEncoder(w).Encode(account)
}

func updateAccountById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	id := r.PathValue("id")
	
	// Update account
	var updatedAccount Account
	err := json.NewDecoder(r.Body).Decode(&updatedAccount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	// Preserve ID
	updatedAccount.ID = id
	
	db.mu.Lock()
	db.accounts[id] = updatedAccount
	db.mu.Unlock()
	
	json.NewEncoder(w).Encode(updatedAccount)
}
		
func deleteAccountById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := r.PathValue("id")
	// Delete account
	db.mu.Lock()
	delete(db.accounts, id)
	db.mu.Unlock()
	
	w.WriteHeader(http.StatusNoContent)
}

func getAllTransactions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	// List all transactions
	db.mu.RLock()
	transactions := make([]Transaction, 0, len(db.transactions))
	for _, txn := range db.transactions {
		transactions = append(transactions, txn)
	}
	db.mu.RUnlock()
	
	json.NewEncoder(w).Encode(transactions)
}

func createTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Create a new transaction
	var newTxn Transaction
	err := json.NewDecoder(r.Body).Decode(&newTxn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	// Verify account exists
	db.mu.RLock()
	_, accountExists := db.accounts[newTxn.AccountID]
	db.mu.RUnlock()
	
	if !accountExists {
		http.Error(w, "Account not found", http.StatusBadRequest)
		return
	}
	
	// Process transaction
	db.mu.Lock()
	db.transactionID++
	newTxn.ID = strconv.Itoa(db.transactionID)
	newTxn.Status = "pending" // Set initial status
	db.transactions[newTxn.ID] = newTxn
	
	// Update account balance for completed transactions
	if newTxn.Type == "deposit" {
		account := db.accounts[newTxn.AccountID]
		account.Balance += newTxn.Amount
		newTxn.Status = "completed"
		db.accounts[newTxn.AccountID] = account
		db.transactions[newTxn.ID] = newTxn
	} else if newTxn.Type == "withdrawal" {
		account := db.accounts[newTxn.AccountID]
		if account.Balance >= newTxn.Amount {
			account.Balance -= newTxn.Amount
			newTxn.Status = "completed"
			db.accounts[newTxn.AccountID] = account
			db.transactions[newTxn.ID] = newTxn
		} else {
			newTxn.Status = "failed"
			db.transactions[newTxn.ID] = newTxn
		}
	}
	db.mu.Unlock()
	
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTxn)
}

func getTransactionById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	// Extract transaction ID from URL
	id := r.PathValue("id")
	
	db.mu.RLock()
	txn, exists := db.transactions[id]
	db.mu.RUnlock()
	
	if !exists {
		http.Error(w, "Transaction not found", http.StatusNotFound)
		return
	}
	
	json.NewEncoder(w).Encode(txn)
}


// Update transaction status (e.g., for workflow purposes)
func updateTransactionById(w http.ResponseWriter, r * http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Extract transaction ID from URL
	id := r.PathValue("id")

	var updatedTxn Transaction
	txn, exists := db.transactions[id]
	if !exists {
		http.Error(w, "Transaction not found", http.StatusNotFound)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&updatedTxn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Only allow status updates
	txn.Status = updatedTxn.Status
		
	db.mu.Lock()
	db.transactions[id] = txn
	db.mu.Unlock()
	
	json.NewEncoder(w).Encode(txn)
}

