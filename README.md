# Question

We want you to implement a **REST API** endpoint that given a list of products, applies some discounts to them
and can be filtered.

### What we expect

- Code structure/architecture must fit this use case, as simple or as complex needed to complete what
  is asked for.
- Tests are a must. Code must be testable without requiring networking or the filesystem. Tests should
  be runnable with 1 command.
- The project must be runnable with 1 simple command from any machine.
- Explanations on decisions taken

### Given this list of products:

```json
{
  "products": [
    {
      "sku": "000001",
      "name": "BV Lean leather ankle boots",
      "category": "boots",
      "price": 89000
    },
    {
      "sku": "000002",
      "name": "BV Lean leather ankle boots",
      "category": "boots",
      "price": 99000
    },
    {
      "sku": "000003",
      "name": "Ashlington leather ankle boots",
      "category": "boots",
      "price": 71000
    },
    {
      "sku": "000004",
      "name": "Naima embellished suede sandals",
      "category": "sandals",
      "price": 79500
    },
    {
      "sku": "000005",
      "name": "Nathane leather sneakers",
      "category": "sneakers",
      "price": 59000
    }
  ]
}
```

- You must take into account that this list could grow to have 20.000 products.

- The prices are `integers` for example, `100.00€` would be `10000`.
- You can store the products as you see fit ( json file, in memory, rdbms of choice )

### Given that:

- Products in the boots category have a **30%** discount.
- The product with `sku = 000003` has a **15%** discount.
- When multiple discounts collide, the **biggest** discount must be applied.

### GET `/products`

- Can be filtered by `category` as a query string parameter
- (optional) Can be filtered by `priceLessThan` as a query string parameter, this filter applies before
  discounts are applied and will show products with prices lesser than or equal the value provided.
- Returns a list of Product with the given discounts applied when necessary
- Must return at most 5 elements. (The order does not matter)

Example product with a discount of 30% applied:

```json
{
  "sku": "000001",
  "name": "BV Lean leather ankle boots",
  "category": "boots",
  "price": {
    "original": 89000,
    "final": 62300,
    "discount_percentage": "30%",
    "currency": "EUR"
  }
}
```

Example product without a discount:

```json
{
  "sku": "000001",
  "name": "BV Lean leather ankle boots",
  "category": "boots",
  "price": {
    "original": 89000,
    "final": 89000,
    "discount_percentage": null,
    "currency": "EUR"
  }
}
```

---

## Answer :

## 📜 Description

This Project Implemented Based on Clean Architecture in Golang.

🔰 Rule of Clean Architecture by Uncle Bob

* Independent of Frameworks. The architecture does not depend on the existence of some library of feature laden
  software. This allows you to use such frameworks as tools, rather than having to cram your system into their limited
  constraints.
* Testable. The business rules can be tested without the UI, Database, Web Server, or any other external element.
* Independent of UI. The UI can change easily, without changing the rest of the system. A Web UI could be replaced with
  a console UI, for example, without changing the business rules.
* Independent of Database. You can swap out Oracle or SQL Server, for Mongo, BigTable, CouchDB, or something else. Your
  business rules are not bound to the database.
* Independent of any external agency. In fact your business rules simply don’t know anything at all about the outside
  world.

📚 More at [Uncle Bob clean-architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)

I was inspired by the strategy pattern to apply discounts on products, so I wrote a method on the entity to apply the
discount of each product.<br>
Also, instead of using a database, memory was used and a map was used to simulate the database indexing.

### Requirements

- [docker](https://docs.docker.com/engine/install/ubuntu/)
- [docker compose plugin](https://docs.docker.com/compose/install/linux/)

### How To Run This Project

Since the project already use Go Module, I recommend to put the source code in any folder except `GOPATH`.
Please clone the project and run the command below :

```shell
# clone the project 
git clone git@github.com:hamidteimouri/go-shop.git

# change directory 
cd go-shop

# run the project via docker
docker compose up -d
```

Then you can see the products via command below :

```shell
curl --request GET --url 'localhost:8000/products'
```

If you want to use another port, check the `project_path/.env` file.

### How to run tests

I just wrote tests just for domain of the project. To run the tests, you can proceed in two ways:

#### The first one

```shell
# change directory to the src
cd go-shop/src

# run tests
go test ./...
```

#### The second one (if you want check the coverage of domain tests, use command below)
```shell
# change directory
cd go-shop/src/internal/domain

# run the tests
go test -coverprofile cover.out
```