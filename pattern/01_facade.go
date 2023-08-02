package pattern

/*
«Фасад» - это шаблон проектирования предназначен для того, чтобы скрыть сложности базовой
системы и предоставить клиенту простой интерфейс. Он предоставляет унифицированный интерфейс
для множества интерфейсов в системе, что упрощает его использование с точки зрения клиента.
По сути, он обеспечивает более высокий уровень абстракции над сложной системой.

Сам термин «Фасад» означает выходящий на улицу или открытое пространство
Показана только лицевая сторона здания, за которой скрыта основная сложность.

Фасад. Структура.
Клиент – использует фасад вместо прямой работы с объектами сложной подсистемы.
Фасад – предоставляет быстрый доступ к определённой функциональности подсистемы. Он «знает», каким классам нужно переадресовать запрос, и какие данные для этого нужны.
Сложная подсистема – состоит из множества разнообразных классов. Для того, чтобы заставить их что-то делать, нужно знать подробности устройства подсистемы, порядок инициализации объектов и так далее.
Дополнительный фасад – можно ввести, чтобы не «захламлять» единственный фасад разнородной функциональностью. Он может использоваться как клиентом, так и другими фасадами.

Фасад. Применимость.
Когда нужно представить простой или урезанный интерфейс к сложной подсистеме.
Когда надо уменьшить количество зависимостей между клиентом и сложной системой. Фасадные объекты позволяют отделить, изолировать компоненты системы от клиента и развивать и работать с ними независимо.
Когда вы хотите разложить подсистему на отдельные слои.

Преимущество:
изолирует клиентов от компонентов сложной подсистемы.

Недостаток:
фасад рискует стать божественным объектом, привязанным ко всем классам программы.

Пример: цифровой кошельк, когда кто-то фактически делает дебет/кредит кошелька, в фоновом
режиме происходит множество вещей, о которых клиент может не знать. В приведенном ниже
списке показаны некоторые действия, происходящие в процессе кредитования/дебетования.

Проверить счет
Проверить защитный PIN-код
Кредитный/дебетовый баланс
Сделать запись в книге
Отправить уведомление

Здесь уместно использовать паттерн Facade.
Клиенту достаточно ввести Номер кошелька, Защитный PIN-код, Сумма и указать тип операции.
Об остальных вещах заботятся в фоновом режиме.
*/

import (
	"fmt"
	"log"
)

type notification struct{}

type walletFacade struct {
	account      *account
	wallet       *wallet
	securityCode *securityCode
	notification *notification
	ledger       *ledger
}

func newWalletFacade(accountID string, code int) *walletFacade {
	fmt.Println("Starting create account")
	walletFacacde := &walletFacade{
		account:      newAccount(accountID),
		securityCode: newSecurityCode(code),
		wallet:       newWallet(),
		notification: &notification{},
		ledger:       &ledger{},
	}
	fmt.Println("Account created")
	return walletFacacde
}

func (w *walletFacade) addMoneyToWallet(accountID string, securityCode int, amount int) error {
	fmt.Println("Starting add money to wallet")
	err := w.account.checkAccount(accountID)
	if err != nil {
		return err
	}
	err = w.securityCode.checkCode(securityCode)
	if err != nil {
		return err
	}
	w.wallet.creditBalance(amount)
	w.notification.sendWalletCreditNotification()
	w.ledger.makeEntry(accountID, "credit", amount)
	return nil
}

func (w *walletFacade) deductMoneyFromWallet(accountID string, securityCode int, amount int) error {
	fmt.Println("Starting debit money from wallet")
	err := w.account.checkAccount(accountID)
	if err != nil {
		return err
	}
	err = w.securityCode.checkCode(securityCode)
	if err != nil {
		return err
	}
	err = w.wallet.debitBalance(amount)
	if err != nil {
		return err
	}
	w.notification.sendWalletDebitNotification()
	w.ledger.makeEntry(accountID, "credit", amount)
	return nil
}

type account struct {
	name string
}

func newAccount(accountName string) *account {
	return &account{
		name: accountName,
	}
}

func (a *account) checkAccount(accountName string) error {
	if a.name != accountName {
		return fmt.Errorf("Account Name is incorrect")
	}
	fmt.Println("Account Verified")
	return nil
}

type securityCode struct {
	code int
}

func newSecurityCode(code int) *securityCode {
	return &securityCode{
		code: code,
	}
}

func (s *securityCode) checkCode(incomingCode int) error {
	if s.code != incomingCode {
		return fmt.Errorf("Security Code is incorrect")
	}
	fmt.Println("SecurityCode Verified")
	return nil
}

type wallet struct {
	balance int
}

func newWallet() *wallet {
	return &wallet{
		balance: 0,
	}
}

func (w *wallet) creditBalance(amount int) {
	w.balance += amount
	fmt.Println("Wallet balance added successfully")
	return
}

func (w *wallet) debitBalance(amount int) error {
	if w.balance < amount {
		return fmt.Errorf("Balance is not sufficient")
	}
	fmt.Println("Wallet balance is Sufficient")
	w.balance = w.balance - amount
	return nil
}

type ledger struct {
}

func (s *ledger) makeEntry(accountID, txnType string, amount int) {
	fmt.Printf("Make ledger entry for accountId %s with txnType %s for amount %d\n", accountID, txnType, amount)
	return
}

func (n *notification) sendWalletCreditNotification() {
	fmt.Println("Sending wallet credit notification")
}

func (n *notification) sendWalletDebitNotification() {
	fmt.Println("Sending wallet debit notification")
}

func useFacade() {
	fmt.Println()
	walletFacade := newWalletFacade("abc", 1234)
	fmt.Println()
	err := walletFacade.addMoneyToWallet("abc", 1234, 10)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
	fmt.Println()
	err = walletFacade.deductMoneyFromWallet("abc", 1234, 5)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
}
