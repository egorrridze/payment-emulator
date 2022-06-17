# payment-emulator
Payment emulator API

### Launch

```
docker build -t emulator .
docker run -p 8080:8080 -t emulator
```

### Endpoints:
endpoint | method | body | response
--- | --- | --- | ---
/payments|POST|{user_id, user_email, sum, currency}| payment id, payment status 
/payments/status/:id|POST| |updated payment status
/payments/status/:id |GET| |current payment status
/payments|POST|{user_id/user_email}|list of users payments
/payments/:id|DELETE| |status of deletion

