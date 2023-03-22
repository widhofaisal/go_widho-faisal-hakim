```sql
SELECT * FROM transaction_details LEFT JOIN products ON transaction_details.product_id=products.id;
```