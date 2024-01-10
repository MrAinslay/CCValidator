# CCValidator
 This is a creadit card validator and BIN Lookup rest API. I used Bintable's free BIN Lookup API and Luhn's algorithm to retrieve BIN data and check if a CC number is valid.

# REQUIREMENTS

To use this API you need to get an API Key from bintables.com. All you need to do is register an account. Also you need to create a key.env file in the root of the path containing the PORT, API_KEY and BASE_URL keys

## KEY.ENV BODY
PORT="8080"
API_KEY="<YOUR_API_KEY>"
BASE_URL="https://api.bintable.com/v1/"

 # ENDPOINTS
 ## GET http://localhost:8080/v1/validate ENDPOINT

 Checks if the given CC number is valid and returns BIN data if the account range (first 6 numbers) is found in the database
 
 ### REQUEST BODY
 ```json
 {
  "credit_card_number": "324000xxxxxxxxxx"
 }
 ```
 ### RESPONSE BODY
 ```json
 {
   "result": "Valid credit card number",
   "data": {
     "card_info": {
       "scheme": "American express",
       "type": "Debit",
       "category": "Business"
     },
     "bank": {
       "name": "BRANCH BANKING AND TRUST COMPANY",
       "website": "http://www.bbt.com/"
       "phone": "910-914-8250 OR 800-226-5228"
     },
     "country": {
       "name": "United states of america",
       "code": "us"
     }
   }
 }
 ```
 ## GET http://localhost:8080/v1/accountrng ENDPOINT

 Returns all the BIN data of the account if the account range is found in the database

 ### REQUEST BODY
 ```json
{
 "account_range": "324000"
}
```
### RESPONSE BODY
```json
{
   "result": "Valid account range",
   "data": {
     "card_info": {
       "scheme": "American express",
       "type": "Debit",
       "category": "Business"
     },
     "bank": {
       "name": "BRANCH BANKING AND TRUST COMPANY",
       "website": "http://www.bbt.com/"
       "phone": "910-914-8250 OR 800-226-5228"
     },
     "country": {
       "name": "United states of america",
       "code": "us"
     }
   }
 }
 ```
 
