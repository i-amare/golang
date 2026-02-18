GET /<store-id> -> any
GET /<store-id>/<item-id> -> any

POST /<store-id> -> admin
DELETE /<store-id> -> admin

PUT /<store-id>/<item-id> -> owner/admin
DELETE /<store-id>/<item-id> -> owner/admin

POST /signup -> any
POST /login -> any

POST /order -> any
PUT /order -> owner/admin
DELETER /order -> owner/admin