# WPG Ecommerce System
Project page: https://github.com/users/izharishaksa/projects/6/views/1

## Problem description
Simple e-commerce system.

## Run instructions
1. Run `docker-compose up`
2. Please wait until all services running
3. Test services using `wpg-ecommerce-system.postman_collection.json`

## Test Case Scenario
1. Create product `POST /products`
2. Register customer `POST /customers`
3. Create order `POST /orders`, order status is `placed`, event `ORDER_PLACED` is sent
4. `ORDER_PLACED` is consumed by `inventory-service`, if inventory is enough or product is exist, event `ORDER_CREATED` is sent, otherwise `ORDER_REJECTED`. Stock and sold are updated accordingly.
5. `ORDER_CREATED` is consumed by `order-service`, status and total price is updated
6. `ORDER_REJECTED` is consumed by `order-service`, status is updated
