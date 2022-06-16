# go-todo-rest-api
Payment emulator API

### Endpoints:
endpoint | method | body | response
--- | --- | --- | ---
/payments|POST|{user_id, user_email, sum, currency}| payment id, payment status 
/payments/status/:id|POST| |updated payment status
/payments/status/:id |GET| |current payment status
/payments|POST|{user_id/user_email}|list of users payments
/payments/:id|DELETE| |status of deletion

